package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Application Command Line"
	app.Usage = "Find Ips and Internet Server Names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find IPS",
			Flags:  flags,
			Action: findIps,
		},
		{
			Name:   "Servers",
			Usage:  "Find Internet Server Names",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app

}

func findIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host) // name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
