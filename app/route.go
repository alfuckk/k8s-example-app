package app

import (
	"github.com/horzions/k8s-example-app/app/account"
	"github.com/horzions/pkg/database"
	"github.com/horzions/pkg/middleware"
)

func setRoutes(app *App) {
	v1 := app.Engine.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			accountModule := account.NewAccount(app.Config, app.Engine, database.NewDB(&app.Config.Database))
			auth.POST("/login", accountModule.Login)
			auth.POST("/register", accountModule.Register)
			auth.POST("/forget", accountModule.ResetPassword)
		}

		accountV1 := v1.Group("/account")
		accountV1.Use(middleware.Jwt(&app.Config.Server))
		{
			accountModule := account.NewAccount(app.Config, app.Engine, database.NewDB(&app.Config.Database))
			accountV1.POST("/add", accountModule.AddAccount)
			accountV1.POST("/modify", accountModule.ModifyAccount)
			accountV1.POST("/delete", accountModule.DeleteAccount)
			accountV1.POST("/info", accountModule.AccountInfo)
		}
	}
}
