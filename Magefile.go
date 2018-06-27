// +build mage

package main

import (
    "github.com/magefile/mage/sh"
    "fmt"
    "os"
    "path/filepath"
    "github.com/magefile/mage/mg"
)

const VERSION = "v0.3-alpha"    //version
const prefix = "gxd-cli" // app name
const path = "bin"	      // target path
var Default = Build

//Default Build
func Build(){
    mg.Deps(Linux)
    mg.Deps(Darwin)
}


//Build for linux
func Linux() {
    var e= make(map[string]string)
    e["CGO_ENABLE"] = "0"
    e["GOOS"] = "linux"
    e["GOARCH"] = "amd64"
    name := fmt.Sprintf("%s-linux-%s",prefix,VERSION)
    if err := os.Mkdir("bin", 0700); err != nil && !os.IsExist(err) {
        fmt.Errorf("failed to create %s: %v", "bin", err)
        os.Exit(1)
    }
    path := filepath.Join("bin",name)
    fmt.Println("Building app for linux...")
    err := sh.RunWith(e,"go","build","-o",path,"main.go","cfg.go","toml_tmpl.go","MultiCfg.go")
    if err!=nil{
        fmt.Println("Built failed.")
        fmt.Println(err)
    }else{
        fmt.Printf("Built successfully,output:%s\n",path)
    }
}

//Build for macOS
func Darwin() {
    name := fmt.Sprintf("%s-darwin-%s",prefix,VERSION)
    if err := os.Mkdir("bin", 0700); err != nil && !os.IsExist(err) {
        fmt.Errorf("failed to create %s: %v", "bin", err)
        os.Exit(1)
    }
    path := filepath.Join("bin",name)
    fmt.Println("Building app for osx...")
    err := sh.Run("go","build","-o",path,"main.go","cfg.go","toml_tmpl.go","MultiCfg.go")
    if err != nil{
        fmt.Println("Built failed.")
        fmt.Println(err)
    }else{
        fmt.Printf("Built successfully,output:%s\n",path)
    }
}

//clean all files in bin/
func Clean(){
    fmt.Printf("Cleaning %s/...",path)
    os.RemoveAll(path)
}

func InstallDep(){

}
