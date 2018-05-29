/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-May-29
 *
*/

package main

import (
    "os"
    "fmt"
    "io"
)

var toml_tmpl = `## toml file below equals to the command 
## "docker run -d -v /go/src/gxd-cli:/code -w /code --net iot -p 8974:80 -p 8996:8080 alpine ./gin-server"
title = "TOML test"
[service]
 detach = false
 name = "test" # container name
 image = "alpine" 
 work_dir = "/code"
 net = "iot" # container network
 cmd = "./gin-server"
 [[service.ports]]
  host = 8974
  target = 80
 [[service.ports]]
  host = 8996
  target = 8080

 [[service.volumes]]
  host = "/go/src/gxd-cli/"
  target = "/code"
`

func checkFileIsExist(filename string) bool {
    var exist = true
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false
    }
    return exist
}


func WriteInitTOML() error{
    var filename = `service.toml`
    if checkFileIsExist(filename){
        fmt.Println("File service.toml exists.")
        return nil
    }else{
        f,e := os.Create(filename)
        if e != nil{
            return e
        }
        _,err := io.WriteString(f,toml_tmpl)
        if err!= nil{
            return err
        }else{
            fmt.Println("Successfully create service.toml")
            return nil
        }
    }
}