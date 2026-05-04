package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

/* Claims interface
GetExpirationTime() (*NumericDate, error)
GetIssuedAt() (*NumericDate, error)
GetNotBefore() (*NumericDate, error)
GetIssuer() (string, error)
GetSubject() (string, error)
GetAudience() (ClaimStrings, error)
*/

type TokenType byte

const (
	TokenTypeAccessToken  = 1
	TokenTypeRefreshToken = 2
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Type      TokenType `json:"type"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var _ jwt.Claims = (*Payload)(nil)

func NewPayload(tokenType TokenType, username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := Payload{
		ID:        tokenID,
		Type:      tokenType,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return &payload, nil
}

// Valid checks if the token payload is valid or not
func (payload *Payload) Valid(tokenType TokenType) error {
	if payload.Type != tokenType {
		return ErrInvalidToken
	}
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: payload.ExpiredAt}, nil
}
func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: payload.IssuedAt}, nil
}
func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: payload.IssuedAt}, nil
}

func (payload *Payload) GetIssuer() (string, error)             { return "", nil }
func (payload *Payload) GetSubject() (string, error)            { return "", nil }
func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) { return jwt.ClaimStrings{}, nil }
