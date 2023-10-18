package request

type ListRequest struct {
	Keyword  string `json:"keyword,omitempty"`
	PageNum  int32  `josn:"page_num,omitempty"`
	PageSize int32  `json:"page_size,omitempty"`
}

func (req ListRequest) Limit() int32 {
	if req.PageSize > 100 {
		req.PageSize = 100
	}
	if req.PageSize < 10 {
		req.PageSize = 10
	}
	return req.PageSize
}

func (req ListRequest) Offset() int32 {
	return (req.PageNum - 1) * req.Limit()
}
