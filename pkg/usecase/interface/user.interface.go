package interfaces

import "github.com/fazilnbr/project-workey/pkg/domain"

type UserService interface {
	CreateUser(newUser domain.Login) error
	FindUser(email string) (*domain.UserResponse, error)
}
