package news_app

import (
	"eckert/domain/discord"
	"eckert/infrastructure/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

type News struct {
}

func InitNews() {
	//设置路由触发和定时触发
	n := new(News)
	r := router.RT.Group("/news")
	r.POST("/all", n.QuireNews)
}

func (s *News) QuireNews(c *gin.Context) {
	var user News
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	//discord.News2()
	//获取 新闻 获取 discord的 动漫 新闻 刷 新闻
	c.JSON(http.StatusCreated, nil)
}

func (s *News) GetNews(c *gin.Context) {
	discord.GetNews()
}
