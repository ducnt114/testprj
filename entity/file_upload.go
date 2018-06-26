package entity

type FileUploadMetaData struct {
	ID       int64  `json:"id"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	FileURL  string `json:"file_url"`
}
