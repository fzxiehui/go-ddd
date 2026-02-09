package crypto

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(password string) string {
	sum := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", sum)
}
