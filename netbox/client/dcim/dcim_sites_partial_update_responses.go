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

// DcimSitesPartialUpdateReader is a Reader for the DcimSitesPartialUpdate structure.
type DcimSitesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSitesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesPartialUpdateOK creates a DcimSitesPartialUpdateOK with default headers values
func NewDcimSitesPartialUpdateOK() *DcimSitesPartialUpdateOK {
	return &DcimSitesPartialUpdateOK{}
}

/*DcimSitesPartialUpdateOK handles this case with default header values.

DcimSitesPartialUpdateOK dcim sites partial update o k
*/
type DcimSitesPartialUpdateOK struct {
	Payload *models.Site
}

func (o *DcimSitesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/{id}/][%d] dcimSitesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimSitesPartialUpdateOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesPartialUpdateDefault creates a DcimSitesPartialUpdateDefault with default headers values
func NewDcimSitesPartialUpdateDefault(code int) *DcimSitesPartialUpdateDefault {
	return &DcimSitesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimSitesPartialUpdateDefault handles this case with default header values.

DcimSitesPartialUpdateDefault dcim sites partial update default
*/
type DcimSitesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites partial update default response
func (o *DcimSitesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSitesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/{id}/][%d] dcim_sites_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
