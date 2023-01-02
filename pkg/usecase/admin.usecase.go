package usecase

import (
	"github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type adminUseCase struct {
	adminRepo  interfaces.AdminRepository
	mailConfig config.MailConfig
}

// ActivateUser implements interfaces.AdminUseCase
func (*adminUseCase) ActivateUser(id int) (*domain.UserResponse, error) {
	panic("unimplemented")
}

// BlockUser implements interfaces.AdminUseCase
func (*adminUseCase) BlockUser(id int) (*domain.UserResponse, error) {
	panic("unimplemented")
}

// FindAdmin implements interfaces.AdminUseCase
func (c *adminUseCase) FindAdmin(email string) (*domain.AdminResponse, error) {
	admin, err := c.adminRepo.FindAdmin(email)

	return &admin, err
}

// ListBlockedUsers implements interfaces.AdminUseCase
func (*adminUseCase) ListBlockedUsers() (*[]domain.UserResponse, error) {
	panic("unimplemented")
}

// ListNewUsers implements interfaces.AdminUseCase
func (*adminUseCase) ListNewUsers() (*[]domain.UserResponse, error) {
	panic("unimplemented")
}

// ListUsers implements interfaces.AdminUseCase
func (*adminUseCase) ListUsers() (*[]domain.UserResponse, error) {
	panic("unimplemented")
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
