package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/fazilnbr/project-workey/pkg/config"
	"github.com/fazilnbr/project-workey/pkg/domain"
	mock "github.com/fazilnbr/project-workey/pkg/mock/usecaseMock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserVerifyAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock.NewMockAuthUseCase(ctrl)
	authHandler := NewAuthHandler(nil, nil, nil, nil, c, config.Config{})

	testData := []struct {
		name       string
		email      string
		tocken     string
		beforeTest func(userUsecase mock.MockAuthUseCase)
		expectUser domain.UserResponse
		expectErr  error
	}{
		{
			name:   "test sucsess response",
			email:  "jon",
			tocken: "token",
			beforeTest: func(userUsecase mock.MockAuthUseCase) {
				userUsecase.EXPECT().UserVerifyAccount("jon", "token").Return(nil)
			},
			expectUser: domain.UserResponse{UserName: "jon"},
			expectErr:  nil,
		},
		{
			name:   "test sucsess response",
			email:  "jon",
			tocken: "token",
			beforeTest: func(userUsecase mock.MockAuthUseCase) {
				userUsecase.EXPECT().UserVerifyAccount("jon", "token").Return(errors.New("usecase error"))
			},
			expectUser: domain.UserResponse{},
			expectErr:  errors.New("usecase error"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeTest(*c)

			gin := gin.New()
			rec := httptest.NewRecorder()

			gin.POST("user/signup", authHandler.UserVerifyAccount)

			req := httptest.NewRequest("POST", "/user/signup", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			gin.ServeHTTP(rec, req)

		})
	}
}
