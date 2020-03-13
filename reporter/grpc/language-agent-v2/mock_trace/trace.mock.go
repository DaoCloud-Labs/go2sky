// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DaoCloud-Labs/go2sky/reporter/grpc/language-agent-v2 (interfaces: TraceSegmentReportServiceClient)

// Package mock_language_agent_v2 is a generated GoMock package.
package mock_language_agent_v2

import (
	context "context"
	language_agent_v2 "github.com/DaoCloud-Labs/go2sky/reporter/grpc/language-agent-v2"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTraceSegmentReportServiceClient is a mock of TraceSegmentReportServiceClient interface
type MockTraceSegmentReportServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTraceSegmentReportServiceClientMockRecorder
}

// MockTraceSegmentReportServiceClientMockRecorder is the mock recorder for MockTraceSegmentReportServiceClient
type MockTraceSegmentReportServiceClientMockRecorder struct {
	mock *MockTraceSegmentReportServiceClient
}

// NewMockTraceSegmentReportServiceClient creates a new mock instance
func NewMockTraceSegmentReportServiceClient(ctrl *gomock.Controller) *MockTraceSegmentReportServiceClient {
	mock := &MockTraceSegmentReportServiceClient{ctrl: ctrl}
	mock.recorder = &MockTraceSegmentReportServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTraceSegmentReportServiceClient) EXPECT() *MockTraceSegmentReportServiceClientMockRecorder {
	return m.recorder
}

// Collect mocks base method
func (m *MockTraceSegmentReportServiceClient) Collect(arg0 context.Context, arg1 ...grpc.CallOption) (language_agent_v2.TraceSegmentReportService_CollectClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Collect", varargs...)
	ret0, _ := ret[0].(language_agent_v2.TraceSegmentReportService_CollectClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Collect indicates an expected call of Collect
func (mr *MockTraceSegmentReportServiceClientMockRecorder) Collect(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockTraceSegmentReportServiceClient)(nil).Collect), varargs...)
}
