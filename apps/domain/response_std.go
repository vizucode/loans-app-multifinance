package domain

type Metadata struct {
	TotalItem     int64 `json:"total_item"`
	TotalPage     int   `json:"total_page"`
	TotalPageItem int   `json:"total_page_item"`
}

type ResponseJson struct {
	StatusCode string      `json:"status_code"`
	MetaData   Metadata    `json:"metadata"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
}
