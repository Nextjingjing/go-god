package services

import (
	"github.com/Nextjingjing/go-god/11-hexagonal/internal/core/domain"
	"github.com/Nextjingjing/go-god/11-hexagonal/internal/core/ports"
)

// Implementation of UserService
// Best practice: keep implementation privately hidden
type userServiceImpl struct {
	repo ports.UserRepository
}

func NewUserServiceImpl(repo ports.UserRepository) ports.UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) CreateUser(name string) (*domain.User, error) {
	user := &domain.User{Name: name}
	return s.repo.Save(user)
}

func (s *userServiceImpl) GetUserByID(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *userServiceImpl) GetAllUsers() ([]*domain.User, error) {
	return s.repo.FindAll()
}

func (s *userServiceImpl) UpdateUser(id string, name string) (*domain.User, error) {
	return s.repo.Save(&domain.User{ID: id, Name: name})
}
