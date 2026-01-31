package Articles

import (
	"boilerplate-go/Config"
	"boilerplate-go/DTO/Articles"
	"boilerplate-go/DTO/General"
)

type Repository interface {
	GetAllArticles(params General.RequestParamDTO) (resp []Articles.ArticleDTO, total int, err error)
	SaveArticles(items []Articles.ArticleDTO) (err error)
}

// CONSTRUCTOR STRUCT
type article struct {
	dbCon Config.DbConInterface
}

// CONSTRUCTOR FUNCTION FOR USER REPOSITORY
func RepositoryNew(dbCon Config.DbConInterface) Repository {
	return &article{
		dbCon: dbCon,
	}
}
