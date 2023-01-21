package getInfoNode

import (
	"context"
	"errors"
	kitlog "github.com/go-kit/kit/log"
	"github.com/nelsongp/testF/internal/getInfoNodeRed"
)

type showInfoNodeRedService struct {
	log     kitlog.Logger
	getRepo getInfoNodeRed.NodeRedRepository
}

func NewShowInfoNodeRedService(log kitlog.Logger, getRepo getInfoNodeRed.NodeRedRepository) *showInfoNodeRedService {
	return &showInfoNodeRedService{log: log, getRepo: getRepo}
}

func (t *showInfoNodeRedService) ShowNodeProcessedInfo(ctx context.Context, startDate string, endDate string) ([]getInfoNodeRed.ShowResponse, error) {
	//var node getInfoNodeRed.NodeRedResponse
	var lst []getInfoNodeRed.ShowResponse
	lstN, err := t.getRepo.GetNodeRedInfo(ctx)
	if err != nil {
		t.log.Log("Error getting data")
		return nil, errors.New("Error getting data")
	}
	if len(lstN) == 0 {
		t.log.Log("No data to read")
		return lst, nil
	}

	for _, l := range lstN {
		lst = append(lst, getInfoNodeRed.ShowResponse{
			TurbinaUno: l.V1,
			Separador:  l.V2,
			Enfriador:  l.V3,
		})
	}

	return lst, nil
}
