package authControllerPkg

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	authServicePkg "github.com/kuroyamii/golang-webapi/internal/auth/service"
)

var key = "gery"

type AuthController struct {
	as     authServicePkg.AuthService
	router *mux.Router
}

func (ac *AuthController) ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Not Authorized"))
				}
				return authServicePkg.AUTH_KEY, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Not Authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorized"))
		}
	})
}

func (ac *AuthController) GetJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != key {
			return
		} else {
			token, err := ac.as.CreateJWT()
			if err != nil {
				return
			}
			fmt.Fprintln(w, token)
			w.Write([]byte(token))
		}
	}
}

func ProvideAuthController(router *mux.Router, as authServicePkg.AuthService) AuthController {
	return AuthController{
		router: router,
		as:     as,
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
}

func (ac *AuthController) InitializeEndPoint() {
	ac.router.Handle("/auth", ac.ValidateJWT(Home)).Methods("GET")
	ac.router.HandleFunc("/jwt", ac.GetJWT).Methods("GET")
}
