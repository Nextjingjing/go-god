package repository

import (
	"github.com/Nextjingjing/go-god/14-testing/internal/core/domain"
	"github.com/Nextjingjing/go-god/14-testing/internal/core/ports"
)

type myMockRepository struct {
	data []domain.Product
}

func NewMyMockRepository() ports.ProductRepository {
	return &myMockRepository{data: []domain.Product{}}
}

func (m *myMockRepository) Save(p *domain.Product) (*domain.Product, error) {
	p.Id = uint(len(m.data) + 1)
	m.data = append(m.data, *p)
	return p, nil
}

func (m *myMockRepository) FindById(id uint) (*domain.Product, error) {
	return &m.data[id-1], nil
}
