// Copyright © 2019 Matt Konda <mkonda@jemurai.com>
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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionString = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the version of crush",
	Long:  `Displays the version of crush - code review helper`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CRUSH Version", versionString)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}