# บทที่ 14 Unit Testing

## สิ่งที่ต้องรู้มาก่อน
- พื้นฐานภาษา GO
- บทที่ 11 Hexagonal architecture

## Unit test 
เมื่อเราได้ทำ Hexagonal architecture เราจะเห็นประโยชน์ของมันในเรื่องการ Testing เนื่องจากเราสามารถทดสอบที่ Core ที่เป็น Pure Business Logic จริงๆ หมายความว่า (เช่น) เราไม่จำเป็นต้องเชื่อมต่อ Database เพื่อทดสอบเลย หากมีการเชื่อมต่อ Database จะไม่เรียกว่า Unit test แล้ว

### Business logic
```go
type productDomainService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &productDomainService{repo: repo}
}

func (s *productDomainService) Create(p *domain.Product) (*domain.Product, error) {
	if p.Price <= 0 {
		return nil, errors.New("Price Should be more than 0.")
	}
	p.Id = 0
	return s.repo.Save(p)
}

func (s *productDomainService) GetById(id uint) (*domain.Product, error) {
	return s.repo.FindById(id)
}

```

## การ Test ด้วย Testify

```bash
go get github.com/stretchr/testify
```

ไฟล์ `create_test.go`
- สำคัญคือต้องลงท้ายด้วย `_test.go` เพื่อให้ Go ทราบว่าเป็นไฟล์ Test
- ผมได้ `repository.NewMyMockRepository()` เป็น Mock Database นั้นเอง
```go
package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Nextjingjing/go-god/14-testing/internal/adapters/repository"
	"github.com/Nextjingjing/go-god/14-testing/internal/core/domain"
	"github.com/Nextjingjing/go-god/14-testing/internal/core/services"
)

func TestCreateTableDriven(t *testing.T) {
	tests := []struct {
		name      string
		product   *domain.Product
		wantErr   bool
		wantID    uint
		wantName  string
		wantPrice uint
	}{
		{
			name: "create success",
			product: &domain.Product{
				Name:  "Book",
				Price: 100,
			},
			wantErr:   false,
			wantID:    1,
			wantName:  "Book",
			wantPrice: 100,
		},
		{
			name: "price zero",
			product: &domain.Product{
				Name:  "Book",
				Price: 0,
			},
			wantErr: true,
		},
		{
			name: "id should reset before save",
			product: &domain.Product{
				Id:    999,
				Name:  "Pen",
				Price: 10,
			},
			wantErr: false,
			wantID:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewMyMockRepository()
			service := services.NewProductService(repo)

			res, err := service.Create(tt.product)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantID, res.Id)

			if tt.wantName != "" {
				assert.Equal(t, tt.wantName, res.Name)
			}

			if tt.wantPrice != 0 {
				assert.Equal(t, tt.wantPrice, res.Price)
			}
		})
	}
}

func TestCreateMultipleProducts(t *testing.T) {
	repo := repository.NewMyMockRepository()
	service := services.NewProductService(repo)

	products := []*domain.Product{
		{Name: "Book", Price: 100},
		{Name: "Pen", Price: 10},
		{Name: "Bag", Price: 500},
	}

	var results []*domain.Product

	for _, p := range products {
		res, err := service.Create(p)
		assert.NoError(t, err)
		results = append(results, res)
	}

	assert.Equal(t, uint(1), results[0].Id)
	assert.Equal(t, uint(2), results[1].Id)
	assert.Equal(t, uint(3), results[2].Id)
}

```

```bash
go test ./... # รัน Test ทุกไฟล์
go test -v ./... # รัน Test พร้อมดูรายละเอียด (Verbose)
go test -v -run TestFindById ./... # รันเฉพาะฟังก์ชันที่ระบุ
go test -cover ./... # การตรวจสอบประสิทธิภาพ
go test -v ./internal/core/services/product_service_test.go # รันเฉพาะไฟล์
```

## การ Mock ด้วย Testify
```go
package repository

import (
	"github.com/Nextjingjing/go-god/14-testing/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type TestifyMockRepository struct {
	mock.Mock
}

func (m *TestifyMockRepository) Save(p *domain.Product) (*domain.Product, error) {
	args := m.Called(p)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *TestifyMockRepository) FindById(id uint) (*domain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

```

```go
type TestifyMockRepository struct {
	mock.Mock
}
```

- Embed `mock.Mock` เข้าไป

```go
func (m *TestifyMockRepository) Save(p *domain.Product) (*domain.Product, error) {
	args := m.Called(p)
	return args.Get(0).(*domain.Product), args.Error(1)
}
```

- สร้าง Function

ไฟล์ `get_by_id_test.go`
```go
package services_test

import (
	"testing"

	"github.com/Nextjingjing/go-god/14-testing/internal/adapters/repository"
	"github.com/Nextjingjing/go-god/14-testing/internal/core/domain"
	"github.com/Nextjingjing/go-god/14-testing/internal/core/services"
	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	mockRepo := new(repository.TestifyMockRepository)
	service := services.NewProductService(mockRepo)

	expectedProduct := &domain.Product{Id: 1, Name: "Gadget"}
	mockRepo.On("FindById", uint(1)).Return(expectedProduct, nil)

	result, err := service.GetById(1)
	assert.Equal(t, expectedProduct, result)
	assert.NotNil(t, result)
	assert.NoError(t, err)
}

```

```go
mockRepo := new(repository.TestifyMockRepository)
service := services.NewProductService(mockRepo)
expectedProduct := &domain.Product{Id: 1, Name: "Gadget"}
mockRepo.On("FindById", uint(1)).Return(expectedProduct, nil)
```

- เราสามารถกำหนดได้ว่าเมื่อ Function ถูกเรียกเรากำหนดได้ว่าจะให้ Return อะไร