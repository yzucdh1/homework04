package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yzucdh1/homework04/global"
	"github.com/yzucdh1/homework04/model"
	"github.com/yzucdh1/homework04/request"
	"github.com/yzucdh1/homework04/response"
	"gorm.io/gorm"
	"net/http"
)

type CommentController struct{}

func (cc *CommentController) CommentCreate(c *gin.Context) {
	var req request.CommentCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := GetValidMessage(err, &req)
		c.JSON(http.StatusCreated, response.ErrorWithCode(response.FAIL, msg))
		return
	}

	var post model.Post
	if err := global.DB.Where("id = ?", req.PostID).Take(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorWithCode(response.NOTFOUND, "文章信息不存在"))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "查询文章信息出错"+err.Error()))
		}
		return
	}

	userId := c.GetUint("userId")
	comment := model.Comment{
		Content: req.Content,
		PostID:  post.ID,
		UserID:  userId,
	}

	if err := global.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "评论发表失败"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Ok(""))
}

func (cc *CommentController) CommentList(c *gin.Context) {
	var req request.CommentListReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.FAIL, "参数错误"))
		return
	}

	var post model.Post
	if err := global.DB.Where("id = ?", req.PostID).Take(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorWithCode(response.NOTFOUND, "文章信息不存在"))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "查询文章信息出错"+err.Error()))
		}
		return
	}

	var comments []model.Comment
	if err := global.DB.Where("post_id = ?", req.PostID).Preload("User").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "查询评论列表出错"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Ok(comments))
}
