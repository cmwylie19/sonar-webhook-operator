package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cmwylie19/sonar-webhook-operator/webhook/hmac"
	"github.com/dgrijalva/jwt-go"
)

// JWTMiddleware
func JWTMiddleware(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(os.Getenv("SECRET")), nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				fmt.Println(token)
				endpoint(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

// LoggingMiddleware
func LoggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if r.URL.Path == "/webhook/healthz" {
		// 	return
		// }
		log.Println(r.URL.Path)
		f(w, r)
	}
}

// Middleware for validating HMAC SHA256 signature
func ValidHMAC(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Printf("Cannot read the body of the request: %s\n", err.Error())
			return
		}
		fmt.Println("X-Sonar-Webhook-HMAC-SHA256: ", r.Header.Get("X-Sonar-Webhook-HMAC-SHA256"))
		err = hmac.Validate(body, r.Header.Get("X-Sonar-Webhook-HMAC-SHA256"), string(os.Getenv("SECRET")))
		if err != nil {
			log.Println("Error validating signature: ", err.Error())
			return
		}
		defer r.Body.Close()

		// Restore the body
		// The outer handler reads the request body to EOF. When the inner handler is called, there's nothing more to read from the body.
		//To fix the issue, restore the request body with the data read previously in the outer handler:
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
		fmt.Println("Middleware: ", body)
		f(w, r)
	}
}
