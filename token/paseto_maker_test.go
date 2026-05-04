package token

import (
	"testing"
	"time"

	"github.com/jonlittler/ts/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload1, err := maker.CreateToken(TokenTypeAccessToken, username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload1)
	// t.Log("payload from create: ", payload1)

	payload2, err := maker.VerifyToken(TokenTypeAccessToken, token)
	require.NoError(t, err)
	require.NotEmpty(t, payload2)
	// t.Log("payload from verify: ", payload2)

	require.NotZero(t, payload2.ID)
	require.Equal(t, username, payload2.Username)
	require.WithinDuration(t, issuedAt, payload2.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload2.ExpiredAt, time.Second)

	require.Equal(t, payload1.ID, payload2.ID)
	require.Equal(t, payload1.Type, payload2.Type)
	require.Equal(t, payload1.Username, payload2.Username)
	require.WithinDuration(t, payload1.IssuedAt, payload2.IssuedAt, time.Second)
	require.WithinDuration(t, payload1.ExpiredAt, payload2.ExpiredAt, time.Second)
}

func TestExpiredPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	/* create token */
	token, payload, err := maker.CreateToken(TokenTypeAccessToken, util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	/* verify token */
	payload, err = maker.VerifyToken(TokenTypeAccessToken, token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestPasetoWrongTokenType(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(TokenTypeAccessToken, util.RandomOwner(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(TokenTypeRefreshToken, token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
