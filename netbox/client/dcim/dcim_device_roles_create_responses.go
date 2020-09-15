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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Boulet-/go-netbox/netbox/models"
)

// DcimDeviceRolesCreateReader is a Reader for the DcimDeviceRolesCreate structure.
type DcimDeviceRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimDeviceRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimDeviceRolesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesCreateCreated creates a DcimDeviceRolesCreateCreated with default headers values
func NewDcimDeviceRolesCreateCreated() *DcimDeviceRolesCreateCreated {
	return &DcimDeviceRolesCreateCreated{}
}

/*DcimDeviceRolesCreateCreated handles this case with default header values.

DcimDeviceRolesCreateCreated dcim device roles create created
*/
type DcimDeviceRolesCreateCreated struct {
	Payload *models.DeviceRole
}

func (o *DcimDeviceRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/device-roles/][%d] dcimDeviceRolesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimDeviceRolesCreateCreated) GetPayload() *models.DeviceRole {
	return o.Payload
}

func (o *DcimDeviceRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceRolesCreateDefault creates a DcimDeviceRolesCreateDefault with default headers values
func NewDcimDeviceRolesCreateDefault(code int) *DcimDeviceRolesCreateDefault {
	return &DcimDeviceRolesCreateDefault{
		_statusCode: code,
	}
}

/*DcimDeviceRolesCreateDefault handles this case with default header values.

DcimDeviceRolesCreateDefault dcim device roles create default
*/
type DcimDeviceRolesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device roles create default response
func (o *DcimDeviceRolesCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimDeviceRolesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/device-roles/][%d] dcim_device-roles_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
