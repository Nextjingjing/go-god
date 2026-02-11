package dto

type UserRequestDTO struct {
	Name string `json:"name" validate:"required"`
}
