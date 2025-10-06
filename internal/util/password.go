package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// GenerateRandomPassword returns a URL-safe random password of n bytes encoded in base64.
func GenerateRandomPassword(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func GenerateRandomOTP(s string) string {
	rnd, _ := GenerateRandomPassword(20)
	hs := sha256.New()
	hs.Write([]byte(rnd))
	hash := hs.Sum(nil)

	otp := fmt.Sprintf("%s:%s", s, hash)
	return base64.RawURLEncoding.EncodeToString([]byte(otp))
}

