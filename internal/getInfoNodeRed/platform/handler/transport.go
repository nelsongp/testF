package handler

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpShowNodeRedInfoHandler(path string, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Handle(path,
		httptransport.NewServer(endpoints,
			DecodeRequestShowInfoRed,
			EncodeShowInfoRed,
		)).Methods(http.MethodPost)
	return r
}

func DecodeRequestShowInfoRed(ctx context.Context, r *http.Request) (interface{}, error) {
	var showInfoRequest ShowInfoRequest
	err := json.NewDecoder(r.Body).Decode(&showInfoRequest)
	showInfoRequest.ctx = ctx
	return showInfoRequest, err
}

func EncodeShowInfoRed(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	resp, _ := response.(ShowInfoResponse)
	if resp.Err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		return json.NewEncoder(w).Encode(resp.Err.Error())
	}
	return json.NewEncoder(w).Encode(resp.Response)
}
