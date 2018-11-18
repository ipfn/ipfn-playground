// Copyright © 2017-2018 Łukasz Kurowski. All Rights Reserved.
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

// Package logger implements logging helpers.
package logger

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
)

var (
	// Verbose - Enables logger Verbose mode.
	Verbose bool

	logger, _ = zap.NewDevelopment()
	sugar     = logger.Sugar()
)

// Sync - Flushes logs to buffer.
func Sync() error {
	return logger.Sync()
}

// Disable - Disables logger.
func Disable() {
	SetLogger(zap.NewNop())
}

// SetLogger - Sets logger.
func SetLogger(l *zap.Logger) {
	logger = l
	sugar = l.Sugar()
}

// Fatal - Prints fatal error and exits.
func Fatal(format string, args ...interface{}) {
	logger.Fatal(fmt.Sprintf(format, args...))
}

// Error - Prints error and exits.
func Error(err error) {
	logger.Error(fmt.Sprintf("error: %v", err))
	os.Exit(1)
}

// Debugf - Prints new formatted line if Verbose is true.
func Debugf(format string, args ...interface{}) {
	logger.Debug(fmt.Sprintf(format, args...))
}

// Debug - Prints a line if Verbose is true.
func Debug(line string) {
	logger.Debug(line)
}

// Printf - Prints new formatted line.
func Printf(format string, args ...interface{}) {
	logger.Info(fmt.Sprintf(format, args...))
}

// Info - Prints new formatted line.
func Info(msg string) {
	logger.Info(msg)
}

// Infow - Prints new formatted line.
func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

// Debugw - Prints new formatted line.
func Debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues...)
}

// Print - Prints a line.
func Print(a ...interface{}) {
	fmt.Println(a...)
}

// Line - Prints a new line.
func Line() {
	fmt.Printf("\n")
}

// PrintJSON - Prints json to console.
func PrintJSON(v interface{}) {
	body, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", body)
}
