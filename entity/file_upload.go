package entity

type FileUploadMetaData struct {
	FileName  string `json:"file_name"`
	FileSize  int64  `json:"file_size"`
	FileURL   string `json:"file_url"`
	CreatedAt int64  `json:"created_at"`
}
