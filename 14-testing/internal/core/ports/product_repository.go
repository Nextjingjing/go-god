package ports

import "github.com/Nextjingjing/go-god/14-testing/internal/core/domain"

type ProductRepository interface {
	Save(p *domain.Product) (*domain.Product, error)
	FindById(id uint) (*domain.Product, error)
}
