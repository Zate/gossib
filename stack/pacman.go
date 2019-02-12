package stack

import (
	hlp "github.com/Zate/gossib/helper"
	// "github.com/acobaugh/osrelease"
)

func UpdateRepo(verbose bool) {
	hlp.RunExec("apt-get update", verbose)
}

func InstallPkg(pkg string, verbose bool) {
	UpdateRepo(verbose)
	hlp.RunExec("apt-get install -y "+pkg, verbose)
	hlp.RunExec("apt-get autoremove -y", verbose)
	hlp.RunExec("apt-get clean", verbose)
}

func RemovePkg(pkg string, verbose bool) {
	//
}
