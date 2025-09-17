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

	jwt := NewJWT("./private_key.pem")

	token, err := jwt.CreateToken(jwt.CreateClaims(bc, "1d", "70d", "TEST"))
	if err != nil {
		return
	}
	fmt.Println(token)
}

func TestParser(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiS2V5SUQiOiJzdHJpbmciLCJBdWRpZW5jZSI6WyJzdHJpbmciXSwiU3ViamVjdCI6IiIsIlZhbHVlIjpudWxsLCJUeXBlIjoxLCJOaWNrbmFtZSI6InN0cmluZyIsIkdlbmRlciI6InN0cmluZyIsIkNvdmVyIjoic3RyaW5nIiwiUGhvbmUiOiJzdHJpbmciLCJFbWFpbCI6InN0cmluZyIsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJURVNUIiwiYXVkIjpbInN0cmluZyJdLCJleHAiOjE3NjQxNDM2OTYsIm5iZiI6MTc1ODA5NTY5Nn0.Vv8MszfQct_eBK-SQ_9XplkVCzO3hzF87OWOWZ-UbwHCxFltXTvPhajG3QBKj3xqIBsRSCa8PqVEWguGKdsJnQHdrk5r_X4CQfzqQ_HA8uIQG4OL9w6cShlRBe4_1tUGXqgYLclivV7o3kKgb6UihWfzVO9sSnptLBusSmAteT3CwUx7N2aR_oMG3_31yeGqbmDgn2clX9uy-aCkx2YWlPpKxtxzZHmpqxIAtLSzpAArSvo_dh5taqQtYn62QMdK49S6rlOlfyQb30E12AC6Ue3iTxl0BMyCsK3AYQqkiTB7uyfrXQW_JxxOmQZUHv9z0gmG1ugHx3mxoWz3-O24eg"
	jwt, _ := utilo.GetParser("./public_key.pem")
	parseToken, err := jwt.ParseToken(token)
	if err != nil {
		return
	}

	fmt.Println(parseToken)

}

func TestGenePem(t *testing.T) {
	utilo.GenRsaCertPair()
}
