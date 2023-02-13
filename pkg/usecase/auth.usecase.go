package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/fazilnbr/project-workey/pkg/config"
	interfaces "github.com/fazilnbr/project-workey/pkg/repository/interface"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/golang-jwt/jwt/v4"
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

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	// //to generate random code
	// rand.Seed(time.Now().UnixNano())
	// code := rand.Intn(999999)

	// message := fmt.Sprintf(
	// 	"\nThe verification code is:\n\n%d.\nUse to verify your account.\n Thank you for using Workey.\n with regards Team Workey.",
	// 	code,
	// )

	subject := "Account Verification"
	message :=
		"From: Events Radar <eventsRadarversion1@gmail.com>\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
			"<html>" +
			"  <head>" +
			"    <style>" +
			"      .blue-button {" +
			"        background-color: blue;" +
			"        color: white;" +
			"        padding: 10px 20px;" +
			"        border-radius: 5px;" +
			"        text-decoration: none;" +
			"        font-size: 16px;" +
			"      }" +
			"    </style>" +
			"  </head>" +
			"  <body>" +
			"    <p>Click the button on verify your accout:</p>" +
			"    <a class=\"blue-button\" href=\"http://localhost:3000/user/verify-account?token=" + tokenString + "\" target=\"_blank\">Access Credentials</a>" +
			"  </body>" +
			"</html>"

	// send random code to user's email
	if err := c.mailConfig.SendMail(c.config, email, message); err != nil {
		return err
	}

	err = c.userRepo.StoreVerificationDetails(email, tokenString)

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
	if !user.Verification {
		return errors.New("failed to login. you are not verified")
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
