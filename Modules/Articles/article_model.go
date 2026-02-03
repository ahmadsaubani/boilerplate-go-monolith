package Articles

import (
	"boilerplate-go/Config"
	"boilerplate-go/Config/DTO/ConfigStructs/ModuleConfigs"
	"boilerplate-go/DTO/Articles"
	"boilerplate-go/DTO/General"
	"boilerplate-go/Repositories"
	"boilerplate-go/Services/News"
)

type article struct {
	repo        Repositories.Repository
	newsApi     Config.NewsConfig
	newsService News.NewsServices
}

type ArticleModules interface {
	GetAllArticleModule(params General.RequestParamDTO) (resp []Articles.ArticleDTO, total int, err error)
	FetchArticle() (err error)
	MapNormalizedArticle(v Articles.NormalizedArticleDTO) Articles.ArticleDTO
}

func NewModule(moduleConfig ConfigStructs.ModuleConfigs, newsApi Config.NewsConfig, newsService News.NewsServices) ArticleModules {
	return &article{
		repo:        moduleConfig.Repo,
		newsApi:     newsApi,
		newsService: newsService,
	}
}
