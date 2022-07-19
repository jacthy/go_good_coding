package presentation

import (
	"awesomeProject/test-eventhorizon/presentation/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.New()
	r.POST("/v1/custom/create", api.CreateHandler)
	r.Run(":9009")
}
