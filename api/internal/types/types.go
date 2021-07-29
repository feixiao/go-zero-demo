// Code generated by goctl. DO NOT EDIT.
package types

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data:omitempty"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type GetUserRequest struct {
	UserID string `path:"id"`
}

type GetUserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}
