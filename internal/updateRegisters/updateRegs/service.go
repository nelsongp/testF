package updateRegs

import (
	"context"
	"errors"
	kitlog "github.com/go-kit/kit/log"
	"github.com/nelsongp/testF/internal/updateRegisters"
)

type updateRegsSvc struct {
	logger    kitlog.Logger
	repoUpdt  updateRegisters.UpdateNProcess
	selectReg updateRegisters.SelectNotProcess
}

func NewUpdateRegsService(logger kitlog.Logger, repoUpdt updateRegisters.UpdateNProcess, selectReg updateRegisters.SelectNotProcess) *updateRegsSvc {
	return &updateRegsSvc{
		logger:    logger,
		repoUpdt:  repoUpdt,
		selectReg: selectReg,
	}
}

func (u *updateRegsSvc) UpdateRegisters() error {
	u.logger.Log("Initiate Retreive Not Process Registers")
	lstRegs, err := u.selectReg.GetNotProcessed(context.Background())
	if err != nil {
		u.logger.Log("Error retrieving not processed registers")
		return errors.New("Error processing list of not processd Registers")
	}
	if len(lstRegs) == 0 {
		u.logger.Log("No data to update")
		return nil
	}
	for _, r := range lstRegs {
		u.logger.Log("Going to update registers")
		err = u.repoUpdt.UpdateNotProcess(context.Background(), r.Id)
		if err != nil {
			u.logger.Log("Error updating values", "id", r.Id)
			return errors.New("Error updating register")
		}
	}
	return nil
}
