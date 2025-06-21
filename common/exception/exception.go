package exception

import (
	"database/sql"

	"github.com/dapod93/learn-grpc/database/ent"
	"github.com/rs/zerolog/log"
)

func Fatal(err error, msg string) {
	if err != nil {
		log.Fatal().Err(err).Msg(msg)
	}
}

// Call panic because if something raise error, there won't be a need to continue the logic
func General(err error, msg string) {
	if err == nil || ent.IsNotFound(err) || err == sql.ErrNoRows {
		return
	}

	log.Panic().Err(err).Caller().Msg(msg)
}
