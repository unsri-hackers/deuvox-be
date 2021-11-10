package jwt

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
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
	ja := &JWTAuth{alg: jwa.SignatureAlgorithm(os.Getenv("JWT_ALG")), signKey: []byte(os.Getenv("JWT_SIGN")), verifyKey: nil}

	if ja.verifyKey != nil {
		ja.verifier = jwt.WithVerify(ja.alg, []byte(ja.verifyKey.(string)))
	} else {
		ja.verifier = jwt.WithVerify(ja.alg, ja.signKey)
	}

	return ja
}

func (ja *JWTAuth) Verifier() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return ja.Verify(TokenFromHeader, TokenFromCookie)(next)
	}
}

func (ja *JWTAuth) Verify(findTokenFns ...func(r *http.Request) string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token, err := ja.VerifyRequest(r, findTokenFns...)
			ctx = NewContext(ctx, token, err)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(hfn)
	}
}

func (ja *JWTAuth) VerifyRequest(r *http.Request, findTokenFns ...func(r *http.Request) string) (jwt.Token, error) {
	var tokenString string

	for _, fn := range findTokenFns {
		tokenString = fn(r)
		if tokenString != "" {
			break
		}
	}
	if tokenString == "" {
		return nil, ErrNoTokenFound
	}

	return ja.VerifyToken(tokenString)
}

func (ja *JWTAuth) VerifyToken(tokenString string) (jwt.Token, error) {
	token, err := ja.Decode(tokenString)
	if err != nil {
		return token, ErrorReason(err)
	}

	if token == nil {
		return nil, ErrUnauthorized
	}

	if err := jwt.Validate(token); err != nil {
		return token, ErrorReason(err)
	}

	return token, nil
}

func (ja *JWTAuth) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := ja.FromContext(r.Context())

		if err != nil {
			http.Error(w, err.Error(), 401)
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func (ja *JWTAuth) CreateToken(jti string, userId string, minute int) (*Token, error) {
	expires_at := time.Now().Add(time.Duration(minute) * time.Minute)

	accessTokenClaims := make(map[string]interface{})
	accessTokenClaims["id"] = jti
	accessTokenClaims["userId"] = userId
	accessTokenClaims["exp"] = expires_at

	_, tokenString, err := ja.Encode(accessTokenClaims)
	if err != nil {
		return nil, err
	}

	return &Token{
		Value:     tokenString,
		Type:      "BEARER",
		ExpiredAt: expires_at,
	}, nil
}

func NewContext(ctx context.Context, t jwt.Token, err error) context.Context {
	ctx = context.WithValue(ctx, TokenCtxKey, t)
	ctx = context.WithValue(ctx, ErrorCtxKey, err)
	return ctx
}

func (ja *JWTAuth) FromContext(ctx context.Context) (jwt.Token, map[string]interface{}, error) {
	token, _ := ctx.Value(TokenCtxKey).(jwt.Token)

	var err error
	var claims map[string]interface{}

	if token != nil {
		claims, err = token.AsMap(context.Background())
		if err != nil {
			return token, nil, err
		}
	} else {
		claims = map[string]interface{}{}
	}

	err, _ = ctx.Value(ErrorCtxKey).(error)

	return token, claims, err
}

func TokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}

func TokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func ErrorReason(err error) error {
	switch err.Error() {
	case "exp not satisfied", ErrExpired.Error():
		return ErrExpired
	case "iat not satisfied", ErrIATInvalid.Error():
		return ErrIATInvalid
	case "nbf not satisfied", ErrNBFInvalid.Error():
		return ErrNBFInvalid
	default:
		return ErrUnauthorized
	}
}

func (ja *JWTAuth) Encode(claims map[string]interface{}) (t jwt.Token, tokenString string, err error) {
	t = jwt.New()
	for k, v := range claims {
		err := t.Set(k, v)
		if err != nil {
			return nil, "", err
		}
	}

	payload, err := ja.sign(t)
	if err != nil {
		return nil, "", err
	}
	tokenString = string(payload)
	return
}

func (ja *JWTAuth) Decode(tokenString string) (jwt.Token, error) {
	return ja.parse([]byte(tokenString))
}

func (ja *JWTAuth) sign(token jwt.Token) ([]byte, error) {
	return jwt.Sign(token, ja.alg, ja.signKey)
}

func (ja *JWTAuth) parse(payload []byte) (jwt.Token, error) {
	return jwt.Parse(payload, ja.verifier)
}
