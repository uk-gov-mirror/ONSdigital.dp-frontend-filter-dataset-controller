// Code generated by MockGen. DO NOT EDIT.
// Source: clients.go

// Package handlers is a generated GoMock package.
package handlers

import (
	context "context"
	dataset "github.com/ONSdigital/dp-api-clients-go/dataset"
	filter "github.com/ONSdigital/dp-api-clients-go/filter"
	hierarchy "github.com/ONSdigital/dp-api-clients-go/hierarchy"
	search "github.com/ONSdigital/dp-api-clients-go/search"
	healthcheck "github.com/ONSdigital/dp-healthcheck/healthcheck"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

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

// Checker mocks base method
func (m *MockFilterClient) Checker(ctx context.Context, check *healthcheck.CheckState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checker", ctx, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// Checker indicates an expected call of Checker
func (mr *MockFilterClientMockRecorder) Checker(ctx, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checker", reflect.TypeOf((*MockFilterClient)(nil).Checker), ctx, check)
}

// GetDimensions mocks base method
func (m *MockFilterClient) GetDimensions(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, filterID string) ([]filter.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensions", ctx, userAuthToken, serviceAuthToken, collectionID, filterID)
	ret0, _ := ret[0].([]filter.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensions indicates an expected call of GetDimensions
func (mr *MockFilterClientMockRecorder) GetDimensions(ctx, userAuthToken, serviceAuthToken, collectionID, filterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensions", reflect.TypeOf((*MockFilterClient)(nil).GetDimensions), ctx, userAuthToken, serviceAuthToken, collectionID, filterID)
}

// GetDimensionOptions mocks base method
func (m *MockFilterClient) GetDimensionOptions(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, filterID, name string) ([]filter.DimensionOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensionOptions", ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name)
	ret0, _ := ret[0].([]filter.DimensionOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensionOptions indicates an expected call of GetDimensionOptions
func (mr *MockFilterClientMockRecorder) GetDimensionOptions(ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionOptions", reflect.TypeOf((*MockFilterClient)(nil).GetDimensionOptions), ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name)
}

// GetJobState mocks base method
func (m *MockFilterClient) GetJobState(ctx context.Context, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, filterID string) (filter.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobState", ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, filterID)
	ret0, _ := ret[0].(filter.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobState indicates an expected call of GetJobState
func (mr *MockFilterClientMockRecorder) GetJobState(ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, filterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobState", reflect.TypeOf((*MockFilterClient)(nil).GetJobState), ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, filterID)
}

// GetOutput mocks base method
func (m *MockFilterClient) GetOutput(ctx context.Context, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, filterOutputID string) (filter.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutput", ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, filterOutputID)
	ret0, _ := ret[0].(filter.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutput indicates an expected call of GetOutput
func (mr *MockFilterClientMockRecorder) GetOutput(ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, filterOutputID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutput", reflect.TypeOf((*MockFilterClient)(nil).GetOutput), ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, filterOutputID)
}

// GetDimension mocks base method
func (m *MockFilterClient) GetDimension(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, filterID, name string) (filter.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimension", ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name)
	ret0, _ := ret[0].(filter.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimension indicates an expected call of GetDimension
func (mr *MockFilterClientMockRecorder) GetDimension(ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimension", reflect.TypeOf((*MockFilterClient)(nil).GetDimension), ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name)
}

// AddDimensionValue mocks base method
func (m *MockFilterClient) AddDimensionValue(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, filterID, name, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDimensionValue", ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDimensionValue indicates an expected call of AddDimensionValue
func (mr *MockFilterClientMockRecorder) AddDimensionValue(ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDimensionValue", reflect.TypeOf((*MockFilterClient)(nil).AddDimensionValue), ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name, value)
}

// RemoveDimensionValue mocks base method
func (m *MockFilterClient) RemoveDimensionValue(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, filterID, name, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDimensionValue", ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDimensionValue indicates an expected call of RemoveDimensionValue
func (mr *MockFilterClientMockRecorder) RemoveDimensionValue(ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDimensionValue", reflect.TypeOf((*MockFilterClient)(nil).RemoveDimensionValue), ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name, value)
}

// RemoveDimension mocks base method
func (m *MockFilterClient) RemoveDimension(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, filterID, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveDimension", ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveDimension indicates an expected call of RemoveDimension
func (mr *MockFilterClientMockRecorder) RemoveDimension(ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveDimension", reflect.TypeOf((*MockFilterClient)(nil).RemoveDimension), ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name)
}

// AddDimension mocks base method
func (m *MockFilterClient) AddDimension(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, filterID, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDimension", ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDimension indicates an expected call of AddDimension
func (mr *MockFilterClientMockRecorder) AddDimension(ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDimension", reflect.TypeOf((*MockFilterClient)(nil).AddDimension), ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name)
}

// AddDimensionValues mocks base method
func (m *MockFilterClient) AddDimensionValues(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, filterID, name string, options []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDimensionValues", ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDimensionValues indicates an expected call of AddDimensionValues
func (mr *MockFilterClientMockRecorder) AddDimensionValues(ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDimensionValues", reflect.TypeOf((*MockFilterClient)(nil).AddDimensionValues), ctx, userAuthToken, serviceAuthToken, collectionID, filterID, name, options)
}

// UpdateBlueprint mocks base method
func (m_2 *MockFilterClient) UpdateBlueprint(ctx context.Context, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID string, m filter.Model, doSubmit bool) (filter.Model, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "UpdateBlueprint", ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, m, doSubmit)
	ret0, _ := ret[0].(filter.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBlueprint indicates an expected call of UpdateBlueprint
func (mr *MockFilterClientMockRecorder) UpdateBlueprint(ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, m, doSubmit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBlueprint", reflect.TypeOf((*MockFilterClient)(nil).UpdateBlueprint), ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, m, doSubmit)
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

// Checker mocks base method
func (m *MockDatasetClient) Checker(ctx context.Context, check *healthcheck.CheckState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checker", ctx, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// Checker indicates an expected call of Checker
func (mr *MockDatasetClientMockRecorder) Checker(ctx, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checker", reflect.TypeOf((*MockDatasetClient)(nil).Checker), ctx, check)
}

// Get mocks base method
func (m *MockDatasetClient) Get(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, datasetID string) (dataset.DatasetDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, userAuthToken, serviceAuthToken, collectionID, datasetID)
	ret0, _ := ret[0].(dataset.DatasetDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDatasetClientMockRecorder) Get(ctx, userAuthToken, serviceAuthToken, collectionID, datasetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDatasetClient)(nil).Get), ctx, userAuthToken, serviceAuthToken, collectionID, datasetID)
}

// GetVersion mocks base method
func (m *MockDatasetClient) GetVersion(ctx context.Context, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, datasetID, edition, version string) (dataset.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, datasetID, edition, version)
	ret0, _ := ret[0].(dataset.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockDatasetClientMockRecorder) GetVersion(ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, datasetID, edition, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockDatasetClient)(nil).GetVersion), ctx, userAuthToken, serviceAuthToken, downloadServiceToken, collectionID, datasetID, edition, version)
}

// GetVersionDimensions mocks base method
func (m *MockDatasetClient) GetVersionDimensions(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, id, edition, version string) (dataset.VersionDimensions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionDimensions", ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version)
	ret0, _ := ret[0].(dataset.VersionDimensions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersionDimensions indicates an expected call of GetVersionDimensions
func (mr *MockDatasetClientMockRecorder) GetVersionDimensions(ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionDimensions", reflect.TypeOf((*MockDatasetClient)(nil).GetVersionDimensions), ctx, userAuthToken, serviceAuthToken, collectionID, id, edition, version)
}

// GetOptions mocks base method
func (m *MockDatasetClient) GetOptions(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, datasetID, edition, version, dimension string) (dataset.Options, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOptions", ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition, version, dimension)
	ret0, _ := ret[0].(dataset.Options)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOptions indicates an expected call of GetOptions
func (mr *MockDatasetClientMockRecorder) GetOptions(ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition, version, dimension interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOptions", reflect.TypeOf((*MockDatasetClient)(nil).GetOptions), ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition, version, dimension)
}

// GetVersionMetadata mocks base method
func (m *MockDatasetClient) GetVersionMetadata(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, datasetID, edition, version string) (dataset.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersionMetadata", ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition, version)
	ret0, _ := ret[0].(dataset.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersionMetadata indicates an expected call of GetVersionMetadata
func (mr *MockDatasetClientMockRecorder) GetVersionMetadata(ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersionMetadata", reflect.TypeOf((*MockDatasetClient)(nil).GetVersionMetadata), ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition, version)
}

// GetVersionMetadata mocks base method
func (m *MockDatasetClient) GetEdition(ctx context.Context, userAuthToken, serviceAuthToken, collectionID, datasetID, edition string) (dataset.Edition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEdition", ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition)
	ret0, _ := ret[0].(dataset.Edition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersionMetadata indicates an expected call of GetVersionMetadata
func (mr *MockDatasetClientMockRecorder) GetEdition(ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEdition", reflect.TypeOf((*MockDatasetClient)(nil).GetEdition), ctx, userAuthToken, serviceAuthToken, collectionID, datasetID, edition)
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

// Checker mocks base method
func (m *MockHierarchyClient) Checker(ctx context.Context, check *healthcheck.CheckState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checker", ctx, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// Checker indicates an expected call of Checker
func (mr *MockHierarchyClientMockRecorder) Checker(ctx, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checker", reflect.TypeOf((*MockHierarchyClient)(nil).Checker), ctx, check)
}

// GetRoot mocks base method
func (m *MockHierarchyClient) GetRoot(ctx context.Context, instanceID, name string) (hierarchy.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoot", ctx, instanceID, name)
	ret0, _ := ret[0].(hierarchy.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoot indicates an expected call of GetRoot
func (mr *MockHierarchyClientMockRecorder) GetRoot(ctx, instanceID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoot", reflect.TypeOf((*MockHierarchyClient)(nil).GetRoot), ctx, instanceID, name)
}

// GetChild mocks base method
func (m *MockHierarchyClient) GetChild(ctx context.Context, instanceID, name, code string) (hierarchy.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChild", ctx, instanceID, name, code)
	ret0, _ := ret[0].(hierarchy.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChild indicates an expected call of GetChild
func (mr *MockHierarchyClientMockRecorder) GetChild(ctx, instanceID, name, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChild", reflect.TypeOf((*MockHierarchyClient)(nil).GetChild), ctx, instanceID, name, code)
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

// Checker mocks base method
func (m *MockSearchClient) Checker(ctx context.Context, check *healthcheck.CheckState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checker", ctx, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// Checker indicates an expected call of Checker
func (mr *MockSearchClientMockRecorder) Checker(ctx, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checker", reflect.TypeOf((*MockSearchClient)(nil).Checker), ctx, check)
}

// Dimension mocks base method
func (m *MockSearchClient) Dimension(ctx context.Context, datasetID, edition, version, name, query string, params ...search.Config) (*search.Model, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, datasetID, edition, version, name, query}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Dimension", varargs...)
	ret0, _ := ret[0].(*search.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dimension indicates an expected call of Dimension
func (mr *MockSearchClientMockRecorder) Dimension(ctx, datasetID, edition, version, name, query interface{}, params ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, datasetID, edition, version, name, query}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dimension", reflect.TypeOf((*MockSearchClient)(nil).Dimension), varargs...)
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

// Checker mocks base method
func (m *MockRenderer) Checker(ctx context.Context, check *healthcheck.CheckState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checker", ctx, check)
	ret0, _ := ret[0].(error)
	return ret0
}

// Checker indicates an expected call of Checker
func (mr *MockRendererMockRecorder) Checker(ctx, check interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checker", reflect.TypeOf((*MockRenderer)(nil).Checker), ctx, check)
}

// Do mocks base method
func (m *MockRenderer) Do(path string, b []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", path, b)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockRendererMockRecorder) Do(path, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockRenderer)(nil).Do), path, b)
}
