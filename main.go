package main

import (
	"github.com/aSel1x/Gin_Template/api"
	"github.com/aSel1x/Gin_Template/core"
)

func main() {

	newConfig, _ := core.NewConfig()

	container := core.NewAppProvider(newConfig)

	r := api.SetupRouter(container)
	r.Run("0.0.0.0:8000")
}
