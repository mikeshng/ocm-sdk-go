/*
Copyright (c) 2020 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// GenericLabelClient is the client of the 'generic_label' resource.
//
// Manages a specific account label.
type GenericLabelClient struct {
	transport http.RoundTripper
	path      string
}

// NewGenericLabelClient creates a new client for the 'generic_label'
// resource using the given transport to send the requests and receive the
// responses.
func NewGenericLabelClient(transport http.RoundTripper, path string) *GenericLabelClient {
	return &GenericLabelClient{
		transport: transport,
		path:      path,
	}
}

// Delete creates a request for the 'delete' method.
//
// Deletes the account label.
func (c *GenericLabelClient) Delete() *GenericLabelDeleteRequest {
	return &GenericLabelDeleteRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the label.
func (c *GenericLabelClient) Get() *GenericLabelGetRequest {
	return &GenericLabelGetRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Update creates a request for the 'update' method.
//
// Updates the account label.
func (c *GenericLabelClient) Update() *GenericLabelUpdateRequest {
	return &GenericLabelUpdateRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// GenericLabelPollRequest is the request for the Poll method.
type GenericLabelPollRequest struct {
	request    *GenericLabelGetRequest
	interval   time.Duration
	statuses   []int
	predicates []func(interface{}) bool
}

// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
func (r *GenericLabelPollRequest) Parameter(name string, value interface{}) *GenericLabelPollRequest {
	r.request.Parameter(name, value)
	return r
}

// Header adds a request header to all the requests that will be used to retrieve the object.
func (r *GenericLabelPollRequest) Header(name string, value interface{}) *GenericLabelPollRequest {
	r.request.Header(name, value)
	return r
}

// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
func (r *GenericLabelPollRequest) Interval(value time.Duration) *GenericLabelPollRequest {
	r.interval = value
	return r
}

// Status set the expected status of the response. Multiple values can be set calling this method
// multiple times. The response will be considered successful if the status is any of those values.
func (r *GenericLabelPollRequest) Status(value int) *GenericLabelPollRequest {
	r.statuses = append(r.statuses, value)
	return r
}

// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
// predicates can be set calling this method multiple times. The response will be considered successful
// if all the predicates are satisfied.
func (r *GenericLabelPollRequest) Predicate(value func(*GenericLabelGetResponse) bool) *GenericLabelPollRequest {
	r.predicates = append(r.predicates, func(response interface{}) bool {
		return value(response.(*GenericLabelGetResponse))
	})
	return r
}

// StartContext starts the polling loop. Responses will be considered successful if the status is one of
// the values specified with the Status method and if all the predicates specified with the Predicate
// method return nil.
//
// The context must have a timeout or deadline, otherwise this method will immediately return an error.
func (r *GenericLabelPollRequest) StartContext(ctx context.Context) (response *GenericLabelPollResponse, err error) {
	result, err := helpers.PollContext(ctx, r.interval, r.statuses, r.predicates, r.task)
	if result != nil {
		response = &GenericLabelPollResponse{
			response: result.(*GenericLabelGetResponse),
		}
	}
	return
}

// task adapts the types of the request/response types so that they can be used with the generic
// polling function from the helpers package.
func (r *GenericLabelPollRequest) task(ctx context.Context) (status int, result interface{}, err error) {
	response, err := r.request.SendContext(ctx)
	if response != nil {
		status = response.Status()
		result = response
	}
	return
}

// GenericLabelPollResponse is the response for the Poll method.
type GenericLabelPollResponse struct {
	response *GenericLabelGetResponse
}

// Status returns the response status code.
func (r *GenericLabelPollResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.response.Status()
}

// Header returns header of the response.
func (r *GenericLabelPollResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.response.Header()
}

// Error returns the response error.
func (r *GenericLabelPollResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.response.Error()
}

// Body returns the value of the 'body' parameter.
//
//
func (r *GenericLabelPollResponse) Body() *Label {
	return r.response.Body()
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *GenericLabelPollResponse) GetBody() (value *Label, ok bool) {
	return r.response.GetBody()
}

// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
// of states and satisfies a set of predicates.
func (c *GenericLabelClient) Poll() *GenericLabelPollRequest {
	return &GenericLabelPollRequest{
		request: c.Get(),
	}
}

// GenericLabelDeleteRequest is the request for the 'delete' method.
type GenericLabelDeleteRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *GenericLabelDeleteRequest) Parameter(name string, value interface{}) *GenericLabelDeleteRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *GenericLabelDeleteRequest) Header(name string, value interface{}) *GenericLabelDeleteRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *GenericLabelDeleteRequest) Send() (result *GenericLabelDeleteResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *GenericLabelDeleteRequest) SendContext(ctx context.Context) (result *GenericLabelDeleteResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "DELETE",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &GenericLabelDeleteResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(response.Body, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	return
}

// GenericLabelDeleteResponse is the response for the 'delete' method.
type GenericLabelDeleteResponse struct {
	status int
	header http.Header
	err    *errors.Error
}

// Status returns the response status code.
func (r *GenericLabelDeleteResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *GenericLabelDeleteResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *GenericLabelDeleteResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// GenericLabelGetRequest is the request for the 'get' method.
type GenericLabelGetRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *GenericLabelGetRequest) Parameter(name string, value interface{}) *GenericLabelGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *GenericLabelGetRequest) Header(name string, value interface{}) *GenericLabelGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *GenericLabelGetRequest) Send() (result *GenericLabelGetResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *GenericLabelGetRequest) SendContext(ctx context.Context) (result *GenericLabelGetResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "GET",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &GenericLabelGetResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(response.Body, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readGenericLabelGetResponse(result, response.Body)
	if err != nil {
		return
	}
	return
}

// GenericLabelGetResponse is the response for the 'get' method.
type GenericLabelGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Label
}

// Status returns the response status code.
func (r *GenericLabelGetResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *GenericLabelGetResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *GenericLabelGetResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *GenericLabelGetResponse) Body() *Label {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *GenericLabelGetResponse) GetBody() (value *Label, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// GenericLabelUpdateRequest is the request for the 'update' method.
type GenericLabelUpdateRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *Label
}

// Parameter adds a query parameter.
func (r *GenericLabelUpdateRequest) Parameter(name string, value interface{}) *GenericLabelUpdateRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *GenericLabelUpdateRequest) Header(name string, value interface{}) *GenericLabelUpdateRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
//
func (r *GenericLabelUpdateRequest) Body(value *Label) *GenericLabelUpdateRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *GenericLabelUpdateRequest) Send() (result *GenericLabelUpdateResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *GenericLabelUpdateRequest) SendContext(ctx context.Context) (result *GenericLabelUpdateResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeGenericLabelUpdateRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "PATCH",
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &GenericLabelUpdateResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(response.Body, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readGenericLabelUpdateResponse(result, response.Body)
	if err != nil {
		return
	}
	return
}

// GenericLabelUpdateResponse is the response for the 'update' method.
type GenericLabelUpdateResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Label
}

// Status returns the response status code.
func (r *GenericLabelUpdateResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *GenericLabelUpdateResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *GenericLabelUpdateResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *GenericLabelUpdateResponse) Body() *Label {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *GenericLabelUpdateResponse) GetBody() (value *Label, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}
