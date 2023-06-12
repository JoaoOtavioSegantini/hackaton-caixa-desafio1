package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hack-caixa/application/services"
	"github.com/hack-caixa/framework/middlewares"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {

	r.Use(middlewares.CORS())
	rateLimit := middlewares.RateLimit()

	api := r.Group("/api/Simulacao")
	{
		api.POST("/", rateLimit, func(ctx *gin.Context) { services.MakeSimulation(ctx, db) })
	}

	return r
}
