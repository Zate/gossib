/*
MIT License

Copyright (c) 2018 Aris Ripandi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package cmd

import (
	"fmt"
	"os"

	stack "github.com/Zate/gossib/stack"
	"github.com/spf13/cobra"
)

var Verbose bool

func init() {
	rootCmd.AddCommand(stackCmd)
	stackCmd.AddCommand(subStackAll)
	stackCmd.AddCommand(subStackNginx)
	stackCmd.AddCommand(subStackMySQL)
	stackCmd.AddCommand(subStackPgSQL)
	stackCmd.AddCommand(subStackPhp)
	stackCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

var stackCmd = &cobra.Command{
	Use:              "stack",
	TraverseChildren: true,
	Short:            "Install application stack",
	Args:             cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need at least one arg. Type `elsa-cli stack -h` to see the available args")
			os.Exit(1)
		}
	},
}

var subStackAll = &cobra.Command{
	Use:   "all",
	Short: "Install full stack packages",
	Run: func(cmd *cobra.Command, args []string) {
		subStackNginx.Execute()
		subStackMySQL.Execute()
		subStackPgSQL.Execute()
		subStackPhp.Execute()
	},
}

var subStackNginx = &cobra.Command{
	Use:   "nginx",
	Short: "Install Nginx packages for static sites",
	Run: func(cmd *cobra.Command, args []string) {
		stack.InstallNginx(Verbose)
	},
}

var subStackMySQL = &cobra.Command{
	Use:   "mysql",
	Short: "Install MySQL packages",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO")
	},
}

var subStackPgSQL = &cobra.Command{
	Use:   "pgsql",
	Short: "Install PostgreSQL packages",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO")
	},
}

var subStackPhp = &cobra.Command{
	Use:   "php",
	Short: "Install PostgreSQL packages",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO")
	},
}
