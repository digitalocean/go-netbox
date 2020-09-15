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

// DcimInventoryItemsCreateReader is a Reader for the DcimInventoryItemsCreate structure.
type DcimInventoryItemsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInventoryItemsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsCreateCreated creates a DcimInventoryItemsCreateCreated with default headers values
func NewDcimInventoryItemsCreateCreated() *DcimInventoryItemsCreateCreated {
	return &DcimInventoryItemsCreateCreated{}
}

/*DcimInventoryItemsCreateCreated handles this case with default header values.

DcimInventoryItemsCreateCreated dcim inventory items create created
*/
type DcimInventoryItemsCreateCreated struct {
	Payload *models.InventoryItem
}

func (o *DcimInventoryItemsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcimInventoryItemsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInventoryItemsCreateCreated) GetPayload() *models.InventoryItem {
	return o.Payload
}

func (o *DcimInventoryItemsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemsCreateDefault creates a DcimInventoryItemsCreateDefault with default headers values
func NewDcimInventoryItemsCreateDefault(code int) *DcimInventoryItemsCreateDefault {
	return &DcimInventoryItemsCreateDefault{
		_statusCode: code,
	}
}

/*DcimInventoryItemsCreateDefault handles this case with default header values.

DcimInventoryItemsCreateDefault dcim inventory items create default
*/
type DcimInventoryItemsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory items create default response
func (o *DcimInventoryItemsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/inventory-items/][%d] dcim_inventory-items_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
