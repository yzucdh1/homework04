package request

type CommentCreateReq struct {
	PostID  uint   `json:"post_id" binding:"required" msg:"文章ID不能为空"`
	Content string `json:"content" binding:"required" msg:"评论的内容不能为空"`
}
