package interceptor

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func UnaryLoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		start := time.Now().UTC()

		resp, err = handler(ctx, req)
		if err != nil {
			log.Error().
				Ctx(ctx).
				Err(err).
				Str("method", info.FullMethod).
				Msg("Request failed")
		}

		log.Info().
			Ctx(ctx).
			Str("method", info.FullMethod).
			Dur("duration", time.Since(start)).
			Msg("Incoming request")

		return resp, err
	}
}
