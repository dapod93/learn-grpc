package main

import (
	"github.com/dapod93/learn-grpc/common/exception"

	"buf.build/go/protovalidate"
)

func main() {
	validator, err := protovalidate.New()
	exception.Fatal(err, "Failed to create new validator")
}
