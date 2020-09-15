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

// DcimConsoleServerPortsPartialUpdateReader is a Reader for the DcimConsoleServerPortsPartialUpdate structure.
type DcimConsoleServerPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsPartialUpdateOK creates a DcimConsoleServerPortsPartialUpdateOK with default headers values
func NewDcimConsoleServerPortsPartialUpdateOK() *DcimConsoleServerPortsPartialUpdateOK {
	return &DcimConsoleServerPortsPartialUpdateOK{}
}

/*DcimConsoleServerPortsPartialUpdateOK handles this case with default header values.

DcimConsoleServerPortsPartialUpdateOK dcim console server ports partial update o k
*/
type DcimConsoleServerPortsPartialUpdateOK struct {
	Payload *models.ConsoleServerPort
}

func (o *DcimConsoleServerPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsPartialUpdateOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortsPartialUpdateDefault creates a DcimConsoleServerPortsPartialUpdateDefault with default headers values
func NewDcimConsoleServerPortsPartialUpdateDefault(code int) *DcimConsoleServerPortsPartialUpdateDefault {
	return &DcimConsoleServerPortsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimConsoleServerPortsPartialUpdateDefault handles this case with default header values.

DcimConsoleServerPortsPartialUpdateDefault dcim console server ports partial update default
*/
type DcimConsoleServerPortsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server ports partial update default response
func (o *DcimConsoleServerPortsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/console-server-ports/{id}/][%d] dcim_console-server-ports_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
