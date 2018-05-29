/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-May-28
 *
*/

package main

import (
    "github.com/urfave/cli"
    "os"
    "log"
)

func main(){
    app := cli.NewApp()
    app.Commands = []cli.Command{
        {
            Name:  "init",
            Usage: "Create template service.toml",
            Action: func(c *cli.Context) error{
                WriteInitTOML()
                return nil
            },
        },
        {
            Name:  "up",
            Usage: "gxd-cli up <toml file>:Run container from *.toml",
            Action: func(c *cli.Context) error{
                var tom = new(TOMLConfig)
                tom.InitFromFile(c.Args().First())
                RunContainer(tom)
                return nil
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}


