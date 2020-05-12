package jwt

import (
	"github.com/CherryDock/CherryDock/api/jsonutils"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type JwtToken struct {
	Token      string
	Expiration string
}

const (
	JWT_KEY = "CherryDock"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.Form.Get("user")
	password := r.Form.Get("password")

	w.Header().Add("Content-Type", "application/json")

	if user != "admin" || password != "password" {
		w.WriteHeader(http.StatusUnauthorized)
	}

	expiration := time.Now().Add(time.Hour * time.Duration(10000))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  expiration.Unix(),
	})

	tokenStr, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		jwtToken := JwtToken{tokenStr, expiration.String()}
		w.Write(jsonutils.FormatToJson(jwtToken))
		return
	}
}

func CheckToken(next http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_KEY), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddleware.Handler(next)
}
