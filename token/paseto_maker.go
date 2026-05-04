package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

var _ Maker = (*PasetoMaker)(nil)

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}
	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

// CreateToken creates a new token for a specific username and duration
func (m *PasetoMaker) CreateToken(tokenType TokenType, username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(TokenTypeAccessToken, username, duration)
	if err != nil {
		return "", nil, err
	}
	token, err := m.paseto.Encrypt(m.symmetricKey, payload, nil)
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (m *PasetoMaker) VerifyToken(tokenType TokenType, token string) (*Payload, error) {
	payload := &Payload{}
	if err := m.paseto.Decrypt(token, m.symmetricKey, payload, nil); err != nil {
		return nil, ErrInvalidToken
	}
	if err := payload.Valid(tokenType); err != nil {
		return nil, err
	}
	return payload, nil
}
