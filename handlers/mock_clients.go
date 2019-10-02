// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ONSdigital/dp-frontend-filter-dataset-controller/handlers (interfaces: DatasetClient,FilterClient,HierarchyClient,Renderer,SearchClient)

// Package handlers is a generated GoMock package.
package handlers

import (
	context "context"
	dataset "github.com/ONSdigital/dp-api-clients-go/dataset"
	filter "github.com/ONSdigital/dp-api-clients-go/filter"
	hierarchy "github.com/ONSdigital/go-ns/clients/hierarchy"
	search "github.com/ONSdigital/go-ns/clients/search"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDatasetClient is a mock of DatasetClient interface
type MockDatasetClient struct {
	ctrl     *gomock.Controller
	recorder *MockDatasetClientMockRecorder
}

// MockDatasetClientMockRecorder is the mock recorder for MockDatasetClient
type MockDatasetClientMockRecorder struct {
	mock *MockDatasetClient
}

// NewMockDatasetClient creates a new mock instance
func NewMockDatasetClient(ctrl *gomock.Controller) *MockDatasetClient {
	mock := &MockDatasetClient{ctrl: ctrl}
	mock.recorder = &MockDatasetClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatasetClient) EXPECT() *MockDatasetClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockDatasetClient) Get(arg0 context.Context, arg1, arg2, arg3, arg4 string) (dataset.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(dataset.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDatasetClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDatasetClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// GetDimensions mocks base method
func (m *MockDatasetClient) GetDimensions(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6 string) (dataset.Dimensions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensions", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(dataset.Dimensions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensions indicates an expected call of GetDimensions
func (mr *MockDatasetClientMockRecorder) GetDimensions(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensions", reflect.TypeOf((*MockDatasetClient)(nil).GetDimensions), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// GetOptions mocks base method
func (m *MockDatasetClient) GetOptions(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6, arg7 string) (dataset.Options, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOptions", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(dataset.Options)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOptions indicates an expected call of GetOptions
func (mr *MockDatasetClientMockRecorder) GetOptions(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOptions", reflect.TypeOf((*MockDatasetClient)(nil).GetOptions), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GetVersion mocks base method
func (m *MockDatasetClient) GetVersion(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6, arg7 string) (dataset.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(dataset.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockDatasetClientMockRecorder) GetVersion(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockDatasetClient)(nil).GetVersion), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// GetVersionMetadata mocks base method
func (m *MockDatasetClient) GetVersionMetadata(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6 string) (dataset.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionMetadata", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(dataset.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersionMetadata indicates an expected call of GetVersionMetadata
func (mr *MockDatasetClientMockRecorder) GetVersionMetadata(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionMetadata", reflect.TypeOf((*MockDatasetClient)(nil).GetVersionMetadata), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// Healthcheck mocks base method
func (m *MockDatasetClient) Healthcheck() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthcheck")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Healthcheck indicates an expected call of Healthcheck
func (mr *MockDatasetClientMockRecorder) Healthcheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthcheck", reflect.TypeOf((*MockDatasetClient)(nil).Healthcheck))
}

// MockFilterClient is a mock of FilterClient interface
type MockFilterClient struct {
	ctrl     *gomock.Controller
	recorder *MockFilterClientMockRecorder
}

// MockFilterClientMockRecorder is the mock recorder for MockFilterClient
type MockFilterClientMockRecorder struct {
	mock *MockFilterClient
}

// NewMockFilterClient creates a new mock instance
func NewMockFilterClient(ctrl *gomock.Controller) *MockFilterClient {
	mock := &MockFilterClient{ctrl: ctrl}
	mock.recorder = &MockFilterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilterClient) EXPECT() *MockFilterClientMockRecorder {
	return m.recorder
}

// AddDimension mocks base method
func (m *MockFilterClient) AddDimension(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDimension", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDimension indicates an expected call of AddDimension
func (mr *MockFilterClientMockRecorder) AddDimension(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDimension", reflect.TypeOf((*MockFilterClient)(nil).AddDimension), arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddDimensionValue mocks base method
func (m *MockFilterClient) AddDimensionValue(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDimensionValue", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDimensionValue indicates an expected call of AddDimensionValue
func (mr *MockFilterClientMockRecorder) AddDimensionValue(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDimensionValue", reflect.TypeOf((*MockFilterClient)(nil).AddDimensionValue), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// AddDimensionValues mocks base method
func (m *MockFilterClient) AddDimensionValues(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string, arg6 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDimensionValues", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDimensionValues indicates an expected call of AddDimensionValues
func (mr *MockFilterClientMockRecorder) AddDimensionValues(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDimensionValues", reflect.TypeOf((*MockFilterClient)(nil).AddDimensionValues), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// CreateBlueprint mocks base method
func (m *MockFilterClient) CreateBlueprint(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6, arg7 string, arg8 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBlueprint", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBlueprint indicates an expected call of CreateBlueprint
func (mr *MockFilterClientMockRecorder) CreateBlueprint(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBlueprint", reflect.TypeOf((*MockFilterClient)(nil).CreateBlueprint), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// GetDimension mocks base method
func (m *MockFilterClient) GetDimension(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) (filter.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimension", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(filter.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimension indicates an expected call of GetDimension
func (mr *MockFilterClientMockRecorder) GetDimension(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimension", reflect.TypeOf((*MockFilterClient)(nil).GetDimension), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetDimensionOptions mocks base method
func (m *MockFilterClient) GetDimensionOptions(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) ([]filter.DimensionOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensionOptions", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]filter.DimensionOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensionOptions indicates an expected call of GetDimensionOptions
func (mr *MockFilterClientMockRecorder) GetDimensionOptions(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionOptions", reflect.TypeOf((*MockFilterClient)(nil).GetDimensionOptions), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetDimensions mocks base method
func (m *MockFilterClient) GetDimensions(arg0 context.Context, arg1, arg2, arg3, arg4 string) ([]filter.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensions", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]filter.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensions indicates an expected call of GetDimensions
func (mr *MockFilterClientMockRecorder) GetDimensions(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensions", reflect.TypeOf((*MockFilterClient)(nil).GetDimensions), arg0, arg1, arg2, arg3, arg4)
}

// GetJobState mocks base method
func (m *MockFilterClient) GetJobState(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) (filter.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobState", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(filter.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobState indicates an expected call of GetJobState
func (mr *MockFilterClientMockRecorder) GetJobState(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobState", reflect.TypeOf((*MockFilterClient)(nil).GetJobState), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetOutput mocks base method
func (m *MockFilterClient) GetOutput(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) (filter.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutput", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(filter.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutput indicates an expected call of GetOutput
func (mr *MockFilterClientMockRecorder) GetOutput(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutput", reflect.TypeOf((*MockFilterClient)(nil).GetOutput), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetPreview mocks base method
func (m *MockFilterClient) GetPreview(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) (filter.Preview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreview", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(filter.Preview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreview indicates an expected call of GetPreview
func (mr *MockFilterClientMockRecorder) GetPreview(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreview", reflect.TypeOf((*MockFilterClient)(nil).GetPreview), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Healthcheck mocks base method
func (m *MockFilterClient) Healthcheck() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthcheck")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Healthcheck indicates an expected call of Healthcheck
func (mr *MockFilterClientMockRecorder) Healthcheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthcheck", reflect.TypeOf((*MockFilterClient)(nil).Healthcheck))
}

// RemoveDimension mocks base method
func (m *MockFilterClient) RemoveDimension(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDimension", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDimension indicates an expected call of RemoveDimension
func (mr *MockFilterClientMockRecorder) RemoveDimension(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDimension", reflect.TypeOf((*MockFilterClient)(nil).RemoveDimension), arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemoveDimensionValue mocks base method
func (m *MockFilterClient) RemoveDimensionValue(arg0 context.Context, arg1, arg2, arg3, arg4, arg5, arg6 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDimensionValue", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDimensionValue indicates an expected call of RemoveDimensionValue
func (mr *MockFilterClientMockRecorder) RemoveDimensionValue(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDimensionValue", reflect.TypeOf((*MockFilterClient)(nil).RemoveDimensionValue), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// UpdateBlueprint mocks base method
func (m *MockFilterClient) UpdateBlueprint(arg0 context.Context, arg1, arg2, arg3, arg4 string, arg5 filter.Model, arg6 bool) (filter.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBlueprint", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(filter.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBlueprint indicates an expected call of UpdateBlueprint
func (mr *MockFilterClientMockRecorder) UpdateBlueprint(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBlueprint", reflect.TypeOf((*MockFilterClient)(nil).UpdateBlueprint), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// MockHierarchyClient is a mock of HierarchyClient interface
type MockHierarchyClient struct {
	ctrl     *gomock.Controller
	recorder *MockHierarchyClientMockRecorder
}

// MockHierarchyClientMockRecorder is the mock recorder for MockHierarchyClient
type MockHierarchyClientMockRecorder struct {
	mock *MockHierarchyClient
}

// NewMockHierarchyClient creates a new mock instance
func NewMockHierarchyClient(ctrl *gomock.Controller) *MockHierarchyClient {
	mock := &MockHierarchyClient{ctrl: ctrl}
	mock.recorder = &MockHierarchyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHierarchyClient) EXPECT() *MockHierarchyClientMockRecorder {
	return m.recorder
}

// GetChild mocks base method
func (m *MockHierarchyClient) GetChild(arg0, arg1, arg2 string) (hierarchy.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChild", arg0, arg1, arg2)
	ret0, _ := ret[0].(hierarchy.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChild indicates an expected call of GetChild
func (mr *MockHierarchyClientMockRecorder) GetChild(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChild", reflect.TypeOf((*MockHierarchyClient)(nil).GetChild), arg0, arg1, arg2)
}

// GetRoot mocks base method
func (m *MockHierarchyClient) GetRoot(arg0, arg1 string) (hierarchy.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoot", arg0, arg1)
	ret0, _ := ret[0].(hierarchy.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoot indicates an expected call of GetRoot
func (mr *MockHierarchyClientMockRecorder) GetRoot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoot", reflect.TypeOf((*MockHierarchyClient)(nil).GetRoot), arg0, arg1)
}

// Healthcheck mocks base method
func (m *MockHierarchyClient) Healthcheck() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthcheck")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Healthcheck indicates an expected call of Healthcheck
func (mr *MockHierarchyClientMockRecorder) Healthcheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthcheck", reflect.TypeOf((*MockHierarchyClient)(nil).Healthcheck))
}

// MockRenderer is a mock of Renderer interface
type MockRenderer struct {
	ctrl     *gomock.Controller
	recorder *MockRendererMockRecorder
}

// MockRendererMockRecorder is the mock recorder for MockRenderer
type MockRendererMockRecorder struct {
	mock *MockRenderer
}

// NewMockRenderer creates a new mock instance
func NewMockRenderer(ctrl *gomock.Controller) *MockRenderer {
	mock := &MockRenderer{ctrl: ctrl}
	mock.recorder = &MockRendererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRenderer) EXPECT() *MockRendererMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockRenderer) Do(arg0 string, arg1 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockRendererMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockRenderer)(nil).Do), arg0, arg1)
}

// Healthcheck mocks base method
func (m *MockRenderer) Healthcheck() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthcheck")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Healthcheck indicates an expected call of Healthcheck
func (mr *MockRendererMockRecorder) Healthcheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthcheck", reflect.TypeOf((*MockRenderer)(nil).Healthcheck))
}

// MockSearchClient is a mock of SearchClient interface
type MockSearchClient struct {
	ctrl     *gomock.Controller
	recorder *MockSearchClientMockRecorder
}

// MockSearchClientMockRecorder is the mock recorder for MockSearchClient
type MockSearchClientMockRecorder struct {
	mock *MockSearchClient
}

// NewMockSearchClient creates a new mock instance
func NewMockSearchClient(ctrl *gomock.Controller) *MockSearchClient {
	mock := &MockSearchClient{ctrl: ctrl}
	mock.recorder = &MockSearchClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchClient) EXPECT() *MockSearchClientMockRecorder {
	return m.recorder
}

// Dimension mocks base method
func (m *MockSearchClient) Dimension(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string, arg6 ...search.Config) (*search.Model, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4, arg5}
	for _, a := range arg6 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Dimension", varargs...)
	ret0, _ := ret[0].(*search.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dimension indicates an expected call of Dimension
func (mr *MockSearchClientMockRecorder) Dimension(arg0, arg1, arg2, arg3, arg4, arg5 interface{}, arg6 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4, arg5}, arg6...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dimension", reflect.TypeOf((*MockSearchClient)(nil).Dimension), varargs...)
}

// Healthcheck mocks base method
func (m *MockSearchClient) Healthcheck() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Healthcheck")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Healthcheck indicates an expected call of Healthcheck
func (mr *MockSearchClientMockRecorder) Healthcheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthcheck", reflect.TypeOf((*MockSearchClient)(nil).Healthcheck))
}
