package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
}
type Postgresql struct {
	Db     *sql.DB
	Logger Logger
}

func New(databaseConnectString string, logger Logger) (*Postgresql, error) {
	db, err := sql.Open("postgres", databaseConnectString)
	if err != nil {
		return nil, err
	}

	return &Postgresql{
		Db:     db,
		Logger: logger,
	}, nil
}
func (p *Postgresql) Close() error {
	return p.Db.Close()
}
