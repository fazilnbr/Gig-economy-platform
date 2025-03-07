package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fazilnbr/project-workey/pkg/common/response"
	"github.com/fazilnbr/project-workey/pkg/domain"
	services "github.com/fazilnbr/project-workey/pkg/usecase/interface"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
	razorpay "github.com/razorpay/razorpay-go"
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
// @Tags User Profile Management
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

	err:=c.Bind(&userprofile)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	err = cr.userService.AddProfile(userprofile, id)

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
// @Tags User Profile Management
// @Security BearerAuth
// @Produce json
// @Param UserEditProfile body domain.Profile{} true "User Edit Profile"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/edit-profile [patch]
func (cr *UserHandler) UserEditProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	var userprofile domain.Profile

	err:=c.Bind(&userprofile)
	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", userprofile, id)

	err = cr.userService.UserEditProfile(userprofile, id)

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
// @Tags User Profile Management
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
		response := response.ErrorResponse("Failed to fetch your data", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
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
// @Tags User List Worker
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
// @Tags User List Worker
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
// @Tags User Favorite
// @Security BearerAuth
// @Produce json
// @Param addtofavotite body domain.Favorite{} true "User Add To Favorite"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/add-to-favorite [post]
func (cr *UserHandler) UserAddToFavorite(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	var favorite domain.Favorite

	err:=c.Bind(&favorite)
	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	favorite.UserId = id

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", favorite, id)

	_, err = cr.userService.AddToFavorite(favorite)

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
// @Tags User Favorite
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
// @Tags User Address Management
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

	err:=c.Bind(&address)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	address.UserId = id
	_, err = cr.userService.AddAddress(address)

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
// @Tags User Address Management
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
// @Tags User Address Management
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
// @Tags User Job Request
// @Security BearerAuth
// @Produce json
// @Param addtofavotite body domain.Request{} true "User Add To Favorite"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/send-job-request [post]
func (cr *UserHandler) UserSendJobRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	var request domain.Request

	err:=c.Bind(&request)
	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}
	request.UserId = id

	fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", request, id)

	_, err = cr.userService.SendJobRequest(request)

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
// @Tags User Job Request
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

	err := cr.userService.DeleteJobRequest(requestId, userid)

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

// @Summary list send job request for users
// @ID list send job request for users
// @Tags User Job Request
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/list-job-request [get]
func (cr *UserHandler) ListSendRequests(c *gin.Context) {
	id, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	fmt.Printf("\n\nuser : %v\n\nmetea : %v\n\n", page, pageSize)
	log.Println(page, "   ", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	requests, metadata, err := cr.userService.ListSendRequests(pagenation, id)

	if err != nil {
		response := response.ErrorResponse("Failed to list send job request of user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	result := struct {
		Users *[]domain.RequestUserResponse
		Meta  *utils.Metadata
	}{
		Users: requests,
		Meta:  metadata,
	}

	response := response.SuccessResponse(true, "SUCCESS", result)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary View One Job Request
// @ID user view one job request
// @Tags User Job Request
// @Security BearerAuth
// @Produce json
// @Param        requestid   query      string  true  "Request Id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/view-one-job-request [get]
func (cr *UserHandler) ViewSendOneRequest(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Writer.Header().Get("id"))

	requestId, _ := strconv.Atoi(c.Query("requestid"))

	// fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", userprofile, id)

	request, err := cr.userService.ViewSendOneRequest(userId, requestId)

	if err != nil {
		response := response.ErrorResponse("Error while editing profile", err.Error(), nil)
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

// @Summary Update Job Complition
// @ID update job complition
// @Tags User Job Request
// @Security BearerAuth
// @Produce json
// @Param        requestid   query      string  true  "Request Id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/update-job-complition-status [patch]
func (cr *UserHandler) UpdateJobComplition(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Writer.Header().Get("id"))

	requestId, _ := strconv.Atoi(c.Query("requestid"))

	// fmt.Printf("\n\nuser Profile : \n%v\n\n%v\n\n", userprofile, id)

	err := cr.userService.UpdateJobComplition(userId, requestId)

	if err != nil {
		response := response.ErrorResponse("Error while Updating Job Complition", err.Error(), nil)
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

// @Summary To Open Home Page To Razor-Pay Payment
// @ID To open home page to razor-pay payment
// @Tags User Job Payment
// @Security BearerAuth
// @Produce json
// @Param        requestid   query      string  true  "Request Id : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /user/razor-pay-home [get]
func (cr *UserHandler) RazorPayHome(c *gin.Context) {
	// userId, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	userId := 5

	requestId, _ := strconv.Atoi(c.Query("requestId"))

	// Fetch razor pay request data
	razordata, err := cr.userService.FetchRazorPayDetials(userId, requestId)

	if err != nil {
		response := response.ErrorResponse("Error while Fetching Razor-Pay Request data", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	// Create order_id from the Razor-Pay server
	client := razorpay.NewClient("rzp_test_vOsKKSWnOE803Q", "JINdUUpdybhJ707mAu37fH84")

	// Create data to get order id
	data := map[string]interface{}{
		"amount":   razordata.Amount * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	// Make an order in razor pay to payment
	body, err := client.Order.Create(data, nil)
	fmt.Println("////////////////reciept", body)
	if err != nil {
		fmt.Println("Problem getting the repository information", err)
		os.Exit(1)
	}

	value := body["id"]

	orderId := value.(string)
	fmt.Println("str////////////////", orderId)

	// Save the order id
	// In razordata
	razordata.OrderId = orderId
	// In database
	payment := domain.JobPayment{
		RequestId: requestId,
		OrderId:   orderId,
		UserId:    userId,
		Amount:    razordata.Amount,
	}
	_, err = cr.userService.SavePaymentOrderDeatials(payment)

	if err != nil {
		response := response.ErrorResponse("Error while Storing Razor-Pay order id", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	// response := response.SuccessResponse(true, "SUCCESS", razordata)
	// c.Writer.Header().Set("Content-Type", "application/json")
	// c.Writer.WriteHeader(http.StatusOK)
	// utils.ResponseJSON(*c, response)
	// razordata.Amount = razordata.Amount * 100
	// fmt.Printf("\n\namt :%v\n\n",razordata.Amount)

	c.HTML(http.StatusOK, "razor-pay-home.html", razordata)

}


func (cr *UserHandler) RazorPaySuccess(c *gin.Context) {
	// userId, _ := strconv.Atoi(c.Writer.Header().Get("id"))
	userId := 5

	// signature := c.Query("signature")
	orderid := c.Query("orderid")

	fmt.Printf("\n\norderid :%v\n\n", orderid)
	// Fetch razor pay request data
	paymentId, err := cr.userService.CheckOrderId(userId, orderid)
	fmt.Printf("\n\npayment id : %v\n\n", paymentId)

	if err != nil {

		c.HTML(http.StatusOK, "razor-pay-failed.html", err)
		// response := response.ErrorResponse("Error while Fetching Razor-Pay Request data", err.Error(), nil)
		// c.Writer.Header().Set("Content-Type", "application/json")
		// c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		// utils.ResponseJSON(*c, response)
		return
	}

	razorpaymentid := c.Query("paymentid")

	err = cr.userService.UpdatePaymentId(razorpaymentid, paymentId)

	if err != nil {
		response := response.ErrorResponse("Error while Updating Razor-Pay Request Id", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	// response := response.SuccessResponse(true, "SUCCESS", razordata)
	// c.Writer.Header().Set("Content-Type", "application/json")
	// c.Writer.WriteHeader(http.StatusOK)
	// utils.ResponseJSON(*c, response)

	c.HTML(http.StatusOK, "razor-pay-success.html", userId)

}
