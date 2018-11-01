// +build go1.9,linux,cgo go1.10,darwin,cgo
// +build !ppc64le

// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2016-2018 IBM Corp. All Rights Reserved.
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

package factory

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

// raceEnabled is set to true when the race build tag is enabled.
// see race_test.go
var raceEnabled bool

func buildPlugin(lib string, t *testing.T) {
	t.Helper()
	// check to see if the example plugin exists
	if _, err := os.Stat(lib); err != nil {
		// build the example plugin
		cmd := exec.Command("go", "build", "-buildmode=plugin")
		if raceEnabled {
			cmd.Args = append(cmd.Args, "-race")
		}
		cmd.Args = append(cmd.Args, "github.com/hyperledger/fabric/examples/plugins/bccsp")
		err := cmd.Run()
		if err != nil {
			t.Fatalf("Could not build plugin: [%s]", err)
		}
	}
}

func TestPluginFactoryName(t *testing.T) {
	f := &PluginFactory{}
	assert.Equal(t, f.Name(), PluginFactoryName)
}

func TestPluginFactoryInvalidConfig(t *testing.T) {
	f := &PluginFactory{}
	opts := &FactoryOpts{}

	_, err := f.Get(nil)
	assert.Error(t, err)

	_, err = f.Get(opts)
	assert.Error(t, err)

	opts.PluginOpts = &PluginOpts{}
	_, err = f.Get(opts)
	assert.Error(t, err)
}

func TestPluginFactoryValidConfig(t *testing.T) {
	// build plugin
	lib := "./bccsp.so"
	defer os.Remove(lib)
	buildPlugin(lib, t)

	f := &PluginFactory{}
	opts := &FactoryOpts{
		PluginOpts: &PluginOpts{
			Library: lib,
		},
	}

	csp, err := f.Get(opts)
	assert.NoError(t, err)
	assert.NotNil(t, csp)

	_, err = csp.GetKey([]byte{123})
	assert.NoError(t, err)
}

func TestPluginFactoryFromOpts(t *testing.T) {
	// build plugin
	lib := "./bccsp.so"
	defer os.Remove(lib)
	buildPlugin(lib, t)

	opts := &FactoryOpts{
		ProviderName: "PLUGIN",
		PluginOpts: &PluginOpts{
			Library: lib,
		},
	}
	csp, err := GetBCCSPFromOpts(opts)
	assert.NoError(t, err)
	assert.NotNil(t, csp)

	_, err = csp.GetKey([]byte{123})
	assert.NoError(t, err)
}
