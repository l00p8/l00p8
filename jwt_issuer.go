package shield

import "github.com/golang-jwt/jwt/v4"

type Issuer interface {
	Issue(claims *JWTClaims) (string, error)
}

type jwtIssuer struct {
	keyType string
	key     interface{}
}

// NewIssuer create JWT token issuer object
// privKey - must be string containing pem encoded public key
// keyType - type of the public key (ES256, RSA256), default is RSA256
func NewIssuer(privKey []byte, keyType string) (Issuer, error) {
	v := &jwtIssuer{keyType: keyType}
	var err error = nil
	var key interface{} = nil
	switch keyType {
	case "ES256":
		key, err = jwt.ParseECPrivateKeyFromPEM(privKey)
	case "RS256":
		key, err = jwt.ParseRSAPrivateKeyFromPEM(privKey)
	default:
		return nil, ErrKeyUnsupported
	}
	if err != nil {
		return nil, ErrKeyIsInvalid
	}
	v.key = key
	return v, nil
}

func (iss *jwtIssuer) GetSigningMethod() jwt.SigningMethod {
	switch iss.keyType {
	case "ES256":
		return jwt.SigningMethodES256
	case "RS256":
		return jwt.SigningMethodRS256
	default:
		return jwt.SigningMethodRS256
	}
}

func (iss *jwtIssuer) Issue(claims *JWTClaims) (string, error) {
	tkn := jwt.NewWithClaims(iss.GetSigningMethod(), claims)
	return tkn.SignedString(iss.key)
}
