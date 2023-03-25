package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/fazilnbr/project-workey/pkg/config"
// 	"github.com/fazilnbr/project-workey/pkg/domain"
// 	"github.com/fazilnbr/project-workey/pkg/repository"
// 	"github.com/fazilnbr/project-workey/pkg/response"
// 	"github.com/fazilnbr/project-workey/pkg/usecase"
// 	"github.com/fazilnbr/project-workey/pkg/utils"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"

// 	_ "github.com/lib/pq"
// )

// var (
// 	gormDB,mock, mockDB  = utils.MockGormDB()
// 	authRepoMock    = repository.NewUserRepo(gormDB)
// 	authService     = usecase.NewUserService(authRepoMock)
// 	authServiceMock = NewAuthHandler(nil, nil, authService, nil, nil, config.Config{})
// )

// func TestLogin(t *testing.T) {

// 	User := domain.User{
// 		UserName: utils.RandomMail(3),
// 		Password: utils.RandomString(4),
// 	}

// 	t.Run("test normal case login 1", func(t *testing.T) {

// 		gin := gin.New()
// 		rec := httptest.NewRecorder()

// 		gin.POST("user/signup", authServiceMock.UserSignUp)

// 		body, err := json.Marshal(User)
// 		assert.NoError(t, err)
// 		req := httptest.NewRequest("POST", "/user/signup", bytes.NewBuffer(body))
// 		req.Header.Set("Content-Type", "application/json")
// 		gin.ServeHTTP(rec, req)

// 		var actual response.Response
// 		err = json.Unmarshal(rec.Body.Bytes(), &actual)
// 		assert.NoError(t, err)

// 		exp := response.Response{
// 			Status:  true,
// 			Message: "SUCCESS",
// 			Errors:  "",
// 			Data:    User,
// 		}

// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, exp.Status, actual.Status)
// 		assert.Equal(t, exp.Message, actual.Message)
// 		// assert.Equal(t, exp.Data, actual.Data)
// 	})

// 	t.Run("test exist email login 2", func(t *testing.T) {
// 		gin := gin.New()
// 		rec := httptest.NewRecorder()

// 		gin.POST("user/signup", authServiceMock.UserSignUp)

// 		body, err := json.Marshal(User)
// 		assert.NoError(t, err)
// 		req := httptest.NewRequest("POST", "/user/signup", bytes.NewBuffer(body))
// 		req.Header.Set("Content-Type", "application/json")
// 		gin.ServeHTTP(rec, req)

// 		var actual response.Response
// 		err = json.Unmarshal(rec.Body.Bytes(), &actual)
// 		assert.NoError(t, err)

// 		exp := response.Response{
// 			Status:  false,
// 			Message: "Failed to create user",
// 			// Errors:  []interface{}{"Username already exists"},
// 			Errors: "Username already exists",
// 			Data:   nil,
// 		}

// 		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
// 		assert.Equal(t, exp.Status, actual.Status)
// 		assert.Equal(t, exp.Message, actual.Message)

// 	})

// 	_, err := gormDB.Exec("DELETE FROM users WHERE user_name=$1", User.UserName)
// 	assert.NoError(t, err)
// }
