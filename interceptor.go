package errorinterceptor

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(err error) error {
	responseStatus, _ := status.FromError(err)
	if responseStatus.Code() != codes.Unknown {
		return err
	}

	ts := time.Now().UTC()
	id := uuid.New()
	replacement := fmt.Sprintf("%s-%s", ts.String(), id.String())

	// For now, log out the original error
	log.Printf("%s %v", replacement, err)

	// Create replacement error
	return status.Error(codes.Internal, replacement)
}

// StreamInterceptor hides result errors that are unknown.
func StreamInterceptor(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler) error {
	return handleError(handler(srv, ss))
}

// UnaryInterceptor hides result errors that are unknown.
func UnaryInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	res, err := handler(ctx, req)
	return res, handleError(err)
}
