package dto

type FileCreate struct {
	Name     string `json:"name" validate:"required"`
	Path     string `json:"path" validate:"required"`
	Size     int64  `json:"size" validate:"required"`
	FileType string `json:"file_type" validate:"required"`
	Public   bool   `json:"public"`
}

type FileUpdate struct {
	Name     string `json:"name,omitempty"`
	FileType string `json:"file_type,omitempty"`
	Public   bool   `json:"public"`
}

type FileBasicInfo struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	FileType string `json:"file_type"`
	Size     int64  `json:"size"`
	Public   bool   `json:"public"`
}

type FileDetails struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path,omitempty"`
	FileType  string `json:"file_type"`
	Size      int64  `json:"size"`
	Public    bool   `json:"public"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
