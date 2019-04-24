// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableSite writable site
// swagger:model WritableSite
type WritableSite struct {

	// ASN
	// Maximum: 4.294967295e+09
	// Minimum: 1
	Asn *int64 `json:"asn,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Contact E-mail
	// Max Length: 254
	// Format: email
	ContactEmail strfmt.Email `json:"contact_email,omitempty"`

	// Contact name
	// Max Length: 50
	ContactName string `json:"contact_name,omitempty"`

	// Contact phone
	// Max Length: 20
	ContactPhone string `json:"contact_phone,omitempty"`

	// Count circuits
	// Read Only: true
	CountCircuits int64 `json:"count_circuits,omitempty"`

	// Count devices
	// Read Only: true
	CountDevices int64 `json:"count_devices,omitempty"`

	// Count prefixes
	// Read Only: true
	CountPrefixes int64 `json:"count_prefixes,omitempty"`

	// Count racks
	// Read Only: true
	CountRacks int64 `json:"count_racks,omitempty"`

	// Count vlans
	// Read Only: true
	CountVlans int64 `json:"count_vlans,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 100
	Description string `json:"description,omitempty"`

	// Facility
	// Max Length: 50
	Facility string `json:"facility,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Latitude
	Latitude *string `json:"latitude,omitempty"`

	// Longitude
	Longitude *string `json:"longitude,omitempty"`

	// Name
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// Physical address
	// Max Length: 200
	PhysicalAddress string `json:"physical_address,omitempty"`

	// Region
	Region *int64 `json:"region,omitempty"`

	// Shipping address
	// Max Length: 200
	ShippingAddress string `json:"shipping_address,omitempty"`

	// Slug
	// Required: true
	// Max Length: 50
	// Min Length: 1
	// Pattern: ^[-a-zA-Z0-9_]+$
	Slug *string `json:"slug"`

	// Status
	// Enum: [1 2 4]
	Status int64 `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// Time zone
	TimeZone string `json:"time_zone,omitempty"`
}

// Validate validates this writable site
func (m *WritableSite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAsn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactPhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhysicalAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShippingAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableSite) validateAsn(formats strfmt.Registry) error {

	if swag.IsZero(m.Asn) { // not required
		return nil
	}

	if err := validate.MinimumInt("asn", "body", int64(*m.Asn), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("asn", "body", int64(*m.Asn), 4.294967295e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateContactEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactEmail) { // not required
		return nil
	}

	if err := validate.MaxLength("contact_email", "body", string(m.ContactEmail), 254); err != nil {
		return err
	}

	if err := validate.FormatOf("contact_email", "body", "email", m.ContactEmail.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateContactName(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactName) { // not required
		return nil
	}

	if err := validate.MaxLength("contact_name", "body", string(m.ContactName), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateContactPhone(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactPhone) { // not required
		return nil
	}

	if err := validate.MaxLength("contact_phone", "body", string(m.ContactPhone), 20); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateFacility(formats strfmt.Registry) error {

	if swag.IsZero(m.Facility) { // not required
		return nil
	}

	if err := validate.MaxLength("facility", "body", string(m.Facility), 50); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateName(formats strfmt.Registry) error {

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

func (m *WritableSite) validatePhysicalAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.PhysicalAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("physical_address", "body", string(m.PhysicalAddress), 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateShippingAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.ShippingAddress) { // not required
		return nil
	}

	if err := validate.MaxLength("shipping_address", "body", string(m.ShippingAddress), 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	if err := validate.MinLength("slug", "body", string(*m.Slug), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("slug", "body", string(*m.Slug), 50); err != nil {
		return err
	}

	if err := validate.Pattern("slug", "body", string(*m.Slug), `^[-a-zA-Z0-9_]+$`); err != nil {
		return err
	}

	return nil
}

var writableSiteTypeStatusPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,4]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableSiteTypeStatusPropEnum = append(writableSiteTypeStatusPropEnum, v)
	}
}

// prop value enum
func (m *WritableSite) validateStatusEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, writableSiteTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableSite) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableSite) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", string(m.Tags[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableSite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableSite) UnmarshalBinary(b []byte) error {
	var res WritableSite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
