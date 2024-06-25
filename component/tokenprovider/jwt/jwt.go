package jwt

import (
	"food-delivery/component/tokenprovider"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	secret string
}

func NewJwtProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		Payload: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:   myToken,
		Created: time.Now(),
		Expiry:  expiry,
	}, nil
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil{
		return nil,tokenprovider.ErrInvalidToken
	}

	if !res.Valid{
		return nil,tokenprovider.ErrInvalidToken
	}

	claims,ok := res.Claims.(*myClaims)
	if !ok{
		return nil,tokenprovider.ErrInvalidToken
	}
	return &claims.Payload,nil
}
