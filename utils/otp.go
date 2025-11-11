package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

func GenerateOTP() string {
	max := big.NewInt(1000000)
	n, _ := rand.Int(rand.Reader, max)
	otp := n.Int64()
	return fmt.Sprintf("%06d", otp)
}

func HashOTP(otp string) string {
	hash := sha256.Sum256([]byte(otp))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func VerifyOTPHash(plain, hashed string) bool {
	cleanedHash := strings.Trim(hashed, "\"")
	return HashOTP(plain) == cleanedHash
}
