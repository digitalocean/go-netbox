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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/smutel/go-netbox/netbox/models"
)

// NewDcimRackGroupsCreateParams creates a new DcimRackGroupsCreateParams object
// with the default values initialized.
func NewDcimRackGroupsCreateParams() *DcimRackGroupsCreateParams {
	var ()
	return &DcimRackGroupsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackGroupsCreateParamsWithTimeout creates a new DcimRackGroupsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimRackGroupsCreateParamsWithTimeout(timeout time.Duration) *DcimRackGroupsCreateParams {
	var ()
	return &DcimRackGroupsCreateParams{

		timeout: timeout,
	}
}

// NewDcimRackGroupsCreateParamsWithContext creates a new DcimRackGroupsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimRackGroupsCreateParamsWithContext(ctx context.Context) *DcimRackGroupsCreateParams {
	var ()
	return &DcimRackGroupsCreateParams{

		Context: ctx,
	}
}

// NewDcimRackGroupsCreateParamsWithHTTPClient creates a new DcimRackGroupsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimRackGroupsCreateParamsWithHTTPClient(client *http.Client) *DcimRackGroupsCreateParams {
	var ()
	return &DcimRackGroupsCreateParams{
		HTTPClient: client,
	}
}

/*DcimRackGroupsCreateParams contains all the parameters to send to the API endpoint
for the dcim rack groups create operation typically these are written to a http.Request
*/
type DcimRackGroupsCreateParams struct {

	/*Data*/
	Data *models.WritableRackGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim rack groups create params
func (o *DcimRackGroupsCreateParams) WithTimeout(timeout time.Duration) *DcimRackGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack groups create params
func (o *DcimRackGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack groups create params
func (o *DcimRackGroupsCreateParams) WithContext(ctx context.Context) *DcimRackGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack groups create params
func (o *DcimRackGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack groups create params
func (o *DcimRackGroupsCreateParams) WithHTTPClient(client *http.Client) *DcimRackGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack groups create params
func (o *DcimRackGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack groups create params
func (o *DcimRackGroupsCreateParams) WithData(data *models.WritableRackGroup) *DcimRackGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack groups create params
func (o *DcimRackGroupsCreateParams) SetData(data *models.WritableRackGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
