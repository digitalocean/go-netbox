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

// DcimPowerPortsCreateReader is a Reader for the DcimPowerPortsCreate structure.
type DcimPowerPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsCreateCreated creates a DcimPowerPortsCreateCreated with default headers values
func NewDcimPowerPortsCreateCreated() *DcimPowerPortsCreateCreated {
	return &DcimPowerPortsCreateCreated{}
}

/*DcimPowerPortsCreateCreated handles this case with default header values.

DcimPowerPortsCreateCreated dcim power ports create created
*/
type DcimPowerPortsCreateCreated struct {
	Payload *models.PowerPort
}

func (o *DcimPowerPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-ports/][%d] dcimPowerPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerPortsCreateCreated) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortsCreateDefault creates a DcimPowerPortsCreateDefault with default headers values
func NewDcimPowerPortsCreateDefault(code int) *DcimPowerPortsCreateDefault {
	return &DcimPowerPortsCreateDefault{
		_statusCode: code,
	}
}

/*DcimPowerPortsCreateDefault handles this case with default header values.

DcimPowerPortsCreateDefault dcim power ports create default
*/
type DcimPowerPortsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power ports create default response
func (o *DcimPowerPortsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/power-ports/][%d] dcim_power-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
