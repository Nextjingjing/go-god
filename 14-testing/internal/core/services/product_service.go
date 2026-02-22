package services

import (
	"errors"

	"github.com/Nextjingjing/go-god/14-testing/internal/core/domain"
	"github.com/Nextjingjing/go-god/14-testing/internal/core/ports"
)

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
