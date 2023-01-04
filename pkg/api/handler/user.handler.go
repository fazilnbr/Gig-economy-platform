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
// @Produce json
// @Param        name   path      string  true  "User Name : "
// @Param        gender   path      string  true  "Gender : "
// @Param        dateofbirth   path      string  true  "Date Of Birth : "
// @Param        housename   path      string  true  "House Name : "
// @Param        place   path      string  true  "Place : "
// @Param        post   path      string  true  "Post : "
// @Param        pin   path      string  true  "Pin : "
// @Param        contactnumber   path      string  true  "Contact Number : "
// @Param        emailid   path      string  true  "Email Id : "
// @Param        photo   path      string  true  "Photo : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Profile}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/login [post]
func (cr *UserHandler) AddProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
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
// @Produce json
// @Param        name   path      string  true  "User Name : "
// @Param        gender   path      string  true  "Gender : "
// @Param        dateofbirth   path      string  true  "Date Of Birth : "
// @Param        housename   path      string  true  "House Name : "
// @Param        place   path      string  true  "Place : "
// @Param        post   path      string  true  "Post : "
// @Param        pin   path      string  true  "Pin : "
// @Param        contactnumber   path      string  true  "Contact Number : "
// @Param        emailid   path      string  true  "Email Id : "
// @Param        photo   path      string  true  "Photo : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Profile}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/login [post]
func (cr *UserHandler) EditProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var userprofile domain.Profile

	c.Bind(&userprofile)

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", userprofile, id)

	err := cr.userService.EditProfile(userprofile, id)

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
