package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/kit/log"
)

type notProcessed struct {
	db  *sql.DB
	log kitlog.Logger
}

func NewNotProcessed(db *sql.DB, log kitlog.Logger) *notProcessed {
	return &notProcessed{
		db:  db,
		log: log,
	}
}

func (n *notProcessed) GetNotProcessed(ctx context.Context) ([]SqlNoProccesedRepo, error) {
	rows, err := n.db.QueryContext(ctx, "SELECT id, v1, v2, v3 FROM test where readed = false")
	if err != nil {
		n.log.Log("Error reading database", "error", err)
		return nil, err
	}
	defer rows.Close()
	var lstVal []SqlNoProccesedRepo

	for rows.Next() {
		var val SqlNoProccesedRepo
		if err := rows.Scan(&val.Id, &val.V1, &val.V2, &val.V3); err != nil {
			n.log.Log("Error while trying to scan query from DB to internal struct", "error", err.Error())
			return nil, err
		}
		lstVal = append(lstVal, val)
	}
	return lstVal, nil
}
