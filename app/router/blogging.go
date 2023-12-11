package router

import (
	"eckert/domain/discord"
	"eckert/infrastructure/router"
	"github.com/gin-gonic/gin"
)

type Blogging string

func InitBlogging() {
	n := new(Blogging)
	r := router.RT.Group("/blogging")
	r.GET("/article_list", n.GetArticleList)
}

func (b *Blogging) GetArticleList(c *gin.Context) {
	discord.GetNews()
}
