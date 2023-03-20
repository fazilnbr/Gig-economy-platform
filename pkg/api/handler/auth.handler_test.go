package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/domain"
	"github.com/fazilnbr/project-workey/pkg/repository"
	"github.com/fazilnbr/project-workey/pkg/response"
	"github.com/fazilnbr/project-workey/pkg/usecase"
	"github.com/fazilnbr/project-workey/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

var (
	username = "anu@gmail.com"
	pwd      = "123456"
	User     = domain.User{
		UserName: username,
		Password: pwd,
	}
	gormDB, _       = utils.MockGormDB()
	authRepoMock    = repository.NewUserRepo(gormDB)
	authService     = usecase.NewUserService(authRepoMock)
	authServiceMock = NewAuthHandler(nil, nil, authService, nil, nil, config.Config{})
)

func TestLogin(t *testing.T) {

	gormDB, _ := utils.MockGormDB()
	authRepoMock := repository.NewUserRepo(gormDB)
	authService := usecase.NewUserService(authRepoMock)
	authServiceMock := NewAuthHandler(nil, nil, authService, nil, nil, config.Config{})

	t.Run("test normal case login 1", func(t *testing.T) {

		gin := gin.New()
		rec := httptest.NewRecorder()

		gin.POST("user/signup", authServiceMock.UserSignUp)

		body, err := json.Marshal(User)
		fmt.Printf("\n\nbody : %v\n\n", string(body))
		assert.NoError(t, err)
		req := httptest.NewRequest("POST", "/user/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		gin.ServeHTTP(rec, req)

		var newUser response.Response
		err = json.Unmarshal(rec.Body.Bytes(), &newUser)
		assert.NoError(t, err)

		exp := response.Response{
			Status:  false,
			Message: "Failed to create user",
			Errors:  []interface{}{"Username already exists"},
			Data:    nil,
		}

		exp = response.Response{
			Status:  true,
			Message: "SUCCESS",
			Errors:  "",
			Data:    User,
		}

		t.Run("test success response", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, exp.Status, newUser.Status)
			// assert.Equal(t, exp.Data, newUser.Data)
		})

		_, err = gormDB.Exec("DELETE FROM users WHERE user_name=$1", username)
		assert.NoError(t, err)

	})
}
