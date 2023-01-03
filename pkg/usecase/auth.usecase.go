package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fazilnbr/project-workey/pkg/config"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type authUseCase struct {
	adminRepo  interfaces.AdminRepository
	workerRepo interfaces.WorkerRepository
	userRepo   interfaces.UserRepository
	mailConfig config.MailConfig
	config     config.Config
}

// WorkerVerifyAccount implements interfaces.AuthUseCase
func (c *authUseCase) WorkerVerifyAccount(email string, code int) error {
	err := c.workerRepo.VerifyAccount(email, code)
	return err
}

// VerifyAccount implements interfaces.AuthUseCase
func (c *authUseCase) UserVerifyAccount(email string, code int) error {
	err := c.userRepo.VerifyAccount(email, code)

	return err
}

// SendVerificationEmail implements interfaces.AuthUseCase
func (c *authUseCase) SendVerificationEmail(email string) error {

	//to generate random code
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999)

	message := fmt.Sprintf(
		"\nThe verification code is:\n\n%d.\nUse to verify your account.\n Thank you for using Workey.\n with regards Team Workey.",
		code,
	)

	// send random code to user's email
	if err := c.mailConfig.SendMail(c.config, email, message); err != nil {
		return err
	}

	err := c.userRepo.StoreVerificationDetails(email, code)

	if err != nil {
		return err
	}

	return nil
}

// VerifyAdmin implements interfaces.AuthUseCase
func (c *authUseCase) VerifyAdmin(email string, password string) error {
	admin, err := c.adminRepo.FindAdmin(email)

	if err != nil {
		return errors.New("failed to login. check your email")
	}

	// isValidPassword := VerifyPassword(password, user.Password)

	if admin.Password != password {
		return errors.New("failed to login. check your credential")
	}

	return nil
}

// VerifyUser implements interfaces.AuthUseCase
func (c *authUseCase) VerifyUser(email string, password string) error {
	user, err := c.userRepo.FindUser(email)
	fmt.Print("\n\n", user, err)

	if err != nil {
		return errors.New("failed to login. check your email")
	}

	isValidPassword := VerifyPassword(password, user.Password)
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil
}

// VerifyWorker implements interfaces.AuthUseCase
func (c *authUseCase) VerifyWorker(email string, password string) error {
	user, err := c.workerRepo.FindWorker(email)
	fmt.Print("\n\n", user, err)

	if err != nil {
		return errors.New("failed to login. check your email")
	}

	isValidPassword := VerifyPassword(password, user.Password)
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil
}

func NewAuthService(
	adminRepo interfaces.AdminRepository,
	workerRepo interfaces.WorkerRepository,
	userRepo interfaces.UserRepository,
	mailConfig config.MailConfig,
	config config.Config,
) services.AuthUseCase {
	return &authUseCase{
		adminRepo:  adminRepo,
		workerRepo: workerRepo,
		userRepo:   userRepo,
		mailConfig: mailConfig,
		config:     config,
	}
}
