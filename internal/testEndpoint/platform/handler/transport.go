package handler

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpGetTestResponseHandler(path string, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Handle(path,
		httptransport.NewServer(endpoints,
			DecodeGetRequest,
			EncondeGetResponse,
		)).Methods(http.MethodGet)
	return r
}

func DecodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var confRequest GetTestRequest
	return confRequest, nil
}

func EncondeGetResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp, _ := response.(getTestResponse)
	if resp.Err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(resp.Err.Error())
	}
	return json.NewEncoder(w).Encode(resp.Response)
}
