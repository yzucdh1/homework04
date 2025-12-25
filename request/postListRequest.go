package request

import "github.com/yzucdh1/homework04/global"

type PostListReq struct {
	global.PageReq
	Title *string `json:"title"`
}
