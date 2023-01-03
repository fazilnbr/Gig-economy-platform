package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fazilnbr/project-workey/pkg/common/response"
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
// @Produce json
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Login}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/listallusers [get]
func (cr *AdminHandler) ListAllUsers(c *gin.Context) {

	users, err := cr.adminService.ListUsers()

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
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

// @Summary list all new users for admin
// @ID list all new users
// @Produce json
// @Param        username   path      string  true  "User Name : "
// @Param        password   path      string  true  "Password : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Login}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/listnewusers [get]
func (cr *AdminHandler) ListNewUsers(c *gin.Context) {

	users, err := cr.adminService.ListNewUsers()

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)
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

// @Summary list all blocked users for admin
// @ID list all blocked users
// @Produce json
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Login}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/listblockedusers [get]
func (cr *AdminHandler) ListBlockUsers(c *gin.Context) {

	users, err := cr.adminService.ListBlockedUsers()

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)

		c.Writer.WriteHeader(http.StatusUnprocessableEntity)

		utils.ResponseJSON(*c, response)
		return
	}
	response := response.SuccessResponse(true, "SUCCESS", users)

	c.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*c, response)
}

// @Summary activate users for admin
// @ID activate users
// @Produce json
// @Param        id   path      string  true  "Id of User : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Login}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/activateusers [post]
func (cr *AdminHandler) ActivateUsers(c *gin.Context) {

	id := c.Query("id")
	Id, _ := strconv.Atoi(id)

	users, err := cr.adminService.ActivateUser(Id)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)

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
// @Produce json
// @Param        id   path      string  true  "Id of User : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Login}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/activateusers [post]
func (cr *AdminHandler) BlockUsers(c *gin.Context) {

	id := c.Query("id")
	Id, _ := strconv.Atoi(id)

	users, err := cr.adminService.BlockUser(Id)

	if err != nil {
		response := response.ErrorResponse("Failed to create user", err.Error(), nil)

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
// @Produce json
// @Param       category   path      string  true  "category : "
// @Success 200 {object} response.Response{Status=bool,Message=string,Errors=string,Data=domain.Category}
// @Failure 422 {object} response.Response{Status=bool,Message=string,Errors=string,Data=string}
// @Router /admin/block [post]
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
