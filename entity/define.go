package entity

type (
	Page struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
	}
	Result struct {
		Data       interface{}
		Pagination struct {
			Total int64
			Page
		}
	}
	SelectInterface struct {
		Value   int64  `json:"value"`
		Name    string `json:"name"`
		Checked bool   `json:"checked"`
	}
)

func (p *Page) Offset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Limit * (p.Page - 1)
}
