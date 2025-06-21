package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/dapod93/learn-grpc/common/exception"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadDotEnv(filename string) {
	if err := godotenv.Load(filename); err != nil {
		log.Info().
			Err(err).
			Caller().
			Msg(fmt.Sprintf("%s file not found, will use os env instead", filename))
	}
}

func GetRequiredEnv(key string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		exception.Fatal(
			errors.New("required ent not found"),
			fmt.Sprintf("Required %s env not found", key),
		)
	}

	return value
}
