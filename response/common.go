package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type PageResultV2 struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int64       `json:"page"`
	PageSize int64       `json:"pageSize"`
}
