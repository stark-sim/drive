package protos

type CreateDirectoryReq struct {
	Name     string `json:"name" binding:"required"`
	ParentId string `json:"parent_id"`
	IsPublic bool   `json:"is_public"`
}

type UpdateDirectoryReq struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
	IsPublic string `json:"is_public"`
}

type GetDirectoryReq struct {
	Id       int64 `form:"id"`
	ParentId int64 `form:"parent_id"`
}

type DeleteDirectoryReq struct {
	Id string `json:"id" binding:"required"`
}
