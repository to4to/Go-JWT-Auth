package routes


import (
	 "github.com/to4to/Go-JWT-Auth/controller"
	"github.com/gin-gonic/gin"
	

)


func AuthRoutes(incomingRoutes *gin.Engine){
  
incomingRoutes.POST("user/signup",controller.Signup())
incomingRoutes.POST("user/login",controller.Login())



}