// Copyright 2016 The go-netbox Authors.
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

package netbox

// Device is a network device
type Device struct {
	ID           int                   `json:"id"`
	Name         string                `json:"name"`
	DisplayName  string                `json:"display_name"`
	DeviceType   *DeviceTypeIdentifier `json:"device_type"`
	DeviceRole   *SimpleIdentifier     `json:"device_role"`
	Platform     *SimpleIdentifier     `json:"platform"`
	Serial       string                `json:"serial"`
	Rack         *RackIdentifier       `json:"rack"`
	Position     int                   `json:"position"`
	Face         int                   `json:"face"`
	ParentDevice *DeviceIdentifier     `json:"parent_device"`
	Status       bool                  `json:"status"`
	PrimaryIP    *IPAddressIdentifier  `json:"primary_ip"`
	PrimaryIP4   *IPAddressIdentifier  `json:"primary_ip4"`
	PrimaryIP6   *IPAddressIdentifier  `json:"primary_ip6"`
	Comments     string
}

// DeviceTypeIdentifier represents the type of network device it is
type DeviceTypeIdentifier struct {
	ID           int               `json:"id"`
	Manufacturer *SimpleIdentifier `json:"manufacturer"`
	Model        string            `json:"model"`
	Slug         string            `json:"slug"`
}

// SimpleIdentifier represents a simple thing that consists of only an id, name,
// and slug
type SimpleIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// RackIdentifier represents a server rack
type RackIdentifier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FacilityID  string `json:"facility_id"`
	DisplayName string `json:"display_name"`
}
