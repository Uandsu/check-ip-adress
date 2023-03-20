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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Find IPS",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: findIps,
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
