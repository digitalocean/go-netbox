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

// ConsolePortTemplate console port template
//
// swagger:model ConsolePortTemplate
type ConsolePortTemplate struct {

	// device type
	// Required: true
	DeviceType *NestedDeviceType `json:"device_type"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// type
	Type *ConsolePortTemplateType `json:"type,omitempty"`
}

// Validate validates this console port template
func (m *ConsolePortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsolePortTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePortTemplate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsolePortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsolePortTemplate) UnmarshalBinary(b []byte) error {
	var res ConsolePortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsolePortTemplateType Type
//
// swagger:model ConsolePortTemplateType
type ConsolePortTemplateType struct {

	// label
	// Required: true
	// Enum: [DE-9 DB-25 RJ-11 RJ-12 RJ-45 USB Type A USB Type B USB Type C USB Mini A USB Mini B USB Micro A USB Micro B Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b other]
	Value *string `json:"value"`
}

// Validate validates this console port template type
func (m *ConsolePortTemplateType) Validate(formats strfmt.Registry) error {
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

var consolePortTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DE-9","DB-25","RJ-11","RJ-12","RJ-45","USB Type A","USB Type B","USB Type C","USB Mini A","USB Mini B","USB Micro A","USB Micro B","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consolePortTemplateTypeTypeLabelPropEnum = append(consolePortTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// ConsolePortTemplateTypeLabelDE9 captures enum value "DE-9"
	ConsolePortTemplateTypeLabelDE9 string = "DE-9"

	// ConsolePortTemplateTypeLabelDB25 captures enum value "DB-25"
	ConsolePortTemplateTypeLabelDB25 string = "DB-25"

	// ConsolePortTemplateTypeLabelRJ11 captures enum value "RJ-11"
	ConsolePortTemplateTypeLabelRJ11 string = "RJ-11"

	// ConsolePortTemplateTypeLabelRJ12 captures enum value "RJ-12"
	ConsolePortTemplateTypeLabelRJ12 string = "RJ-12"

	// ConsolePortTemplateTypeLabelRJ45 captures enum value "RJ-45"
	ConsolePortTemplateTypeLabelRJ45 string = "RJ-45"

	// ConsolePortTemplateTypeLabelUSBTypeA captures enum value "USB Type A"
	ConsolePortTemplateTypeLabelUSBTypeA string = "USB Type A"

	// ConsolePortTemplateTypeLabelUSBTypeB captures enum value "USB Type B"
	ConsolePortTemplateTypeLabelUSBTypeB string = "USB Type B"

	// ConsolePortTemplateTypeLabelUSBTypeC captures enum value "USB Type C"
	ConsolePortTemplateTypeLabelUSBTypeC string = "USB Type C"

	// ConsolePortTemplateTypeLabelUSBMiniA captures enum value "USB Mini A"
	ConsolePortTemplateTypeLabelUSBMiniA string = "USB Mini A"

	// ConsolePortTemplateTypeLabelUSBMiniB captures enum value "USB Mini B"
	ConsolePortTemplateTypeLabelUSBMiniB string = "USB Mini B"

	// ConsolePortTemplateTypeLabelUSBMicroA captures enum value "USB Micro A"
	ConsolePortTemplateTypeLabelUSBMicroA string = "USB Micro A"

	// ConsolePortTemplateTypeLabelUSBMicroB captures enum value "USB Micro B"
	ConsolePortTemplateTypeLabelUSBMicroB string = "USB Micro B"

	// ConsolePortTemplateTypeLabelOther captures enum value "Other"
	ConsolePortTemplateTypeLabelOther string = "Other"
)

// prop value enum
func (m *ConsolePortTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, consolePortTemplateTypeTypeLabelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConsolePortTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consolePortTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consolePortTemplateTypeTypeValuePropEnum = append(consolePortTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// ConsolePortTemplateTypeValueDe9 captures enum value "de-9"
	ConsolePortTemplateTypeValueDe9 string = "de-9"

	// ConsolePortTemplateTypeValueDb25 captures enum value "db-25"
	ConsolePortTemplateTypeValueDb25 string = "db-25"

	// ConsolePortTemplateTypeValueRj11 captures enum value "rj-11"
	ConsolePortTemplateTypeValueRj11 string = "rj-11"

	// ConsolePortTemplateTypeValueRj12 captures enum value "rj-12"
	ConsolePortTemplateTypeValueRj12 string = "rj-12"

	// ConsolePortTemplateTypeValueRj45 captures enum value "rj-45"
	ConsolePortTemplateTypeValueRj45 string = "rj-45"

	// ConsolePortTemplateTypeValueUsba captures enum value "usb-a"
	ConsolePortTemplateTypeValueUsba string = "usb-a"

	// ConsolePortTemplateTypeValueUsbb captures enum value "usb-b"
	ConsolePortTemplateTypeValueUsbb string = "usb-b"

	// ConsolePortTemplateTypeValueUsbc captures enum value "usb-c"
	ConsolePortTemplateTypeValueUsbc string = "usb-c"

	// ConsolePortTemplateTypeValueUsbMinia captures enum value "usb-mini-a"
	ConsolePortTemplateTypeValueUsbMinia string = "usb-mini-a"

	// ConsolePortTemplateTypeValueUsbMinib captures enum value "usb-mini-b"
	ConsolePortTemplateTypeValueUsbMinib string = "usb-mini-b"

	// ConsolePortTemplateTypeValueUsbMicroa captures enum value "usb-micro-a"
	ConsolePortTemplateTypeValueUsbMicroa string = "usb-micro-a"

	// ConsolePortTemplateTypeValueUsbMicrob captures enum value "usb-micro-b"
	ConsolePortTemplateTypeValueUsbMicrob string = "usb-micro-b"

	// ConsolePortTemplateTypeValueOther captures enum value "other"
	ConsolePortTemplateTypeValueOther string = "other"
)

// prop value enum
func (m *ConsolePortTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, consolePortTemplateTypeTypeValuePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConsolePortTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsolePortTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsolePortTemplateType) UnmarshalBinary(b []byte) error {
	var res ConsolePortTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
