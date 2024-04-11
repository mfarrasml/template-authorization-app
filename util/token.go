package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUtil interface {
	NewAuthToken(email string) (string, error)
	ParseAuthToken(tokenString string) (*UserAuthClaims, error)
}

type jwtTokenUtil struct {
	secret     string
	issuer     string
	expMinutes int
}

type JwtTokenOpts struct {
	Secret     string
	Issuer     string
	ExpMinutes int
}

func NewJwtTokenUtil(opt JwtTokenOpts) *jwtTokenUtil {
	return &jwtTokenUtil{
		secret:     opt.Secret,
		issuer:     opt.Issuer,
		expMinutes: opt.ExpMinutes,
	}
}

type UserAuthClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (t *jwtTokenUtil) NewAuthToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserAuthClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    t.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t.expMinutes) * time.Minute)),
		},
	})

	signed, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", errors.New("error signing JWT claims")
	}

	return signed, nil
}

func (t *jwtTokenUtil) ParseAuthToken(tokenString string) (*UserAuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserAuthClaims{}, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(t.issuer), nil
	}, jwt.WithIssuer(t.issuer),
		jwt.WithIssuedAt(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserAuthClaims)
	if !ok {
		return nil, errors.New("unknown claims")
	}

	return claims, err
}
