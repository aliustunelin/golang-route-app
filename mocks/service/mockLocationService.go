// Code generated by MockGen. DO NOT EDIT.
// Source: golang-route-app/services (interfaces: LocationService)

// Package services is a generated GoMock package.
package services

import (
	dto "golang-route-app/dto"
	models "golang-route-app/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockLocationService is a mock of LocationService interface.
type MockLocationService struct {
	ctrl     *gomock.Controller
	recorder *MockLocationServiceMockRecorder
}

// MockLocationServiceMockRecorder is the mock recorder for MockLocationService.
type MockLocationServiceMockRecorder struct {
	mock *MockLocationService
}

// NewMockLocationService creates a new mock instance.
func NewMockLocationService(ctrl *gomock.Controller) *MockLocationService {
	mock := &MockLocationService{ctrl: ctrl}
	mock.recorder = &MockLocationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocationService) EXPECT() *MockLocationServiceMockRecorder {
	return m.recorder
}

// LocationDelete mocks base method.
func (m *MockLocationService) LocationDelete(arg0 primitive.ObjectID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationDelete", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationDelete indicates an expected call of LocationDelete.
func (mr *MockLocationServiceMockRecorder) LocationDelete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationDelete", reflect.TypeOf((*MockLocationService)(nil).LocationDelete), arg0)
}

// LocationGetAll mocks base method.
func (m *MockLocationService) LocationGetAll() ([]models.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationGetAll")
	ret0, _ := ret[0].([]models.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationGetAll indicates an expected call of LocationGetAll.
func (mr *MockLocationServiceMockRecorder) LocationGetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationGetAll", reflect.TypeOf((*MockLocationService)(nil).LocationGetAll))
}

// LocationGetByNameWithData mocks base method.
func (m *MockLocationService) LocationGetByNameWithData(arg0 primitive.ObjectID) ([]models.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationGetByNameWithData", arg0)
	ret0, _ := ret[0].([]models.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationGetByNameWithData indicates an expected call of LocationGetByNameWithData.
func (mr *MockLocationServiceMockRecorder) LocationGetByNameWithData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationGetByNameWithData", reflect.TypeOf((*MockLocationService)(nil).LocationGetByNameWithData), arg0)
}

// LocationInsert mocks base method.
func (m *MockLocationService) LocationInsert(arg0 models.Location) (*dto.LocationDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationInsert", arg0)
	ret0, _ := ret[0].(*dto.LocationDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationInsert indicates an expected call of LocationInsert.
func (mr *MockLocationServiceMockRecorder) LocationInsert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationInsert", reflect.TypeOf((*MockLocationService)(nil).LocationInsert), arg0)
}

// LocationRouting mocks base method.
func (m *MockLocationService) LocationRouting(arg0 models.Location) ([]primitive.M, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationRouting", arg0)
	ret0, _ := ret[0].([]primitive.M)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationRouting indicates an expected call of LocationRouting.
func (mr *MockLocationServiceMockRecorder) LocationRouting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationRouting", reflect.TypeOf((*MockLocationService)(nil).LocationRouting), arg0)
}

// LocationUpdateByID mocks base method.
func (m *MockLocationService) LocationUpdateByID(arg0 models.Location) (*dto.LocationDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocationUpdateByID", arg0)
	ret0, _ := ret[0].(*dto.LocationDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocationUpdateByID indicates an expected call of LocationUpdateByID.
func (mr *MockLocationServiceMockRecorder) LocationUpdateByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocationUpdateByID", reflect.TypeOf((*MockLocationService)(nil).LocationUpdateByID), arg0)
}