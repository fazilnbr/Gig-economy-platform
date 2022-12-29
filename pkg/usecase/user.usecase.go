package usecase

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type userService struct {
	userRepo interfaces.UserRepository
}

func NewUserService(
	userRepo interfaces.UserRepository) services.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// CreateUser implements interfaces.UserService
func (c *userService) CreateUser(newUser domain.Login) error {
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

func HashPassword(password string) string {
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password

}

// FindUser implements interfaces.UserService
func (c *userService) FindUser(email string) (*domain.UserResponse, error) {
	user, err := c.userRepo.FindUser(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
