/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package flog_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"

	"github.com/ipfn/ipfn/src/go/utils/flog"
	"github.com/ipfn/ipfn/src/go/utils/flog/fabenc"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"google.golang.org/grpc/grpclog"
)

func TestFabricLoggerEncoding(t *testing.T) {
	formatters, err := fabenc.ParseFormat("%{color}[%{module}] %{shortfunc} -> %{level:.4s}%{color:reset} %{message}")
	assert.NoError(t, err)
	enc := fabenc.NewFormatEncoder(formatters...)

	buf := &bytes.Buffer{}
	core := zapcore.NewCore(enc, zapcore.AddSync(buf), zap.NewAtomicLevel())
	zl := flog.NewZapLogger(core).Named("test").With(zap.String("extra", "field"))
	fl := flog.NewFabricLogger(zl)

	buf.Reset()
	fl.Info("string value", 0, 1.23, struct{}{})
	assert.Equal(t, "\x1b[34m[test] TestFabricLoggerEncoding -> INFO\x1b[0m string value 0 1.23 {} {\"extra\": \"field\"}\n", buf.String())

	buf.Reset()
	fl.Infof("string %s, %d, %.3f, %v", "strval", 0, 1.23, struct{}{})
	assert.Equal(t, "\x1b[34m[test] TestFabricLoggerEncoding -> INFO\x1b[0m string strval, 0, 1.230, {} {\"extra\": \"field\"}\n", buf.String())

	buf.Reset()
	fl.Infow("this is a message", "int", 0, "float", 1.23, "struct", struct{}{})
	assert.Equal(t, "\x1b[34m[test] TestFabricLoggerEncoding -> INFO\x1b[0m this is a message {\"extra\": \"field\", \"int\": 0, \"float\": 1.23, \"struct\": {}}\n", buf.String())
}

