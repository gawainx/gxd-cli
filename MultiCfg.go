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
//用于配置同时启动多个容器
type ContainerConfigs []ContainerConfig

//  Title    配置名称
//	Net      网络名称(设定是否新建网络)
//	Services 容器配置文件集合
//	ctx      上下文信息
//	cli      用于连接 docker 的客户端载体
type MultiTOMLConfig struct {
	Title    string
	Net      NetworkConfig `toml:"network"`
	Services ContainerConfigs
	ctx      context.Context
	cli      *client.Client
}

// Init toml from *.toml
//InitFromFile: 从配置文件中读取信息
//filename: 文件名
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

//MultiTOMLConfig.RunContainers:运行多容器
func (m *MultiTOMLConfig) RunContainers() {
	for serIndex, service := range m.Services {
		log.Printf("Creating %d service...\n", serIndex)
		//对每个服务调用 RunContainer 方法
		service.RunContainer(m.cli,m.ctx)
	}
}
//MultiTOMLConfig.CreateNet 创建网络
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
    //创建网络并输出创建结果
    response,err := m.cli.NetworkCreate(m.ctx,m.Net.Name,types.NetworkCreate{})
    if err != nil{
        fmt.Println(err)
    }else{
        fmt.Printf("Created Network %s with warning %s.\n",response.ID,response.Warning)
    }

}
