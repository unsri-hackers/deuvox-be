package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/lestrrat-go/jwx/jwt"
)

const (
	RefreshToken = "refresh"
	AccessToken  = "access"
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
	Value     string `json:"value"`
	Type      string `json:"type"`
	ExpiredAt int64  `json:"expired_at"`
	JTI       string `json:"-"`
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

func CreateAccessToken(userID string) (Token, string, error) {
	return createToken(AccessToken, userID)
}
func CreateRefreshToken(userID string) (Token, string, error) {
	return createToken(RefreshToken, userID)
}

func VerifyToken(token []byte) (jwt.Token, error) {
	sign, err := jws.Verify(token, jwa.HS256, []byte(os.Getenv("JWT_SIGN")))
	if err != nil {
		return nil, errors.New("Invalid token")
	}
	signedToken, err := jwt.Parse(sign)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse signed token: %v", err)
	}
	return signedToken, nil
}

func createToken(tokenType string, userID string) (Token, string, error) {
	var tok Token
	var jti string
	duration := time.Now().Add(24 * time.Hour).Unix()
	if tokenType == RefreshToken {
		duration = time.Now().Add(7 * 24 * time.Hour).Unix()
	}
	jwtID, err := uuid.NewRandom()
	if err != nil {
		return tok, jti, fmt.Errorf("Failed to create random uuid: %v", err)
	}
	jti = jwtID.String()
	t := jwt.New()
	t.Set(jwt.AudienceKey, "deuvox") // TODO: set audience based on server (development, production)
	t.Set(jwt.ExpirationKey, duration)
	t.Set(jwt.IssuedAtKey, time.Now().Unix())
	t.Set(jwt.IssuerKey, "Deuvox Backend")
	t.Set(jwt.JwtIDKey, jti)
	t.Set(jwt.SubjectKey, userID)

	sign, err := jwt.Sign(t, jwa.HS256, []byte(os.Getenv("JWT_SIGN")))
	if err != nil {
		return tok, jti, err
	}
	tok = Token{Value: string(sign), Type: "BEARER", ExpiredAt: duration}
	return tok, jti, nil
}
