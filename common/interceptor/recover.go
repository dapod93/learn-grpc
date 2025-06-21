package interceptor

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryRecoverInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Error().
					Ctx(ctx).
					Str("method", info.FullMethod).
					Str("panic", fmt.Sprintf("%v", r)).
					Str("stack", string(debug.Stack())).
					Msg("Recovered from panic")

				switch v := r.(type) {
				case error:
					err = extractError(v)

				case string:
					err = extractError(errors.New(v))
				}
			}

		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}

func extractError(err error) error {
	return status.Error(codes.Internal, err.Error())
}
