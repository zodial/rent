package auth

import (
	"time"
)

// JWTManager handles token signing & verification using HMAC.
type JWTManager struct {
	Secret []byte
	TTL    time.Duration
}

// NewJWTManager creates a new manager with secret and ttl.
func NewJWTManager(secret string, ttl time.Duration) *JWTManager {
	return &JWTManager{
		Secret: []byte(secret),
		TTL:    ttl,
	}
}

// GenerateToken generates a signed JWT containing admin id and tenant.
func (j *JWTManager) GenerateToken(adminID uint, tenant string) (string, error) {
	claims := jwt.MapClaims{
		"sub": adminID,
		"tid": tenant,
		"exp": time.Now().Add(j.TTL).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Secret)
}

// VerifyToken parses and validates a JWT and returns claims map.
func (j *JWTManager) VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenInvalidMethod
		}
		return j.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
