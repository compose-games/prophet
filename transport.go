package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
	
	"github.com/go-kit/kit/endpoint"
)

func makeConsultEndpoint(svc ProphetService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := svc.Consult()
		if err != nil {
			return consultResponse{v, err.Error()}, nil
		}
		return consultResponse{v, ""}, nil
	}
}

func decodeConsultRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request consultRequest
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type consultRequest struct { }

type consultResponse struct {
	S 	string `json:"v"`
	Err	string `json:"err,omitempty"`
}
