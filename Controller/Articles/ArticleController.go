package Articles

import (
	"boilerplate-go/Config/DTO/ConfigStructs/Utilities"
	"boilerplate-go/DTO/General"
	"boilerplate-go/Libraries/Helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleControllerInterface interface {
	GetAllArticle(g *gin.Context)
	FetchArticle(g *gin.Context)
}

func NewController(u ConfigStructs.Utilities) ArticleControllerInterface {
	return &article{u}
}

func (a article) GetAllArticle(g *gin.Context) {
	var params General.RequestParamDTO
	params.Page, _ = strconv.Atoi(g.DefaultQuery("page", "1"))
	params.PerPage, _ = strconv.Atoi(g.DefaultQuery("per_page", "10"))
	params.OrderBy = g.DefaultQuery("order_by", "desc")
	res, total, err := a.Modules.ArticleModule.GetAllArticleModule(params)
	if err != nil {
		Helpers.HttpResponseError(g, nil, 500)
		return
	}
	Helpers.HttpResponseSuccessWithPagination(g, res, total)
}

func (a article) FetchArticle(g *gin.Context) {
	err := a.Modules.ArticleModule.FetchArticle()
	if err != nil {
		Helpers.HttpResponseError(g, nil, 500)
		return
	}
	Helpers.HttpResponseSuccess(g, http.StatusText(http.StatusOK))
}
