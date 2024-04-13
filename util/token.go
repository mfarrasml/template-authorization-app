package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

type TokenUtil interface {
	NewAuthToken(id int, email string) (string, error)
	ParseAuthToken(tokenString string) (*UserAuthClaims, error)
	NewRefreshToken(id int) (refToken string, jti string, err error)
	ParseRefreshToken(tokenString string) (*UserRefreshClaims, error)
}

type jwtTokenUtil struct {
	secret           string
	issuer           string
	accTknExpMinutes int
	refTknExpMinutes int
}

type JwtTokenOpts struct {
	Secret           string
	Issuer           string
	AccTknExpMinutes int
	RefTknExpMinutes int
}

func NewJwtTokenUtil(opt JwtTokenOpts) *jwtTokenUtil {
	return &jwtTokenUtil{
		secret:           opt.Secret,
		issuer:           opt.Issuer,
		accTknExpMinutes: opt.AccTknExpMinutes,
		refTknExpMinutes: opt.RefTknExpMinutes,
	}
}

type UserAuthClaims struct {
	UserId int    `json:"id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
type UserRefreshClaims struct {
	UserId int `json:"id"`
	jwt.RegisteredClaims
}

func (t *jwtTokenUtil) NewAuthToken(id int, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserAuthClaims{
		UserId: id,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    t.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t.accTknExpMinutes) * time.Minute)),
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
		return []byte(t.secret), nil
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

func (t *jwtTokenUtil) NewRefreshToken(id int) (refToken string, jti string, err error) {
	jti = uuid.NewV4().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserRefreshClaims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    t.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t.refTknExpMinutes) * time.Minute)),
		},
	})

	signed, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", "", errors.New("error signing JWT claims")
	}

	return signed, jti, nil
}

func (t *jwtTokenUtil) ParseRefreshToken(tokenString string) (*UserRefreshClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserRefreshClaims{}, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	}, jwt.WithIssuer(t.issuer),
		jwt.WithIssuedAt(),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserRefreshClaims)
	if !ok {
		return nil, errors.New("unknown claims")
	}

	return claims, err
}
