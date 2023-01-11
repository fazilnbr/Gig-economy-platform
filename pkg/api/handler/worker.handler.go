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

// @Summary Add profile for Worker
// @ID worker add profile
// @Tags Worker
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
// @Router /worker/addprofile [post]
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

// @Summary Edit profile for Worker
// @ID worker edit profile
// @Tags Worker
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
// @Router /worker/editprofile [patch]
func (cr *WorkerHandler) WorkerEditProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	fmt.Printf("\n\n%v\n\n", id)
	var userprofile domain.Profile

	c.Bind(&userprofile)

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", userprofile, id)

	err := cr.workerService.WorkerEditProfile(userprofile, id)

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

// @Summary Change Password for worker
// @ID worker change password
// @Tags Worker
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
// @Router /worker/changepassword [patch]
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

// @Summary list all job categories for Worker
// @ID list all job category for worker
// @Tags Worker
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/listjobcategory [get]
func (cr *WorkerHandler) ListJobCategoryUser(c *gin.Context) {

	// page, err := strconv.Atoi(c.Query("page"))

	// pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	// fmt.Printf("\n\nuser : %v\n\nmetea : %v\n\n", page, c.Query("page"))
	// log.Println(page, "   ", pageSize)

	// pagenation := utils.Filter{
	// 	Page:     page,
	// 	PageSize: pageSize,
	// }

	categories, err := cr.workerService.ListJobCategoryUser()

	if err != nil {
		response := response.ErrorResponse("Failed To List Job Category", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	// result := struct {
	// 	Users *[]domain.UserResponse
	// 	Meta  *utils.Metadata
	// }{
	// 	Users: users,
	// 	Meta:  metadata,
	// }

	response := response.SuccessResponse(true, "SUCCESS", categories)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}
