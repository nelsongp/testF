package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/kit/log"
	"github.com/nelsongp/testF/internal/getInfoNodeRed"
)

type nodeRedRepo struct {
	db  *sql.DB
	log kitlog.Logger
}

func NewNodeRedRepo(db *sql.DB, log kitlog.Logger) *nodeRedRepo {
	return &nodeRedRepo{
		db:  db,
		log: log,
	}
}

func (n *nodeRedRepo) GetNodeRedInfo(ctx context.Context) ([]getInfoNodeRed.NodeRedResponse, error) {
	rows, err := n.db.QueryContext(ctx, "SELECT id, v1, v2, v3 FROM test where readed = true")
	if err != nil {
		n.log.Log("Error reading database", "error", err)
		return nil, err
	}
	defer rows.Close()
	var lstVal []getInfoNodeRed.NodeRedResponse

	for rows.Next() {
		var resp getInfoNodeRed.NodeRedResponse
		var val sqlNodeRedMainRepo
		if err := rows.Scan(&val.Id, &val.V1, &val.V2, &val.V3); err != nil {
			n.log.Log("Error while trying to scan query from DB to internal struct", "error", err.Error())
			return nil, err
		}
		resp = getInfoNodeRed.NodeRedResponse{
			Id: val.Id,
			V1: val.V1,
			V2: val.V2,
			V3: val.V3,
		}
		lstVal = append(lstVal, resp)
	}
	return lstVal, nil
}
