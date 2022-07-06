package controller

import (
	"net/http"
	"replapp/logic"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func GetMarvalCharacter(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		return
	}
	offset := c.Query("offset")
	if offset == "" {
		return
	}
	marvals, err := logic.GetAllMarvalCharacter(name, offset)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, marvals)
}
