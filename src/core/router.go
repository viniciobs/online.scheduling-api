package core

import (
	"github.com/gin-gonic/gin"
	user_api "github.com/online.scheduling-api/src/api"
)

func ConfigureRouter() *gin.Engine {
	router := gin.Default()

	configureUserRoutes(router)
	configureModalityRoutes(router)

	return router
}

func configureUserRoutes(router *gin.Engine) {
	router.GET("/api/users", user_api.GetAll)
	router.GET("/api/users/:id", user_api.GetById)
	router.POST("/api/users", user_api.Create)
	router.PATCH("/api/users/:id", user_api.Update)
	router.DELETE("/api/users/:id", user_api.Delete)
}

func configureModalityRoutes(router *gin.Engine) {
	// TODO
}
