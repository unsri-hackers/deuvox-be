package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

type contextKey struct {
	name string
}

type JWTAuth struct {
	alg       jwa.SignatureAlgorithm
	signKey   interface{} // private-key
	verifyKey interface{} // public-key, only used by RSA and ECDSA algorithms
	verifier  jwt.ParseOption
}

type Token struct {
	Value     string    `json:"value"`
	Type      string    `json:"type"`
	ExpiredAt time.Time `json:"expired_at"`
	JTI       string    `json:"-"`
}

var (
	TokenCtxKey            = &contextKey{"Token"}
	ErrorCtxKey            = &contextKey{"Error"}
	AccessTokenExpiration  = 15
	RefreshTokenExpiration = 60
)

var (
	ErrUnauthorized = errors.New("token is unauthorized")
	ErrExpired      = errors.New("token is expired")
	ErrNBFInvalid   = errors.New("token nbf validation failed")
	ErrIATInvalid   = errors.New("token iat validation failed")
	ErrNoTokenFound = errors.New("no token found")
	ErrAlgoInvalid  = errors.New("algorithm mismatch")
)

func New() *JWTAuth {
	ja := &JWTAuth{alg: jwa.SignatureAlgorithm(os.Getenv("JWT_ALG")), signKey: []byte(os.Getenv("JWT_SIGN"))}

	if ja.verifyKey != nil {
		ja.verifier = jwt.WithVerify(ja.alg, []byte(ja.verifyKey.(string)))
	} else {
		ja.verifier = jwt.WithVerify(ja.alg, ja.signKey)
	}

	return ja
}
