package hello

import (
	"api-server/handler"

	"github.com/gin-gonic/gin"
)

//Hello 测试
func Hello(ctx *gin.Context) {
	handler.SendResponse(ctx, nil, map[string]interface{}{
		"id": 1,
	})
}
