package helper

import (
    "os"
    "fmt"
    "runtime"
    "github.com/acobaugh/osrelease"
)

var (
	osList          = []string{"linux"}
	archList        = []string{"386", "amd64"}
	distroList      = []string{"Debian GNU/Linux", "Ubuntu"}
	debianVersion   = []string{"8", "9"}
	ubuntuVersion   = []string{"16.04", "18.04"}
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
