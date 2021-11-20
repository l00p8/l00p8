package shield

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	jwt.RegisteredClaims
	//Audience interface{} `json:"aud,omitempty"`
}

func (c JWTClaims) Valid() error {
	err := c.RegisteredClaims.Valid()
	return err
}
