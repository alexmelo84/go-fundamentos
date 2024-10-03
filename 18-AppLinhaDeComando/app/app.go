package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Gera a aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca de IPs e servidores na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca de IPs",
			Flags: flags,
			Action: buscaIPs,
		},
		{
			Name: "servidores",
			Usage: "Busca de servidores",
			Flags: flags,
			Action: buscaServidores,
		},
	}

	return app
}

func buscaIPs(c *cli.Context) {
	host := c.String("host")
	
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaServidores(c *cli.Context) {
	host := c.String("host")
	
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
