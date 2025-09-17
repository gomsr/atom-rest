package utilr

import (
	"fmt"
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	ak, sk := GenerateKeyPair()
	fmt.Println(ak, sk)
}
