package app

import (
	"github.com/gin-gonic/gin"
	"github.com/horzions/pkg/config"
)

type App struct {
	Config *config.Config
	Engine *gin.Engine
}

func NewApp(c *config.Config) *App {
	return &App{Engine: gin.New(), Config: c}
}
