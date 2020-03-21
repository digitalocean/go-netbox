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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CableStatus Status
//
// swagger:model cableStatus
type CableStatus struct {

	// label
	// Required: true
	// Enum: [Connected Planned Decommissioning]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [connected planned decommissioning]
	Value *string `json:"value"`
}

// Validate validates this cable status
func (m *CableStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cableStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Connected","Planned","Decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableStatusTypeLabelPropEnum = append(cableStatusTypeLabelPropEnum, v)
	}
}

const (

	// CableStatusLabelConnected captures enum value "Connected"
	CableStatusLabelConnected string = "Connected"

	// CableStatusLabelPlanned captures enum value "Planned"
	CableStatusLabelPlanned string = "Planned"

	// CableStatusLabelDecommissioning captures enum value "Decommissioning"
	CableStatusLabelDecommissioning string = "Decommissioning"
)

// prop value enum
func (m *CableStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cableStatusTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CableStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var cableStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["connected","planned","decommissioning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cableStatusTypeValuePropEnum = append(cableStatusTypeValuePropEnum, v)
	}
}

const (

	// CableStatusValueConnected captures enum value "connected"
	CableStatusValueConnected string = "connected"

	// CableStatusValuePlanned captures enum value "planned"
	CableStatusValuePlanned string = "planned"

	// CableStatusValueDecommissioning captures enum value "decommissioning"
	CableStatusValueDecommissioning string = "decommissioning"
)

// prop value enum
func (m *CableStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, cableStatusTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CableStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CableStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CableStatus) UnmarshalBinary(b []byte) error {
	var res CableStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
