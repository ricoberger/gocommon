package jwt

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var (
	// ErrJwtValidationFailed is used if the validation of a jwt token fails.
	ErrJwtValidationFailed = errors.New("jwt validation failed")
	// ErrUnexpectedSigningMethod is used if the signing method of a provided jwt token is invalid.
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
)

// Create creates, signs, and encodes a JWT token using the HMAC signing method.
func Create(claims jwt.MapClaims, signingSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(signingSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Parse parses and validates a token using the HMAC signing method.
func Parse(tokenString, signingSecret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}

		return []byte(signingSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrJwtValidationFailed
}
