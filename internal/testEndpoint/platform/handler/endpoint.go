package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/nelsongp/testF/internal/testEndpoint"
)

func MakeTestEndpoint(t testEndpoint.TestEndpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTestRequest)
		resp, err := t.ResponseTestService(req.Name)
		return getTestResponse{
			Response: resp,
			Err:      err,
		}, nil
	}
}

type GetTestRequest struct {
	Name string
}

type getTestResponse struct {
	Response interface{}
	Err      error
}
