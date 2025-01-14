package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"songLibrary/internal/logger"
)

func StoreJsonRequest[Req any, Res any](db *sql.DB, logger logger.Logger, funcName string, req *Req) (*Res, error) {
	j, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("Select * FROM public.%s($1)", funcName)
	logger.Debugf("| SQL query | %s | %s |", query, string(j))
	row := db.QueryRow(query, j)
	if err := row.Scan(&j); err != nil {
		return nil, err
	}
	res := new(Res)
	err = json.Unmarshal(j, res)
	return res, err
}
