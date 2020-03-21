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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/models"
)

// CircuitsCircuitTypesListReader is a Reader for the CircuitsCircuitTypesList structure.
type CircuitsCircuitTypesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCircuitsCircuitTypesListOK creates a CircuitsCircuitTypesListOK with default headers values
func NewCircuitsCircuitTypesListOK() *CircuitsCircuitTypesListOK {
	return &CircuitsCircuitTypesListOK{}
}

/*CircuitsCircuitTypesListOK handles this case with default header values.

CircuitsCircuitTypesListOK circuits circuit types list o k
*/
type CircuitsCircuitTypesListOK struct {
	Payload *models.CircuitsCircuitTypesListOKBody
}

func (o *CircuitsCircuitTypesListOK) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-types/][%d] circuitsCircuitTypesListOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesListOK) GetPayload() *models.CircuitsCircuitTypesListOKBody {
	return o.Payload
}

func (o *CircuitsCircuitTypesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitsCircuitTypesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
