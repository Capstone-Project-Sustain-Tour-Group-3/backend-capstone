package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncodeRedisKey(key string) string {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
