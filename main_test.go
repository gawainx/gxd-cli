/*
 *Gawain Open Source Project
 *Author: Gawain Antarx
 *Create Date: 2018-May-29
 *
*/

package main

import "testing"

func TestRunContainer(t *testing.T) {
    var tom = new(TOMLConfig)
    tom.InitFromFile("service.toml")
    RunContainer(tom)
}

func TestWriteInitTOML(t *testing.T) {
    WriteInitTOML()
}