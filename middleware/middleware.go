package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ricoberger/gocommon/jwt"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// Handler is our custom HTTP handler.
// See: http://blog.golang.org/error-handling-and-go
type Handler func(http.ResponseWriter, *http.Request, httprouter.Params) *Error

// Error is our custom error type.
// See: http://blog.golang.org/error-handling-and-go
type Error struct {
	Error   error  `json:"-"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// HandleError display the error message to the user with the correct HTTP
// status code and log the full error to the developer console.
func HandleError(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if e := h(w, r, ps); e != nil { // e is *Error, not os.Error.
			log.WithFields(log.Fields{
				"message":        e.Message,
				"status_code":    e.Code,
				"underlying_err": e.Error.Error(),
			}).Error("HTTP handler returned an error")

			js, err := json.Marshal(e)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(e.Code)
			w.Write(js)
		}
	}
}

// Errorf is our custom error function.
func Errorf(err error, code int, format string, v ...interface{}) *Error {
	return &Error{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    code,
	}
}

// Write return json data.
func Write(w http.ResponseWriter, r *http.Request, data interface{}) *Error {
	js, err := json.Marshal(data)
	if err != nil {
		return Errorf(err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	w.WriteHeader(http.StatusOK)
	if data != nil {
		w.Write(js)
	}

	return nil
}

// BasicAuth handles basic authentication.
// It checkes if the the sent basic auth crendentials matchs the required user
// and password.
func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			h(w, r, ps)
		} else {
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

// Cors sets cors headers.
func Cors(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT")

		h(w, r, ps)
	}
}

// Log handles logging for http requests.
func Log(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		log.WithFields(log.Fields{
			"host":         r.Host,
			"address":      r.RemoteAddr,
			"method":       r.Method,
			"requestURI":   r.RequestURI,
			"proto":        r.Proto,
			"useragent":    r.UserAgent(),
			"x-request-id": r.Header.Get("X-Request-ID"),
		}).Info("HTTP request")

		h(w, r, ps)
	}
}

// ContextKey implements type for context key.
type ContextKey string

// ContextJWTKey is the key for the jwt context value.
const ContextJWTKey ContextKey = "jwt"

// GetExp returns the expire date from an given jwt claims map.
func getExp(claims jwtgo.MapClaims) int64 {
	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0
	}

	return int64(exp)
}

// BearerAuth handles bearer token authentication.
// The claims from the JWT token can be accessed in the handler function as
// follows:
//   claims, ok := r.Context()
//     .Value(middleware.ContextJWTKey)
//     .(middleware.MapClaims)
func BearerAuth(h httprouter.Handle, signingSecret string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get authentication header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Check if authentication token is present
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Validate authentication token.
		claims, err := jwt.Parse(authHeaderParts[1], signingSecret)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Validate the expire date of the jwt token.
		exp := getExp(claims)
		if time.Unix(exp, 0).Before(time.Now()) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextJWTKey, claims)
		h(w, r.WithContext(ctx), ps)
	}
}
