package sqlite

import (
	"github.com/Nextjingjing/go-god/11-hexagonal/internal/core/domain"
	"gorm.io/gorm"
)

// Implementation of UserRepository
// Best practice: keep implementation privately hidden
type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Save(user *domain.User) (*domain.User, error) {
	userModel := &UserModel{
		Name: user.Name,
	}
	if err := r.db.Create(userModel).Error; err != nil {
		return nil, err
	}
	return &domain.User{
		ID:   userModel.ID,
		Name: userModel.Name,
	}, nil
}

func (r *userRepositoryImpl) FindByID(id uint) (*domain.User, error) {
	var user UserModel
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &domain.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *userRepositoryImpl) FindAll() ([]*domain.User, error) {
	var users []*UserModel
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	var domainUsers []*domain.User
	for _, user := range users {
		domainUsers = append(domainUsers, &domain.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return domainUsers, nil
}
