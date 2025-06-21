package util

import "github.com/dapod93/learn-grpc/common/exception"

func MustSet[T any](input T, err error) T {
	if err != nil {
		exception.General(err, "Failed to run MustSet function")
	}

	return input
}
