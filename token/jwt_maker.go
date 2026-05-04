package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

var _ Maker = (*JWTMaker)(nil)

const minSecretKeySize = 32

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (m *JWTMaker) CreateToken(tokenType TokenType, username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(tokenType, username, duration)
	if err != nil {
		return "", nil, err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(m.secretKey))
	return token, payload, nil
}

// VerifyToken checks if the token is valid or not
func (m *JWTMaker) VerifyToken(tokenType TokenType, token string) (*Payload, error) {

	/* very important to prevent trivial attacks */
	fn := func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, fn)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	if err := payload.Valid(tokenType); err != nil {
		return nil, err
	}

	return payload, nil
}
