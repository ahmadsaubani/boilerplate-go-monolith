package Articles

import (
	"boilerplate-go/DTO/Articles"
	"boilerplate-go/DTO/Categories"
	"boilerplate-go/DTO/General"
	"boilerplate-go/Libraries/Helpers"
	"strings"
	"time"
)

func (a article) GetAllArticleModule(params General.RequestParamDTO) (resp []Articles.ArticleDTO, total int, err error) {
	resp, total, err = a.repo.Article.GetAllArticles(params)
	if err != nil {
		return
	}
	return
}

func (a article) FetchArticle() (err error) {

	var articles []Articles.ArticleDTO
	items, err := a.newsService.FetchTheNews()
	if err != nil {
		return
	}

	for _, item := range items {
		normalized := a.newsService.NormalizeTheNews(item)
		articles = append(articles, a.MapNormalizedArticle(normalized))
	}

	meshItems, err := a.newsService.FetchNewsMesh()
	if err != nil {
		return
	}

	for _, itemTwo := range meshItems {
		normalized := a.newsService.NormalizeNewsMesh(itemTwo)
		articles = append(articles, a.MapNormalizedArticle(normalized))
	}

	if err = a.repo.Article.SaveArticles(articles); err != nil {
		return
	}

	return
}

func (a article) MapNormalizedArticle(v Articles.NormalizedArticleDTO) Articles.ArticleDTO {
	pubTime, _ := time.Parse(time.RFC3339, v.PublishedAt)

	cats := make([]Categories.CategoryDTO, 0, len(v.Categories))
	for _, c := range v.Categories {
		cats = append(cats, Categories.CategoryDTO{
			Label: Helpers.CapitalizeFirst(c),
			Value: strings.ToLower(c),
		})
	}

	return Articles.ArticleDTO{
		Title:       v.Title,
		Description: v.Description,
		URL:         v.URL,
		ImageURL:    v.ImageURL,
		PublishedAt: pubTime,
		Source:      v.Source,
		Categories:  cats,
	}
}
