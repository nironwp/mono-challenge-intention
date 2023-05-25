// Code generated by MockGen. DO NOT EDIT.
// Source: domain/entities/intent-product.go

package mock_entities

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	"intent-service/domain/entities"
)

// MockIntentProductInterface is a mock of IntentProductInterface interface
type MockIntentProductInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIntentProductInterfaceMockRecorder
}

// MockIntentProductInterfaceMockRecorder is the mock recorder for MockIntentProductInterface
type MockIntentProductInterfaceMockRecorder struct {
	mock *MockIntentProductInterface
}

// NewMockIntentProductInterface creates a new mock instance
func NewMockIntentProductInterface(ctrl *gomock.Controller) *MockIntentProductInterface {
	mock := &MockIntentProductInterface{ctrl: ctrl}
	mock.recorder = &MockIntentProductInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIntentProductInterface) EXPECT() *MockIntentProductInterfaceMockRecorder {
	return _m.recorder
}

// GetId mocks base method
func (_m *MockIntentProductInterface) GetId() int {
	ret := _m.ctrl.Call(_m, "GetId")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetId indicates an expected call of GetId
func (_mr *MockIntentProductInterfaceMockRecorder) GetId() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetId", reflect.TypeOf((*MockIntentProductInterface)(nil).GetId))
}

// GetTitle mocks base method
func (_m *MockIntentProductInterface) GetTitle() string {
	ret := _m.ctrl.Call(_m, "GetTitle")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTitle indicates an expected call of GetTitle
func (_mr *MockIntentProductInterfaceMockRecorder) GetTitle() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetTitle", reflect.TypeOf((*MockIntentProductInterface)(nil).GetTitle))
}

// GetPrice mocks base method
func (_m *MockIntentProductInterface) GetPrice() float64 {
	ret := _m.ctrl.Call(_m, "GetPrice")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetPrice indicates an expected call of GetPrice
func (_mr *MockIntentProductInterfaceMockRecorder) GetPrice() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetPrice", reflect.TypeOf((*MockIntentProductInterface)(nil).GetPrice))
}

// GetCategory mocks base method
func (_m *MockIntentProductInterface) GetCategory() string {
	ret := _m.ctrl.Call(_m, "GetCategory")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCategory indicates an expected call of GetCategory
func (_mr *MockIntentProductInterfaceMockRecorder) GetCategory() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetCategory", reflect.TypeOf((*MockIntentProductInterface)(nil).GetCategory))
}

// GetDescription mocks base method
func (_m *MockIntentProductInterface) GetDescription() string {
	ret := _m.ctrl.Call(_m, "GetDescription")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDescription indicates an expected call of GetDescription
func (_mr *MockIntentProductInterfaceMockRecorder) GetDescription() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetDescription", reflect.TypeOf((*MockIntentProductInterface)(nil).GetDescription))
}

// GetImage mocks base method
func (_m *MockIntentProductInterface) GetImage() string {
	ret := _m.ctrl.Call(_m, "GetImage")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetImage indicates an expected call of GetImage
func (_mr *MockIntentProductInterfaceMockRecorder) GetImage() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetImage", reflect.TypeOf((*MockIntentProductInterface)(nil).GetImage))
}

// MockIntentProductServiceInterface is a mock of IntentProductServiceInterface interface
type MockIntentProductServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIntentProductServiceInterfaceMockRecorder
}

// MockIntentProductServiceInterfaceMockRecorder is the mock recorder for MockIntentProductServiceInterface
type MockIntentProductServiceInterfaceMockRecorder struct {
	mock *MockIntentProductServiceInterface
}

// NewMockIntentProductServiceInterface creates a new mock instance
func NewMockIntentProductServiceInterface(ctrl *gomock.Controller) *MockIntentProductServiceInterface {
	mock := &MockIntentProductServiceInterface{ctrl: ctrl}
	mock.recorder = &MockIntentProductServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIntentProductServiceInterface) EXPECT() *MockIntentProductServiceInterfaceMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockIntentProductServiceInterface) Get(id string) (entities.IntentProductInterface, error) {
	ret := _m.ctrl.Call(_m, "Get", id)
	ret0, _ := ret[0].(entities.IntentProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (_mr *MockIntentProductServiceInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockIntentProductServiceInterface)(nil).Get), arg0)
}

