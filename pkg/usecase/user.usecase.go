package usecase

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserService(
	userRepo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

// AddProfile implements interfaces.UserUseCase
func (c *userUseCase) AddProfile(userProfile domain.Profile, id int) error {
	_, err := c.userRepo.AddProfile(userProfile, id)

	return err

}

// CreateUser implements interfaces.UserService
func (c *userUseCase) CreateUser(newUser domain.Login) error {
	_, err := c.userRepo.FindUser(newUser.UserName)

	if err == nil {
		return errors.New("Username already exists")
	}

	//hashing password
	newUser.Password = HashPassword(newUser.Password)

	_, err = c.userRepo.InsertUser(newUser)
	// fmt.Printf("\n\n\nerr2 : %v\n\n\n", err)

	return err
}

// FindUser implements interfaces.UserService
func (c *userUseCase) FindUser(email string) (*domain.UserResponse, error) {
	user, err := c.userRepo.FindUser(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// VerifyUser implements interfaces.UserService
func (c *userUseCase) VerifyUser(email string, password string) error {
	user, err := c.userRepo.FindUser(email)

	if err != nil {
		return errors.New("failed to login. check your email")
	}

	isValidPassword := VerifyPassword(password, user.Password)
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
	}

	return nil
}

func HashPassword(password string) string {
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password

}

func VerifyPassword(requestPassword, dbPassword string) bool {

	requestPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestPassword)))
	return requestPassword == dbPassword
}
