package shared

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type SimpleAuthenticator struct {
	SecretKey []byte
}

func (a *SimpleAuthenticator) GenerateToken(userID string) (string, error) {
	claims := map[string]interface{}{
		"user_id": userID,
		"exp":     time.Now().Add(168 * time.Hour).Unix(),
	}
	jsonBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	sig := hmac.New(sha256.New, a.SecretKey)
	sig.Write(jsonBytes)
	signature := sig.Sum(nil)

	token := base64.StdEncoding.EncodeToString(jsonBytes) + "." + base64.StdEncoding.EncodeToString(signature)
	return token, nil
}

func NewSimpleAuthenticator(secretKey string) *SimpleAuthenticator {
	return &SimpleAuthenticator{
		SecretKey: []byte(secretKey),
	}
}

func (a *SimpleAuthenticator) DecodeToken(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid token format")
	}

	encodedClaims, encodedSig := parts[0], parts[1]

	claimsBytes, err := base64.StdEncoding.DecodeString(encodedClaims)
	if err != nil {
		return "", fmt.Errorf("failed to decode claims: %v", err)
	}

	sig := hmac.New(sha256.New, a.SecretKey)
	sig.Write(claimsBytes)
	expectedSig := sig.Sum(nil)

	providedSig, err := base64.StdEncoding.DecodeString(encodedSig)
	if err != nil {
		return "", fmt.Errorf("failed to decode signature: %v", err)
	}

	if !hmac.Equal(expectedSig, providedSig) {
		return "", fmt.Errorf("invalid signature")
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(claimsBytes, &claims); err != nil {
		return "", fmt.Errorf("failed to parse claims: %v", err)
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return "", fmt.Errorf("invalid or missing 'exp' claim")
	}

	if time.Now().Unix() > int64(exp) {
		return "", fmt.Errorf("token expired")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fmt.Errorf("invalid or missing 'user_id' claim")
	}

	return userID, nil
}
