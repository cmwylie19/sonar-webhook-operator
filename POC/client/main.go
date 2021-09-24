package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SonarResult struct {
	Status string `json:"status"`
	// Project SonarProject `bson:"project"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var status SonarResult
	err := json.NewDecoder(r.Body).Decode(&status)
	json_data, err := json.Marshal(status)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Status: ", status.Status)
	log.Println("Marshalled Data: ", string(json_data))
	if err != nil {
		log.Println("Error ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8080/webhook/post", bytes.NewBuffer(json_data))
	req.Header.Set("X-Sonar-Webhook-HMAC-SHA256", GenerateHMACSignature(string(json_data)))
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(responseBody))
}
func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token: ", err.Error())
	}

	client := &http.Client{}

	// data
	values := map[string]string{"status": "SUCCESS"}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err.Error())
	}
	req, _ := http.NewRequest("POST", "http://localhost:8080/webhook/store", bytes.NewBuffer(json_data))
	req.Header.Set("Token", validToken)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(body))
}
func GenerateHMACSignature(message string) string {
	key := []byte(os.Getenv("SECRET"))
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
	//return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Casey Wylie"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/post", handlePost)
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	handleRequests()
}
