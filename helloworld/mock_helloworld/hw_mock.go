// Automatically generated by MockGen. DO NOT EDIT!
// Source: devfest/helloworld/helloworld (interfaces: GreeterClient)

package mock_helloworld

import (
	helloworld "devfest/helloworld/helloworld"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Mock of GreeterClient interface
type MockGreeterClient struct {
	ctrl     *gomock.Controller
	recorder *_MockGreeterClientRecorder
}

// Recorder for MockGreeterClient (not exported)
type _MockGreeterClientRecorder struct {
	mock *MockGreeterClient
}

func NewMockGreeterClient(ctrl *gomock.Controller) *MockGreeterClient {
	mock := &MockGreeterClient{ctrl: ctrl}
	mock.recorder = &_MockGreeterClientRecorder{mock}
	return mock
}

func (_m *MockGreeterClient) EXPECT() *_MockGreeterClientRecorder {
	return _m.recorder
}

func (_m *MockGreeterClient) SayHello(_param0 context.Context, _param1 *helloworld.HelloRequest, _param2 ...grpc.CallOption) (*helloworld.HelloReply, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "SayHello", _s...)
	ret0, _ := ret[0].(*helloworld.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockGreeterClientRecorder) SayHello(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SayHello", _s...)
}
