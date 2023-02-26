package usecase

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/fazilnbr/project-workey/pkg/domain"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

// UpdateJobComplition implements interfaces.UserUseCase
func (c *userUseCase) UpdateJobComplition(userId int, requestId int) error {
	return c.userRepo.UpdateJobComplition(userId,requestId)
}

// ViewSendRequest implements interfaces.UserUseCase
func (c *userUseCase) ViewSendOneRequest(userId int, requestId int) (*domain.RequestUserResponse, error) {
	request, err := c.userRepo.ViewSendOneRequest(userId, requestId)
	return &request, err
}

// ListSendRequests implements interfaces.UserUseCase
func (c *userUseCase) ListSendRequests(pagenation utils.Filter, id int) (*[]domain.RequestUserResponse, *utils.Metadata, error) {
	requests, metadata, err := c.userRepo.ListSendRequests(pagenation, id)
	return &requests, &metadata, err
}

// DeleteJobRequest implements interfaces.UserUseCase
func (c *userUseCase) DeleteJobRequest(requestId int, userid int) error {
	return c.userRepo.DeleteJobRequest(requestId, userid)
}

// SendJobRequest implements interfaces.UserUseCase
func (c *userUseCase) SendJobRequest(request domain.Request) (int, error) {
	id, err := c.userRepo.CheckInRequest(request)
	if err != nil {
		return id, err
	}
	if id != 0 {
		return id, errors.New("You already added this job request")
	}
	id, err = c.userRepo.SendJobRequest(request)
	return id, err
}

// DeleteAddress implements interfaces.UserUseCase
func (c *userUseCase) DeleteAddress(id int, userid int) error {
	err := c.userRepo.DeleteAddress(id, userid)

	return err
}

// ListAddress implements interfaces.UserUseCase
func (c *userUseCase) ListAddress(id int) (*[]domain.Address, error) {
	address, err := c.userRepo.ListAddress(id)

	return &address, err
}

// AddAddress implements interfaces.UserUseCase
func (c *userUseCase) AddAddress(address domain.Address) (int, error) {
	id, err := c.userRepo.AddAddress(address)

	return id, err
}

// ListFevorite implements interfaces.UserUseCase
func (c *userUseCase) ListFevorite(pagenation utils.Filter, id int) (*[]domain.ListFavorite, *utils.Metadata, error) {
	favorites, metadata, err := c.userRepo.ListFevorite(pagenation, id)

	return &favorites, &metadata, err
}

// AddToFavorite implements interfaces.UserUseCase
func (c *userUseCase) AddToFavorite(favorite domain.Favorite) (int, error) {
	id, err := c.userRepo.CheckInFevorite(favorite)
	if err != nil {
		return id, err
	}
	if id != 0 {
		return id, errors.New("You already added this job to your list")
	}
	id, err = c.userRepo.AddToFavorite(favorite)
	return id, err
}

// SearchWorkersWithJob implements interfaces.UserUseCase
func (c *userUseCase) SearchWorkersWithJob(pagenation utils.Filter, key string) (*[]domain.ListJobsWithWorker, *utils.Metadata, error) {
	jobs, metadata, err := c.userRepo.SearchWorkersWithJob(pagenation, key)

	return &jobs, &metadata, err
}

// ListWorkersWithJob implements interfaces.UserUseCase
func (c *userUseCase) ListWorkersWithJob(pagenation utils.Filter) (*[]domain.ListJobsWithWorker, *utils.Metadata, error) {
	jobs, metadata, err := c.userRepo.ListWorkersWithJob(pagenation)

	return &jobs, &metadata, err
}

// VerifyPassword implements interfaces.UserUseCase
func (c *userUseCase) UserVerifyPassword(changepassword domain.ChangePassword, id int) error {
	user, err := c.userRepo.FindUser(changepassword.Email)
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
func (c *userUseCase) UserChangePassword(changepassword string, id int) error {
	//hashing password
	changepassword = HashPassword(changepassword)
	_, err := c.userRepo.UserChangePassword(changepassword, id)

	return err

}

// EditProfile implements interfaces.UserUseCase
func (c *userUseCase) UserEditProfile(userProfile domain.Profile, id int) error {
	_, err := c.userRepo.UserEditProfile(userProfile, id)

	return err

}

// AddProfile implements interfaces.UserUseCase
func (c *userUseCase) AddProfile(userProfile domain.Profile, id int) error {
	_, err := c.userRepo.UserAddProfile(userProfile, id)

	return err

}

// CreateUser implements interfaces.UserService
func (c *userUseCase) CreateUser(newUser domain.User) error {
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

func NewUserService(
	userRepo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
