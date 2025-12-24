package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yzucdh1/homework04/global"
	"github.com/yzucdh1/homework04/model"
	"github.com/yzucdh1/homework04/request"
	"github.com/yzucdh1/homework04/response"
	"net/http"
)

var PostGroup *gin.RouterGroup

func PostHandler() {
	PostGroup.Use(JWTTokenMiddleware)
	PostGroup.POST("/create", Create)
	PostGroup.POST("/list", List)
	PostGroup.POST("/detail", Detail)
	PostGroup.POST("/update", Update)
	PostGroup.POST("/delete", Delete)
}

func Create(c *gin.Context) {
	var req = request.PostReq{}
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

func List(c *gin.Context) {
	var req = request.ListReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		msg := GetValidMessage(err, &req)
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.ERROR, msg))
		return
	}

	var posts []response.Post
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

func Detail(c *gin.Context) {

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
