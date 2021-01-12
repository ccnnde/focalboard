// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-octo-tasks/server/services/store (interfaces: Store)

// Package mockstore is a generated GoMock package.
package mockstore

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/mattermost/mattermost-octo-tasks/server/model"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CleanUpSessions mocks base method
func (m *MockStore) CleanUpSessions(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanUpSessions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanUpSessions indicates an expected call of CleanUpSessions
func (mr *MockStoreMockRecorder) CleanUpSessions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanUpSessions", reflect.TypeOf((*MockStore)(nil).CleanUpSessions), arg0)
}

// CreateSession mocks base method
func (m *MockStore) CreateSession(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of CreateSession
func (mr *MockStoreMockRecorder) CreateSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0)
}

// CreateUser mocks base method
func (m *MockStore) CreateUser(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockStoreMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0)
}

// DeleteBlock mocks base method
func (m *MockStore) DeleteBlock(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBlock indicates an expected call of DeleteBlock
func (mr *MockStoreMockRecorder) DeleteBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBlock", reflect.TypeOf((*MockStore)(nil).DeleteBlock), arg0, arg1)
}

// DeleteSession mocks base method
func (m *MockStore) DeleteSession(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession
func (mr *MockStoreMockRecorder) DeleteSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStore)(nil).DeleteSession), arg0)
}

// GetAllBlocks mocks base method
func (m *MockStore) GetAllBlocks() ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBlocks")
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBlocks indicates an expected call of GetAllBlocks
func (mr *MockStoreMockRecorder) GetAllBlocks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBlocks", reflect.TypeOf((*MockStore)(nil).GetAllBlocks))
}

// GetBlocksWithParent mocks base method
func (m *MockStore) GetBlocksWithParent(arg0 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithParent", arg0)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithParent indicates an expected call of GetBlocksWithParent
func (mr *MockStoreMockRecorder) GetBlocksWithParent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithParent", reflect.TypeOf((*MockStore)(nil).GetBlocksWithParent), arg0)
}

// GetBlocksWithParentAndType mocks base method
func (m *MockStore) GetBlocksWithParentAndType(arg0, arg1 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithParentAndType", arg0, arg1)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithParentAndType indicates an expected call of GetBlocksWithParentAndType
func (mr *MockStoreMockRecorder) GetBlocksWithParentAndType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithParentAndType", reflect.TypeOf((*MockStore)(nil).GetBlocksWithParentAndType), arg0, arg1)
}

// GetBlocksWithType mocks base method
func (m *MockStore) GetBlocksWithType(arg0 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlocksWithType", arg0)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlocksWithType indicates an expected call of GetBlocksWithType
func (mr *MockStoreMockRecorder) GetBlocksWithType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlocksWithType", reflect.TypeOf((*MockStore)(nil).GetBlocksWithType), arg0)
}

// GetParentID mocks base method
func (m *MockStore) GetParentID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParentID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParentID indicates an expected call of GetParentID
func (mr *MockStoreMockRecorder) GetParentID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParentID", reflect.TypeOf((*MockStore)(nil).GetParentID), arg0)
}

// GetSession mocks base method
func (m *MockStore) GetSession(arg0 string, arg1 int64) (*model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetSharing mocks base method
func (m *MockStore) GetSharing(arg0 string) (*model.Sharing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharing", arg0)
	ret0, _ := ret[0].(*model.Sharing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharing indicates an expected call of GetSharing
func (mr *MockStoreMockRecorder) GetSharing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharing", reflect.TypeOf((*MockStore)(nil).GetSharing), arg0)
}

// GetSubTree2 mocks base method
func (m *MockStore) GetSubTree2(arg0 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubTree2", arg0)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubTree2 indicates an expected call of GetSubTree2
func (mr *MockStoreMockRecorder) GetSubTree2(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubTree2", reflect.TypeOf((*MockStore)(nil).GetSubTree2), arg0)
}

// GetSubTree3 mocks base method
func (m *MockStore) GetSubTree3(arg0 string) ([]model.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubTree3", arg0)
	ret0, _ := ret[0].([]model.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubTree3 indicates an expected call of GetSubTree3
func (mr *MockStoreMockRecorder) GetSubTree3(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubTree3", reflect.TypeOf((*MockStore)(nil).GetSubTree3), arg0)
}

// GetSystemSettings mocks base method
func (m *MockStore) GetSystemSettings() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSystemSettings")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSystemSettings indicates an expected call of GetSystemSettings
func (mr *MockStoreMockRecorder) GetSystemSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSystemSettings", reflect.TypeOf((*MockStore)(nil).GetSystemSettings))
}

// GetUserByEmail mocks base method
func (m *MockStore) GetUserByEmail(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (mr *MockStoreMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockStore)(nil).GetUserByEmail), arg0)
}

// GetUserById mocks base method
func (m *MockStore) GetUserById(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById
func (mr *MockStoreMockRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockStore)(nil).GetUserById), arg0)
}

// GetUserByUsername mocks base method
func (m *MockStore) GetUserByUsername(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0)
}

// InsertBlock mocks base method
func (m *MockStore) InsertBlock(arg0 model.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBlock indicates an expected call of InsertBlock
func (mr *MockStoreMockRecorder) InsertBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBlock", reflect.TypeOf((*MockStore)(nil).InsertBlock), arg0)
}

// RefreshSession mocks base method
func (m *MockStore) RefreshSession(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshSession indicates an expected call of RefreshSession
func (mr *MockStoreMockRecorder) RefreshSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshSession", reflect.TypeOf((*MockStore)(nil).RefreshSession), arg0)
}

// SetSystemSetting mocks base method
func (m *MockStore) SetSystemSetting(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSystemSetting", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSystemSetting indicates an expected call of SetSystemSetting
func (mr *MockStoreMockRecorder) SetSystemSetting(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSystemSetting", reflect.TypeOf((*MockStore)(nil).SetSystemSetting), arg0, arg1)
}

// Shutdown mocks base method
func (m *MockStore) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockStoreMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockStore)(nil).Shutdown))
}

// UpdateSession mocks base method
func (m *MockStore) UpdateSession(arg0 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSession", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSession indicates an expected call of UpdateSession
func (mr *MockStoreMockRecorder) UpdateSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSession", reflect.TypeOf((*MockStore)(nil).UpdateSession), arg0)
}

// UpdateUser mocks base method
func (m *MockStore) UpdateUser(arg0 *model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockStoreMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0)
}

// UpsertSharing mocks base method
func (m *MockStore) UpsertSharing(arg0 model.Sharing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertSharing", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertSharing indicates an expected call of UpsertSharing
func (mr *MockStoreMockRecorder) UpsertSharing(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertSharing", reflect.TypeOf((*MockStore)(nil).UpsertSharing), arg0)
}
