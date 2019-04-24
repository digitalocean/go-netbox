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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimPowerPortsListParams creates a new DcimPowerPortsListParams object
// with the default values initialized.
func NewDcimPowerPortsListParams() *DcimPowerPortsListParams {
	var ()
	return &DcimPowerPortsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsListParamsWithTimeout creates a new DcimPowerPortsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPortsListParamsWithTimeout(timeout time.Duration) *DcimPowerPortsListParams {
	var ()
	return &DcimPowerPortsListParams{

		timeout: timeout,
	}
}

// NewDcimPowerPortsListParamsWithContext creates a new DcimPowerPortsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPortsListParamsWithContext(ctx context.Context) *DcimPowerPortsListParams {
	var ()
	return &DcimPowerPortsListParams{

		Context: ctx,
	}
}

// NewDcimPowerPortsListParamsWithHTTPClient creates a new DcimPowerPortsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPortsListParamsWithHTTPClient(client *http.Client) *DcimPowerPortsListParams {
	var ()
	return &DcimPowerPortsListParams{
		HTTPClient: client,
	}
}

/*DcimPowerPortsListParams contains all the parameters to send to the API endpoint
for the dcim power ports list operation typically these are written to a http.Request
*/
type DcimPowerPortsListParams struct {

	/*Cabled*/
	Cabled *string
	/*ConnectionStatus*/
	ConnectionStatus *string
	/*Device*/
	Device *string
	/*DeviceID*/
	DeviceID *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Tag*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithTimeout(timeout time.Duration) *DcimPowerPortsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithContext(ctx context.Context) *DcimPowerPortsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithHTTPClient(client *http.Client) *DcimPowerPortsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithCabled(cabled *string) *DcimPowerPortsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithConnectionStatus adds the connectionStatus to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithConnectionStatus(connectionStatus *string) *DcimPowerPortsListParams {
	o.SetConnectionStatus(connectionStatus)
	return o
}

// SetConnectionStatus adds the connectionStatus to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetConnectionStatus(connectionStatus *string) {
	o.ConnectionStatus = connectionStatus
}

// WithDevice adds the device to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDevice(device *string) *DcimPowerPortsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDeviceID(deviceID *string) *DcimPowerPortsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithLimit adds the limit to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithLimit(limit *int64) *DcimPowerPortsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithName(name *string) *DcimPowerPortsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithOffset(offset *int64) *DcimPowerPortsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithQ(q *string) *DcimPowerPortsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithTag(tag *string) *DcimPowerPortsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cabled != nil {

		// query param cabled
		var qrCabled string
		if o.Cabled != nil {
			qrCabled = *o.Cabled
		}
		qCabled := qrCabled
		if qCabled != "" {
			if err := r.SetQueryParam("cabled", qCabled); err != nil {
				return err
			}
		}

	}

	if o.ConnectionStatus != nil {

		// query param connection_status
		var qrConnectionStatus string
		if o.ConnectionStatus != nil {
			qrConnectionStatus = *o.ConnectionStatus
		}
		qConnectionStatus := qrConnectionStatus
		if qConnectionStatus != "" {
			if err := r.SetQueryParam("connection_status", qConnectionStatus); err != nil {
				return err
			}
		}

	}

	if o.Device != nil {

		// query param device
		var qrDevice string
		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {
			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}

	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string
		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {
			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
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

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
