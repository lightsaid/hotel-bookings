package reps

import "math"

type Metadata struct {
	CurrentPage  int32 `json:"current_page"`
	PageSize     int32 `json:"page_size"`
	FirstPage    int32 `json:"first_page"`
	LastPage     int32 `json:"last_page"`
	TotalRecords int64 `json:"total_records"`
}

func CalculateMetadata(totalRecords int64, PageNum, pageSize int32) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  PageNum,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int32(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}
