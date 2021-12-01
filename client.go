package gohttpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
	newrelic "github.com/newrelic/go-agent"
)

// New func returns a Client interface
func New(baseUrl string) Client {
	return &client{baseUrl: baseUrl, client: &http.Client{Transport: newrelic.NewRoundTripper(nil, http.DefaultTransport)}}
}

// Get func returns a request
func (h client) Get(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", h.baseUrl, endpoint), bytes.NewBuffer([]byte{}))
}

// GetWith func returns a request
func (h client) GetWith(endpoint string, params interface{}) (*http.Request, error) {
	queryString, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?%s", h.baseUrl, endpoint, queryString.Encode()), bytes.NewBuffer([]byte{}))
}

// Post func returns a request
func (h client) Post(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodPost, h.baseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

// PostWith func returns a request
func (h client) PostWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPost, h.baseUrl+endpoint, bytes.NewBuffer(json))
}

// Put func returns a request
func (h client) Put(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodPut, h.baseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

// PutWith func returns a request
func (h client) PutWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPut, h.baseUrl+endpoint, bytes.NewBuffer(json))
}

// Patch func returns a request
func (h client) Patch(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodPatch, h.baseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

// PatchWith func returns a request
func (h client) PatchWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodPatch, h.baseUrl+endpoint, bytes.NewBuffer(json))
}

// Delete func returns a request
func (h client) Delete(endpoint string) (*http.Request, error) {
	return http.NewRequest(http.MethodDelete, h.baseUrl+endpoint, bytes.NewBuffer([]byte{}))
}

// DeleteWith func returns a request
func (h client) DeleteWith(endpoint string, params interface{}) (*http.Request, error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(http.MethodDelete, h.baseUrl+endpoint, bytes.NewBuffer(json))
}

// Do func returns a response with your data
func (h client) Do(ctx context.Context, request *http.Request) (Response, error) {
	req := newrelic.RequestWithTransactionContext(request, newrelic.FromContext(ctx))

	response, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &ResponseStruct{
		Status:        response.Status,
		StatusCode:    response.StatusCode,
		Header:        response.Header,
		ContentLength: response.ContentLength,
		Body:          body,
	}, nil
}

// Get func returns ResponseStruct struct of request
func (r ResponseStruct) Get() ResponseStruct {
	return r
}

// To func returns converts string to struct
func (r ResponseStruct) To(value interface{}) {
	err := json.Unmarshal(r.Body, &value)
	if err != nil {
		value = nil
	}
}
