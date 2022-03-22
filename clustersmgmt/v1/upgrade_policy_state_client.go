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

package v1 // github.com/openshift-online/ocm-sdk-go/v2/clustersmgmt/v1

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/openshift-online/ocm-sdk-go/v2/errors"
	"github.com/openshift-online/ocm-sdk-go/v2/helpers"
)

// UpgradePolicyStateClient is the client of the 'upgrade_policy_state' resource.
//
// Manages a specific upgrade policy state.
type UpgradePolicyStateClient struct {
	transport http.RoundTripper
	path      string
}

// NewUpgradePolicyStateClient creates a new client for the 'upgrade_policy_state'
// resource using the given transport to send the requests and receive the
// responses.
func NewUpgradePolicyStateClient(transport http.RoundTripper, path string) *UpgradePolicyStateClient {
	return &UpgradePolicyStateClient{
		transport: transport,
		path:      path,
	}
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the upgrade policy state.
func (c *UpgradePolicyStateClient) Get() *UpgradePolicyStateGetRequest {
	return &UpgradePolicyStateGetRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Update creates a request for the 'update' method.
//
// Update the upgrade policy state.
func (c *UpgradePolicyStateClient) Update() *UpgradePolicyStateUpdateRequest {
	return &UpgradePolicyStateUpdateRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// UpgradePolicyStatePollRequest is the request for the Poll method.
type UpgradePolicyStatePollRequest struct {
	request    *UpgradePolicyStateGetRequest
	interval   time.Duration
	statuses   []int
	predicates []func(interface{}) bool
}

// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
func (r *UpgradePolicyStatePollRequest) Parameter(name string, value interface{}) *UpgradePolicyStatePollRequest {
	r.request.Parameter(name, value)
	return r
}

// Header adds a request header to all the requests that will be used to retrieve the object.
func (r *UpgradePolicyStatePollRequest) Header(name string, value interface{}) *UpgradePolicyStatePollRequest {
	r.request.Header(name, value)
	return r
}

// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
func (r *UpgradePolicyStatePollRequest) Interval(value time.Duration) *UpgradePolicyStatePollRequest {
	r.interval = value
	return r
}

// Status set the expected status of the response. Multiple values can be set calling this method
// multiple times. The response will be considered successful if the status is any of those values.
func (r *UpgradePolicyStatePollRequest) Status(value int) *UpgradePolicyStatePollRequest {
	r.statuses = append(r.statuses, value)
	return r
}

// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
// predicates can be set calling this method multiple times. The response will be considered successful
// if all the predicates are satisfied.
func (r *UpgradePolicyStatePollRequest) Predicate(value func(*UpgradePolicyStateGetResponse) bool) *UpgradePolicyStatePollRequest {
	r.predicates = append(r.predicates, func(response interface{}) bool {
		return value(response.(*UpgradePolicyStateGetResponse))
	})
	return r
}

// Start starts the polling loop. Responses will be considered successful if the status is one of
// the values specified with the Status method and if all the predicates specified with the Predicate
// method return nil.
//
// The context must have a timeout or deadline, otherwise this method will immediately return an error.
func (r *UpgradePolicyStatePollRequest) Start(ctx context.Context) (response *UpgradePolicyStatePollResponse, err error) {
	result, err := helpers.Poll(ctx, r.interval, r.statuses, r.predicates, r.task)
	if result != nil {
		response = &UpgradePolicyStatePollResponse{
			response: result.(*UpgradePolicyStateGetResponse),
		}
	}
	return
}

// task adapts the types of the request/response types so that they can be used with the generic
// polling function from the helpers package.
func (r *UpgradePolicyStatePollRequest) task(ctx context.Context) (status int, result interface{}, err error) {
	response, err := r.request.Send(ctx)
	if response != nil {
		status = response.Status()
		result = response
	}
	return
}

// UpgradePolicyStatePollResponse is the response for the Poll method.
type UpgradePolicyStatePollResponse struct {
	response *UpgradePolicyStateGetResponse
}

// Status returns the response status code.
func (r *UpgradePolicyStatePollResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.response.Status()
}

// Header returns header of the response.
func (r *UpgradePolicyStatePollResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.response.Header()
}

// Error returns the response error.
func (r *UpgradePolicyStatePollResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.response.Error()
}

// Body returns the value of the 'body' parameter.
//
//
func (r *UpgradePolicyStatePollResponse) Body() *UpgradePolicyState {
	return r.response.Body()
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *UpgradePolicyStatePollResponse) GetBody() (value *UpgradePolicyState, ok bool) {
	return r.response.GetBody()
}

// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
// of states and satisfies a set of predicates.
func (c *UpgradePolicyStateClient) Poll() *UpgradePolicyStatePollRequest {
	return &UpgradePolicyStatePollRequest{
		request: c.Get(),
	}
}

// UpgradePolicyStateGetRequest is the request for the 'get' method.
type UpgradePolicyStateGetRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *UpgradePolicyStateGetRequest) Parameter(name string, value interface{}) *UpgradePolicyStateGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *UpgradePolicyStateGetRequest) Header(name string, value interface{}) *UpgradePolicyStateGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Impersonate wraps requests on behalf of another user.
// Note: Services that do not support this feature may silently ignore this call.
func (r *UpgradePolicyStateGetRequest) Impersonate(user string) *UpgradePolicyStateGetRequest {
	helpers.AddImpersonationHeader(&r.header, user)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *UpgradePolicyStateGetRequest) Send(ctx context.Context) (result *UpgradePolicyStateGetResponse, err error) {
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
	result = &UpgradePolicyStateGetResponse{}
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
	err = readUpgradePolicyStateGetResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// UpgradePolicyStateGetResponse is the response for the 'get' method.
type UpgradePolicyStateGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *UpgradePolicyState
}

// Status returns the response status code.
func (r *UpgradePolicyStateGetResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *UpgradePolicyStateGetResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *UpgradePolicyStateGetResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *UpgradePolicyStateGetResponse) Body() *UpgradePolicyState {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *UpgradePolicyStateGetResponse) GetBody() (value *UpgradePolicyState, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// UpgradePolicyStateUpdateRequest is the request for the 'update' method.
type UpgradePolicyStateUpdateRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *UpgradePolicyState
}

// Parameter adds a query parameter.
func (r *UpgradePolicyStateUpdateRequest) Parameter(name string, value interface{}) *UpgradePolicyStateUpdateRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *UpgradePolicyStateUpdateRequest) Header(name string, value interface{}) *UpgradePolicyStateUpdateRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Impersonate wraps requests on behalf of another user.
// Note: Services that do not support this feature may silently ignore this call.
func (r *UpgradePolicyStateUpdateRequest) Impersonate(user string) *UpgradePolicyStateUpdateRequest {
	helpers.AddImpersonationHeader(&r.header, user)
	return r
}

// Body sets the value of the 'body' parameter.
//
//
func (r *UpgradePolicyStateUpdateRequest) Body(value *UpgradePolicyState) *UpgradePolicyStateUpdateRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *UpgradePolicyStateUpdateRequest) Send(ctx context.Context) (result *UpgradePolicyStateUpdateResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeUpgradePolicyStateUpdateRequest(r, buffer)
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
	result = &UpgradePolicyStateUpdateResponse{}
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
	err = readUpgradePolicyStateUpdateResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// UpgradePolicyStateUpdateResponse is the response for the 'update' method.
type UpgradePolicyStateUpdateResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *UpgradePolicyState
}

// Status returns the response status code.
func (r *UpgradePolicyStateUpdateResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *UpgradePolicyStateUpdateResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *UpgradePolicyStateUpdateResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *UpgradePolicyStateUpdateResponse) Body() *UpgradePolicyState {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *UpgradePolicyStateUpdateResponse) GetBody() (value *UpgradePolicyState, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}
