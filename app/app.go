package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// gerar vai retornar o comando pronta para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplication search server"
	app.Usage = "Buscar ips e nomes de servidor na internet"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands = []cli.Command{
		{Name: "ip",
			Usage:  "Buscar ips e nomes de servidor na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "Servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}
func buscarIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println()
		fmt.Println(ip)
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
