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

package helper

import (
	"fmt"
	"os"
	"runtime"

	"github.com/acobaugh/osrelease"
)

var (
	osList        = []string{"linux"}
	archList      = []string{"386", "amd64"}
	distroList    = []string{"Debian GNU/Linux", "Ubuntu"}
	debianVersion = []string{"9"}
	ubuntuVersion = []string{"16.04", "18.04"}
)

func CheckOS() {
	// Check OS type
	osType := runtime.GOOS
	osTypeCheck := inArray(osType, osList)
	if osTypeCheck == false {
		exitFailedCheck(osType)
	}

	// Check architecture
	osArch := runtime.GOARCH
	osArchCheck := inArray(osArch, archList)
	if osArchCheck == false {
		exitFailedCheck(osArch)
	}

	osrelease, err := osrelease.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Check distro name
	osName := osrelease["NAME"]
	osVersion := osrelease["VERSION_ID"]
	osNameCheck := inArray(osName, distroList)
	if osNameCheck == false {
		exitFailedCheck(osName)
	}

	// Check Debian version
	if osName == distroList[0] {
		versionCheck := inArray(osVersion, debianVersion)
		if versionCheck == false {
			exitFailedCheck(osName + " " + osVersion)
		}
	}

	// Check Ubuntu version
	if osName == distroList[1] {
		versionCheck := inArray(osrelease["VERSION_ID"], ubuntuVersion)
		if versionCheck == false {
			exitFailedCheck(osName + " " + osVersion)
		}
	}
}

func exitFailedCheck(msg string) {
	fmt.Printf("%v currently is not supported!\n", msg)
	os.Exit(0)
}
