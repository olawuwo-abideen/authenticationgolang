package routes

import (
	controller "github.com/olawuwo-abideen/authenticationgolang/controllers"
	"github.com/olawuwo-abideen/authenticationgolang/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:id", controller.GetUser())
}
