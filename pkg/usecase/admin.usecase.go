package usecase

import (
	"github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type adminUseCase struct {
	adminRepo  interfaces.AdminRepository
	mailConfig config.MailConfig
}

// AddJobCategory implements interfaces.AdminUseCase
func (c *adminUseCase) AddJobCategory(category string) error {
	err := c.adminRepo.AddJobCategory(category)
	return err
}

// ActivateUser implements interfaces.AdminUseCase
func (c *adminUseCase) ActivateWorker(id int) (*domain.UserResponse, error) {
	user, err := c.adminRepo.ActivateWorker(id)

	return &user, err
}

// BlockUser implements interfaces.AdminUseCase
func (c *adminUseCase) BlockWorker(id int) (*domain.UserResponse, error) {
	user, err := c.adminRepo.BlockWorker(id)

	return &user, err
}

// ListBlockedUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListBlockedWorkers() (*[]domain.UserResponse, error) {
	users, err := c.adminRepo.ListBlockedWorkers()

	return &users, err
}

// ListNewUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListNewWorkers() (*[]domain.UserResponse, error) {
	users, err := c.adminRepo.ListNewWorkers()

	return &users, err
}

// ListUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListWorkers() (*[]domain.UserResponse, error) {
	users, err := c.adminRepo.ListWorkers()

	return &users, err
}

// ActivateUser implements interfaces.AdminUseCase
func (c *adminUseCase) ActivateUser(id int) (*domain.UserResponse, error) {
	user, err := c.adminRepo.ActivateUser(id)

	return &user, err
}

// BlockUser implements interfaces.AdminUseCase
func (c *adminUseCase) BlockUser(id int) (*domain.UserResponse, error) {
	user, err := c.adminRepo.BlockUser(id)

	return &user, err
}

// FindAdmin implements interfaces.AdminUseCase
func (c *adminUseCase) FindAdmin(email string) (*domain.AdminResponse, error) {
	admin, err := c.adminRepo.FindAdmin(email)

	return &admin, err
}

// ListBlockedUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListBlockedUsers() (*[]domain.UserResponse, error) {
	users, err := c.adminRepo.ListBlockedUsers()

	return &users, err
}

// ListNewUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListNewUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {
	users, metadata, err := c.adminRepo.ListNewUsers(pagenation)

	return &users, &metadata, err
}

// ListUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListUsers() (*[]domain.UserResponse, error) {
	users, err := c.adminRepo.ListUsers()

	return &users, err
}

// SendVerificationEmail implements interfaces.AdminUseCase
func (*adminUseCase) SendVerificationEmail(email string) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.AdminUseCase
func (*adminUseCase) VerifyAccount(email string, code int) error {
	panic("unimplemented")
}

func NewAdminService(
	adminRepo interfaces.AdminRepository,
	mailConfig config.MailConfig) services.AdminUseCase {
	return &adminUseCase{
		adminRepo:  adminRepo,
		mailConfig: mailConfig,
	}
}
