package dto

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	UserID string `json:"user_id"`
	Token  string `json:"token"`
}
