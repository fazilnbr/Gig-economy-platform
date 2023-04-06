package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/fazilnbr/project-workey/pkg/common/response"
	"github.com/fazilnbr/project-workey/pkg/domain"
	mock "github.com/fazilnbr/project-workey/pkg/mock/usecaseMock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	c := mock.NewMockUserUseCase(ctrl)

	UserHandler := NewUserHandler(c)

	testData := []struct {
		name           string
		password       domain.ChangePassword
		beforeTest     func(userUsecase mock.MockUserUseCase)
		expectCode     int
		expectResponse response.Response
		expectErr      error
	}{
		{
			name:       "test binding error",
			password:   domain.ChangePassword{},
			beforeTest: func(userUsecase mock.MockUserUseCase) {},
			expectCode: 400,
			expectResponse: response.Response{
				Status:  false,
				Message: "Failed to fetch your data",
				Errors:  "",
				Data:    nil,
			},
			expectErr: nil,
		},
		{
			name: "test failed to verify user error",
			password: domain.ChangePassword{
				NewPassword: "12345",
				OldPassword: "54321",
			},
			beforeTest: func(userUsecase mock.MockUserUseCase) {
				userUsecase.EXPECT().UserVerifyPassword(domain.ChangePassword{
					NewPassword: "12345",
					OldPassword: "54321",
				}, 0).Return(errors.New("UserVerifyPassword usecase error"))
			},
			expectCode: 422,
			expectResponse: response.Response{
				Status:  false,
				Message: "Faild to verify user password",
				Errors:  "UserVerifyPassword usecase error",
				Data:    nil,
			},
			expectErr: nil,
		},
		{
			name: "test failed to change password error",
			password: domain.ChangePassword{
				NewPassword: "12345",
				OldPassword: "54321",
			},
			beforeTest: func(userUsecase mock.MockUserUseCase) {
				userUsecase.EXPECT().UserVerifyPassword(domain.ChangePassword{
					NewPassword: "12345",
					OldPassword: "54321",
				}, 0).Return(nil)
				userUsecase.EXPECT().UserChangePassword("12345", 0).Return(errors.New("UserChangePassword usecase error"))
			},
			expectCode: 422,
			expectResponse: response.Response{
				Status:  false,
				Message: "Error while changing Password",
				Errors:  "UserChangePassword usecase error",
				Data:    nil,
			},
			expectErr: nil,
		},
		{
			name: "test success response",
			password: domain.ChangePassword{
				NewPassword: "12345",
				OldPassword: "54321",
			},
			beforeTest: func(userUsecase mock.MockUserUseCase) {
				userUsecase.EXPECT().UserVerifyPassword(domain.ChangePassword{
					NewPassword: "12345",
					OldPassword: "54321",
				}, 0).Return(nil)
				userUsecase.EXPECT().UserChangePassword("12345", 0).Return(nil)
			},
			expectCode: 200,
			expectResponse: response.Response{
				Status:  true,
				Message: "SUCCESS",
				Errors:  nil,
				Data: domain.ChangePassword{
					NewPassword: "12345",
					OldPassword: "54321",
				},
			},
			expectErr: nil,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(*c)

			gin := gin.New()
			rec := httptest.NewRecorder()

			gin.PATCH("/change-password", UserHandler.UserChangePassword)

			var body []byte
			body, err := json.Marshal(tt.password)
			assert.NoError(t, err)
			req := httptest.NewRequest("PATCH", "/change-password", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			

			// Set a query parameter named "gury" with a value of "param"
			q := req.URL.Query()
			q.Add("email", "")
			q.Add("code", "")
			req.URL.RawQuery = q.Encode()

			gin.ServeHTTP(rec, req)

			var actual response.Response
			err = json.Unmarshal(rec.Body.Bytes(), &actual)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectCode, rec.Code)
			assert.Equal(t, tt.expectResponse.Message, actual.Message)
			// assert.Equal(t, tt.expectResponse.Data, actual.Data)

		})
	}
}
