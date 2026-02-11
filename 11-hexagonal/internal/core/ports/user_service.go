package ports

import "github.com/Nextjingjing/go-god/11-hexagonal/internal/core/domain"

// Inbound port
type UserService interface {
	CreateUser(name string) (*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
	UpdateUser(id string, name string) (*domain.User, error)
}
