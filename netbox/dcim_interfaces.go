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

// Code generated by generate_functions.go. DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// InterfacesService is used in a Client to access NetBox's dcim/interfaces API methods.
type InterfacesService struct {
	c *Client
}

// Get retrieves an Interface object from NetBox by its ID.
func (s *InterfacesService) Get(id int) (*Interface, error) {
	req, err := s.c.NewRequest(
		http.MethodGet,
		fmt.Sprintf("api/dcim/interfaces/%d/", id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	t := new(Interface)
	err = s.c.Do(req, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// List returns a Page associated with an NetBox API Endpoint.
func (s *InterfacesService) List(options *ListInterfaceOptions) *Page {
	return NewPage(s.c, "api/dcim/interfaces/", options)
}

// Extract retrives a list of Interface objects from page.
func (s *InterfacesService) Extract(page *Page) ([]*Interface, error) {
	if err := page.Err(); err != nil {
		return nil, err
	}

	var groups []*Interface
	if err := json.Unmarshal(page.data.Results, &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

// Create creates a new Interface object in NetBox and returns the ID of the new object.
func (s *InterfacesService) Create(data *Interface) (int, error) {
	req, err := s.c.NewJSONRequest(http.MethodPost, "api/dcim/interfaces/", nil, data)
	if err != nil {
		return 0, err
	}

	g := new(writableInterface)
	err = s.c.Do(req, g)
	if err != nil {
		return 0, err
	}
	return g.ID, nil
}

// Update changes an existing Interface object in NetBox, and returns the ID of the new object.
func (s *InterfacesService) Update(data *Interface) (int, error) {
	req, err := s.c.NewJSONRequest(
		http.MethodPatch,
		fmt.Sprintf("api/dcim/interfaces/%d/", data.ID),
		nil,
		data,
	)
	if err != nil {
		return 0, err
	}

	// g is just used to verify correct api result.
	// data is not changed, because the g is not the full representation that one would
	// get with Get. But if the response was unmarshaled into writableInterface correctly,
	// everything went fine, and we do not need to update data.
	g := new(writableInterface)
	err = s.c.Do(req, g)
	if err != nil {
		return 0, err
	}
	return g.ID, nil
}

// Delete deletes an existing Interface object from NetBox.
func (s *InterfacesService) Delete(data *Interface) error {
	req, err := s.c.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("api/dcim/interfaces/%d/", data.ID),
		nil,
	)
	if err != nil {
		return err
	}

	return s.c.Do(req, nil)
}
