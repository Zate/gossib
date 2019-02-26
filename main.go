package main

import (
	"github.com/Zate/gossib/cmd"
	hlp "github.com/Zate/gossib/helper"
)

// This variables generated during build
const (
	VERSION    = "0.0.1"
	BUILD_DATE = "Tue Feb 12 01:02:17 STD 2019"
	BUILD_ARCH = ""
)

func main() {
	hlp.CheckOS()
	hlp.CheckPrivileges()
	cmd.Execute(VERSION, BUILD_DATE, BUILD_ARCH)
}
