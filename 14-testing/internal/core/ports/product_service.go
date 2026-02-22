package ports

import "github.com/Nextjingjing/go-god/14-testing/internal/core/domain"

type ProductService interface {
	Create(p *domain.Product) (*domain.Product, error)
	GetById(id uint) (*domain.Product, error)
}
