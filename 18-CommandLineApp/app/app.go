package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Returns command line applciation ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e nomes de servidor"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços da internet",
			Flags:  flags,
			Action: getIPs,
		},
		{
			Name:   "servers",
			Usage:  "Get servers' name",
			Flags:  flags,
			Action: getServers,
		},
	}

	return app
}

func getIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
