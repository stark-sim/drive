package protos

type UserAddReq struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone,omitempty"`
}
