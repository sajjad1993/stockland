package presenter

import (
	"stockland/service/product/domain"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsAdmin   bool   `json:"is_admin"`
	Role      string `json:"role"`
	Username  string `json:"username"`
}

// swagger:model LoginRequest
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessRequest struct {
	Refresh string `json:"refresh"`
}

type LoginResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	IsAdmin bool   `json:"is_admin"`
}

type GetUserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsAdmin   bool   `json:"is_admin"`
	Role      string `json:"role"`
	Username  string `json:"username"`
}

type GetAllUsersResponse struct {
	Users []*GetUserResponse `json:"users"`
}

// swagger:model CreateProductRequest
type RequestProduct struct {
	Prompt string `json:"Prompt"`
	Image  string `json:"Image"`
}

type CreateUserResponse struct {
	ID uint `json:"id"`
}

// swagger:model UpdateUserRequest
type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsAdmin   bool   `json:"is_admin"`
	Password  string `json:"password"`
}

type UpdateUserResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsAdmin   bool   `json:"is_admin"`
	Role      string `json:"role"`
	Username  string `json:"username"`
}

func NewUserFromEntity(user *domain.User) *User {
	return &User{
		ID: user.ID,
	}
}
