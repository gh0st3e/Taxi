package service

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type TokenClaims struct {
	jwt.StandardClaims
	UserID string
}

type Token struct {
}

// NewToken func that allows to get a new access token
func NewToken(userID, signingKey string, tokenTTL time.Duration) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	stringToken, err := TokenToString(token, signingKey)

	return stringToken, err
}

// TokenToString func that allow parse token to string
func TokenToString(token *jwt.Token, signingKey string) (string, error) {
	stringToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", errors.Wrap(err, "jwt.tokenToString.SignedString: Couldn't parse token to string")
	}

	return stringToken, nil
}

// ParseToken func that allow to parse token from string to *jwt.Token
func ParseToken(stringToken, signingKey string) (string, error) {
	token, err := jwt.ParseWithClaims(stringToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", errors.Wrap(err, "service.ParseToken.ParseWithClaims: couldn't parse token")
	}

	if !token.Valid {
		return "", errors.New("token isn't valid")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return "", errors.Wrap(err, "service.ParseToken.Claims: token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func CheckTokenForExpired(token string) error {
	return nil
}
