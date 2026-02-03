package News

import (
	"boilerplate-go/Config"
	"boilerplate-go/DTO/Articles"
)

type NewsSetting struct {
	TheNewsApi  *Config.NewsConfig
	MeshNewsApi *Config.MeshNewsConfig
}

type NewsServices interface {
	FetchTheNews() (res []Articles.TheNewsItemDTO, err error)
	NormalizeTheNews(v Articles.TheNewsItemDTO) Articles.NormalizedArticleDTO
	FetchNewsMesh() (res []Articles.NewsMeshItemDTO, err error)
	NormalizeNewsMesh(v Articles.NewsMeshItemDTO) Articles.NormalizedArticleDTO
}
