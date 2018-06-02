package errorinterceptor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	MaskedData = "Private Unknown"
)

var (
	ErrOK                 = status.Error(codes.OK, "OK")
	ErrCanceled           = status.Error(codes.Canceled, "Canceled")
	ErrUnknown            = status.Error(codes.Unknown, MaskedData)
	ErrInvalidArgument    = status.Error(codes.InvalidArgument, "InvalidArgument")
	ErrDeadlineExceeded   = status.Error(codes.DeadlineExceeded, "DeadlineExceeded")
	ErrNotFound           = status.Error(codes.NotFound, "NotFound")
	ErrAlreadyExists      = status.Error(codes.AlreadyExists, "AlreadyExists")
	ErrPermissionDenied   = status.Error(codes.PermissionDenied, "PermissionDenied")
	ErrResourceExhausted  = status.Error(codes.ResourceExhausted, "ResourceExhausted")
	ErrFailedPrecondition = status.Error(codes.FailedPrecondition, "FailedPrecondition")
	ErrAborted            = status.Error(codes.Aborted, "Aborted")
	ErrOutOfRange         = status.Error(codes.OutOfRange, "OutOfRange")
	ErrUnimplemented      = status.Error(codes.Unimplemented, "Unimplemented")
	ErrInternal           = status.Error(codes.Internal, "Internal")
	ErrUnavailable        = status.Error(codes.Unavailable, "Unavailable")
	ErrDataLoss           = status.Error(codes.DataLoss, "DataLoss")
	ErrUnauthenticated    = status.Error(codes.Unauthenticated, "Unauthenticated")
)

func TestStreamInterceptorUnknown(t *testing.T) {
	err := StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrUnknown
	})
	assert.NotEqual(t, ErrUnknown, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.NotContains(t, MaskedData, err.Error())
}

func TestStreamInterceptorPassthrough(t *testing.T) {
	var err error

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return nil
	})
	assert.NoError(t, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrInvalidArgument
	})
	assert.Equal(t, ErrInvalidArgument, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrDeadlineExceeded
	})
	assert.Equal(t, ErrDeadlineExceeded, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrAlreadyExists
	})
	assert.Equal(t, ErrAlreadyExists, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrPermissionDenied
	})
	assert.Equal(t, ErrPermissionDenied, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrResourceExhausted
	})
	assert.Equal(t, ErrResourceExhausted, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrFailedPrecondition
	})
	assert.Equal(t, ErrFailedPrecondition, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrAborted
	})
	assert.Equal(t, ErrAborted, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrOutOfRange
	})
	assert.Equal(t, ErrOutOfRange, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrUnimplemented
	})
	assert.Equal(t, ErrUnimplemented, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrInternal
	})
	assert.Equal(t, ErrInternal, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrUnavailable
	})
	assert.Equal(t, ErrUnavailable, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrDataLoss
	})
	assert.Equal(t, ErrDataLoss, err)

	err = StreamInterceptor(nil, nil, nil, func(srv interface{}, stream grpc.ServerStream) error {
		return ErrUnauthenticated
	})
	assert.Equal(t, ErrUnauthenticated, err)
}

func TestUnaryInterceptorUnknown(t *testing.T) {
	_, err := UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrUnknown
	})
	assert.NotEqual(t, ErrUnknown, err)
	assert.Equal(t, codes.Internal, status.Code(err))
	assert.NotContains(t, MaskedData, err.Error())
}

func TestUnaryInterceptorPassthrough(t *testing.T) {
	var err error
	var res interface{}

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return "data", nil
	})
	assert.NoError(t, err)
	assert.Equal(t, "data", res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrInvalidArgument
	})
	assert.Equal(t, ErrInvalidArgument, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrDeadlineExceeded
	})
	assert.Equal(t, ErrDeadlineExceeded, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrAlreadyExists
	})
	assert.Equal(t, ErrAlreadyExists, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrPermissionDenied
	})
	assert.Equal(t, ErrPermissionDenied, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrResourceExhausted
	})
	assert.Equal(t, ErrResourceExhausted, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrFailedPrecondition
	})
	assert.Equal(t, ErrFailedPrecondition, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrAborted
	})
	assert.Equal(t, ErrAborted, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrOutOfRange
	})
	assert.Equal(t, ErrOutOfRange, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrUnimplemented
	})
	assert.Equal(t, ErrUnimplemented, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrInternal
	})
	assert.Equal(t, ErrInternal, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrUnavailable
	})
	assert.Equal(t, ErrUnavailable, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrDataLoss
	})
	assert.Equal(t, ErrDataLoss, err)
	assert.Nil(t, res)

	res, err = UnaryInterceptor(nil, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, ErrUnauthenticated
	})
	assert.Equal(t, ErrUnauthenticated, err)
	assert.Nil(t, res)
}
