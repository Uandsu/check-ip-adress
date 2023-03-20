package app

import "github.com/urfave/cli"

//Gerar vai retornar a aplicação
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Application Command Line"
	app.Usage = "Find Ips and Internet Server Names"

	return app

}
