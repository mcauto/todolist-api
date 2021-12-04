package todo

// Params is a struct for the parameters of the todo module
type Params struct {
	Page  *int `query:"page"  json:"page,omitempty"  validate:"gte=1" default:"1"  example:"1"`
	Limit *int `query:"limit" json:"limit,omitempty" validate:"gte=1" default:"10" example:"10"`
}

// Pagination is a struct for the pagination of the todo module
func (p Params) Pagination() (page, limit int) {

	if p.Limit == nil {
		limit = 10
	} else {
		limit = *p.Limit
	}

	if p.Page == nil {
		page = 0
	} else {
		page = limit * (*p.Page - 1)
	}
	return page, limit
}

// Body is a struct for the body of the todo module
type Body struct {
	Title string `json:"title"`
}
