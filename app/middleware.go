package app

import (
	"github.com/gin-gonic/gin"
	"github.com/horzions/pkg/middleware"
)

func setMiddlewares(r *gin.Engine) {
	r.Use(middleware.DefaultStructuredLogger())
	r.Use(gin.Recovery()) // adds the default recovery middleware
}
