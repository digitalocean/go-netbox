// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewExtrasObjectChangesListParams creates a new ExtrasObjectChangesListParams object
// with the default values initialized.
func NewExtrasObjectChangesListParams() *ExtrasObjectChangesListParams {
	var ()
	return &ExtrasObjectChangesListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasObjectChangesListParamsWithTimeout creates a new ExtrasObjectChangesListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasObjectChangesListParamsWithTimeout(timeout time.Duration) *ExtrasObjectChangesListParams {
	var ()
	return &ExtrasObjectChangesListParams{

		timeout: timeout,
	}
}

// NewExtrasObjectChangesListParamsWithContext creates a new ExtrasObjectChangesListParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasObjectChangesListParamsWithContext(ctx context.Context) *ExtrasObjectChangesListParams {
	var ()
	return &ExtrasObjectChangesListParams{

		Context: ctx,
	}
}

// NewExtrasObjectChangesListParamsWithHTTPClient creates a new ExtrasObjectChangesListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasObjectChangesListParamsWithHTTPClient(client *http.Client) *ExtrasObjectChangesListParams {
	var ()
	return &ExtrasObjectChangesListParams{
		HTTPClient: client,
	}
}

/*ExtrasObjectChangesListParams contains all the parameters to send to the API endpoint
for the extras object changes list operation typically these are written to a http.Request
*/
type ExtrasObjectChangesListParams struct {

	/*Action*/
	Action *string
	/*ChangedObjectID*/
	ChangedObjectID *float64
	/*ChangedObjectType*/
	ChangedObjectType *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*ObjectRepr*/
	ObjectRepr *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*RequestID*/
	RequestID *string
	/*Time*/
	Time *string
	/*User*/
	User *string
	/*UserName*/
	UserName *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithTimeout(timeout time.Duration) *ExtrasObjectChangesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithContext(ctx context.Context) *ExtrasObjectChangesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithHTTPClient(client *http.Client) *ExtrasObjectChangesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithAction(action *string) *ExtrasObjectChangesListParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetAction(action *string) {
	o.Action = action
}

// WithChangedObjectID adds the changedObjectID to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectID(changedObjectID *float64) *ExtrasObjectChangesListParams {
	o.SetChangedObjectID(changedObjectID)
	return o
}

// SetChangedObjectID adds the changedObjectId to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectID(changedObjectID *float64) {
	o.ChangedObjectID = changedObjectID
}

// WithChangedObjectType adds the changedObjectType to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithChangedObjectType(changedObjectType *string) *ExtrasObjectChangesListParams {
	o.SetChangedObjectType(changedObjectType)
	return o
}

// SetChangedObjectType adds the changedObjectType to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetChangedObjectType(changedObjectType *string) {
	o.ChangedObjectType = changedObjectType
}

// WithLimit adds the limit to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithLimit(limit *int64) *ExtrasObjectChangesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithObjectRepr adds the objectRepr to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithObjectRepr(objectRepr *string) *ExtrasObjectChangesListParams {
	o.SetObjectRepr(objectRepr)
	return o
}

// SetObjectRepr adds the objectRepr to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetObjectRepr(objectRepr *string) {
	o.ObjectRepr = objectRepr
}

// WithOffset adds the offset to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithOffset(offset *int64) *ExtrasObjectChangesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithQ(q *string) *ExtrasObjectChangesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRequestID adds the requestID to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithRequestID(requestID *string) *ExtrasObjectChangesListParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetRequestID(requestID *string) {
	o.RequestID = requestID
}

// WithTime adds the time to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithTime(time *string) *ExtrasObjectChangesListParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetTime(time *string) {
	o.Time = time
}

// WithUser adds the user to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUser(user *string) *ExtrasObjectChangesListParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUser(user *string) {
	o.User = user
}

// WithUserName adds the userName to the extras object changes list params
func (o *ExtrasObjectChangesListParams) WithUserName(userName *string) *ExtrasObjectChangesListParams {
	o.SetUserName(userName)
	return o
}

// SetUserName adds the userName to the extras object changes list params
func (o *ExtrasObjectChangesListParams) SetUserName(userName *string) {
	o.UserName = userName
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasObjectChangesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Action != nil {

		// query param action
		var qrAction string
		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {
			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}

	}

	if o.ChangedObjectID != nil {

		// query param changed_object_id
		var qrChangedObjectID float64
		if o.ChangedObjectID != nil {
			qrChangedObjectID = *o.ChangedObjectID
		}
		qChangedObjectID := swag.FormatFloat64(qrChangedObjectID)
		if qChangedObjectID != "" {
			if err := r.SetQueryParam("changed_object_id", qChangedObjectID); err != nil {
				return err
			}
		}

	}

	if o.ChangedObjectType != nil {

		// query param changed_object_type
		var qrChangedObjectType string
		if o.ChangedObjectType != nil {
			qrChangedObjectType = *o.ChangedObjectType
		}
		qChangedObjectType := qrChangedObjectType
		if qChangedObjectType != "" {
			if err := r.SetQueryParam("changed_object_type", qChangedObjectType); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.ObjectRepr != nil {

		// query param object_repr
		var qrObjectRepr string
		if o.ObjectRepr != nil {
			qrObjectRepr = *o.ObjectRepr
		}
		qObjectRepr := qrObjectRepr
		if qObjectRepr != "" {
			if err := r.SetQueryParam("object_repr", qObjectRepr); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.RequestID != nil {

		// query param request_id
		var qrRequestID string
		if o.RequestID != nil {
			qrRequestID = *o.RequestID
		}
		qRequestID := qrRequestID
		if qRequestID != "" {
			if err := r.SetQueryParam("request_id", qRequestID); err != nil {
				return err
			}
		}

	}

	if o.Time != nil {

		// query param time
		var qrTime string
		if o.Time != nil {
			qrTime = *o.Time
		}
		qTime := qrTime
		if qTime != "" {
			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}

	}

	if o.User != nil {

		// query param user
		var qrUser string
		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {
			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}

	}

	if o.UserName != nil {

		// query param user_name
		var qrUserName string
		if o.UserName != nil {
			qrUserName = *o.UserName
		}
		qUserName := qrUserName
		if qUserName != "" {
			if err := r.SetQueryParam("user_name", qUserName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
