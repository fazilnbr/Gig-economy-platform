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

type AdminHandler struct {
	adminService services.AdminUseCase
}

func NewAdminHandler(adminService services.AdminUseCase) AdminHandler {
	return AdminHandler{
		adminService: adminService,
	}
}

// @Summary list all active users for admin
// @ID list all active users
// @Tags Admin
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/listallusers [get]
func (cr *AdminHandler) ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	fmt.Printf("\n\nuser : %v\n\nmetea : %v\n\n", page, pageSize)
	log.Println(page, "   ", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	users, metadata, err := cr.adminService.ListUsers(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	result := struct {
		Users *[]domain.UserResponse
		Meta  *utils.Metadata
	}{
		Users: users,
		Meta:  metadata,
	}

	response := response.SuccessResponse(true, "SUCCESS", result)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary list all new users for admin
// @ID list all new users
// @Tags Admin
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/listnewusers [get]
func (cr *AdminHandler) ListNewUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	log.Println(page, "   ", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	users, metadata, err := cr.adminService.ListNewUsers(pagenation)

	result := struct {
		Users *[]domain.UserResponse
		Meta  *utils.Metadata
	}{
		Users: users,
		Meta:  metadata,
	}

	if err != nil {
		response := response.ErrorResponse("Failed to list user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", result)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary list all blocked users for admin
// @ID list all blocked users
// @Tags Admin
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/listblockedusers [get]
func (cr *AdminHandler) ListBlockUsers(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	log.Println(page, "   ", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}
	users, metadata, err := cr.adminService.ListBlockedUsers(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list user", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	result := struct {
		Users *[]domain.UserResponse
		Meta  *utils.Metadata
	}{
		Users: users,
		Meta:  metadata,
	}

	response := response.SuccessResponse(true, "SUCCESS", result)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary activate users for admin
// @ID activate users
// @Tags Admin
// @Produce json
// @Param        id   path      string  true  "Id of User : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/activateusers [put]
func (cr *AdminHandler) ActivateUsers(c *gin.Context) {

	id := c.Query("id")
	Id, _ := strconv.Atoi(id)

	users, err := cr.adminService.ActivateUser(Id)

	if err != nil {
		response := response.ErrorResponse("Failed to list user", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", users)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary block users for admin
// @ID block users
// @Tags Admin
// @Produce json
// @Param        id   path      string  true  "Id of User : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/blockusers [put]
func (cr *AdminHandler) BlockUsers(c *gin.Context) {

	id := c.Query("id")
	Id, _ := strconv.Atoi(id)

	users, err := cr.adminService.BlockUser(Id)

	if err != nil {
		response := response.ErrorResponse("Failed to block user", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", users)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary list all active workers for admin
// @ID list all active workers
// @Tags Admin
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/listallworkers [get]
func (cr *AdminHandler) ListAllWorkers(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	log.Println(page, "   ", pageSize)

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	users, metadata, err := cr.adminService.ListWorkers(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list worker", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	result := struct {
		Users *[]domain.UserResponse
		Meta  *utils.Metadata
	}{
		Users: users,
		Meta:  metadata,
	}

	response := response.SuccessResponse(true, "SUCCESS", result)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary list all new workers for admin
// @ID list all new workers
// @Tags Admin
// @Produce json
// @Param        username   path      string  true  "User Name : "
// @Param        password   path      string  true  "Password : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/listnewworkers [get]
func (cr *AdminHandler) ListNewWorkers(c *gin.Context) {

	users, err := cr.adminService.ListNewWorkers()

	if err != nil {
		response := response.ErrorResponse("Failed to list worker", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", users)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary list all blocked workers for admin
// @ID list all blocked workers
// @Tags Admin
// @Produce json
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/listblockedworkers [get]
func (cr *AdminHandler) ListBlockWorkers(c *gin.Context) {

	users, err := cr.adminService.ListBlockedWorkers()

	if err != nil {
		response := response.ErrorResponse("Failed to list worker", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", users)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary activate workers for admin
// @ID activate workers
// @Tags Admin
// @Produce json
// @Param        id   path      string  true  "Id of User : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/activateworkers [patch]
func (cr *AdminHandler) ActivateWorkers(c *gin.Context) {

	id := c.Query("id")
	Id, _ := strconv.Atoi(id)

	users, err := cr.adminService.ActivateWorker(Id)

	if err != nil {
		response := response.ErrorResponse("Failed to activate worker", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", users)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary block workers for admin
// @ID block workers
// @Tags Admin
// @Produce json
// @Param        id   path      string  true  "Id of User : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/blockworkers [patch]
func (cr *AdminHandler) BlockWorkers(c *gin.Context) {

	id := c.Query("id")
	Id, _ := strconv.Atoi(id)

	users, err := cr.adminService.BlockWorker(Id)

	if err != nil {
		response := response.ErrorResponse("Failed to block worker", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", users)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary add job category for admin
// @ID add category
// @Tags Admin
// @Produce json
// @Param       category   path      string  true  "category : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/jobcategory [post]
func (cr *AdminHandler) AddJobCategory(c *gin.Context) {

	category := c.Query("category")

	fmt.Printf("\n\ncat : %v\n\n", category)

	err := cr.adminService.AddJobCategory(category)

	if err != nil {
		response := response.ErrorResponse("Failed to add category", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", nil)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}
