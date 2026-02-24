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
