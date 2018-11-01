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
	"errors"
	"fmt"
	"os"
	"plugin"

	"github.com/ipfn/ipfn/src/go/bccsp"
)

const (
	// PluginFactoryName is the factory name for BCCSP plugins
	PluginFactoryName = "PLUGIN"
)

// PluginOpts contains the options for the PluginFactory
type PluginOpts struct {
	// Path to plugin library
	Library string
	// Config map for the plugin library
	Config map[string]interface{}
}

// PluginFactory is the factory for BCCSP plugins
type PluginFactory struct{}

// Name returns the name of this factory
func (f *PluginFactory) Name() string {
	return PluginFactoryName
}

// Get returns an instance of BCCSP using Opts.
func (f *PluginFactory) Get(config *FactoryOpts) (bccsp.BCCSP, error) {
	// check for valid config
	if config == nil || config.PluginOpts == nil {
		return nil, errors.New("Invalid config. It must not be nil.")
	}

	// Library is required property
	if config.PluginOpts.Library == "" {
		return nil, errors.New("Invalid config: missing property 'Library'")
	}

	// make sure the library exists
	if _, err := os.Stat(config.PluginOpts.Library); err != nil {
		return nil, fmt.Errorf("Could not find library '%s' [%s]", config.PluginOpts.Library, err)
	}

	// attempt to load the library as a plugin
	plug, err := plugin.Open(config.PluginOpts.Library)
	if err != nil {
		return nil, fmt.Errorf("Failed to load plugin '%s' [%s]", config.PluginOpts.Library, err)
	}

	// lookup the required symbol 'New'
	sym, err := plug.Lookup("New")
	if err != nil {
		return nil, fmt.Errorf("Could not find required symbol 'CryptoServiceProvider' [%s]", err)
	}

	// check to make sure symbol New meets the required function signature
	new, ok := sym.(func(config map[string]interface{}) (bccsp.BCCSP, error))
	if !ok {
		return nil, fmt.Errorf("Plugin does not implement the required function signature for 'New'")
	}

	return new(config.PluginOpts.Config)
}
