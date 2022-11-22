package auth

import (
	"crypto/rand"
	"encoding/hex"
)

type TokenInfo struct {
	IsValid   bool
	Requester string
	Perms     Permissions
}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
