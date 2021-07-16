// Copyright 2020 The Atlas Authors. All rights reserved.
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

package atlas

import "strings"

// NewAddressList creates new AddressList
func NewAddressList() *AddressList {
	return &AddressList{}
}

// NewAddressListFromString creates new AddressList from string
func NewAddressListFromString(str string) *AddressList {
	parts := strings.Split(str, separatorAddressList)
	if len(parts) == 0 {
		return nil
	}

	list := NewAddressList()
	for _, address := range parts {
		if address := NewAddressFromString(address); address != nil {
			list.Append(address)
		}
	}

	if list.Len() > 0 {
		return list
	}

	return nil
}

// Ensure returns new or existing AddressList
func (m *AddressList) Ensure() *AddressList {
	if m == nil {
		return NewAddressList()
	}
	return m
}

// Len return len of specified nested domain or len of the whole AddressList in case no domain specified
func (m *AddressList) Len(domains ...*Domain) int {
	if len(domains) > 0 {
		return m.LenDomain(domains[0])
	}
	return len(m.GetAddresses())
}

// LenDomain return len of specified nested domain
func (m *AddressList) LenDomain(domain *Domain) int {
	res := 0
	for _, address := range m.GetAddresses() {
		if address.GetAddressDomain().Equals(domain) {
			res++
		}
	}
	return res
}

// All wraps GetAddresses and Select and returns all Addresses in specified domains
// May return nil
func (m *AddressList) All(domains ...*Domain) []*Address {
	if len(domains) > 0 {
		return m.Select(domains...).GetAddresses()
	}
	return m.GetAddresses()
}

// Slice returns slice of addresses
// May return nil
func (m *AddressList) Slice(a, b int) []*Address {
	// Sanity check
	if (a < 0) || (b < 0) {
		return nil
	}
	if b > a {
		return nil
	}

	// Boundaries looks like sane, continue with addresses

	addresses := m.GetAddresses()
	if len(addresses) == 0 {
		return nil
	}
	if (a > len(addresses)) || (b > len(addresses)) {
		return nil
	}

	return addresses[a:b]
}

// Has checks whether AddressList has something (in case do domain specified) or specified nested domain exists
func (m *AddressList) Has(domains ...*Domain) bool {
	if len(domains) > 0 {
		return m.HasDomain(domains[0])
	}
	// Deal with AddressList itself
	return m.Len() > 0
}

// First return first address in the list of specified nested domain
// May return nil
func (m *AddressList) First(domains ...*Domain) *Address {
	if len(domains) > 0 {
		return m.FirstDomain(domains[0])
	}
	// Deal with AddressList itself
	if addresses := m.GetAddresses(); len(addresses) > 0 {
		return addresses[0]
	}
	return nil
}

// Last returns last address in the list of specified nested domain
func (m *AddressList) Last(domains ...*Domain) *Address {
	if len(domains) > 0 {
		return m.LastDomain(domains[0])
	}
	// Deal with AddressList itself
	if addresses := m.GetAddresses(); len(addresses) > 0 {
		return addresses[len(addresses)-1]
	}
	return nil
}

// HasDomain checks whether specified nested domain exists
func (m *AddressList) HasDomain(domain *Domain) bool {
	return m.LenDomain(domain) > 0
}

// FirstDomain gets the first address of specified nested domains
func (m *AddressList) FirstDomain(domains ...*Domain) *Address {
	for _, address := range m.GetAddresses() {
		for _, domain := range domains {
			if address.GetAddressDomain().Equals(domain) {
				return address
			}
		}
	}
	return nil
}

// LastDomain gets the last address of specified nested domains
func (m *AddressList) LastDomain(domains ...*Domain) *Address {
	var res *Address = nil
	for _, address := range m.GetAddresses() {
		for _, domain := range domains {
			if address.GetAddressDomain().Equals(domain) {
				res = address
			}
		}
	}
	return res
}

// Select selects all addresses with specified domains into new AddressList
// May return nil in case nothing selected
func (m *AddressList) Select(domains ...*Domain) *AddressList {
	var res *AddressList = nil
	for _, address := range m.GetAddresses() {
		for _, domain := range domains {
			if address.GetAddressDomain().Equals(domain) {
				if res == nil {
					res = NewAddressList()
				}
				res.Append(address)
			}
		}
	}
	return res
}

// Exclude selects all addresses without specified domains into new AddressList
// May return nil in case result is empty
func (m *AddressList) Exclude(domains ...*Domain) *AddressList {
	var res *AddressList = nil
	for _, address := range m.GetAddresses() {
		for _, domain := range domains {
			if address.GetAddressDomain().Equals(domain) {
				// Skip this address, it is in delete list
			} else {
				// Keep this address
				if res == nil {
					res = NewAddressList()
				}
				res.Append(address)
			}
		}
	}
	return res
}

// Delete deletes from the AddressList all addresses with specified domains
func (m *AddressList) Delete(domains ...*Domain) *AddressList {
	var keep []*Address = nil
	for _, address := range m.GetAddresses() {
		for _, domain := range domains {
			if address.GetAddressDomain().Equals(domain) {
				// Skip this address, it is in delete list
			} else {
				// Keep this address
				keep = append(keep, address)
			}
		}
	}
	m.Replace(keep...)
	return m
}

// Append appends addresses to the AddressList
func (m *AddressList) Append(addresses ...*Address) *AddressList {
	if m != nil {
		m.Addresses = append(m.Addresses, addresses...)
		return m
	}
	return nil
}

// Replace replaces existing list with provided list of addresses
func (m *AddressList) Replace(addresses ...*Address) *AddressList {
	if m != nil {
		m.Addresses = append([]*Address{}, addresses...)
		return m
	}
	return nil
}

const separatorAddressList = ","

// String
func (m *AddressList) String() string {
	res := ""
	for _, address := range m.GetAddresses() {
		if len(res) > 0 {
			res += separatorAddressList
		}
		res += address.String()
	}
	return res
}

// FullString
func (m *AddressList) FullString() string {
	res := ""
	for _, address := range m.GetAddresses() {
		if len(res) > 0 {
			res += separatorAddressList
		}
		res += address.FullString()
	}
	return res
}
