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
	"fmt"
	"os"
)

var (
	// Verbose - Enables logger Verbose mode.
	Verbose bool
)

// Fatal - Prints fatal error and exits.
func Fatal(format string, args ...interface{}) {
	fmt.Printf("error: %s\n", fmt.Sprintf(format, args...))
	os.Exit(1)
}

// Error - Prints error and exits.
func Error(err error) {
	fmt.Printf("error: %v\n", err)
	os.Exit(1)
}

// Debugf - Prints new formatted line if Verbose is true.
func Debugf(format string, args ...interface{}) {
	if Verbose {
		fmt.Println(fmt.Sprintf(format, args...))
	}
}

// Debug - Prints a line if Verbose is true.
func Debug(line string) {
	if Verbose {
		fmt.Println(line)
	}
}

// Printf - Prints new formatted line.
func Printf(format string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(format, args...))
}

// Print - Prints a line.
func Print(line string) {
	fmt.Println(line)
}

// Line - Prints a new line.
func Line() {
	fmt.Printf("\n")
}
