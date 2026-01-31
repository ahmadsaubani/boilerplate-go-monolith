package Articles

import (
	"boilerplate-go/DTO/Articles"
	"boilerplate-go/DTO/Categories"
	"boilerplate-go/DTO/General"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
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

	baseURL := a.newsApi.Host
	apiToken := a.newsApi.ApiToken
	if apiToken == "" {
		return errors.New("API_TOKEN is missing in environment")
	}

	totalPages := 3
	limit := 3

	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		allErr error
	)

	for page := 1; page < totalPages; page++ {
		wg.Add(1)

		go func(currentPage int) {
			defer wg.Done()

			q := url.Values{}
			q.Set("api_token", apiToken)
			q.Set("categories", "general,science,business,tech,sports,travel")
			q.Set("language", "en")
			q.Set("page", strconv.Itoa(currentPage))
			q.Set("limit", strconv.Itoa(limit))

			reqURL := baseURL + "?" + q.Encode()

			resp, e := http.Get(reqURL)
			if e != nil {
				mu.Lock()
				allErr = e
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				mu.Lock()
				allErr = fmt.Errorf("page %d: status %s", currentPage, resp.Status)
				mu.Unlock()
				return
			}

			var apiResp struct {
				Data []struct {
					Title       string   `json:"title"`
					Description string   `json:"description"`
					URL         string   `json:"url"`
					ImageURL    string   `json:"image_url"`
					PublishedAt string   `json:"published_at"`
					Source      string   `json:"source"`
					Categories  []string `json:"categories"`
				} `json:"data"`
			}

			if e := json.NewDecoder(resp.Body).Decode(&apiResp); e != nil {
				mu.Lock()
				allErr = e

				mu.Unlock()
				return
			}

			var temp []Articles.ArticleDTO
			for _, v := range apiResp.Data {
				pubTime, _ := time.Parse(time.RFC3339, v.PublishedAt)

				var cats []Categories.CategoryDTO
				for _, c := range v.Categories {
					cats = append(cats, Categories.CategoryDTO{
						Label: c,
						Value: strings.ToLower(c),
					})
				}

				temp = append(temp, Articles.ArticleDTO{
					Title:       v.Title,
					Description: v.Description,
					URL:         v.URL,
					ImageURL:    v.ImageURL,
					PublishedAt: pubTime,
					Source:      v.Source,
					Categories:  cats,
				})
			}

			mu.Lock()
			articles = append(articles, temp...)
			mu.Unlock()
		}(page)
	}

	wg.Wait()

	if allErr != nil {
		return allErr
	}

	err = a.repo.Article.SaveArticles(articles)
	if err != nil {
		return
	}

	return
}
