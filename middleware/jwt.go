package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gomsr/atom-rest/response"
	"github.com/kongmsr/oneid-core/utilo"
)

func JWTAuth(abPubKeyPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := RetrieveToken(c)
		if err != nil || len(token) == 0 {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		j, err := utilo.GetParserOfAbPath(abPubKeyPath)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

func RetrieveToken(c *gin.Context) (string, error) {
	// retrieve token from header
	token := c.Request.Header.Get("Authorization")

	if len(token) == 0 || len(token) < 7 {
		return "", errors.New("invalid access token")
	}

	// remove token Bearer prefix if necessary
	return utilo.RemoveBearer(token), nil
}
