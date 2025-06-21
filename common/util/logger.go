package util

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Logger = zerolog.
		New(
			zerolog.MultiLevelWriter(
				[]io.Writer{
					zerolog.ConsoleWriter{Out: os.Stdout},
				}...,
			),
		).
		With().
		Stack().
		Timestamp().
		Logger()
}
