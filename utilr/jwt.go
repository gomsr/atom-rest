package utilr

import (
	"crypto/rsa"
	"github.com/kongmsr/oneid-core/modelo"
	"github.com/kongmsr/oneid-core/utilo"
	"golang.org/x/sync/singleflight"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey *rsa.PrivateKey
}

var (
	sg = &singleflight.Group{}
)

// NewJWT 创建一个新的jwt实例
func NewJWT(abPrivateKeyPath string) *JWT {
	if pk, err := utilo.ReadPriKeyOfAbPath(abPrivateKeyPath); err != nil {
		panic("读取私钥失败: " + err.Error())
	} else {
		return &JWT{pk}
	}
}

func (j *JWT) CreateClaims(baseClaims modelo.BaseClaims, bufferTime, expiresTime, issuer string) modelo.CustomClaims {
	bf, _ := ParseDuration(bufferTime)
	ep, _ := ParseDuration(expiresTime)
	claims := modelo.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  baseClaims.Audience,                       // 受众
			Subject:   baseClaims.Subject,                        // Subject(appid)
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间
			Issuer:    issuer,                                    // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims modelo.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims modelo.CustomClaims) (string, error) {
	v, err, _ := sg.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// Deprecated: ParseToken 解析 token(rsa 需要使用公钥进行签名验证) use ParserToken
// utilo.GetParserOfAbPath(abPubKeyPath).ParseToken(xx)
func (j *JWT) ParseToken(tokenString string) (*modelo.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &modelo.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, utilo.TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, utilo.TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, utilo.TokenNotValidYet
			} else {
				return nil, utilo.TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*modelo.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, utilo.TokenInvalid

	} else {
		return nil, utilo.TokenInvalid
	}
}

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
