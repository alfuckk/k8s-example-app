package app

import (
	"github.com/gin-gonic/gin"
	"github.com/horzions/pkg/config"
	"github.com/horzions/pkg/serve"
)

type App struct {
	Config *config.Config
	Engine *gin.Engine
}

func Start() {
	config := config.NewYamlConfig("./app.yaml")
	app := NewApp(config)
	setMiddlewares(app.Engine)
	setRoutes(app)
	serve.NewServe(&config.Server, app.Engine).Run()
}

func NewApp(c *config.Config) *App {
	return &App{Engine: gin.New(), Config: c}
}
