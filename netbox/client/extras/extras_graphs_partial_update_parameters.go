// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/smutel/go-netbox/netbox/models"
)

// NewExtrasGraphsPartialUpdateParams creates a new ExtrasGraphsPartialUpdateParams object
// with the default values initialized.
func NewExtrasGraphsPartialUpdateParams() *ExtrasGraphsPartialUpdateParams {
	var ()
	return &ExtrasGraphsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGraphsPartialUpdateParamsWithTimeout creates a new ExtrasGraphsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasGraphsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasGraphsPartialUpdateParams {
	var ()
	return &ExtrasGraphsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasGraphsPartialUpdateParamsWithContext creates a new ExtrasGraphsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasGraphsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasGraphsPartialUpdateParams {
	var ()
	return &ExtrasGraphsPartialUpdateParams{

		Context: ctx,
	}
}

// NewExtrasGraphsPartialUpdateParamsWithHTTPClient creates a new ExtrasGraphsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasGraphsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasGraphsPartialUpdateParams {
	var ()
	return &ExtrasGraphsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasGraphsPartialUpdateParams contains all the parameters to send to the API endpoint
for the extras graphs partial update operation typically these are written to a http.Request
*/
type ExtrasGraphsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableGraph
	/*ID
	  A unique integer value identifying this graph.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasGraphsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasGraphsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasGraphsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) WithData(data *models.WritableGraph) *ExtrasGraphsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) SetData(data *models.WritableGraph) {
	o.Data = data
}

// WithID adds the id to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) WithID(id int64) *ExtrasGraphsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras graphs partial update params
func (o *ExtrasGraphsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGraphsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
