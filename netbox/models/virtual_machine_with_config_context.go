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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VirtualMachineWithConfigContext virtual machine with config context
//
// swagger:model VirtualMachineWithConfigContext
type VirtualMachineWithConfigContext struct {

	// cluster
	// Required: true
	Cluster *NestedCluster `json:"cluster"`

	// Comments
	Comments string `json:"comments,omitempty"`

	// Config context
	// Read Only: true
	ConfigContext map[string]string `json:"config_context,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Disk (GB)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Disk *int64 `json:"disk,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Local context data
	LocalContextData *string `json:"local_context_data,omitempty"`

	// Memory (MB)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	Memory *int64 `json:"memory,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// platform
	Platform *NestedPlatform `json:"platform,omitempty"`

	// primary ip
	PrimaryIP *NestedIPAddress `json:"primary_ip,omitempty"`

	// primary ip4
	PrimaryIp4 *NestedIPAddress `json:"primary_ip4,omitempty"`

	// primary ip6
	PrimaryIp6 *NestedIPAddress `json:"primary_ip6,omitempty"`

	// role
	Role *NestedDeviceRole `json:"role,omitempty"`

	// site
	Site *NestedSite `json:"site,omitempty"`

	// status
	Status *VirtualMachineWithConfigContextStatus `json:"status,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// VCPUs
	// Maximum: 32767
	// Minimum: 0
	Vcpus *int64 `json:"vcpus,omitempty"`
}

// Validate validates this virtual machine with config context
func (m *VirtualMachineWithConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp4(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryIp6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcpus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VirtualMachineWithConfigContext) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if err := validate.MinimumInt("disk", "body", int64(*m.Disk), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("disk", "body", int64(*m.Disk), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateLastUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if err := validate.MinimumInt("memory", "body", int64(*m.Memory), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("memory", "body", int64(*m.Memory), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validatePlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	if m.Platform != nil {
		if err := m.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validatePrimaryIP(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryIP) { // not required
		return nil
	}

	if m.PrimaryIP != nil {
		if err := m.PrimaryIP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validatePrimaryIp4(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryIp4) { // not required
		return nil
	}

	if m.PrimaryIp4 != nil {
		if err := m.PrimaryIp4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip4")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validatePrimaryIp6(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryIp6) { // not required
		return nil
	}

	if m.PrimaryIp6 != nil {
		if err := m.PrimaryIp6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primary_ip6")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateSite(formats strfmt.Registry) error {

	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateTags(formats strfmt.Registry) error {

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

func (m *VirtualMachineWithConfigContext) validateTenant(formats strfmt.Registry) error {

	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *VirtualMachineWithConfigContext) validateVcpus(formats strfmt.Registry) error {

	if swag.IsZero(m.Vcpus) { // not required
		return nil
	}

	if err := validate.MinimumInt("vcpus", "body", int64(*m.Vcpus), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vcpus", "body", int64(*m.Vcpus), 32767, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineWithConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineWithConfigContext) UnmarshalBinary(b []byte) error {
	var res VirtualMachineWithConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VirtualMachineWithConfigContextStatus Status
//
// swagger:model VirtualMachineWithConfigContextStatus
type VirtualMachineWithConfigContextStatus struct {

	// label
	// Required: true
	Label *string `json:"label"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this virtual machine with config context status
func (m *VirtualMachineWithConfigContextStatus) Validate(formats strfmt.Registry) error {
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

func (m *VirtualMachineWithConfigContextStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *VirtualMachineWithConfigContextStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VirtualMachineWithConfigContextStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VirtualMachineWithConfigContextStatus) UnmarshalBinary(b []byte) error {
	var res VirtualMachineWithConfigContextStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
