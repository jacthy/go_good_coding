package api

import (
	"awesomeProject/test-eventhorizon/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHandler(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	println(string(data))
	result := application.HandCustomCommand(data, "createCustom")
	c.JSON(http.StatusOK, result)
}
