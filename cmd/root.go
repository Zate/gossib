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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultConfigFile = "elsacp"
const defaultConfigPath = "/etc/elsacp/"

var (
	VERSION    string
	BUILD_DATE string
	BUILD_ARCH string
)

var rootCmd = &cobra.Command{
	Use:   "elsa-cli",
	Short: "A simplified Open Source Linux server administration tool",
	Long: "\nElsaCP is a robust and simplified Open Source Linux server administration tool." +
		"\nBuilt in Indonesia based on the author's experience in managing Linux servers." +
		"\n\nComplete documentation available at https://elsacp.github.io",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need at least one arg!")
			fmt.Println("Use `help` flag to see the available commands")
		}
	},
}

var cfgDir string
var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is "+defaultConfigPath+defaultConfigFile+".yaml)")
	viper.SetDefault("config", defaultConfigPath+defaultConfigFile+".yaml")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		if _, err := os.Stat(defaultConfigPath); os.IsNotExist(err) {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config file in etc directory
		viper.SetConfigName(defaultConfigFile)
		viper.AddConfigPath(defaultConfigPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func Execute(ver string, build string, arch string) {
	VERSION = ver
	BUILD_DATE = build
	BUILD_ARCH = arch
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
