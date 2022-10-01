package protos

import "time"

type CreateDirectoryReq struct {
	Name     string `json:"name" binding:"required"`
	ParentID string `json:"parent_id"`
	IsPublic bool   `json:"is_public"`
}

type UpdateDirectoryReq struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
	IsPublic bool `json:"is_public"`
}

type GetDirectoryReq struct {
	PageReq PageReq `form:"page_req"`
	ID       int64 `form:"id"`
	ParentID int64 `form:"parent_id"`
}

type DeleteDirectoryReq struct {
	ID string `json:"id" binding:"required"`
}

type DirectoryWithContentResp struct {
	ID       string              `json:"id"`
	Name     string              `json:"name"`
	ParentID string              `json:"parent_id"`
	IsPublic bool                `json:"is_public"`
	Contents []*DirectoryContent `json:"contents"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type DirectoryContent struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    string    `json:"user_id"`
	IsPublic  bool      `json:"is_public"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}