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

package secrets

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
)

// NewSecretsGetSessionKeyCreateParams creates a new SecretsGetSessionKeyCreateParams object
// with the default values initialized.
func NewSecretsGetSessionKeyCreateParams() *SecretsGetSessionKeyCreateParams {

	return &SecretsGetSessionKeyCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsGetSessionKeyCreateParamsWithTimeout creates a new SecretsGetSessionKeyCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsGetSessionKeyCreateParamsWithTimeout(timeout time.Duration) *SecretsGetSessionKeyCreateParams {

	return &SecretsGetSessionKeyCreateParams{

		timeout: timeout,
	}
}

// NewSecretsGetSessionKeyCreateParamsWithContext creates a new SecretsGetSessionKeyCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsGetSessionKeyCreateParamsWithContext(ctx context.Context) *SecretsGetSessionKeyCreateParams {

	return &SecretsGetSessionKeyCreateParams{

		Context: ctx,
	}
}

// NewSecretsGetSessionKeyCreateParamsWithHTTPClient creates a new SecretsGetSessionKeyCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsGetSessionKeyCreateParamsWithHTTPClient(client *http.Client) *SecretsGetSessionKeyCreateParams {

	return &SecretsGetSessionKeyCreateParams{
		HTTPClient: client,
	}
}

/*SecretsGetSessionKeyCreateParams contains all the parameters to send to the API endpoint
for the secrets get session key create operation typically these are written to a http.Request
*/
type SecretsGetSessionKeyCreateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets get session key create params
func (o *SecretsGetSessionKeyCreateParams) WithTimeout(timeout time.Duration) *SecretsGetSessionKeyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets get session key create params
func (o *SecretsGetSessionKeyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets get session key create params
func (o *SecretsGetSessionKeyCreateParams) WithContext(ctx context.Context) *SecretsGetSessionKeyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets get session key create params
func (o *SecretsGetSessionKeyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets get session key create params
func (o *SecretsGetSessionKeyCreateParams) WithHTTPClient(client *http.Client) *SecretsGetSessionKeyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets get session key create params
func (o *SecretsGetSessionKeyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsGetSessionKeyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
