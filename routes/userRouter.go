package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/to4to/Go-JWT-Auth/controllers"
	"github.com/to4to/Go-JWT-Auth/middleware"
)


func UserRoutes(incomingRoutes *gin.Engine){

incomingRoutes.Use(middleware.Authenticate())

incomingRoutes.GET("/users",controller.GetUsers())
incomingRoutes.GET("/users/:user_id",controller.GetUser())


}