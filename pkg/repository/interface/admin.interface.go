package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
)

type AdminRepository interface {
	FindAdmin(email string) (domain.AdminResponse, error)
	StoreVerificationDetails(string, int) error
	ListNewUsers() ([]domain.UserResponse, error)
	ListBlockedUsers() ([]domain.UserResponse, error)
	ListUsers() ([]domain.UserResponse, error)
	ActivateUser(id int) (domain.UserResponse, error)
	BlockUser(id int) (domain.UserResponse, error)
}
