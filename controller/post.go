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

type PostController struct{}

func (pc *PostController) PostCreate(c *gin.Context) {
	var req = request.PostCreateReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := GetValidMessage(err, &req)
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.FAIL, msg))
		return
	}
	// 获取用户ID
	userId := c.GetUint("userId")

	post := model.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userId,
	}
	if err := global.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "创建文章失败"+err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Ok(""))
}

func (pc *PostController) PostList(c *gin.Context) {
	var req = request.PostListReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := GetValidMessage(err, &req)
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.ERROR, msg))
		return
	}

	var posts []model.Post
	var result global.PageRes
	var err error
	var pageQuery = global.PageReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if req.Title != nil {
		result, err = global.Paginate(&posts, &pageQuery, "title like ?", "%"+*req.Title+"%")
	} else {
		result, err = global.Paginate(&posts, &pageQuery, nil, nil)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("查询失败："+err.Error()))
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, response.Ok(result))
}

func (pc *PostController) PostDetail(c *gin.Context) {
	var req request.PostDetailReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.FAIL, "参数错误"))
		return
	}

	var post model.Post
	if err := global.DB.Where("id = ?", req.ID).Preload("User").Take(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorWithCode(response.NOTFOUND, "文章信息不存在"))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "查询详情信息出错"+err.Error()))
		}
		return
	}
	c.JSON(http.StatusOK, response.Ok(post))
}

func (pc *PostController) PostUpdate(c *gin.Context) {
	var req request.PostUpdateReq
	if err := c.ShouldBind(&req); err != nil {
		msg := GetValidMessage(err, &req)
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.FAIL, msg))
		return
	}

	var post model.Post
	if err := global.DB.Where("id = ?", req.ID).Take(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorWithCode(response.NOTFOUND, "文章信息不存在"))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "查询文章信息出错"+err.Error()))
		}
		return
	}

	// 只有文章的作者才可以更新
	userId := c.GetUint("userId")
	if userId != post.UserID {
		c.JSON(http.StatusUnauthorized, response.ErrorWithCode(response.UNAUTHORIZED, "只有文章的作者才能更新自己的文章信息"))
		return
	}

	post.Title = req.Title
	post.Content = req.Content
	if err := global.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "修改文章信息出错"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Ok(""))
}

func (pc *PostController) PostDelete(c *gin.Context) {
	var req request.PostDetailReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.FAIL, "参数错误"))
		return
	}

	var post model.Post
	if err := global.DB.Where("id = ?", req.ID).Take(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, response.ErrorWithCode(response.NOTFOUND, "文章信息不存在"))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "查询文章信息出错"+err.Error()))
		}
		return
	}

	// 只有文章的作者才可以删除
	userId := c.GetUint("userId")
	if userId != post.UserID {
		c.JSON(http.StatusUnauthorized, response.ErrorWithCode(response.UNAUTHORIZED, "只有文章的作者才能删除自己的文章信息"))
		return
	}

	if err := global.DB.Select("Comments").Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "删除文章信息出错"+err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Ok(""))
}
