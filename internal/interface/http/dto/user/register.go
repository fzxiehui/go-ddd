package dto

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResp struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}
