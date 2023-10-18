package request

type ListRequest struct {
	Keyword  string `form:"keyword,omitempty"`
	PageNum  int32  `form:"page_num" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5"`
}

func (req ListRequest) Limit() int32 {
	if req.PageSize > 100 {
		req.PageSize = 100
	}
	if req.PageSize < 5 {
		req.PageSize = 5
	}
	return req.PageSize
}

func (req ListRequest) Offset() int32 {
	return (req.PageNum - 1) * req.Limit()
}
