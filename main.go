package main

import (
	"net"

	"buf.build/go/protovalidate"
	"github.com/dapod93/learn-grpc/common/config"
	"github.com/dapod93/learn-grpc/common/exception"
	"github.com/dapod93/learn-grpc/common/interceptor"
	"github.com/dapod93/learn-grpc/common/util"
	"github.com/dapod93/learn-grpc/database"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	router "github.com/dapod93/learn-grpc/src/router/grpc"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
)

func init() {
	util.LoadDotEnv(".env")
	util.InitLogger()

	config.Server = config.ServerConfig{
		DatabaseLocation: util.GetRequiredEnv("DATABASE_LOCATION"),
	}

}

func main() {
	database.DbConn = database.NewConn()
	defer database.DbConn.Close()

	validator, err := protovalidate.New()
	exception.Fatal(err, "Failed to create new validator")

	listenPort, err := net.Listen("tcp", ":50051")
	exception.Fatal(err, "Failed to listen on port :50051")
	defer listenPort.Close()

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryLoggerInterceptor(),
			interceptor.UnaryRecoverInterceptor(),
			protovalidate_middleware.UnaryServerInterceptor(validator),
		),
	)

	router.RegisterService(server)

	log.Info().Msg("gRPC server starting")
	exception.Fatal(server.Serve(listenPort), "Failed to serve server")
}
