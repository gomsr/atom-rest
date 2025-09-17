package utilr

import (
	"fmt"
	"github.com/kongmsr/oneid-core/modelo"
	"github.com/kongmsr/oneid-core/utilo"
	"testing"
)

func TestCreateToken(t *testing.T) {
	bc := modelo.BaseClaims{
		ID:       1,
		KeyID:    "string",
		Audience: []string{"string"},

		Type:     1,
		Nickname: "string",
		Gender:   "string",
		Cover:    "string",
		Phone:    "string",
		Email:    "string",
	}

	jwt := NewJWT("../source/cert/private_key.pem")

	token, err := jwt.CreateToken(jwt.CreateClaims(bc, "1d", "70d", "TEST"))
	if err != nil {
		return
	}
	fmt.Println(token)
}

func TestParser(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiS2V5SUQiOiJzYXMiLCJBdWRpZW5jZSI6WyIvb25laWQvYWFjcy8qIiwiL29uZWlkL2FhY3Mvd3BwLWFkbWluIiwiL29uZWlkL2FhY3Mvd3BwLWFkbWluL2EiXSwiU3ViamVjdCI6IndwcC1hZG1pbiIsIlZhbHVlIjpudWxsLCJUeXBlIjoxLCJOaWNrbmFtZSI6IuiMg-engOiLsSIsIkdlbmRlciI6IiIsIkNvdmVyIjoiaHR0cHM6Ly9xbXBsdXNpbWcuaGVucm9uZ3lpLnRvcC9ndmFfaGVhZGVyLmpwZyIsIlBob25lIjoiMTg2MDU1NDg4ODQiLCJFbWFpbCI6IiIsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJPTkVJRF9BQUNTIiwic3ViIjoid3BwLWFkbWluIiwiYXVkIjpbIi9vbmVpZC9hYWNzLyoiLCIvb25laWQvYWFjcy93cHAtYWRtaW4iLCIvb25laWQvYWFjcy93cHAtYWRtaW4vYSJdLCJleHAiOjE3MTg3MjEyNjMsIm5iZiI6MTcxMjY3MzI2M30.dNzrIlWf_FcoxCGLZW9f0qDt8HSqU1_xhzUnD48DpUC6ccBmvT9V1dNmLxqU6OOZYNccyadro71GbJuVTHmuX6f7FrjEkxr3xb4gcFuAie9vSBqDa5RxZY9WaqgG5GGtxozQOZn2DW6-LVudn1s_VLe5ieqvUCy5FZFPNXEw7aOWDKRjtsW-ZyWMrRZ4jvX2p_KKvg34IdF1cg_PSNM1UIXvJsCV8S-pIJZ_a3qjIaCuk-foEaJbGH_yCWkg7e_nWpeArhSKDs4cEfXdm2D7JMW_SLX4u86rRotXv_b-UaTUd7F1UaGFteap4tW9UFP7WhuM5qOnpnCBpbs7HyMKxQ"
	jwt, _ := utilo.GetParser("../source/cert/public_key.pem")
	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		return
	}

	fmt.Println(parseToken)

}
