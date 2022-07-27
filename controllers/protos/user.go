package protos

type UserAddReq struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type UserGetReq struct {
	Id int64 `form:"id"`
}

type UserLoginReq struct {
	Phone    string `form:"phone" json:"phone"`
	Password string `form:"password" json:"password"`
}
