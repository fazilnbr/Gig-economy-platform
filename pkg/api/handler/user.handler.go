package handler

import (
	"fmt"
	"log"
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
// @Security BearerAuth
// @Produce json
// @Param UserAddProfile body domain.Profile{} true "User Add Profile"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/add-profile [post]
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
// @Security BearerAuth
// @Produce json
// @Param UserEditProfile body domain.Profile{} true "User Edit Profile"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/edit-profile [patch]
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
// @Security BearerAuth
// @Produce json
// @Param UserChangePassword body domain.ChangePassword{} true "User Change Password"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/change-password [patch]
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
		response := response.ErrorResponse("Faild to verify user password", err.Error(), nil)
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

// @Summary list all job with workers for users
// @ID list all job with workers for users
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/list-workers-with-job [get]
func (cr *UserHandler) ListWorkersWithJob(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	fmt.Printf("\n\nuser : %v\n\nmetea : %v\n\n", page, pageSize)
	log.Println(page, "   ", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	jobs, metadata, err := cr.userService.ListWorkersWithJob(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list worker with job", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	result := struct {
		Users *[]domain.ListJobsWithWorker
		Meta  *utils.Metadata
	}{
		Users: jobs,
		Meta:  metadata,
	}

	response := response.SuccessResponse(true, "SUCCESS", result)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary search job with workers for users
// @ID search job with workers for users
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param        search   query      string  true  "search : "
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/search-workers-with-job [get]
func (cr *UserHandler) SearchWorkersWithJob(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	searchkey := c.Query("search")
	fmt.Printf("\n\nuser : %v\n\nmetea : %v\n\n", page, pageSize)
	log.Println(page, "   ", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	jobs, metadata, err := cr.userService.SearchWorkersWithJob(pagenation, searchkey)

	if err != nil {
		response := response.ErrorResponse("Failed to search worker with job", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	result := struct {
		Users *[]domain.ListJobsWithWorker
		Meta  *utils.Metadata
	}{
		Users: jobs,
		Meta:  metadata,
	}

	response := response.SuccessResponse(true, "SUCCESS", result)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary user could add to favoroite list of worker
// @ID user add to favorite list
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param addtofavotite body domain.Favorite{} true "User Add To Favorite"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/add-to-favorite [post]
func (cr *UserHandler) UserAddToFavorite(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	var favorite domain.Favorite

	c.Bind(&favorite)
	favorite.UserId = id

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", favorite, id)

	_, err := cr.userService.AddToFavorite(favorite)

	if err != nil {
		response := response.ErrorResponse("Error while add to favorite worker", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", favorite)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary list favorite list of workers for users
// @ID list favorite list of workers for users
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/list-favorite-list [get]
func (cr *UserHandler) ListFavorite(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	fmt.Printf("\n\nuser : %v\n\nmetea : %v\n\n", page, pageSize)
	log.Println(page, "   ", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	favorites, metadata, err := cr.userService.ListFevorite(pagenation, id)

	if err != nil {
		response := response.ErrorResponse("Failed to list favorite worker of user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	result := struct {
		Users *[]domain.ListFavorite
		Meta  *utils.Metadata
	}{
		Users: favorites,
		Meta:  metadata,
	}

	response := response.SuccessResponse(true, "SUCCESS", result)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Add address for User
// @ID user add address
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param UserAddProfile body domain.Address{} true "User Add Profile"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/add-address [post]
func (cr *UserHandler) UserAddAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))

	fmt.Printf("\n\nidea : %v\n\n", id)
	var address domain.Address

	c.Bind(&address)
	address.UserId = id
	_, err := cr.userService.AddAddress(address)

	if err != nil {
		response := response.ErrorResponse("Error while adding user address", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", address)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary List address for User
// @ID user list address
// @Tags User
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/list-address [get]
func (cr *UserHandler) UserListAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))

	address, err := cr.userService.ListAddress(id)

	if err != nil {
		response := response.ErrorResponse("Error while list corrent addresses of users", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", address)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary Delete address for user
// @ID user delete address
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param        addressid   query      string  true  "Job Id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/delete-address [delete]
func (cr *UserHandler) DeleteAddress(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	id, _ := strconv.Atoi(c.Query("addressid"))

	// c.Bind(&userprofile)

	fmt.Printf("\n\nuser Profile : \n%v\n\n\n\n", id)

	err := cr.userService.DeleteAddress(id, userid)

	if err != nil {
		response := response.ErrorResponse("Error while deleting current address of user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", id)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary user could send job request to worker
// @ID user send job request to worker
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param addtofavotite body domain.Request{} true "User Add To Favorite"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/send-job-request [post]
func (cr *UserHandler) UserSendJobRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	var request domain.Request

	c.Bind(&request)
	request.UserId = id

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", request, id)

	_, err := cr.userService.SendJobRequest(request)

	if err != nil {
		response := response.ErrorResponse("Error while sending job request ", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", request)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}


// @Summary Cancel request for user
// @ID user cancel request
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param        requestId   query      string  true  "Request Id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/cancel-job-request [delete]
func (cr *UserHandler) DeleteJobRequest(c *gin.Context) {
	userid, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	requestId, _ := strconv.Atoi(c.Query("requestId"))

	// c.Bind(&userprofile)

	// fmt.Printf("\n\nuser Profile : \n%v\n\n\n\n", id)

	err := cr.userService.DeleteJobRequest(requestId,userid)

	if err != nil {
		response := response.ErrorResponse("Error while deleting current address of user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	response := response.SuccessResponse(true, "SUCCESS", requestId)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}