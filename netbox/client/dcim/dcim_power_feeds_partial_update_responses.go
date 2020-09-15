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

// DcimPowerFeedsPartialUpdateReader is a Reader for the DcimPowerFeedsPartialUpdate structure.
type DcimPowerFeedsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsPartialUpdateOK creates a DcimPowerFeedsPartialUpdateOK with default headers values
func NewDcimPowerFeedsPartialUpdateOK() *DcimPowerFeedsPartialUpdateOK {
	return &DcimPowerFeedsPartialUpdateOK{}
}

/*DcimPowerFeedsPartialUpdateOK handles this case with default header values.

DcimPowerFeedsPartialUpdateOK dcim power feeds partial update o k
*/
type DcimPowerFeedsPartialUpdateOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-feeds/{id}/][%d] dcimPowerFeedsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsPartialUpdateOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsPartialUpdateDefault creates a DcimPowerFeedsPartialUpdateDefault with default headers values
func NewDcimPowerFeedsPartialUpdateDefault(code int) *DcimPowerFeedsPartialUpdateDefault {
	return &DcimPowerFeedsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimPowerFeedsPartialUpdateDefault handles this case with default header values.

DcimPowerFeedsPartialUpdateDefault dcim power feeds partial update default
*/
type DcimPowerFeedsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds partial update default response
func (o *DcimPowerFeedsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerFeedsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-feeds/{id}/][%d] dcim_power-feeds_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
