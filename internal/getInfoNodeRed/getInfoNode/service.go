package getInfoNode

import (
	"context"
	kitlog "github.com/go-kit/kit/log"
	"github.com/nelsongp/testF/internal/getInfoNodeRed"
)

type showInfoNodeRedService struct {
	log kitlog.Logger
}

func NewShowInfoNodeRedService(log kitlog.Logger) *showInfoNodeRedService {
	return &showInfoNodeRedService{log: log}
}

func (t *showInfoNodeRedService) ShowNodeProcessedInfo(ctx context.Context, startDate string, endDate string) ([]getInfoNodeRed.ShowResponse, error) {
	//var node getInfoNodeRed.NodeRedResponse
	var lst []getInfoNodeRed.ShowResponse

	lst = append(lst, getInfoNodeRed.ShowResponse{
		TurbinaUno: 10,
		Separador:  20,
		Enfriador:  30,
	})
	return lst, nil
}
