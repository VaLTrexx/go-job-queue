package job

import (
	"crypto/rand"
	"encoding/hex"
)

func NewID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
