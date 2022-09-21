package main

import (
	"github.com/hasanhakkaev/chgdns"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hasanhakkaev/chgdns/ui"
	"github.com/urfave/cli/v2"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Name = "chgdns"
	app.Version = version
	app.Authors = []*cli.Author{{
		Name:  "Hasan Hakkaev",
		Email: "me@hhakkaev.com",
	}}
	app.Usage = "Quickly set the DNS servers of your computer"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			EnvVars: []string{"GITHUB_TOKEN"},
			Name:    "list",
			Usage:   "List current DNS server settings",
			Aliases: []string{"t"},
		},
	}

	app.Action = func(c *cli.Context) error {
		log.SetFlags(0)
		f, err := tea.LogToFile("chgdns.log", "")
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
		defer func() { _ = f.Close() }()

		dnsServers := chgdns.DnsServers{}
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
		p := tea.NewProgram(ui.NewInitialModel(dnsServers), tea.WithAltScreen())
		if err = p.Start(); err != nil {
			return cli.Exit(err.Error(), 1)
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
