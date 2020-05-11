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

package config_service

import (
	"bytes"
	"fmt"

	conf "github.com/spf13/viper"
)

type ConfigService struct {
	Verbose bool     `json:"verbose" yaml:"verbose"`
	Brokers []string `json:"brokers" yaml:"brokers"`
	Topic   string   `json:"topic"   yaml:"topic"`
}

var Config ConfigService

func ReadIn() {
	conf.Unmarshal(&Config)
}

func (c *ConfigService) String() string {
	b := &bytes.Buffer{}

	_, _ = fmt.Fprintf(b, "Verbose: %v\n", c.Verbose)
	_, _ = fmt.Fprintf(b, "Brokers: %v\n", c.Brokers)
	_, _ = fmt.Fprintf(b, "Topic: %v\n", c.Topic)

	return b.String()
}
