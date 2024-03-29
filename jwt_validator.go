package shield

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenIsInvalid   = errors.New("token is invalid")
	ErrTokenExpired     = errors.New("token expired")
	ErrTokenIssuedAt    = errors.New("token issued at error")
	ErrKeyIsInvalid     = errors.New("key is invalid")
	ErrKeyUnsupported   = errors.New("unsupported signing algorithm")
	ErrSignatureInvalid = errors.New("signature is invalid")
)

type Validator interface {
	Valid(token string) (*JWTClaims, error)
}

type jwtValidator struct {
	pubKey  []byte
	keyType string
	key     interface{}
}

func (v *jwtValidator) Valid(jwtToken string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return v.key, nil
	})
	if err != nil {
		return nil, processValidationErr(err)
	}
	return claims, nil
}

func processValidationErr(err error) error {
	switch err {
	case jwt.ErrTokenExpired:
		return ErrTokenExpired
	case jwt.ErrTokenNotValidYet:
		return ErrTokenIsInvalid
	case jwt.ErrTokenMalformed:
		return ErrTokenIsInvalid
	case jwt.ErrTokenSignatureInvalid:
		return ErrSignatureInvalid
	case jwt.ErrTokenInvalidIssuer:
		return ErrTokenIssuedAt
	}
	if err != nil {
		return err
	}
	return nil
}

// NewValidator create JWT token validator object
// pubKey - must be string containing pem encoded public key
// keyType - type of the public key (ES256, RSA256), default is RSA256
func NewValidator(pubKey []byte, keyType string) (Validator, error) {
	v := &jwtValidator{pubKey: pubKey, keyType: keyType}
	var err error = nil
	var key interface{} = nil
	switch keyType {
	case "ES256":
		key, err = jwt.ParseECPublicKeyFromPEM(pubKey)
	case "RS256":
		key, err = jwt.ParseRSAPublicKeyFromPEM(pubKey)
	default:
		return nil, ErrKeyUnsupported
	}
	if err != nil {
		return nil, ErrKeyIsInvalid
	}
	v.key = key
	return v, nil
}
