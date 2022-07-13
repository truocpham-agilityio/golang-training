package pagination

import (
	"net/http"
	"strconv"
)

// Pagination is a struct that holds pagination information.
type Pagination struct {
	Limit      int           `json:"limit,omitempty;query:limit"`
	Page       int           `json:"page,omitempty;query:page"`
	Sort       string        `json:"sort,omitempty;query:sort"`
	TotalRows  int64         `json:"total_rows"`
	TotalPages int           `json:"total_pages"`
	Rows       []interface{} `json:"rows"`
}

// GetOffset returns the offset of the pagination.
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetLimit returns the limit of the pagination.
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		return 10
	}

	return p.Limit
}

// GetPage returns the page of the pagination.
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		return 1
	}

	return p.Page
}

// GetSort returns the sort term of the pagination.
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	}

	return p.Sort
}

// GeneratePaginationFromRequest generates pagination information from a request.
func GeneratePaginationFromRequest(r *http.Request) Pagination {
	limit := 10
	page := 1
	sort := "created_at desc"
	query := r.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}

	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
