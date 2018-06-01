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
[[services]]
 detach = false
 name = "tg" # container name
 image = "alpine" 
 work_dir = "/code"
 net = "iot" # container network
 cmd = "echo hello"
 [[services.ports]]
  host = 8974
  target = 80
 [[services.ports]]
  host = 8996
  target = 8080

 [[services.volumes]]
  host = "pwd"
  target = "/code"

 [[services.volumes]]
  host = "pwd/tmp"
  target = "/cg"
 [[services]]
 name = "tb"
 image = "alpine"
 cmd = "echo hello world."
`

func checkFileIsExist(filename string) bool {
    var exist = true
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false
    }
    return exist
}


func WriteInitTOML() error{
    var filename = `services.toml`
    if checkFileIsExist(filename){
        fmt.Println("File services.toml exists.")
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