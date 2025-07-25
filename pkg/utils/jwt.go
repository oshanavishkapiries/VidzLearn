package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

// JWTHeader represents the header part of a JWT
// [INFO] @better-comments:info
// {"alg":"HS256","typ":"JWT"}
type JWTHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// JWTPayload represents the payload part of a JWT
// [INFO] @better-comments:info
// You can add more fields as needed
// Example: UserID, Expiry, etc.
type JWTPayload struct {
	Sub   string `json:"sub"` // Subject (user id)
	Exp   int64  `json:"exp"` // Expiry (unix timestamp)
	Email string `json:"email,omitempty"`
}

// secretKey should be loaded from config in production
var secretKey = []byte("your-very-secret-key")

// SetSecretKey allows setting the JWT secret key at runtime
func SetSecretKey(key string) {
	secretKey = []byte(key)
}

// SignJWT creates a JWT string for the given payload
func SignJWT(payload JWTPayload) (string, error) {
	header := JWTHeader{Alg: "HS256", Typ: "JWT"}
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	headerB64 := base64.RawURLEncoding.EncodeToString(headerJSON)
	payloadB64 := base64.RawURLEncoding.EncodeToString(payloadJSON)

	unsignedToken := headerB64 + "." + payloadB64
	sig := signHS256(unsignedToken, secretKey)
	sigB64 := base64.RawURLEncoding.EncodeToString(sig)

	return unsignedToken + "." + sigB64, nil
}

// VerifyJWT verifies the JWT string and returns the payload if valid
func VerifyJWT(token string) (*JWTPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}
	unsignedToken := parts[0] + "." + parts[1]
	sig, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, errors.New("invalid signature encoding")
	}
	if !hmac.Equal(sig, signHS256(unsignedToken, secretKey)) {
		return nil, errors.New("invalid token signature")
	}
	payloadJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, errors.New("invalid payload encoding")
	}
	var payload JWTPayload
	if err := json.Unmarshal(payloadJSON, &payload); err != nil {
		return nil, errors.New("invalid payload json")
	}
	if time.Now().Unix() > payload.Exp {
		return nil, errors.New("token expired")
	}
	return &payload, nil
}

// signHS256 creates an HMAC SHA256 signature
func signHS256(data string, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return h.Sum(nil)
}
