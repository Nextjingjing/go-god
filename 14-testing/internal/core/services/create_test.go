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
