package grpc_sentry

import (
	"context"
	"google.golang.org/grpc"
)

type EventID string

// ExceptionCapturer specifies the implementation of a method to capture the given error
type ExceptionCapturer interface {
	CaptureException(err error) *EventID
}

// UnaryServerInterceptor creates an interceptor which catches the errors from each service method and reports them to Sentry
func UnaryServerInterceptor(ec ExceptionCapturer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			ec.CaptureException(err)
		}

		return resp, err
	}
}
