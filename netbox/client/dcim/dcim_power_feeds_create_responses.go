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

// DcimPowerFeedsCreateReader is a Reader for the DcimPowerFeedsCreate structure.
type DcimPowerFeedsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerFeedsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsCreateCreated creates a DcimPowerFeedsCreateCreated with default headers values
func NewDcimPowerFeedsCreateCreated() *DcimPowerFeedsCreateCreated {
	return &DcimPowerFeedsCreateCreated{}
}

/*DcimPowerFeedsCreateCreated handles this case with default header values.

DcimPowerFeedsCreateCreated dcim power feeds create created
*/
type DcimPowerFeedsCreateCreated struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-feeds/][%d] dcimPowerFeedsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerFeedsCreateCreated) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsCreateDefault creates a DcimPowerFeedsCreateDefault with default headers values
func NewDcimPowerFeedsCreateDefault(code int) *DcimPowerFeedsCreateDefault {
	return &DcimPowerFeedsCreateDefault{
		_statusCode: code,
	}
}

/*DcimPowerFeedsCreateDefault handles this case with default header values.

DcimPowerFeedsCreateDefault dcim power feeds create default
*/
type DcimPowerFeedsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds create default response
func (o *DcimPowerFeedsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerFeedsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/power-feeds/][%d] dcim_power-feeds_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
