package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Create() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI application"
	app.Usage = "Get IP and server names from the Internet"

	flags := []cli.Flag{ //Create a specific flag
		cli.StringFlag{
			Name:  "host",            //Flag name
			Value: "alwayscatch.com", //Flag default value
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Get IP from an internet address",
			Flags:  flags,
			Action: getIPs,
		},
		{
			Name:   "servers",
			Usage:  "Get servers from the internet",
			Flags:  flags,
			Action: getServers,
		},
	}
	return app
}

func getIPs(action *cli.Context) {
	host := action.String("host")
	ips, error := net.LookupIP(host) //Get IPs from the internet
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips { //List all public IPs found
		fmt.Println(ip)
	}
}

func getServers(action *cli.Context) {
	host := action.String("host")
	servers, error := net.LookupNS(host) //Get servers from the internet
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers { //List all public servers found
		fmt.Println(server.Host)
	}
}
