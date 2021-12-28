package shield

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	jwt.RegisteredClaims
	//Audience interface{} `json:"aud,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

func (c JWTClaims) Valid() error {
	err := c.RegisteredClaims.Valid()
	return err
}
