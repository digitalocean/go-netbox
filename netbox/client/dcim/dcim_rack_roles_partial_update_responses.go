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

// DcimRackRolesPartialUpdateReader is a Reader for the DcimRackRolesPartialUpdate structure.
type DcimRackRolesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackRolesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackRolesPartialUpdateOK creates a DcimRackRolesPartialUpdateOK with default headers values
func NewDcimRackRolesPartialUpdateOK() *DcimRackRolesPartialUpdateOK {
	return &DcimRackRolesPartialUpdateOK{}
}

/*DcimRackRolesPartialUpdateOK handles this case with default header values.

DcimRackRolesPartialUpdateOK dcim rack roles partial update o k
*/
type DcimRackRolesPartialUpdateOK struct {
	Payload *models.RackRole
}

func (o *DcimRackRolesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-roles/{id}/][%d] dcimRackRolesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesPartialUpdateOK) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesPartialUpdateDefault creates a DcimRackRolesPartialUpdateDefault with default headers values
func NewDcimRackRolesPartialUpdateDefault(code int) *DcimRackRolesPartialUpdateDefault {
	return &DcimRackRolesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimRackRolesPartialUpdateDefault handles this case with default header values.

DcimRackRolesPartialUpdateDefault dcim rack roles partial update default
*/
type DcimRackRolesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack roles partial update default response
func (o *DcimRackRolesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackRolesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rack-roles/{id}/][%d] dcim_rack-roles_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
