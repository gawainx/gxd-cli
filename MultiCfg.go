/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-May-30
 *
 */

package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/docker/docker/client"
	"context"
    "github.com/docker/docker/api/types"
    "fmt"
)

//ContainerConfigs : Multi-Container Configs
type ContainerConfigs []ContainerConfig

type MultiTOMLConfig struct {
	Title    string
	Net      NetworkConfig `toml:"network"`
	Services ContainerConfigs
	ctx      context.Context
	cli      *client.Client
}

// Init toml from *.toml
func (m *MultiTOMLConfig) InitFromFile(filename string) {
	_, e := toml.DecodeFile(filename, m)
	if e != nil {
		log.Fatal(e)
	}
    m.ctx = context.Background()
    cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
    if err != nil {
        panic(err)
    }else{
        m.cli = cli
    }
}

func (m *MultiTOMLConfig) RunContainers() {
	for serIndex, service := range m.Services {
		log.Printf("Creating %d service...\n", serIndex)
		service.RunContainer(m.cli,m.ctx)
	}
}

func (m *MultiTOMLConfig) CreateNet(){
    netList,e := m.cli.NetworkList(m.ctx,types.NetworkListOptions{})
    if e!= nil{
        fmt.Println(e)
    }
    //Check if network exists.
    for _,item := range netList{
        if item.Name == m.Net.Name{
            fmt.Printf("Network %s exists.\n",m.Net.Name)
            return
        }
    }
    response,err := m.cli.NetworkCreate(m.ctx,m.Net.Name,types.NetworkCreate{})
    if err != nil{
        fmt.Println(err)
    }else{
        fmt.Printf("Created Network %s with warning %s.\n",response.ID,response.Warning)
    }

}
