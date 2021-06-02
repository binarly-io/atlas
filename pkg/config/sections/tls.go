// Copyright 2021 The Atlas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sections

import (
	"fmt"
	"github.com/binarly-io/atlas/pkg/config/items"
)

// TLSConfigurator
type TLSConfigurator interface {
	GetTLSEnabled() bool
	GetTLSCAFile() string
	GetTLSServerHostOverride() string
}

// Interface compatibility
var _ TLSConfigurator = TLS{}

// TLS
type TLS struct {
	TLS *items.TLS `mapstructure:"tls"`
}

// TLSNormalize
func (c TLS) TLSNormalize() {
	if c.TLS == nil {
		c.TLS = items.NewTLS()
	}
}

// GetTLSEnabled
func (c TLS) GetTLSEnabled() bool {
	return c.TLS.GetEnabled()
}

// GetTLSCAFile
func (c TLS) GetTLSCAFile() string {
	return c.TLS.GetCAFile()
}

// GetTLSServerHostOverride
func (c TLS) GetTLSServerHostOverride() string {
	return c.TLS.GetServerHostOverride()
}

// String
func (c TLS) String() string {
	return fmt.Sprintf("TLS=%s", c.TLS)
}
