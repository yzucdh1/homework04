package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yzucdh1/homework04/global"
	"github.com/yzucdh1/homework04/response"
	"net/http"
	"strings"
)

func JWTTokenMiddleware(c *gin.Context) {
	// 1. 从请求头获取Token
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, response.ErrorWithCode(response.UNAUTHORIZED, "请先登录"))
		c.Abort()
		return
	}

	// 2. 解析Token格式（切割 Bearer 前缀）
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusUnauthorized, response.ErrorWithCode(response.UNAUTHORIZED, "Token格式错误"))
		c.Abort()
		return
	}
	tokenString := parts[1] // 提取纯Token字符串

	// 3. 解析并验证Token
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否为HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名算法")
		}
		return []byte(global.Cfg.Server.SecretKey), nil
	})

	// 4. 处理验证错误
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorWithCode(response.UNAUTHORIZED, "Token验证失败："+err.Error()))
		c.Abort()
		return
	}

	// 5. 验证Token有效性
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, response.ErrorWithCode(response.UNAUTHORIZED, "Token无效"))
			c.Abort()
			return
		}
		c.Set("userId", uint(uint64(userId)))
	} else {
		c.JSON(http.StatusUnauthorized, response.ErrorWithCode(response.UNAUTHORIZED, "Token无效"))
		c.Abort()
		return
	}

	// 6. 验证通过，执行后续接口逻辑
	c.Next()
}
