package main

import (
	"flag"
	"fmt"
	"github.com/pritunl/pritunl-link/cmd"
	"github.com/pritunl/pritunl-link/logger"
	"github.com/pritunl/pritunl-link/requires"
)

const help = `
Usage: pritunl-link COMMAND

Commands:
  start                     Start link service
  add                       Add a Pritunl server URI
  remove                    Remove a Pritunl server URI
  clear                     Clear all configured Pritunl server URIs
  local-address             Manually set local IP address
  public-address            Manually set public IP address
  verify-on                 Enable HTTPS certificate verification when connecting to Pritunl server
  verify-off                Disable HTTPS certificate verification when connecting to Pritunl server
  disconnected-timeout-on   Enable restart when disconnected for duration of timeout
  disconnected-timeout-off  Disable restart when disconnected for duration of timeout
  provider                  Manually set network provider
  unifi-username            Set Unifi username
  unifi-password            Set Unifi password
  unifi-controller          Set URL of Unifi controller
  unifi-site                Set the Unifi site if different then default
`

func Init() {
	logger.Init()
	requires.Init()
}

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "start":
		Init()
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		break
	case "add":
		Init()
		err := cmd.Add(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "remove":
		Init()
		err := cmd.Remove(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "clear":
		Init()
		err := cmd.Clear()
		if err != nil {
			panic(err)
		}
		break
	case "local-address":
		Init()
		err := cmd.LocalAddress(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "public-address":
		Init()
		err := cmd.PublicAddress(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "verify-on":
		Init()
		err := cmd.VerifyOn()
		if err != nil {
			panic(err)
		}
		break
	case "verify-off":
		Init()
		err := cmd.VerifyOff()
		if err != nil {
			panic(err)
		}
		break
	case "disconnected-timeout-on":
		Init()
		err := cmd.DisconnectedTimeoutOn()
		if err != nil {
			panic(err)
		}
		break
	case "disconnected-timeout-off":
		Init()
		err := cmd.DisconnectedTimeoutOff()
		if err != nil {
			panic(err)
		}
		break
	case "provider":
		Init()
		err := cmd.Provider(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "unifi-username":
		Init()
		err := cmd.UnifiUsername(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "unifi-password":
		Init()
		err := cmd.UnifiPassword(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "unifi-controller":
		Init()
		err := cmd.UnifiController(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "unifi-site":
		Init()
		err := cmd.UnifiSite(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	default:
		fmt.Println(help)
	}
}
