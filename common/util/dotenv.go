package util

import (
	"fmt"

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
