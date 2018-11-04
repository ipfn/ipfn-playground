/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package flog_test

import (
	"bytes"
	"testing"

	"github.com/ipfn/ipfn/src/go/utils/flog"
	"github.com/stretchr/testify/assert"
)

func TestGlobalReset(t *testing.T) {
	flog.Reset()
	flog.SetModuleLevel("module", "DEBUG")
	flog.Global.SetFormat("json")

	system, err := flog.New(flog.Config{})
	assert.NoError(t, err)
	assert.NotEqual(t, flog.Global.ModuleLevels, system.ModuleLevels)
	assert.NotEqual(t, flog.Global.Encoding(), system.Encoding())

	flog.Reset()
	assert.Equal(t, flog.Global.ModuleLevels, system.ModuleLevels)
	assert.Equal(t, flog.Global.Encoding(), system.Encoding())
}

func TestGlobalInitConsole(t *testing.T) {
	flog.Reset()
	defer flog.Reset()

	buf := &bytes.Buffer{}
	flog.Init(flog.Config{
		Format:  "%{message}",
		LogSpec: "DEBUG",
		Writer:  buf,
	})

	logger := flog.MustGetLogger("testlogger")
	logger.Debug("this is a message")

	assert.Equal(t, "this is a message\n", buf.String())
}

func TestGlobalInitJSON(t *testing.T) {
	flog.Reset()
	defer flog.Reset()

	buf := &bytes.Buffer{}
	flog.Init(flog.Config{
		Format:  "json",
		LogSpec: "DEBUG",
		Writer:  buf,
	})

	logger := flog.MustGetLogger("testlogger")
	logger.Debug("this is a message")

	assert.Regexp(t, `{"level":"debug","ts":\d+.\d+,"name":"testlogger","caller":"flog/global_test.go:\d+","msg":"this is a message"}\s+`, buf.String())
}

func TestGlobalInitPanic(t *testing.T) {
	flog.Reset()
	defer flog.Reset()

	assert.Panics(t, func() {
		flog.Init(flog.Config{
			Format: "%{color:evil}",
		})
	})
}

func TestGlobalGetAndRestoreLevels(t *testing.T) {
	flog.Reset()

	flog.SetModuleLevel("test-1", "DEBUG")
	flog.SetModuleLevel("test-2", "ERROR")
	flog.SetModuleLevel("test-3", "WARN")
	levels := flog.GetModuleLevels()

	assert.Equal(t, "DEBUG", flog.GetModuleLevel("test-1"))
	assert.Equal(t, "ERROR", flog.GetModuleLevel("test-2"))
	assert.Equal(t, "WARN", flog.GetModuleLevel("test-3"))

	flog.Reset()
	assert.Equal(t, "INFO", flog.GetModuleLevel("test-1"))
	assert.Equal(t, "INFO", flog.GetModuleLevel("test-2"))
	assert.Equal(t, "INFO", flog.GetModuleLevel("test-3"))

	flog.RestoreLevels(levels)
	assert.Equal(t, "DEBUG", flog.GetModuleLevel("test-1"))
	assert.Equal(t, "ERROR", flog.GetModuleLevel("test-2"))
	assert.Equal(t, "WARN", flog.GetModuleLevel("test-3"))
}

func TestGlobalDefaultLevel(t *testing.T) {
	flog.Reset()

	assert.Equal(t, "INFO", flog.DefaultLevel())
}

func TestGlobalSetModuleLevels(t *testing.T) {
	flog.Reset()

	flog.SetModuleLevel("a-module", "DEBUG")
	flog.SetModuleLevel("another-module", "DEBUG")
	assert.Equal(t, "DEBUG", flog.GetModuleLevel("a-module"))
	assert.Equal(t, "DEBUG", flog.GetModuleLevel("another-module"))

	flog.SetModuleLevels("^a-", "INFO")
	assert.Equal(t, "INFO", flog.GetModuleLevel("a-module"))
	assert.Equal(t, "DEBUG", flog.GetModuleLevel("another-module"))

	flog.SetModuleLevels("module", "WARN")
	assert.Equal(t, "WARN", flog.GetModuleLevel("a-module"))
	assert.Equal(t, "WARN", flog.GetModuleLevel("another-module"))
}

func TestGlobalSetModuleLevelsBadRegex(t *testing.T) {
	flog.Reset()

	err := flog.SetModuleLevels("((", "DEBUG")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing regexp: ")
}

func TestGlobalMustGetLogger(t *testing.T) {
	flog.Reset()

	l := flog.MustGetLogger("module-name")
	assert.NotNil(t, l)
}

func TestFlogginInitPanic(t *testing.T) {
	defer flog.Reset()

	assert.Panics(t, func() {
		flog.Init(flog.Config{
			Format: "%{color:broken}",
		})
	})
}

func TestGlobalInitFromSpec(t *testing.T) {
	defer flog.Reset()

	tests := []struct {
		name           string
		spec           string
		expectedResult string
		expectedLevels map[string]string
	}{
		{
			name:           "SingleModuleLevel",
			spec:           "a=info",
			expectedResult: "INFO",
			expectedLevels: map[string]string{"a": "INFO"},
		},
		{
			name:           "MultipleModulesMultipleLevels",
			spec:           "a=info:b=debug",
			expectedResult: "INFO",
			expectedLevels: map[string]string{"a": "INFO", "b": "DEBUG"},
		},
		{
			name:           "MultipleModulesSameLevel",
			spec:           "a,b=warning",
			expectedResult: "INFO",
			expectedLevels: map[string]string{"a": "WARN", "b": "WARN"},
		},
		{
			name:           "DefaultAndModules",
			spec:           "ERROR:a=warning",
			expectedResult: "ERROR",
			expectedLevels: map[string]string{"a": "WARN"},
		},
		{
			name:           "ModuleAndDefault",
			spec:           "a=debug:info",
			expectedResult: "INFO",
			expectedLevels: map[string]string{"a": "DEBUG"},
		},
		{
			name:           "EmptyModuleEqualsLevel",
			spec:           "=info",
			expectedResult: "INFO",
			expectedLevels: map[string]string{},
		},
		{
			name:           "InvalidSyntax",
			spec:           "a=b=c",
			expectedResult: "INFO",
			expectedLevels: map[string]string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			flog.Reset()

			l := flog.InitFromSpec(tc.spec)
			assert.Equal(t, tc.expectedResult, l)

			for k, v := range tc.expectedLevels {
				assert.Equal(t, v, flog.GetModuleLevel(k))
			}
		})
	}
}
