package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yzucdh1/homework04/global"
	"github.com/yzucdh1/homework04/model"
	"github.com/yzucdh1/homework04/request"
	"github.com/yzucdh1/homework04/response"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var UserGroup *gin.RouterGroup

func UserHandler() {
	UserGroup.POST("/register", Register)
	UserGroup.POST("/login", Login)
}

func Register(c *gin.Context) {
	var user request.User
	if err := c.ShouldBindJSON(&user); err != nil {
		msg := global.GetValidMessage(err, &user)
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.FAIL, msg))
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.FAIL, "密码加密异常"))
		return
	}

	dbUser := model.User{
		Username:  user.Username,
		Password:  string(hashedPassword),
		Email:     user.Email,
		SecretKey: uuid.New().String(),
	}

	if err := global.DB.Create(&dbUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("用户注册失败"))
		return
	}

	c.JSON(http.StatusOK, response.Ok(""))
}

func Login(c *gin.Context) {
	var loginReq request.LoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		msg := global.GetValidMessage(err, &loginReq)
		c.JSON(http.StatusBadRequest, response.ErrorWithCode(response.FAIL, msg))
		return
	}

	var storedUser model.User
	if err := global.DB.Where("username = ?", loginReq.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorWithCode(response.UNAUTHORIZED, "用户名或者密码错误"))
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginReq.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorWithCode(response.UNAUTHORIZED, "用户名或者密码错误"))
		return
	}

	// 生成JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(storedUser.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorWithCode(response.ERROR, "系统错误"))
		return
	}

	c.JSON(http.StatusOK, response.Ok(tokenString))
}
