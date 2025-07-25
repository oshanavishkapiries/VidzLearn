package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// GenerateRandomOtp generates a 6-digit numeric OTP as a string
func GenerateRandomOtp() string {
	max := big.NewInt(1000000) // 6 digits
	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "000000" // fallback
	}
	return fmt.Sprintf("%06d", num.Int64())
}

// GenerateOtpReference generates a unique reference string (UUID)
func GenerateOtpReference() string {
	return uuid.New().String()
}

// HashOtp hashes the OTP using bcrypt
func HashOtp(otp string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.DefaultCost)
	return string(hash), err
}

// HashOtpReference hashes the OTP reference using bcrypt
func HashOtpReference(ref string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(ref), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareOtpHash compares a plain OTP with its hash
func CompareOtpHash(otp, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(otp)) == nil
}

// CompareOtpReferenceHash compares a plain reference with its hash
func CompareOtpReferenceHash(ref, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(ref)) == nil
}
