package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/nelsongp/testF/internal/getInfoNodeRed"
)

func MakeRetrieveNodeRedInfo(t getInfoNodeRed.ShowProcessedInfo) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ShowInfoRequest)
		resp, err := t.ShowNodeProcessedInfo(req.ctx, req.StartDate, req.EndDate)
		return ShowInfoResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type ShowInfoRequest struct {
	ctx       context.Context
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type ShowInfoResponse struct {
	Response interface{}
	Err      error
}
