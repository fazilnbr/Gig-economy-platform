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

// UpdateJobCategory implements interfaces.AdminUseCase
func (c *adminUseCase) UpdateJobCategory(category domain.Category) (int, error) {
	id, err := c.adminRepo.UpdateJobCategory(category)

	return id, err
}

// ListJobCategory implements interfaces.AdminUseCase
func (c *adminUseCase) ListJobCategory(pagenation utils.Filter) (*[]domain.Category, utils.Metadata, error) {
	categories, metadata, err := c.adminRepo.ListJobCategory(pagenation)

	return &categories, metadata, err
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
func (c *adminUseCase) ListBlockedWorkers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {
	users, metadata, err := c.adminRepo.ListBlockedWorkers(pagenation)

	return &users, &metadata, err
}

// ListNewUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListNewWorkers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {

	users, metadata, err := c.adminRepo.ListNewWorkers(pagenation)

	return &users, &metadata, err
}

// ListUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListWorkers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {
	users, metadata, err := c.adminRepo.ListWorkers(pagenation)

	return &users, &metadata, err
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
func (c *adminUseCase) ListBlockedUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {
	users, metadata, err := c.adminRepo.ListBlockedUsers(pagenation)

	return &users, &metadata, err
}

// ListNewUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListNewUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {
	users, metadata, err := c.adminRepo.ListNewUsers(pagenation)

	return &users, &metadata, err
}

// ListUsers implements interfaces.AdminUseCase
func (c *adminUseCase) ListUsers(pagenation utils.Filter) (*[]domain.UserResponse, *utils.Metadata, error) {
	users, metadata, err := c.adminRepo.ListUsers(pagenation)

	return &users, &metadata, err
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
