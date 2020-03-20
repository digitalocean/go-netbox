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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableCircuitTermination writable circuit termination
// swagger:model WritableCircuitTermination
type WritableCircuitTermination struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Circuit
	// Required: true
	Circuit *int64 `json:"circuit"`

	// Connected endpoint
	//
	//
	//         Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Connection status
	// Enum: [false true]
	ConnectionStatus bool `json:"connection_status,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Port speed (Kbps)
	// Required: true
	// Maximum: 2.147483647e+09
	// Minimum: 0
	PortSpeed *int64 `json:"port_speed"`

	// Patch panel/port(s)
	// Max Length: 100
	PpInfo string `json:"pp_info,omitempty"`

	// Site
	// Required: true
	Site *int64 `json:"site"`

	// Termination
	// Required: true
	// Enum: [A Z]
	TermSide *string `json:"term_side"`

	// Upstream speed (Kbps)
	//
	// Upstream speed, if different from port speed
	// Maximum: 2.147483647e+09
	// Minimum: 0
	UpstreamSpeed *int64 `json:"upstream_speed,omitempty"`

	// Cross-connect ID
	// Max Length: 50
	XconnectID string `json:"xconnect_id,omitempty"`
}

// Validate validates this writable circuit termination
func (m *WritableCircuitTermination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCircuit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePpInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermSide(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpstreamSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXconnectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCircuitTermination) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableCircuitTermination) validateCircuit(formats strfmt.Registry) error {

	if err := validate.Required("circuit", "body", m.Circuit); err != nil {
		return err
	}

	return nil
}

var writableCircuitTerminationTypeConnectionStatusPropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCircuitTerminationTypeConnectionStatusPropEnum = append(writableCircuitTerminationTypeConnectionStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableCircuitTermination) validateConnectionStatusEnum(path, location string, value bool) error {
	if err := validate.Enum(path, location, value, writableCircuitTerminationTypeConnectionStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableCircuitTermination) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStatusEnum("connection_status", "body", m.ConnectionStatus); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validatePortSpeed(formats strfmt.Registry) error {

	if err := validate.Required("port_speed", "body", m.PortSpeed); err != nil {
		return err
	}

	if err := validate.MinimumInt("port_speed", "body", int64(*m.PortSpeed), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port_speed", "body", int64(*m.PortSpeed), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validatePpInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.PpInfo) { // not required
		return nil
	}

	if err := validate.MaxLength("pp_info", "body", string(m.PpInfo), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validateSite(formats strfmt.Registry) error {

	if err := validate.Required("site", "body", m.Site); err != nil {
		return err
	}

	return nil
}

var writableCircuitTerminationTypeTermSidePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","Z"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCircuitTerminationTypeTermSidePropEnum = append(writableCircuitTerminationTypeTermSidePropEnum, v)
	}
}

const (

	// WritableCircuitTerminationTermSideA captures enum value "A"
	WritableCircuitTerminationTermSideA string = "A"

	// WritableCircuitTerminationTermSideZ captures enum value "Z"
	WritableCircuitTerminationTermSideZ string = "Z"
)

// prop value enum
func (m *WritableCircuitTermination) validateTermSideEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableCircuitTerminationTypeTermSidePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableCircuitTermination) validateTermSide(formats strfmt.Registry) error {

	if err := validate.Required("term_side", "body", m.TermSide); err != nil {
		return err
	}

	// value enum
	if err := m.validateTermSideEnum("term_side", "body", *m.TermSide); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validateUpstreamSpeed(formats strfmt.Registry) error {

	if swag.IsZero(m.UpstreamSpeed) { // not required
		return nil
	}

	if err := validate.MinimumInt("upstream_speed", "body", int64(*m.UpstreamSpeed), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("upstream_speed", "body", int64(*m.UpstreamSpeed), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validateXconnectID(formats strfmt.Registry) error {

	if swag.IsZero(m.XconnectID) { // not required
		return nil
	}

	if err := validate.MaxLength("xconnect_id", "body", string(m.XconnectID), 50); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableCircuitTermination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableCircuitTermination) UnmarshalBinary(b []byte) error {
	var res WritableCircuitTermination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
