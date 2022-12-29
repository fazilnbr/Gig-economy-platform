package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
)

type UserRepository interface {
	FindUser(email string) (domain.UserResponse, error)
	InsertUser(login domain.Login) (int, error)
}
