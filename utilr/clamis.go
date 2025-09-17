package utilr

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kongmsr/oneid-core/modelo"
	"github.com/kongmsr/oneid-core/utilo"
)

func GetClaims(c *gin.Context, publicKeyPath string) (*modelo.CustomClaims, error) {
	token := c.Request.Header.Get("Authentication")
	if len(token) == 0 {
		return nil, errors.New("Authentication is empty")
	}

	if parser, err := utilo.GetParser(publicKeyPath); err != nil {
		return nil, err
	} else {
		return parser.ParseToken(token)
	}
}

// GetAccessKeyID 从Gin的Context中获取从jwt解析出来的用户ID
func GetAccessKeyID(c *gin.Context, publicKeyPath string) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c, publicKeyPath); err != nil {
			return ""
		} else {
			return cl.BaseClaims.KeyID
		}
	} else {
		waitUse := claims.(*modelo.CustomClaims)
		return waitUse.BaseClaims.KeyID
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户
func GetUserInfo(c *gin.Context, publicKeyPath string) *modelo.CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c, publicKeyPath); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*modelo.CustomClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context, publicKeyPath string) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c, publicKeyPath); err != nil {
			return ""
		} else {
			return cl.Nickname
		}
	} else {
		waitUse := claims.(*modelo.CustomClaims)
		return waitUse.Nickname
	}
}
