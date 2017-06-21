// Copyright 2017 The go-netbox Authors.
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

// Describe Interface Status
const (
	InterfaceStatusOffline   = 0
	InterfaceStatusActive    = 1
	InterfaceStatusPlanned   = 2
	InterfaceStatusStaged    = 3
	InterfaceStatusFailed    = 4
	InterfaceStatusInventory = 5
)

// Describe the face of a device in the rack
const (
	RackFaceFront = 0
	RackFaceRear  = 1
)
