package types

type MsgResponse struct {
    Status      string  `json:"status"`
    StatusCode  int     `json:"statusCode"`
	Message     string  `json:"message"`
}

type Result struct {
    Status      string      `json:"status"`
    StatusCode  int         `json:"statusCode"`
    Data        interface{} `json:"data"`
}

type Filter struct {
    Page int
    Limit int
    SortBy string
    SortDir string
}
