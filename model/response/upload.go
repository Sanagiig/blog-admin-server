package response

type UploadAuthData struct {
	UpToken    string `json:"upToken"`
	Key        string `json:"key"`
	RefreshUrl string `json:"refreshUrl"`
}
