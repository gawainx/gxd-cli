/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-May-30
 *
*/

package main

import (
    "github.com/BurntSushi/toml"
    "log"
)
type ContainerConfigs []ContainerConfig

type MultiTOMLConfig struct {
    Title           string
    Net             NetworkConfig    `toml:"network"`
    Services        ContainerConfigs
}

// Init toml from *.toml
func (m *MultiTOMLConfig) InitFromFile(filename string){
    _,e := toml.DecodeFile(filename,m)
    if e != nil{
        log.Fatal(e)
    }
}

func (m *MultiTOMLConfig) RunContainers(){
    for serIndex,service := range m.Services{
        log.Printf("Creating %d service...\n",serIndex)
        service.RunContainer()
    }
}
