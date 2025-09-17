package utilr

import "github.com/google/uuid"

// GenerateKeyPair generate ak, sk
func GenerateKeyPair() (string, string) {

	return uuid.New().String(), uuid.New().String()
}
