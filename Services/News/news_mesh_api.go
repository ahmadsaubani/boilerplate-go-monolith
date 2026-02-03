package News

import (
	"boilerplate-go/DTO/Articles"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

func (n NewsSetting) FetchNewsMesh() (res []Articles.NewsMeshItemDTO, err error) {
	baseURL := n.MeshNewsApi.Host
	apiToken := n.MeshNewsApi.ApiToken
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		allErr error
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

		q := url.Values{}
		q.Set("apiKey", apiToken)
		q.Set("country", "id")
		q.Set("limit", "25")

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
			allErr = fmt.Errorf("status %s", resp.Status)
			mu.Unlock()
			return
		}

		var result struct {
			Data []Articles.NewsMeshItemDTO `json:"data"`
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
	}()

	wg.Wait()

	if allErr != nil {
		return
	}

	return
}

func (n NewsSetting) NormalizeNewsMesh(v Articles.NewsMeshItemDTO) Articles.NormalizedArticleDTO {
	return Articles.NormalizedArticleDTO{
		Title:       v.Title,
		Description: v.Description,
		URL:         v.Link,
		ImageURL:    v.MediaURL,
		PublishedAt: v.PublishedDate,
		Source:      v.Source,
		Categories:  []string{v.Category},
	}
}
