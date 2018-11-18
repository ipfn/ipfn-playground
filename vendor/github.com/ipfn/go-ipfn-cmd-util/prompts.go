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

package cmdutil

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	prompt "github.com/segmentio/go-prompt"

	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

// PromptLine - Prompts for entire line.
func PromptLine(entity string) (_ string, err error) {
	fmt.Printf("%s: ", entity)
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		return
	}
	return strings.TrimSpace(line), nil
}

// PromptPassword - Prompts for masked password.
func PromptPassword(entity string) string {
	for {
		value := prompt.PasswordMasked(entity)
		if value != "" {
			return value
		}
		logger.Printf("%s cannot be empty", entity)
	}
}

// PromptPasswordRepeated - Prompts for masked repeated password.
func PromptPasswordRepeated(entity string) string {
	for {
		value := prompt.PasswordMasked(fmt.Sprintf("Choose %s", entity))
		if value == "" {
			logger.Fatal(fmt.Sprintf("Canceled %s input", entity))
			continue
		}
		repeated := prompt.PasswordMasked(fmt.Sprintf("Repeat %s", entity))
		if repeated != value {
			logger.Printf("Repeated %s does not match", entity)
			continue
		}
		return value
	}
}

// PromptConfirmed - Prompts for a value confirmed with function.
func PromptConfirmed(entity string, fn func(string) error) string {
	for {
		value, err := PromptLine(fmt.Sprintf("Choose %s", entity))
		if err != nil {
			logger.Error(err)
		}
		err = fn(value)
		if err == nil {
			return value
		}
		logger.Printf("Error: %v", err)
	}
}
