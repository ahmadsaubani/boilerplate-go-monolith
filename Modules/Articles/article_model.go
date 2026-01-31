package Articles

import (
	"boilerplate-go/Config"
	"boilerplate-go/Config/DTO/ConfigStructs/ModuleConfigs"
	"boilerplate-go/DTO/Articles"
	"boilerplate-go/DTO/General"
	"boilerplate-go/Repositories"
)

type article struct {
	repo    Repositories.Repository
	newsApi Config.NewsConfig
}

type ArticleModules interface {
	GetAllArticleModule(params General.RequestParamDTO) (resp []Articles.ArticleDTO, total int, err error)
	FetchArticle() (err error)
}

func NewModule(moduleConfig ConfigStructs.ModuleConfigs, newsApi Config.NewsConfig) ArticleModules {
	return &article{
		repo:    moduleConfig.Repo,
		newsApi: newsApi,
	}
}