// Create mocks base method
func (_m *MockIntentProductServiceInterface) Create(product entities.IntentProductInterface) (entities.IntentProductInterface, error) {
	ret := _m.ctrl.Call(_m, "Create", product)
	ret0, _ := ret[0].(entities.IntentProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (_mr *MockIntentProductServiceInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Create", reflect.TypeOf((*MockIntentProductServiceInterface)(nil).Create), arg0)
}

// MockIntentProductReader is a mock of IntentProductReader interface
type MockIntentProductReader struct {
	ctrl     *gomock.Controller
	recorder *MockIntentProductReaderMockRecorder
}

// MockIntentProductReaderMockRecorder is the mock recorder for MockIntentProductReader
type MockIntentProductReaderMockRecorder struct {
	mock *MockIntentProductReader
}

// NewMockIntentProductReader creates a new mock instance
func NewMockIntentProductReader(ctrl *gomock.Controller) *MockIntentProductReader {
	mock := &MockIntentProductReader{ctrl: ctrl}
	mock.recorder = &MockIntentProductReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIntentProductReader) EXPECT() *MockIntentProductReaderMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockIntentProductReader) Get(id string) (entities.IntentProductInterface, error) {
	ret := _m.ctrl.Call(_m, "Get", id)
	ret0, _ := ret[0].(entities.IntentProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (_mr *MockIntentProductReaderMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockIntentProductReader)(nil).Get), arg0)
}

// MockIntentProductWriter is a mock of IntentProductWriter interface
type MockIntentProductWriter struct {
	ctrl     *gomock.Controller
	recorder *MockIntentProductWriterMockRecorder
}

// MockIntentProductWriterMockRecorder is the mock recorder for MockIntentProductWriter
type MockIntentProductWriterMockRecorder struct {
	mock *MockIntentProductWriter
}

// NewMockIntentProductWriter creates a new mock instance
func NewMockIntentProductWriter(ctrl *gomock.Controller) *MockIntentProductWriter {
	mock := &MockIntentProductWriter{ctrl: ctrl}
	mock.recorder = &MockIntentProductWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIntentProductWriter) EXPECT() *MockIntentProductWriterMockRecorder {
	return _m.recorder
}

// Save mocks base method
func (_m *MockIntentProductWriter) Save(product entities.IntentProductInterface) (entities.IntentProductInterface, error) {
	ret := _m.ctrl.Call(_m, "Save", product)
	ret0, _ := ret[0].(entities.IntentProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (_mr *MockIntentProductWriterMockRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Save", reflect.TypeOf((*MockIntentProductWriter)(nil).Save), arg0)
}

// MockIntentProductPersistenceInterface is a mock of IntentProductPersistenceInterface interface
type MockIntentProductPersistenceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIntentProductPersistenceInterfaceMockRecorder
}

// MockIntentProductPersistenceInterfaceMockRecorder is the mock recorder for MockIntentProductPersistenceInterface
type MockIntentProductPersistenceInterfaceMockRecorder struct {
	mock *MockIntentProductPersistenceInterface
}

// NewMockIntentProductPersistenceInterface creates a new mock instance
func NewMockIntentProductPersistenceInterface(ctrl *gomock.Controller) *MockIntentProductPersistenceInterface {
	mock := &MockIntentProductPersistenceInterface{ctrl: ctrl}
	mock.recorder = &MockIntentProductPersistenceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockIntentProductPersistenceInterface) EXPECT() *MockIntentProductPersistenceInterfaceMockRecorder {
	return _m.recorder
}

// Get mocks base method
func (_m *MockIntentProductPersistenceInterface) Get(id string) (entities.IntentProductInterface, error) {
	ret := _m.ctrl.Call(_m, "Get", id)
	ret0, _ := ret[0].(entities.IntentProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (_mr *MockIntentProductPersistenceInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockIntentProductPersistenceInterface)(nil).Get), arg0)
}

// Save mocks base method
func (_m *MockIntentProductPersistenceInterface) Save(product entities.IntentProductInterface) (entities.IntentProductInterface, error) {
	ret := _m.ctrl.Call(_m, "Save", product)
	ret0, _ := ret[0].(entities.IntentProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (_mr *MockIntentProductPersistenceInterfaceMockRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Save", reflect.TypeOf((*MockIntentProductPersistenceInterface)(nil).Save), arg0)
}
