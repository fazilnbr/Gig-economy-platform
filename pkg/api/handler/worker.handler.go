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

type WorkerHandler struct {
	workerService services.WorkerUseCase
}

func NewWorkerHandler(workerService services.WorkerUseCase) WorkerHandler {
	return WorkerHandler{
		workerService: workerService,
	}
}

func (cr *WorkerHandler) AddProfile(c *gin.Context) {
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
