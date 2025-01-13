package postgresql

import (
	"database/sql"
)

type Logger interface {
	Debug(string)
	Debugf(string, []interface{})
	Info(string)
	Infof(string, []interface{})
}
type Postgresql struct {
	Db     *sql.DB
	Logger Logger
}

func New(db *sql.DB, logger Logger) *Postgresql {
	return &Postgresql{
		Db:     db,
		Logger: logger,
	}

}
