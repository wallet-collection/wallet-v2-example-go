package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet/api/res"
	"wallet/pkg/util"
)

// AuthRequired 认证中间件
func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("x-token")
		aid, err := util.VerifyToken(token)
		if err != nil {
			c.Abort()
			res.APIResponse(c, res.ErrToken, nil)
			return
		}
		c.Set("aid", aid)
	}

}

// AuthId 认证中间件
func AuthId() gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("x-token")
		aid, err := util.VerifyToken(token)
		if err != nil {
			aid = int64(0)
		}
		c.Set("aid", aid)
	}

}

// Cors 跨域设置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, x-token")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
