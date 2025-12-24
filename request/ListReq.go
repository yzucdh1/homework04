package request

import "github.com/yzucdh1/homework04/global"

type ListReq struct {
	global.PageReq
	Title *string `json:"title"`
}
