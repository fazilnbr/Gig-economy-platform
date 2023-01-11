package usecase

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
)

type workerService struct {
	workerRepo interfaces.WorkerRepository
}

// ListJobCategoryUser implements interfaces.WorkerUseCase
func (c *workerService) ListJobCategoryUser() ([]domain.Category, error) {
	id, err := c.workerRepo.ListJobCategoryUser()

	return id, err
}

// VerifyPassword implements interfaces.UserUseCase
func (c *workerService) WorkerVerifyPassword(changepassword domain.ChangePassword, id int) error {
	user, err := c.workerRepo.FindWorker(changepassword.Email)
	if err != nil {
		return errors.New("Invalid User")
	}

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", user, changepassword.OldPassword)

	isValidPassword := VerifyPassword(changepassword.OldPassword, user.Password)
	if !isValidPassword {
		return errors.New("Invalid Password")
	}
	return nil
}

// ChangePassword implements interfaces.UserUseCase
func (c *workerService) WorkerChangePassword(changepassword string, id int) error {
	//hashing password
	changepassword = HashPassword(changepassword)
	_, err := c.workerRepo.WorkerChangePassword(changepassword, id)

	return err

}

// EditProfile implements interfaces.UserUseCase
func (c *workerService) WorkerEditProfile(userProfile domain.Profile, id int) error {
	_, err := c.workerRepo.WorkerEditProfile(userProfile, id)

	return err

}

// AddProfile implements interfaces.WorkerUseCase
func (c *workerService) AddProfile(workerProfile domain.Profile, id int) error {
	_, err := c.workerRepo.WorkerAddProfile(workerProfile, id)

	return err
}

// CreateUser implements interfaces.WorkerUseCase
func (c *workerService) CreateUser(newWorker domain.Login) error {
	_, err := c.workerRepo.FindWorker(newWorker.UserName)

	if err == nil {
		return errors.New("Username Already Exists")
	}
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// Hashing the password

	newWorker.Password = HashPassword(newWorker.Password)

	_, err = c.workerRepo.InsertWorker(newWorker)

	return err

}

// FindWorker implements interfaces.WorkerUseCase
func (c *workerService) FindWorker(email string) (*domain.WorkerResponse, error) {
	user, err := c.workerRepo.FindWorker(email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// SendVerificationEmail implements interfaces.WorkerUseCase
func (*workerService) SendVerificationEmail(email string) error {
	panic("unimplemented")
}

// VerifyAccount implements interfaces.WorkerUseCase
func (*workerService) VerifyAccount(email string, code int) error {
	panic("unimplemented")
}

func NewWorkerService(workerRepo interfaces.WorkerRepository) services.WorkerUseCase {
	return &workerService{
		workerRepo: workerRepo,
	}
}
