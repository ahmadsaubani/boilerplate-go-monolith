package Articles

import (
	"boilerplate-go/DTO/Categories"
	"time"
)

type ArticleDTO struct {
	ID          int                      `json:"id"`
	UUID        string                   `json:"uuid"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	URL         string                   `json:"url"`
	ImageURL    string                   `json:"image_url"`
	Categories  []Categories.CategoryDTO `json:"categories"`
	PublishedAt time.Time                `json:"published_at"`
	Source      string                   `json:"source"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
}

type TheNewsItemDTO struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	ImageURL    string   `json:"image_url"`
	PublishedAt string   `json:"published_at"`
	Source      string   `json:"source"`
	Categories  []string `json:"categories"`
}

type NewsMeshItemDTO struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Link          string `json:"link"`
	MediaURL      string `json:"media_url"`
	PublishedDate string `json:"published_date"`
	Source        string `json:"source"`
	Category      string `json:"category"`
}

type NormalizedArticleDTO struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	ImageURL    string   `json:"image_url"`
	PublishedAt string   `json:"published_at"`
	Source      string   `json:"source"`
	Categories  []string `json:"categories"`
}
