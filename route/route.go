package route

import (
	"github.com/gin-gonic/gin"
)

func Routes()  {
	route := gin.Default()
	route.GET("/", )

	route.Run(":8081")
}