func TestFabricLogger(t *testing.T) {
	var enabler zap.LevelEnablerFunc = func(l zapcore.Level) bool { return true }

	var tests = []struct {
		desc    string
		f       func(fl *flog.FabricLogger)
		level   zapcore.Level
		message string
		fields  []zapcore.Field
		panics  bool
	}{
		{
			desc:    "DPanic",
			f:       func(fl *flog.FabricLogger) { fl.DPanic("arg1", "arg2") },
			level:   zapcore.DPanicLevel,
			message: "arg1 arg2",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "DPanicf",
			f:       func(fl *flog.FabricLogger) { fl.DPanicf("panic: %s, %d", "reason", 99) },
			level:   zapcore.DPanicLevel,
			message: "panic: reason, 99",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "DPanicw",
			f:       func(fl *flog.FabricLogger) { fl.DPanicw("I'm in a panic", "reason", "something", "code", 99) },
			level:   zapcore.DPanicLevel,
			message: "I'm in a panic",
			fields:  []zapcore.Field{zap.String("reason", "something"), zap.Int("code", 99)},
		},
		{
			desc:    "Debug",
			f:       func(fl *flog.FabricLogger) { fl.Debug("arg1", "arg2") },
			level:   zapcore.DebugLevel,
			message: "arg1 arg2",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Debugf",
			f:       func(fl *flog.FabricLogger) { fl.Debugf("debug: %s, %d", "goo", 99) },
			level:   zapcore.DebugLevel,
			message: "debug: goo, 99",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Debugw",
			f:       func(fl *flog.FabricLogger) { fl.Debugw("debug data", "key", "value") },
			level:   zapcore.DebugLevel,
			message: "debug data",
			fields:  []zapcore.Field{zap.String("key", "value")},
		},
		{
			desc:    "Error",
			f:       func(fl *flog.FabricLogger) { fl.Error("oh noes", errors.New("bananas")) },
			level:   zapcore.ErrorLevel,
			message: "oh noes bananas",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Errorf",
			f:       func(fl *flog.FabricLogger) { fl.Errorf("error: %s", errors.New("bananas")) },
			level:   zapcore.ErrorLevel,
			message: "error: bananas",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Errorw",
			f:       func(fl *flog.FabricLogger) { fl.Errorw("something failed", "err", errors.New("bananas")) },
			level:   zapcore.ErrorLevel,
			message: "something failed",
			fields:  []zapcore.Field{zap.NamedError("err", errors.New("bananas"))},
		},
		{
			desc:    "Info",
			f:       func(fl *flog.FabricLogger) { fl.Info("fyi", "things are great") },
			level:   zapcore.InfoLevel,
			message: "fyi things are great",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Infof",
			f:       func(fl *flog.FabricLogger) { fl.Infof("fyi: %s", "things are great") },
			level:   zapcore.InfoLevel,
			message: "fyi: things are great",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Infow",
			f:       func(fl *flog.FabricLogger) { fl.Infow("fyi", "fish", "are smelly", "fruit", "is sweet") },
			level:   zapcore.InfoLevel,
			message: "fyi",
			fields:  []zapcore.Field{zap.String("fish", "are smelly"), zap.String("fruit", "is sweet")},
		},
		{
			desc:    "Panic",
			f:       func(fl *flog.FabricLogger) { fl.Panic("oh noes", errors.New("platypus")) },
			level:   zapcore.PanicLevel,
			message: "oh noes platypus",
			fields:  []zapcore.Field{},
			panics:  true,
		},
		{
			desc:    "Panicf",
			f:       func(fl *flog.FabricLogger) { fl.Panicf("error: %s", errors.New("platypus")) },
			level:   zapcore.PanicLevel,
			message: "error: platypus",
			fields:  []zapcore.Field{},
			panics:  true,
		},
		{
			desc:    "Panicw",
			f:       func(fl *flog.FabricLogger) { fl.Panicw("something failed", "err", errors.New("platypus")) },
			level:   zapcore.PanicLevel,
			message: "something failed",
			fields:  []zapcore.Field{zap.NamedError("err", errors.New("platypus"))},
			panics:  true,
		},
		{
			desc:    "Warn",
			f:       func(fl *flog.FabricLogger) { fl.Warn("oh noes", errors.New("monkeys")) },
			level:   zapcore.WarnLevel,
			message: "oh noes monkeys",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Warnf",
			f:       func(fl *flog.FabricLogger) { fl.Warnf("error: %s", errors.New("monkeys")) },
			level:   zapcore.WarnLevel,
			message: "error: monkeys",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Warnw",
			f:       func(fl *flog.FabricLogger) { fl.Warnw("something is weird", "err", errors.New("monkeys")) },
			level:   zapcore.WarnLevel,
			message: "something is weird",
			fields:  []zapcore.Field{zap.NamedError("err", errors.New("monkeys"))},
		},
		{
			desc:    "Warning",
			f:       func(fl *flog.FabricLogger) { fl.Warning("oh noes", errors.New("monkeys")) },
			level:   zapcore.WarnLevel,
			message: "oh noes monkeys",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Warningf",
			f:       func(fl *flog.FabricLogger) { fl.Warningf("error: %s", errors.New("monkeys")) },
			level:   zapcore.WarnLevel,
			message: "error: monkeys",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "With",
			f:       func(fl *flog.FabricLogger) { fl.With("key", "value").Debug("cool messages", "and stuff") },
			level:   zapcore.DebugLevel,
			message: "cool messages and stuff",
			fields:  []zapcore.Field{zap.String("key", "value")},
		},
		{
			desc: "WithOptions",
			f: func(fl *flog.FabricLogger) {
				fl.WithOptions(zap.Fields(zap.String("optionkey", "optionvalue"))).Debug("cool messages", "and stuff")
			},
			level:   zapcore.DebugLevel,
			message: "cool messages and stuff",
			fields:  []zapcore.Field{zap.String("optionkey", "optionvalue")},
		},
		{
			desc:    "Critical",
			f:       func(fl *flog.FabricLogger) { fl.Critical("critical as error", errors.New("kiwi")) },
			level:   zapcore.ErrorLevel,
			message: "critical as error kiwi",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Criticalf",
			f:       func(fl *flog.FabricLogger) { fl.Criticalf("critical: %s", errors.New("kiwi")) },
			level:   zapcore.ErrorLevel,
			message: "critical: kiwi",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Notice",
			f:       func(fl *flog.FabricLogger) { fl.Notice("notice", "as info") },
			level:   zapcore.InfoLevel,
			message: "notice as info",
			fields:  []zapcore.Field{},
		},
		{
			desc:    "Noticef",
			f:       func(fl *flog.FabricLogger) { fl.Noticef("notice: %s", "this is info") },
			level:   zapcore.InfoLevel,
			message: "notice: this is info",
			fields:  []zapcore.Field{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			core, logs := observer.New(enabler)
			fl := flog.NewFabricLogger(zap.New(core)).Named("lname")

			if tc.panics {
				assert.Panics(t, func() { tc.f(fl) })
			} else {
				tc.f(fl)
			}

			err := fl.Sync()
			assert.NoError(t, err)

			entries := logs.All()
			assert.Len(t, entries, 1)
			entry := entries[0]

			assert.Equal(t, tc.level, entry.Level)
			assert.Equal(t, tc.message, entry.Message)
			assert.Equal(t, tc.fields, entry.Context)
			assert.Equal(t, "lname", entry.LoggerName)
		})
	}
}

func TestIsEnabledFor(t *testing.T) {
	formatters, err := fabenc.ParseFormat("%{color}[%{module}] %{shortfunc} -> %{level:.4s}%{color:reset} %{message}")
	assert.NoError(t, err)
	enc := fabenc.NewFormatEncoder(formatters...)

	enablerCallCount := 0
	enabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		enablerCallCount++
		return l == zapcore.ErrorLevel
	})

	core := zapcore.NewCore(enc, zapcore.AddSync(ioutil.Discard), enabler)
	zl := zap.New(core).Named("test")
	fl := flog.NewFabricLogger(zl)

	assert.True(t, fl.IsEnabledFor(zapcore.ErrorLevel))
	assert.False(t, fl.IsEnabledFor(zapcore.PanicLevel))
	assert.Equal(t, 2, enablerCallCount)
}

func logCaller(l grpclog.Logger, msg string)   { l.Println(msg) }
func callWrapper(l grpclog.Logger, msg string) { logCaller(l, msg) }

func TestGRPCLogger(t *testing.T) {
	// ensure it includes the name as module, logs at debug level, and the caller with appropriate skip level
	formatters, err := fabenc.ParseFormat("%{module} %{level} %{shortfunc} %{message}")
	assert.NoError(t, err)
	enc := fabenc.NewFormatEncoder(formatters...)

	buf := &bytes.Buffer{}
	enabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool { return true })
	core := zapcore.NewCore(enc, zapcore.AddSync(buf), enabler)
	zl := zap.New(core).Named("grpc")
	gl := flog.NewGRPCLogger(zl)

	callWrapper(gl, "message")
	assert.Equal(t, "grpc DEBUG TestGRPCLogger message\n", buf.String())
}
