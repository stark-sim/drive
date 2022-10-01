package protos

/*
PageReq 通用分页请求
 */
type PageReq struct {
	PageNo int32 `form:"page_no" json:"page_no"`
	PageSize int32 `form:"page_size" json:"page_size"`
}

/*
PageResponse 通用分页结果
 */
type PageResponse struct {
	PageNo int32
	PageSize int32
	Total int32
	PageCount int32
}