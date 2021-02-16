package grpc_sentry_test

import (
	"context"
	"errors"
	"testing"
  "github.com/danielpieper/go-grpc-sentry"
)

type ExceptionCapturerMock struct {
	CaughtError error
}

func (ec *ExceptionCapturerMock) CaptureException(err error) {
	ec.CaughtError = err
}

func Test_UnaryServerInterceptor(t *testing.T) {
	expectedRes := "result"
	expectedErr := errors.New("error")
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return expectedRes, expectedErr
	}
	mock := ExceptionCapturerMock{}
	interceptor := grpc_sentry.UnaryServerInterceptor(&mock)

	actualRes, actualErr := interceptor(context.Background(), nil, nil, handler)
	if mock.CaughtError != expectedErr {
		t.Errorf("expected '%v', got '%v'", expectedErr, mock.CaughtError)
	}

	if actualRes != expectedRes {
		t.Errorf("expected '%v', got '%v'", expectedRes, actualRes)
	}

	if actualErr != expectedErr {
		t.Errorf("expected '%v', got '%v'", expectedErr, actualErr)
	}
}
