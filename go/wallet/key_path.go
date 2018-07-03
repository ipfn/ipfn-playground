// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
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

package wallet

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ipfn/ipfn/go/keypair"
)

// KeyPath - Wallet key path.
type KeyPath struct {
	// Address - Public account address.
	Address string `json:"address,omitempty"`

	// SeedName - Name of the master key.
	SeedName string `json:"seed_name,omitempty"`

	// DerivationPath - Key derivation path.
	DerivationPath string `json:"path,omitempty"`
}

// NewKeyPath - Creates new key path from wallet and path.
func NewKeyPath(wallet, path string, hash bool) *KeyPath {
	if hash && !strings.HasPrefix(path, "x/") {
		return &KeyPath{SeedName: wallet, DerivationPath: fmt.Sprintf("x/%s", path)}
	}
	return &KeyPath{SeedName: wallet, DerivationPath: path}
}

// ParseKeyPath - Parses key path, returns seed name and path.
func ParseKeyPath(key string) (_ *KeyPath, err error) {
	seed, path, err := parseKeyPath(key)
	if err != nil {
		return
	}
	return &KeyPath{SeedName: seed, DerivationPath: path}, nil
}

// UnpackPath - Unpacks derivation path.
func (path *KeyPath) UnpackPath() string {
	if t := strings.TrimPrefix(path.DerivationPath, "x/"); t != path.DerivationPath {
		return keypair.HashPath(t)
	}
	return path.DerivationPath
}

// IsHashPath - Returns true if derivation path is hashpath.
func (path *KeyPath) IsHashPath() bool {
	return strings.HasPrefix(path.DerivationPath, "x/")
}

// String - Returns joined key path.
func (path *KeyPath) String() string {
	return strings.Join([]string{path.SeedName, path.DerivationPath}, "/")
}

// Extend - Returns extended key path.
func (path *KeyPath) Extend(extra string) *KeyPath {
	return &KeyPath{
		SeedName:       path.SeedName,
		DerivationPath: strings.Join([]string{path.DerivationPath, extra}, "/"),
	}
}

// MarshalJSON - Marshals key path to JSON.
func (path *KeyPath) MarshalJSON() (res []byte, _ error) {
	res = append(res, '"')
	res = append(res, []byte(path.String())...)
	res = append(res, '"')
	return
}

// UnmarshalJSON - Unmarshals key path from JSON.
func (path *KeyPath) UnmarshalJSON(data []byte) (err error) {
	data = bytes.Trim(data, "\"")
	path.SeedName, path.DerivationPath, err = parseKeyPath(string(data))
	return
}

func parseKeyPath(path string) (seed, res string, err error) {
	if path == "" {
		path = "default"
	}
	if !strings.Contains(path, "/") {
		seed = path
		return
	}
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		err = fmt.Errorf("path %q is invalid", path)
		return
	}
	switch parts[1] {
	case "x", "m":
		seed = parts[0]
		res = strings.Join(parts[1:], "/")
		return
	}
	err = fmt.Errorf("unknown part %q in path %q (supported: x, m)", parts[1], path)
	return
}
