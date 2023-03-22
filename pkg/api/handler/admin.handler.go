package handler

import (
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
// @Tags Admin User Management
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-all-users [get]
func (cr *AdminHandler) ListAllUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	users, metadata, err := cr.adminService.ListUsers(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list all active user", err.Error(), nil)
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
// @Tags Admin User Management
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-new-users [get]
func (cr *AdminHandler) ListNewUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

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
		response := response.ErrorResponse("Failed to list new user", err.Error(), nil)
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
// @Tags Admin User Management
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-blocked-users [get]
func (cr *AdminHandler) ListBlockUsers(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}
	users, metadata, err := cr.adminService.ListBlockedUsers(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list blocked user", err.Error(), nil)

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
// @Tags Admin User Management
// @Security BearerAuth
// @Produce json
// @Param        id   query      string  true  "Id of User : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/activate-users [patch]
func (cr *AdminHandler) ActivateUsers(c *gin.Context) {

	id := c.Query("id")
	Id, _ := strconv.Atoi(id)

	users, err := cr.adminService.ActivateUser(Id)

	if err != nil {
		response := response.ErrorResponse("Failed to activate user", err.Error(), nil)

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
// @Tags Admin User Management
// @Security BearerAuth
// @Produce json
// @Param        id   query      string  true  "Id of User : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/block-users [patch]
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
// @Tags Admin Worker Management
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-all-workers [get]
func (cr *AdminHandler) ListAllWorkers(c *gin.Context) {

	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	users, metadata, err := cr.adminService.ListWorkers(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list all active worker", err.Error(), nil)
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
// @Tags Admin Worker Management
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-new-workers [get]
func (cr *AdminHandler) ListNewWorkers(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	users, metadata, err := cr.adminService.ListNewWorkers(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list new worker", err.Error(), nil)
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

// @Summary list all blocked workers for admin
// @ID list all blocked workers
// @Tags Admin Worker Management
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-blocked-workers [get]
func (cr *AdminHandler) ListBlockWorkers(c *gin.Context) {

	page, err := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	users, metadata, err := cr.adminService.ListBlockedWorkers(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list blocked worker", err.Error(), nil)

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

// @Summary activate workers for admin
// @ID activate workers
// @Tags Admin Worker Management
// @Security BearerAuth
// @Produce json
// @Param        id   query      string  true  "Id of User : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/activate-workers [patch]
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
// @Tags Admin Worker Management
// @Security BearerAuth
// @Produce json
// @Param        id   query      string  true  "Id of User : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/block-workers [patch]
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
// @Tags Admin Job Management
// @Security BearerAuth
// @Produce json
// @Param       category   query      string  true  "category : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/add-job-category [post]
func (cr *AdminHandler) AddJobCategory(c *gin.Context) {

	category := c.Query("category")

	err := cr.adminService.AddJobCategory(category)

	if err != nil {
		response := response.ErrorResponse("Failed to add job category", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", category)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary list all job categories for admin
// @ID list all job category
// @Tags Admin Job Management
// @Security BearerAuth
// @Produce json
// @Param        page   query      string  true  "Page : "
// @Param        pagesize   query      string  true  "Pagesize : "
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/list-job-category [get]
func (cr *AdminHandler) ListJobCategory(c *gin.Context) {

	page, err := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	pagenation := utils.Filter{
		Page:     page,
		PageSize: pageSize,
	}

	categories, metadata, err := cr.adminService.ListJobCategory(pagenation)

	if err != nil {
		response := response.ErrorResponse("Failed to list job categories", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}

	result := struct {
		Users *[]domain.Category
		Meta  *utils.Metadata
	}{
		Users: categories,
		Meta:  &metadata,
	}

	response := response.SuccessResponse(true, "SUCCESS", result)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary update job category for admin
// @ID update category
// @Tags Admin Job Management
// @Security BearerAuth
// @Produce json
// @Param AdminLogin body domain.Category{} true "admin login"
// @Success 200 {object} response.Response{}
// @Failure 422 {object} response.Response{}
// @Router /admin/update-job-category [patch]
func (cr *AdminHandler) UpdateJobCategory(c *gin.Context) {

	// category := c.Query("category")
	// category := c.Param("category")
	var category domain.Category

	err := c.Bind(&category)
	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*c, response)
		return
	}

	_, err = cr.adminService.UpdateJobCategory(category)

	if err != nil {
		response := response.ErrorResponse("Failed to update job category", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", category)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}
