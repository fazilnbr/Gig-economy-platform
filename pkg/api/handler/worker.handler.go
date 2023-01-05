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

type WorkerHandler struct {
	workerService services.WorkerUseCase
}

func NewWorkerHandler(workerService services.WorkerUseCase) WorkerHandler {
	return WorkerHandler{
		workerService: workerService,
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
func (cr *WorkerHandler) WorkerAddProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))

	fmt.Printf("\n\nidea : %v\n\n", id)
	var userprofile domain.Profile

	c.Bind(&userprofile)

	err := cr.workerService.AddProfile(userprofile, id)

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
func (cr *WorkerHandler) WorkerEditProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	fmt.Printf("\n\n%v\n\n", id)
	var userprofile domain.Profile

	c.Bind(&userprofile)

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", userprofile, id)

	err := cr.workerService.WorkerEditProfile(userprofile, id)

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
func (cr *WorkerHandler) WorkerChangePassword(c *gin.Context) {

	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))

	fmt.Println("id : ", id)
	var changepassword domain.ChangePassword

	err := c.Bind(&changepassword)
	if err != nil {
		fmt.Println("pooooo : ", err)
	}

	err = cr.workerService.WorkerVerifyPassword(changepassword, id)

	if err != nil {
		response := response.ErrorResponse("Wrong Email id or Password", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	err = cr.workerService.WorkerChangePassword(changepassword.NewPassword, id)

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
