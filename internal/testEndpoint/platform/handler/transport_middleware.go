package handler

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
)

type Middleware func(endpoint endpoint.Endpoint) endpoint.Endpoint

func GetTestResponseMiddleware(log kitlog.Logger) Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(GetTestRequest)
			defer log.Log("process finished", "request", req)
			return e(ctx, request)
		}
	}
}
