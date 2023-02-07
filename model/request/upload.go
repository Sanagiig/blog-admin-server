package request

type UploadAuthBody struct {
	UploadType string `form:"uploadType"`
	Username   string `form:"username"`
}

type UploadRefreshCDNBody struct {
	URLs []string `json:"urls"`
}
