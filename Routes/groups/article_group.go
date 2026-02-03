package groups

import (
	"boilerplate-go/Controller/Articles"

	"github.com/gin-gonic/gin"
)

type iArticleGroup interface {
	ArticleApiGroup(group *gin.RouterGroup, api Articles.ArticleControllerInterface)
}

func (r articleGroup) ArticleApiGroup(group *gin.RouterGroup, api Articles.ArticleControllerInterface) {
	group.GET("/articles", api.GetAllArticle)
	group.GET("/article/pulls", api.FetchArticle)
}

type articleGroup struct{}

func newArticleRouterGroup() iArticleGroup {
	return &articleGroup{}
}
