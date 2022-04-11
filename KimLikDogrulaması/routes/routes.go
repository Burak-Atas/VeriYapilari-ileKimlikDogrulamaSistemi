package routes

import (
	"KimLikDogrulaması/controller"

	"github.com/gin-gonic/gin"
)

func GirisRoutes(rts *gin.Engine) {
	rts.GET("/signup", controller.SignUp())
	rts.GET("/login", controller.LogIn())
	rts.GET("/users", controller.GetUSers())
}
