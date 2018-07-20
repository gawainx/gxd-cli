/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-May-28
 *
 */

package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	//版本号
	app.Version = "1.0-rc"
	//作者信息
	app.Author = "Gawain Antarx"
	app.Email = "liangyixp@live.cn"
	app.Description = `pz-cli is PiZzahub' Docker CLIent.`
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "Create template service.toml",
			Action: func(c *cli.Context) error {
				WriteInitTOML()
				return nil
			},
		},
		{
			Name:  "up",
			Usage: "pz-cli up <toml file>:Run containers from *.toml",
			Action: func(c *cli.Context) error {
				var tom = new(MultiTOMLConfig)
				tom.InitFromFile(c.Args().First())
				if tom.Net.Name != "" {
					tom.CreateNet()
				}
				tom.RunContainers()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Panic(err)
	}
}
