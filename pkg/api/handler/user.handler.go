package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fazilnbr/project-workey/pkg/common/response"
	"github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserUseCase
}

func NewUserHandler(userService services.UserUseCase) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

// @Summary Add profile for User
// @ID user add profile
// @Tags User
// @Produce json
// @Param        name   query      string  true  "User Name : "
// @Param        gender   query      string  true  "Gender : "
// @Param        dateofbirth   query      string  true  "Date Of Birth : "
// @Param        housename   query      string  true  "House Name : "
// @Param        place   query      string  true  "Place : "
// @Param        post   query      string  true  "Post : "
// @Param        pin   query      string  true  "Pin : "
// @Param        contactnumber   query      string  true  "Contact Number : "
// @Param        emailid   query      string  true  "Email Id : "
// @Param        photo   query      string  true  "Photo : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/addprofile [post]
func (cr *UserHandler) UserAddProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))

	fmt.Printf("\n\nidea : %v\n\n", id)
	var userprofile domain.Profile

	c.Bind(&userprofile)

	err := cr.userService.AddProfile(userprofile, id)

	if err != nil {
		response := response.ErrorResponse("Error while adding profile", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", userprofile)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Edit profile for User
// @ID user edit profile
// @Tags User
// @Produce json
// @Param        name   query      string  true  "User Name : "
// @Param        gender   query      string  true  "Gender : "
// @Param        dateofbirth   query      string  true  "Date Of Birth : "
// @Param        housename   query      string  true  "House Name : "
// @Param        place   query      string  true  "Place : "
// @Param        post   query      string  true  "Post : "
// @Param        pin   query      string  true  "Pin : "
// @Param        contactnumber   query      string  true  "Contact Number : "
// @Param        emailid   query      string  true  "Email Id : "
// @Param        photo   query      string  true  "Photo : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/editprofile [patch]
func (cr *UserHandler) UserEditProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	var userprofile domain.Profile

	c.Bind(&userprofile)

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", userprofile, id)

	err := cr.userService.UserEditProfile(userprofile, id)

	if err != nil {
		response := response.ErrorResponse("Error while editing profile", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", userprofile)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Change Password for User
// @ID user change password
// @Tags User
// @Produce json
// @Param        name   query      string  true  "User Name : "
// @Param        gender   query      string  true  "Gender : "
// @Param        dateofbirth   query      string  true  "Date Of Birth : "
// @Param        housename   query      string  true  "House Name : "
// @Param        place   query      string  true  "Place : "
// @Param        post   query      string  true  "Post : "
// @Param        pin   query      string  true  "Pin : "
// @Param        contactnumber   query      string  true  "Contact Number : "
// @Param        emailid   query      string  true  "Email Id : "
// @Param        photo   query      string  true  "Photo : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/changepassword [patch]
func (cr *UserHandler) UserChangePassword(c *gin.Context) {

	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))

	fmt.Println("id : ", id)
	var changepassword domain.ChangePassword

	err := c.Bind(&changepassword)
	if err != nil {
		fmt.Println("pooooo : ", err)
	}

	err = cr.userService.UserVerifyPassword(changepassword, id)

	if err != nil {
		response := response.ErrorResponse("Wrong Email id or Password", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	err = cr.userService.UserChangePassword(changepassword.NewPassword, id)

	if err != nil {
		response := response.ErrorResponse("Error while changing Password", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	changepassword.NewPassword = ""
	changepassword.OldPassword = ""
	response := response.SuccessResponse(true, "SUCCESS", changepassword)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}
