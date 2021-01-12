package dto

import "github.com/gustvision/backend-interview/pkg/user"

type GetUserReq struct {
	ID string
}

type GetUserResp struct {
	user.User

	Total float64
}
