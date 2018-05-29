/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-May-28
 *
*/

package main

import "flag"

var f = flag.String("f","service.toml","Set config *.toml file.")

func main(){
    flag.Parse()
    var tom = new(TOMLConfig)
    tom.InitFromFile(*f)
    RunContainer(tom)
}


