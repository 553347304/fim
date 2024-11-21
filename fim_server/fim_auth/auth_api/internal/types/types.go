// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type OpenLoginInfoResponse struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Href string `json:"href"` // 跳转地址
}
