// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	slack "github.com/slack-go/slack"
)

// MockSlackClient is a mock of SlackClient interface.
type MockSlackClient struct {
	ctrl     *gomock.Controller
	recorder *MockSlackClientMockRecorder
}

// MockSlackClientMockRecorder is the mock recorder for MockSlackClient.
type MockSlackClientMockRecorder struct {
	mock *MockSlackClient
}

// NewMockSlackClient creates a new mock instance.
func NewMockSlackClient(ctrl *gomock.Controller) *MockSlackClient {
	mock := &MockSlackClient{ctrl: ctrl}
	mock.recorder = &MockSlackClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlackClient) EXPECT() *MockSlackClientMockRecorder {
	return m.recorder
}

// SendMessageContext mocks base method.
func (m *MockSlackClient) SendMessageContext(ctx context.Context, channelID string, options ...slack.MsgOption) (string, string, string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, channelID}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageContext", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// SendMessageContext indicates an expected call of SendMessageContext.
func (mr *MockSlackClientMockRecorder) SendMessageContext(ctx, channelID interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, channelID}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageContext", reflect.TypeOf((*MockSlackClient)(nil).SendMessageContext), varargs...)
}