package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuroyamii/golang-webapi/pkg/entity/response"
)

func ContentTypeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		handler.ServeHTTP(rw, r)
	})
}
func CorsMiddleware(whitelistedUrls map[string]bool) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")

			requestOriginUrl := r.Header.Get("Origin")
			log.Printf("INFO CorsMiddleware: received request from %s %v", requestOriginUrl, whitelistedUrls[requestOriginUrl])
			if whitelistedUrls[requestOriginUrl] {
				rw.Header().Set("Access-Control-Allow-Origin", requestOriginUrl)
			}
			if r.Method != http.MethodOptions {
				handler.ServeHTTP(rw, r)
				return
			}
			rw.Write([]byte("Anjay Mabar"))
		})
	}
}

func ErrorHandlingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		defer func() {
			recover := recover()
			if recover != nil {
				switch v := recover.(type) {
				case *response.BaseResponse:
					rw.WriteHeader(v.Code)
					response.NewBaseResponse(
						v.Code,
						v.Message,
						v.Errors,
						v.Data,
					).ToJSON(rw)
				case error:
					rw.WriteHeader(500)
					response.NewBaseResponse(
						500,
						response.RESPONSE_ERROR_RUNTIME_MESSAGE,
						response.NewErrorResponseData(response.NewErrorResponseValue("msg", v.Error())),
						nil,
					).ToJSON(rw)
				default:
					rw.WriteHeader(500)
					response.NewBaseResponse(
						500,
						response.RESPONSE_ERROR_RUNTIME_MESSAGE,
						response.NewErrorResponseData(response.NewErrorResponseValue("msg", "Runtime Error")),
						nil,
					).ToJSON(rw)
				}
			}
		}()
		handler.ServeHTTP(rw, r)
	})
}
