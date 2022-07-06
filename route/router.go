package route

import (
	ctrl "replapp/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", ctrl.Ping)
		v1.GET("/characters", ctrl.GetMarvalCharacter)
	}
	router.Run(":8080")
}
