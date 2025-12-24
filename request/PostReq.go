package request

type PostReq struct {
	Title   string `json:"title" binding:"required" msg:"文章标题不能为空"`
	Content string `json:"content" binding:"required" msg:"文章内容不能为空"`
}
