package main

import (
	"github.com/Zate/gossib/cmd"
	hlp "github.com/Zate/gossib/helper"
)

// This variables generated during build
const (
	VERSION    = ""
	BUILD_DATE = ""
	BUILD_ARCH = ""
)

func main() {
	hlp.CheckOS()
	hlp.CheckPrivileges()
	cmd.Execute(VERSION, BUILD_DATE, BUILD_ARCH)
}
