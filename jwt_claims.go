package shield

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	//Audience interface{} `json:"aud,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

func (c JWTClaims) Valid() error {
	//err := c.RegisteredClaims.Valid()
	if c.ExpiresAt.Before(time.Now()) {
		return ErrTokenExpired
	}
	return nil
}
