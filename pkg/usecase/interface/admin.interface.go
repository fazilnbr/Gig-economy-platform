package interfaces

import (
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type AdminUseCase interface {
	FindAdmin(email string) (*domain.AdminResponse, error)
	SendVerificationEmail(email string) error
	VerifyAccount(email string, code int) error
	ListNewUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	ListBlockedUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	ListUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	ActivateUser(id int) (*domain.UserResponse, error)
	BlockUser(id int) (*domain.UserResponse, error)
	ListNewWorkers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	ListBlockedWorkers() (*[]domain.UserResponse, error)
	ListWorkers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error)
	ActivateWorker(id int) (*domain.UserResponse, error)
	BlockWorker(id int) (*domain.UserResponse, error)
	AddJobCategory(category string) error
}
