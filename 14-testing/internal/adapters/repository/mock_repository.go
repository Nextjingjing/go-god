package repository

import (
	"github.com/Nextjingjing/go-god/14-testing/internal/core/domain"
	"github.com/Nextjingjing/go-god/14-testing/internal/core/ports"
)

type mockRepository struct {
	data []domain.Product
}

func NewMockRepository() ports.ProductRepository {
	return &mockRepository{data: []domain.Product{}}
}

func (m *mockRepository) Save(p *domain.Product) (*domain.Product, error) {
	p.Id = uint(len(m.data) + 1)
	m.data = append(m.data, *p)
	return p, nil
}

func (m *mockRepository) FindById(id uint) (*domain.Product, error) {
	return &m.data[id-1], nil
}
