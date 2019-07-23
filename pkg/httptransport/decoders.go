package httptransport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"microservice/pkg/endpoints"
)

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// decodeHTTPSumRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded sum request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.SumRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPConcatRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded concat request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPConcatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ConcatRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// decodeHTTPSumResponse is a transport/http.DecodeResponseFunc that decodes a
// JSON-encoded sum response from the HTTP response body. If the response has a
// non-200 status code, we will interpret that as an error and attempt to decode
// the specific error message from the response body. Primarily useful in a
// client.
func decodeHTTPSumResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp endpoints.SumResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

// decodeHTTPConcatResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// has a non-200 status code, we will interpret that as an error and attempt to
// decode the specific error message from the response body. Primarily useful in
// a client.
func decodeHTTPConcatResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp endpoints.ConcatResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
