package user

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	GetUserByID(id int) (*User, error)
}

// UserService struct that uses UserRepository
type UserService struct {
	Repo UserRepository
}

// GetUser method to fetch user by ID
func (s *UserService) GetUser(id int) (*User, error) {
	return s.Repo.GetUserByID(id)
}

// Mock implementation of UserRepository for testing
type MockUserRepository struct {
	Users map[int]User
}

func (m *MockUserRepository) GetUserByID(id int) (*User, error) {
	user, exists := m.Users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

type PostgresUserRepository struct {
	// db *sql.DB // Assume we have a database connection here
}

func (p *PostgresUserRepository) GetUserByID(id int) (*User, error) {
	// Simulate database fetch
	return &User{ID: id, Name: "DB User", Email: "dbuser@example.com"}, nil
}
