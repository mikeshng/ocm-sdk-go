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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// EnvironmentClient is the client of the 'environment' resource.
//
// Manages a specific environment.
type EnvironmentClient struct {
	transport http.RoundTripper
	path      string
}

// NewEnvironmentClient creates a new client for the 'environment'
// resource using the given transport to send the requests and receive the
// responses.
func NewEnvironmentClient(transport http.RoundTripper, path string) *EnvironmentClient {
	return &EnvironmentClient{
		transport: transport,
		path:      path,
	}
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the environment.
func (c *EnvironmentClient) Get() *EnvironmentGetRequest {
	return &EnvironmentGetRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Patch creates a request for the 'patch' method.
//
// Updates the environment.
//
// Attributes that can be updated are:
//
// - `last_upgrade_available_check`
// - `last_limited_support_check`
func (c *EnvironmentClient) Patch() *EnvironmentPatchRequest {
	return &EnvironmentPatchRequest{
		transport: c.transport,
		path:      path.Join(c.path, "patch"),
	}
}

// EnvironmentPollRequest is the request for the Poll method.
type EnvironmentPollRequest struct {
	request    *EnvironmentGetRequest
	interval   time.Duration
	statuses   []int
	predicates []func(interface{}) bool
}

// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
func (r *EnvironmentPollRequest) Parameter(name string, value interface{}) *EnvironmentPollRequest {
	r.request.Parameter(name, value)
	return r
}

// Header adds a request header to all the requests that will be used to retrieve the object.
func (r *EnvironmentPollRequest) Header(name string, value interface{}) *EnvironmentPollRequest {
	r.request.Header(name, value)
	return r
}

// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
func (r *EnvironmentPollRequest) Interval(value time.Duration) *EnvironmentPollRequest {
	r.interval = value
	return r
}

// Status set the expected status of the response. Multiple values can be set calling this method
// multiple times. The response will be considered successful if the status is any of those values.
func (r *EnvironmentPollRequest) Status(value int) *EnvironmentPollRequest {
	r.statuses = append(r.statuses, value)
	return r
}

// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
// predicates can be set calling this method multiple times. The response will be considered successful
// if all the predicates are satisfied.
func (r *EnvironmentPollRequest) Predicate(value func(*EnvironmentGetResponse) bool) *EnvironmentPollRequest {
	r.predicates = append(r.predicates, func(response interface{}) bool {
		return value(response.(*EnvironmentGetResponse))
	})
	return r
}

// StartContext starts the polling loop. Responses will be considered successful if the status is one of
// the values specified with the Status method and if all the predicates specified with the Predicate
// method return nil.
//
// The context must have a timeout or deadline, otherwise this method will immediately return an error.
func (r *EnvironmentPollRequest) StartContext(ctx context.Context) (response *EnvironmentPollResponse, err error) {
	result, err := helpers.PollContext(ctx, r.interval, r.statuses, r.predicates, r.task)
	if result != nil {
		response = &EnvironmentPollResponse{
			response: result.(*EnvironmentGetResponse),
		}
	}
	return
}

// task adapts the types of the request/response types so that they can be used with the generic
// polling function from the helpers package.
func (r *EnvironmentPollRequest) task(ctx context.Context) (status int, result interface{}, err error) {
	response, err := r.request.SendContext(ctx)
	if response != nil {
		status = response.Status()
		result = response
	}
	return
}

// EnvironmentPollResponse is the response for the Poll method.
type EnvironmentPollResponse struct {
	response *EnvironmentGetResponse
}

// Status returns the response status code.
func (r *EnvironmentPollResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.response.Status()
}

// Header returns header of the response.
func (r *EnvironmentPollResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.response.Header()
}

// Error returns the response error.
func (r *EnvironmentPollResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.response.Error()
}

// Body returns the value of the 'body' parameter.
//
//
func (r *EnvironmentPollResponse) Body() *Environment {
	return r.response.Body()
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EnvironmentPollResponse) GetBody() (value *Environment, ok bool) {
	return r.response.GetBody()
}

// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
// of states and satisfies a set of predicates.
func (c *EnvironmentClient) Poll() *EnvironmentPollRequest {
	return &EnvironmentPollRequest{
		request: c.Get(),
	}
}

// EnvironmentGetRequest is the request for the 'get' method.
type EnvironmentGetRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *EnvironmentGetRequest) Parameter(name string, value interface{}) *EnvironmentGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *EnvironmentGetRequest) Header(name string, value interface{}) *EnvironmentGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *EnvironmentGetRequest) Send() (result *EnvironmentGetResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *EnvironmentGetRequest) SendContext(ctx context.Context) (result *EnvironmentGetResponse, err error) {
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
	result = &EnvironmentGetResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readEnvironmentGetResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// EnvironmentGetResponse is the response for the 'get' method.
type EnvironmentGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Environment
}

// Status returns the response status code.
func (r *EnvironmentGetResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *EnvironmentGetResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *EnvironmentGetResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *EnvironmentGetResponse) Body() *Environment {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EnvironmentGetResponse) GetBody() (value *Environment, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// EnvironmentPatchRequest is the request for the 'patch' method.
type EnvironmentPatchRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *Environment
}

// Parameter adds a query parameter.
func (r *EnvironmentPatchRequest) Parameter(name string, value interface{}) *EnvironmentPatchRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *EnvironmentPatchRequest) Header(name string, value interface{}) *EnvironmentPatchRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
//
func (r *EnvironmentPatchRequest) Body(value *Environment) *EnvironmentPatchRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *EnvironmentPatchRequest) Send() (result *EnvironmentPatchResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *EnvironmentPatchRequest) SendContext(ctx context.Context) (result *EnvironmentPatchResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeEnvironmentPatchRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "POST",
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
	result = &EnvironmentPatchResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readEnvironmentPatchResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// EnvironmentPatchResponse is the response for the 'patch' method.
type EnvironmentPatchResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Environment
}

// Status returns the response status code.
func (r *EnvironmentPatchResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *EnvironmentPatchResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *EnvironmentPatchResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *EnvironmentPatchResponse) Body() *Environment {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *EnvironmentPatchResponse) GetBody() (value *Environment, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}