package usecase

import (
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type authUseCase struct {
	userRepo interfaces.UserRepository
}

// VerifyAdmin implements interfaces.AuthUseCase
func (*authUseCase) VerifyAdmin(email string, password string) error {
	panic("unimplemented")
}

// VerifyUser implements interfaces.AuthUseCase
func (*authUseCase) VerifyUser(email string, password string) error {
	panic("unimplemented")
}

// VerifyWorker implements interfaces.AuthUseCase
func (*authUseCase) VerifyWorker(email string, password string) error {
	panic("unimplemented")
}

func NewAuthService(
	userRepo interfaces.UserRepository,
) services.AuthUseCase {
	return &authUseCase{
		userRepo: userRepo,
	}
}
