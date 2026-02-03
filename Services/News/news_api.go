package News

import (
	"boilerplate-go/DTO/Articles"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func (n NewsSetting) FetchTheNews() (res []Articles.TheNewsItemDTO, err error) {

	baseURL := n.TheNewsApi.Host
	apiToken := n.TheNewsApi.ApiToken
	totalPages := 3
	limit := 3

	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		allErr error
	)

	for page := 1; page <= totalPages; page++ {
		wg.Add(1)

		go func(p int) {
			defer wg.Done()

			q := url.Values{}
			q.Set("api_token", apiToken)
			q.Set("categories", "general,science,business,tech,sports,travel")
			q.Set("language", "en")
			q.Set("page", strconv.Itoa(p))
			q.Set("limit", strconv.Itoa(limit))

			reqURL := baseURL + "?" + q.Encode()

			resp, err := http.Get(reqURL)
			if err != nil {
				mu.Lock()
				allErr = err
				mu.Unlock()
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				mu.Lock()
				allErr = fmt.Errorf("page %d: status %s", p, resp.Status)
				mu.Unlock()
				return
			}

			var result struct {
				Data []Articles.TheNewsItemDTO `json:"data"`
			}

			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				mu.Lock()
				allErr = err
				mu.Unlock()
				return
			}

			mu.Lock()
			res = append(res, result.Data...)
			mu.Unlock()
		}(page)
	}

	wg.Wait()

	if allErr != nil {
		return
	}
	return
}

func (n NewsSetting) NormalizeTheNews(v Articles.TheNewsItemDTO) Articles.NormalizedArticleDTO {
	return Articles.NormalizedArticleDTO{
		Title:       v.Title,
		Description: v.Description,
		URL:         v.URL,
		ImageURL:    v.ImageURL,
		PublishedAt: v.PublishedAt,
		Source:      v.Source,
		Categories:  v.Categories,
	}
}
