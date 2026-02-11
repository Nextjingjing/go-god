package ports

import "github.com/Nextjingjing/go-god/11-hexagonal/internal/core/domain"

// Outbound port
type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
	FindAll() ([]*domain.User, error)
}
