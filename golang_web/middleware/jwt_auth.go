package middleware


import (
	"errors"
	"kubeclub/scaffold_go_web/pkg/tool"
	"kubeclub/scaffold_go_web/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("token") != "" {
			DeToken(c.Query("token"), c)
		} else {
			// 获取authorization header
			tokenString := c.GetHeader("token")
			// token 前缀为 jwt 用空格分开
			if tokenString == "" || !strings.HasPrefix(tokenString, "jwt") {
				c.JSON(http.StatusUnauthorized, gin.H{"errcode": 401, "errmsg": "未登录或非法访问"})
				c.Abort()
				return
			}
			tokenString = tokenString[4:] // 截取token 从jwt开始
			DeToken(tokenString, c)
		}
	}
}

// DeToken 解析token
func DeToken(t string, c *gin.Context) {

	token, claims, err := tool.ParseToken(t)

	if err != nil || !token.Valid {
		ResponseError(c, InvalidRequestErrorCode, errors.New("授权已过期"))
		c.Abort()
		return
	}

	// 验证通过之后 获取 claim中的userId
	userId := claims.ID
	//var user dao.User
	//common.GVA_DB.First(&user, userId)
	user, err := (&dao.User{}).Find(c, int64(userId))
	// 判断用户是否存在
	if err != nil &&  user.Name == ""{
		ResponseError(c, InvalidRequestErrorCode, errors.New("认证失败"))
		c.Abort()
		return
	}

	// 用户存在, 将用户的信息写入 context
	c.Set("user", user)
	c.Set("claims", claims)
	c.Next()
}
