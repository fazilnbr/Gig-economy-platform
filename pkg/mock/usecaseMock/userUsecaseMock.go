// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/user.interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	domain "github.com/fazilnbr/project-workey/pkg/domain"
	utils "github.com/fazilnbr/project-workey/pkg/utils"
	gomock "github.com/golang/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserUseCase) AddAddress(address domain.Address) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", address)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserUseCaseMockRecorder) AddAddress(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserUseCase)(nil).AddAddress), address)
}

// AddProfile mocks base method.
func (m *MockUserUseCase) AddProfile(userProfile domain.Profile, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProfile", userProfile, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProfile indicates an expected call of AddProfile.
func (mr *MockUserUseCaseMockRecorder) AddProfile(userProfile, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProfile", reflect.TypeOf((*MockUserUseCase)(nil).AddProfile), userProfile, id)
}

// AddToFavorite mocks base method.
func (m *MockUserUseCase) AddToFavorite(favorite domain.Favorite) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToFavorite", favorite)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddToFavorite indicates an expected call of AddToFavorite.
func (mr *MockUserUseCaseMockRecorder) AddToFavorite(favorite interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToFavorite", reflect.TypeOf((*MockUserUseCase)(nil).AddToFavorite), favorite)
}

// CheckOrderId mocks base method.
func (m *MockUserUseCase) CheckOrderId(userId int, orderId string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckOrderId", userId, orderId)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckOrderId indicates an expected call of CheckOrderId.
func (mr *MockUserUseCaseMockRecorder) CheckOrderId(userId, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckOrderId", reflect.TypeOf((*MockUserUseCase)(nil).CheckOrderId), userId, orderId)
}

// CreateUser mocks base method.
func (m *MockUserUseCase) CreateUser(newUser domain.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserUseCaseMockRecorder) CreateUser(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserUseCase)(nil).CreateUser), newUser)
}

// DeleteAddress mocks base method.
func (m *MockUserUseCase) DeleteAddress(id, userid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", id, userid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserUseCaseMockRecorder) DeleteAddress(id, userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserUseCase)(nil).DeleteAddress), id, userid)
}

// DeleteJobRequest mocks base method.
func (m *MockUserUseCase) DeleteJobRequest(requestId, userid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJobRequest", requestId, userid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJobRequest indicates an expected call of DeleteJobRequest.
func (mr *MockUserUseCaseMockRecorder) DeleteJobRequest(requestId, userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJobRequest", reflect.TypeOf((*MockUserUseCase)(nil).DeleteJobRequest), requestId, userid)
}

// FetchRazorPayDetials mocks base method.
func (m *MockUserUseCase) FetchRazorPayDetials(userId, requestId int) (*domain.RazorPayVariables, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRazorPayDetials", userId, requestId)
	ret0, _ := ret[0].(*domain.RazorPayVariables)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRazorPayDetials indicates an expected call of FetchRazorPayDetials.
func (mr *MockUserUseCaseMockRecorder) FetchRazorPayDetials(userId, requestId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRazorPayDetials", reflect.TypeOf((*MockUserUseCase)(nil).FetchRazorPayDetials), userId, requestId)
}

// FindUser mocks base method.
func (m *MockUserUseCase) FindUser(email string) (*domain.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", email)
	ret0, _ := ret[0].(*domain.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockUserUseCaseMockRecorder) FindUser(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockUserUseCase)(nil).FindUser), email)
}

