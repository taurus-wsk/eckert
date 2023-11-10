package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

var RT *gin.Engine

func InitRouter() {
	RT = gin.Default()
}
