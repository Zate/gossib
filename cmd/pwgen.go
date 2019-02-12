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
	"strconv"

	hlp "github.com/Zate/gossib/helper"
	"github.com/spf13/cobra"
)

var SafeUrl bool

func init() {
	rootCmd.AddCommand(pwgenCmd)
	pwgenCmd.PersistentFlags().BoolVarP(&SafeUrl, "safe-url", "s", false, "return a URL-safe, base64 encoded")
}

var pwgenCmd = &cobra.Command{
	Use:   "pwgen",
	Short: "Generate random password",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if SafeUrl {
				PrintPwgenSafe(i)
			} else {
				PrintPwgen(i)
			}
		} else {
			if SafeUrl {
				PrintPwgenSafe(8)
			} else {
				PrintPwgen(8)
			}
		}
	},
}

func PrintPwgen(n int) {
	p, _ := hlp.StrRand(n)
	fmt.Println(p)
}

func PrintPwgenSafe(n int) {
	p, _ := hlp.StrRandURLSafe(n)
	fmt.Println(p)
}
