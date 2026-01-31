package General

type RequestParamDTO struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	SortBy  string `json:"sortBy"`
}
