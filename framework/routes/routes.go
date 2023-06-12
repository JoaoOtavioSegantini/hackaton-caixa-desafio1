package routes

import (
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/hack-caixa/application/services"
	"github.com/hack-caixa/framework/middlewares"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {

	r.Use(middlewares.CORS())

	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})

	ratelimit := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	api := r.Group("/api/Simulacao")
	{
		api.POST("/", ratelimit, func(ctx *gin.Context) { services.MakeSimulation(ctx, db) })
	}

	return r
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}
