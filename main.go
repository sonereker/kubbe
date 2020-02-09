package main

import (
	"github.com/sonereker/kubbe/app"
	"github.com/sonereker/kubbe/config"
)

func main() {
	c := config.GetConfig()

	a := &app.App{}
	a.Init(c)
	a.Run(c.App.Port)
}
