package route

import (
	"github.com/Iretoms/hng-task-one/controller"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	r.GET("/hello", controller.HelloCall())
}