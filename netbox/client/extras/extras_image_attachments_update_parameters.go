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

// NewExtrasImageAttachmentsUpdateParams creates a new ExtrasImageAttachmentsUpdateParams object
// with the default values initialized.
func NewExtrasImageAttachmentsUpdateParams() *ExtrasImageAttachmentsUpdateParams {
	var ()
	return &ExtrasImageAttachmentsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsUpdateParamsWithTimeout creates a new ExtrasImageAttachmentsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasImageAttachmentsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsUpdateParams {
	var ()
	return &ExtrasImageAttachmentsUpdateParams{

		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsUpdateParamsWithContext creates a new ExtrasImageAttachmentsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasImageAttachmentsUpdateParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsUpdateParams {
	var ()
	return &ExtrasImageAttachmentsUpdateParams{

		Context: ctx,
	}
}

// NewExtrasImageAttachmentsUpdateParamsWithHTTPClient creates a new ExtrasImageAttachmentsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasImageAttachmentsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsUpdateParams {
	var ()
	return &ExtrasImageAttachmentsUpdateParams{
		HTTPClient: client,
	}
}

/*ExtrasImageAttachmentsUpdateParams contains all the parameters to send to the API endpoint
for the extras image attachments update operation typically these are written to a http.Request
*/
type ExtrasImageAttachmentsUpdateParams struct {

	/*Data*/
	Data *models.ImageAttachment
	/*ID
	  A unique integer value identifying this image attachment.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithData(data *models.ImageAttachment) *ExtrasImageAttachmentsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetData(data *models.ImageAttachment) {
	o.Data = data
}

// WithID adds the id to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithID(id int64) *ExtrasImageAttachmentsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
