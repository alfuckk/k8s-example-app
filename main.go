package main

import (
	"github.com/horzions/pkg/config"
	"github.com/horzions/pkg/serve"
)

func main() {
	config := config.NewYamlConfig("./app.yaml")
	app := NewApp(config)
	setMiddlewares(app.Engine)
	setRoutes(app)
	serve.NewServe(&config.Server, app.Engine).Run()
}
