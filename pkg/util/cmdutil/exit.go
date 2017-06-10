// Licensed to Pulumi Corporation ("Pulumi") under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// Pulumi licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
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
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/pulumi/lumi/pkg/diag"
)

// RunFunc wraps an error-returning run func with standard Lumi error handling.  All Lumi commands should wrap
// themselves in this to ensure consistent and appropriate error behavior.  In particular, we want to avoid any calls to
// os.Exit in the middle of a callstack which might prohibit reaping of child processes, resources, etc.  And we wish to
// avoid the default Cobra unhandled error behavior, because it is formatted incorrectly and needlessly prints usage.
func RunFunc(run func(cmd *cobra.Command, args []string) error) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := run(cmd, args); err != nil {
			ExitError(err.Error())
		}
	}
}

// ExitError issues an error and exits with a standard error exit code.
func ExitError(msg string, args ...interface{}) {
	ExitErrorCode(-1, msg, args...)
}

// ExitErrorCode issues an error and exists with the given error exit code.
func ExitErrorCode(code int, msg string, args ...interface{}) {
	Diag().Errorf(diag.Message(fmt.Sprintf(msg, args...)))
	os.Exit(code)
}
