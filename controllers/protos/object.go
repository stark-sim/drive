package protos

type ObjectCreateReq struct {
	DirectoryID string `json:"directory_id" binding:"required"`
	URL         string `json:"url" binding:"required"`
	Name        string `json:"name" binding:"required"`
	IsPublic    bool   `json:"is_public"`
}

type ObjectUpdateReq struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	IsPublic bool   `json:"is_public"`
}

type ObjectDeleteReq struct {
	ID string `json:"id" binding:"required"`
}
