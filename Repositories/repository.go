package Repositories

import (
	"boilerplate-go/Config"
	"boilerplate-go/Repositories/Articles"
)

// CONSTRUCTOR STRUCT FOR ALL REPOSITORY
type Repository struct {
	Article Articles.Repository
}

// REPOSITORY INITIALIZATION
func InitRepo(dbCon Config.DbConInterface) Repository {
	return Repository{
		Article: Articles.RepositoryNew(dbCon),
	}
}
