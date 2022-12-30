package repository

import (
	"database/sql"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
)

type adminRepo struct {
	db *sql.DB
}

// ActivateUser implements interfaces.AdminRepository
func (*adminRepo) ActivateUser(id int) (domain.UserResponse, error) {
	panic("unimplemented")
}

// BlockUser implements interfaces.AdminRepository
func (*adminRepo) BlockUser(id int) (domain.UserResponse, error) {
	panic("unimplemented")
}

// FindAdmin implements interfaces.AdminRepository
func (*adminRepo) FindAdmin(email string) (domain.AdminResponse, error) {
	panic("unimplemented")
}

// ListBlockedUsers implements interfaces.AdminRepository
func (*adminRepo) ListBlockedUsers() ([]domain.UserResponse, error) {
	panic("unimplemented")
}

// ListNewUsers implements interfaces.AdminRepository
func (*adminRepo) ListNewUsers() ([]domain.UserResponse, error) {
	panic("unimplemented")
}

// ListUsers implements interfaces.AdminRepository
func (*adminRepo) ListUsers() ([]domain.UserResponse, error) {
	panic("unimplemented")
}

// StoreVerificationDetails implements interfaces.AdminRepository
func (*adminRepo) StoreVerificationDetails(string, int) error {
	panic("unimplemented")
}

func NewAdminRepo(db *sql.DB) interfaces.AdminRepository {
	return &adminRepo{
		db: db,
	}
}
