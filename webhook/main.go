package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cmwylie19/sonar-webhook-operator/webhook/handlers"
	"github.com/cmwylie19/sonar-webhook-operator/webhook/middleware"
)

func HandleRequests() {
	http.HandleFunc("/webhook/healthz", middleware.LoggingMiddleware(handlers.GetRemote))
	http.Handle("/webhook/store", middleware.JWTMiddleware(handlers.SonarEndpoint))
	http.HandleFunc("/webhook/post", middleware.ValidHMAC(handlers.GetPost))
	http.HandleFunc("/webhook/x/post", middleware.LoggingMiddleware(handlers.GetPostTest))
	http.HandleFunc("/webhook/x/results", middleware.LoggingMiddleware(handlers.GetResults))
	http.HandleFunc("/webhook/views", handlers.ViewHandler)
	http.HandleFunc("/webhook/create", handlers.CreateUser)
	fmt.Println("Serving from :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	HandleRequests()
}