// FindUserWithId mocks base method.
func (m *MockUserUseCase) FindUserWithId(id int) (*domain.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserWithId", id)
	ret0, _ := ret[0].(*domain.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserWithId indicates an expected call of FindUserWithId.
func (mr *MockUserUseCaseMockRecorder) FindUserWithId(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserWithId", reflect.TypeOf((*MockUserUseCase)(nil).FindUserWithId), id)
}

// ListAddress mocks base method.
func (m *MockUserUseCase) ListAddress(id int) (*[]domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAddress", id)
	ret0, _ := ret[0].(*[]domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAddress indicates an expected call of ListAddress.
func (mr *MockUserUseCaseMockRecorder) ListAddress(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAddress", reflect.TypeOf((*MockUserUseCase)(nil).ListAddress), id)
}

// ListFevorite mocks base method.
func (m *MockUserUseCase) ListFevorite(pagenation utils.Filter, id int) (*[]domain.ListFavorite, *utils.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFevorite", pagenation, id)
	ret0, _ := ret[0].(*[]domain.ListFavorite)
	ret1, _ := ret[1].(*utils.Metadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFevorite indicates an expected call of ListFevorite.
func (mr *MockUserUseCaseMockRecorder) ListFevorite(pagenation, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFevorite", reflect.TypeOf((*MockUserUseCase)(nil).ListFevorite), pagenation, id)
}

// ListSendRequests mocks base method.
func (m *MockUserUseCase) ListSendRequests(pagenation utils.Filter, id int) (*[]domain.RequestUserResponse, *utils.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSendRequests", pagenation, id)
	ret0, _ := ret[0].(*[]domain.RequestUserResponse)
	ret1, _ := ret[1].(*utils.Metadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSendRequests indicates an expected call of ListSendRequests.
func (mr *MockUserUseCaseMockRecorder) ListSendRequests(pagenation, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSendRequests", reflect.TypeOf((*MockUserUseCase)(nil).ListSendRequests), pagenation, id)
}

// ListWorkersWithJob mocks base method.
func (m *MockUserUseCase) ListWorkersWithJob(pagenation utils.Filter) (*[]domain.ListJobsWithWorker, *utils.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkersWithJob", pagenation)
	ret0, _ := ret[0].(*[]domain.ListJobsWithWorker)
	ret1, _ := ret[1].(*utils.Metadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListWorkersWithJob indicates an expected call of ListWorkersWithJob.
func (mr *MockUserUseCaseMockRecorder) ListWorkersWithJob(pagenation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkersWithJob", reflect.TypeOf((*MockUserUseCase)(nil).ListWorkersWithJob), pagenation)
}

// SavePaymentOrderDeatials mocks base method.
func (m *MockUserUseCase) SavePaymentOrderDeatials(payment domain.JobPayment) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePaymentOrderDeatials", payment)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SavePaymentOrderDeatials indicates an expected call of SavePaymentOrderDeatials.
func (mr *MockUserUseCaseMockRecorder) SavePaymentOrderDeatials(payment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePaymentOrderDeatials", reflect.TypeOf((*MockUserUseCase)(nil).SavePaymentOrderDeatials), payment)
}

// SearchWorkersWithJob mocks base method.
func (m *MockUserUseCase) SearchWorkersWithJob(pagenation utils.Filter, key string) (*[]domain.ListJobsWithWorker, *utils.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWorkersWithJob", pagenation, key)
	ret0, _ := ret[0].(*[]domain.ListJobsWithWorker)
	ret1, _ := ret[1].(*utils.Metadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SearchWorkersWithJob indicates an expected call of SearchWorkersWithJob.
func (mr *MockUserUseCaseMockRecorder) SearchWorkersWithJob(pagenation, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWorkersWithJob", reflect.TypeOf((*MockUserUseCase)(nil).SearchWorkersWithJob), pagenation, key)
}

// SendJobRequest mocks base method.
func (m *MockUserUseCase) SendJobRequest(request domain.Request) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendJobRequest", request)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendJobRequest indicates an expected call of SendJobRequest.
func (mr *MockUserUseCaseMockRecorder) SendJobRequest(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendJobRequest", reflect.TypeOf((*MockUserUseCase)(nil).SendJobRequest), request)
}

// UpdateJobComplition mocks base method.
func (m *MockUserUseCase) UpdateJobComplition(userId, requestId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateJobComplition", userId, requestId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateJobComplition indicates an expected call of UpdateJobComplition.
func (mr *MockUserUseCaseMockRecorder) UpdateJobComplition(userId, requestId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateJobComplition", reflect.TypeOf((*MockUserUseCase)(nil).UpdateJobComplition), userId, requestId)
}

// UpdatePaymentId mocks base method.
func (m *MockUserUseCase) UpdatePaymentId(razorPaymentId string, idPayment int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePaymentId", razorPaymentId, idPayment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePaymentId indicates an expected call of UpdatePaymentId.
func (mr *MockUserUseCaseMockRecorder) UpdatePaymentId(razorPaymentId, idPayment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePaymentId", reflect.TypeOf((*MockUserUseCase)(nil).UpdatePaymentId), razorPaymentId, idPayment)
}

// UserChangePassword mocks base method.
func (m *MockUserUseCase) UserChangePassword(changepassword string, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserChangePassword", changepassword, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserChangePassword indicates an expected call of UserChangePassword.
func (mr *MockUserUseCaseMockRecorder) UserChangePassword(changepassword, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserChangePassword", reflect.TypeOf((*MockUserUseCase)(nil).UserChangePassword), changepassword, id)
}

// UserEditProfile mocks base method.
func (m *MockUserUseCase) UserEditProfile(userProfile domain.Profile, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserEditProfile", userProfile, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserEditProfile indicates an expected call of UserEditProfile.
func (mr *MockUserUseCaseMockRecorder) UserEditProfile(userProfile, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserEditProfile", reflect.TypeOf((*MockUserUseCase)(nil).UserEditProfile), userProfile, id)
}

// UserVerifyPassword mocks base method.
func (m *MockUserUseCase) UserVerifyPassword(changepassword domain.ChangePassword, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserVerifyPassword", changepassword, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserVerifyPassword indicates an expected call of UserVerifyPassword.
func (mr *MockUserUseCaseMockRecorder) UserVerifyPassword(changepassword, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserVerifyPassword", reflect.TypeOf((*MockUserUseCase)(nil).UserVerifyPassword), changepassword, id)
}

// VerifyUser mocks base method.
func (m *MockUserUseCase) VerifyUser(email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUser indicates an expected call of VerifyUser.
func (mr *MockUserUseCaseMockRecorder) VerifyUser(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockUserUseCase)(nil).VerifyUser), email, password)
}

// ViewSendOneRequest mocks base method.
func (m *MockUserUseCase) ViewSendOneRequest(userId, requestId int) (*domain.RequestUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewSendOneRequest", userId, requestId)
	ret0, _ := ret[0].(*domain.RequestUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewSendOneRequest indicates an expected call of ViewSendOneRequest.
func (mr *MockUserUseCaseMockRecorder) ViewSendOneRequest(userId, requestId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewSendOneRequest", reflect.TypeOf((*MockUserUseCase)(nil).ViewSendOneRequest), userId, requestId)
}
