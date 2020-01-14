package logpgx

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type Logger struct {
}

func (l *Logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	log.Println("level:", level, "msg:", msg, "data", data)
}
