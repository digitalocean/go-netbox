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

package netbox

import (
	"fmt"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	client "github.com/netbox-community/go-netbox/plumbing"
)

// NewNetboxAt returns a client which will connect to the given
// hostname, which can optionally include a port, e.g. localhost:8000
func NewNetboxAt(host string) *client.Netbox {
	t := client.DefaultTransportConfig().WithHost(host)
	return client.NewHTTPClientWithConfig(strfmt.Default, t)
}

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

// NewNetboxWithAPIKey returna client which will connect to the given
// hostname (and optionally port), and will set the expected Authorization
// header on each request
func NewNetboxWithAPIKey(host string, apiToken string) *client.Netbox {
	t := runtimeclient.New(host, client.DefaultBasePath, client.DefaultSchemes)
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, apiToken))
	return client.New(t, strfmt.Default)
}
