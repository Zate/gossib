package stack

import (
  "fmt"
  "github.com/acobaugh/osrelease"
  hlp "github.com/elsacp/elsa-cli/helper"
)

func SetupNginxRepo(verbose bool) {
  os, err := osrelease.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
  }

  osName := os["NAME"]
  switch osName {
    case "Debian GNU/Linux":
      osCode := hlp.RunExecV("dpkg --status tzdata|grep Provides|cut -f2 -d'-'")
      content := "deb http://nginx.org/packages/debian/ " + osCode + " nginx"
      hlp.RunExec("curl -sS https://nginx.org/keys/nginx_signing.key | apt-key add -", verbose)
      hlp.WriteFile("/etc/apt/sources.list.d/nginx.list", content)
    case "Ubuntu":
      content := `deb http://nginx.org/packages/ubuntu/ ` + os["VERSION_CODENAME"] + ` nginx`
      hlp.RunExec("curl -sS https://nginx.org/keys/nginx_signing.key | apt-key add -", verbose)
      hlp.WriteFile("/etc/apt/sources.list.d/nginx.list", content)
    default:
      fmt.Printf("Unsupported operating system\n", osName)
  }
}

func InstallNginx(verbose bool) {
  if verbose == true {
    c := hlp.ConfirmAction("Confirm to install nginx?")
    if c {
      SetupNginxRepo(verbose)
      // InstallPkg("nginx", verbose)
    }
  } else {
      SetupNginxRepo(verbose)
      // InstallPkg("nginx", verbose)
  }
}
