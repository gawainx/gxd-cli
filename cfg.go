/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-May-28
 *
*/

package main

import (
	"github.com/BurntSushi/toml"
	"log"
	"fmt"
    "encoding/json"
    "github.com/docker/docker/client"
    "context"
    "github.com/docker/go-connections/nat"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/api/types"
)

type PortInt int
func (pi PortInt) String() string{
    return fmt.Sprintf("%d",pi)
}

type TOMLConfig struct {
    Title string
    Net   NetworkConfig    `toml:"network"`
    Service ContainerConfig
}

// To create network
type NetworkConfig struct {
    Name string
}
func (n *NetworkConfig) String() string{
    return n.Name
}

// Config container
type ContainerConfig struct {
    Name 			string
    Image			string
    Detached 		bool
    WorkDir			string `toml:"work_dir"`
    CMD             string `toml:"cmd"`
    Net             string
    Ports 			[]Port
    Volumes			[]Vol
}

func (c *ContainerConfig) JSONStr() string{
    res,e := json.Marshal(c)
    if e != nil{
        return ""
    }else{
        return string(res)
    }
}

type Port struct {
    Host			PortInt
    Target			PortInt
}
func (p *Port) String()string{
	return fmt.Sprintf("%d:%d",p.Host,p.Target)
}

type Vol struct {
    Host 			string
    Target 			string
}
func(v *Vol) String()string{
	return fmt.Sprintf("%s:%s",v.Host,v.Target)
}


// Init toml from *.toml
func (t *TOMLConfig) InitFromFile(filename string){
    _,e := toml.DecodeFile(filename,t)
    if e != nil{
    	log.Fatal(e)
    }
}

func RunContainer(t *TOMLConfig){
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
    if err != nil {
        panic(err)
    }

    //set mount volumes
    vols := make([]string,len(t.Service.Volumes))
    for index,item := range t.Service.Volumes{
        vols[index] = item.String()
    }

    //set exposed ports for containers and publish ports
    exports := make(nat.PortSet)
    pts := make(nat.PortMap)
    for _,p := range t.Service.Ports{
        tmpPort, _ := nat.NewPort("tcp",p.Target.String())
        pb := make([]nat.PortBinding,0)
        pb = append(pb,nat.PortBinding{
            HostPort:p.Host.String(),
        })
        exports[tmpPort] = struct{}{}
        pts[tmpPort] = pb
    }

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image:t.Service.Image,
        ExposedPorts:exports,
        Cmd:[]string{t.Service.CMD},
        WorkingDir:t.Service.WorkDir,
    },&container.HostConfig{
        Binds:vols,
        PortBindings:pts,
        NetworkMode:container.NetworkMode(t.Service.Net),
    },nil,t.Service.Name)
    if err != nil{
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }else{
        log.Printf("Container %s is created and started.\n",resp.ID)
    }
}
