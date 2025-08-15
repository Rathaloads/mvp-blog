package dto

// omitempty
type BasePageQuery struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}
