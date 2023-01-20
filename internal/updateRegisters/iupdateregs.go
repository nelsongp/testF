package updateRegisters

import (
	"context"
	"github.com/nelsongp/testF/internal/updateRegisters/platform/storage/mysql"
)

type UpdateRegs interface {
	UpdateRegisters() error
}

type UpdateNProcess interface {
	UpdateNotProcess(ctx context.Context, idNot int) error
}

type SelectNotProcess interface {
	GetNotProcessed(ctx context.Context) ([]mysql.SqlNoProccesedRepo, error)
}
