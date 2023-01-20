package mysql

import (
	"context"
	"database/sql"
	kitlog "github.com/go-kit/kit/log"
)

type updateRepo struct {
	db  *sql.DB
	log kitlog.Logger
}

func NewUpdateRepo(db *sql.DB, log kitlog.Logger) *updateRepo {
	return &updateRepo{
		db:  db,
		log: log,
	}
}

func (n *updateRepo) UpdateNotProcess(ctx context.Context, idNot int) error {
	query := "UPDATE test SET readed = ? WHERE id = ?"
	statement, err := n.db.PrepareContext(ctx, query)
	//query, err := ecr.db.ExecContext(ctx, , ecr.config.GetString("appProperties.subscription.expireStatus"), accountToExpire.SusId, accountToExpire.SusMembershipNumber)
	if err != nil {
		n.log.Log("Error while trying to update subscription status", "error:", err)
		return err
	}
	rows, errExec := statement.ExecContext(ctx, true, idNot)
	if errExec != nil {
		n.log.Log("Error while trying to update subscription status", "error:", errExec)
		return errExec
	}
	rowsAffected, rowsAffectedError := rows.RowsAffected()
	if rowsAffectedError != nil {
		n.log.Log("Error obtaining number of rows affected by executed query", "error: ", rowsAffectedError)
		return rowsAffectedError
	}

	if rowsAffected == 0 {
		n.log.Log("No rows updated after executing update query", "No", "update")
	}

	n.log.Log("Update account status process finished", "process", "finish")
	return nil
}
