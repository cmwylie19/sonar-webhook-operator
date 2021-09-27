package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cmwylie19/sonar-webhook-operator/webhook/helper"
	"github.com/cmwylie19/sonar-webhook-operator/webhook/models"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Generate a JWT Token
func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

// Hash plaintext password
func HashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(bytes), err
}

// Check plaintext against hashed password
func CheckPasswordHash(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}

// Return all SonarResults
func ReadAllResults() ([]models.WebHook, error) {
	wh := []models.WebHook{}
	filter := bson.M{}
	client, err := helper.GetMongoClient()
	if err != nil {
		log.Println("Error GetMongoClient: ", err.Error())
		return wh, err
	}

	collection := client.Database(helper.DB).Collection(helper.RESULTS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return wh, findError
	}

	for cur.Next(context.TODO()) {
		h := models.WebHook{}
		err := cur.Decode(&h)
		if err != nil {
			return wh, err
		}
		wh = append(wh, h)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(wh) == 0 {
		return wh, mongo.ErrNoDocuments
	}

	//Return result without any error.
	return wh, nil

}

func GetUserByEmail(email string) (models.User, error) {
	result := models.User{}
	filter := bson.M{"email": email}
	client, err := helper.GetMongoClient()
	if err != nil {
		return result, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(helper.DB).Collection(helper.USERS)

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	//Return result without any error.
	return result, nil
}

// Create User and Hash Password
func CreateUser(user models.User) error {
	client, err := helper.GetMongoClient()
	if err != nil {
		return err
	}
	u, err := GetUserByEmail(user.Email)
	if err != nil {
		user.Password, _ = HashPassword(user.Password)

		collection := client.Database(helper.DB).Collection(helper.USERS)
		_, err = collection.InsertOne(context.TODO(), user)
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("User with email %v already exists", u.Email)
	}

}

// POST from sonarqube webhook.
// Store exact data into mongo
func StoreSonarResults(wh models.WebHook) error {
	client, err := helper.GetMongoClient()
	if err != nil {
		log.Println("Error GetMongoClient: ", err.Error())
		return err
	}
	collection := client.Database(helper.DB).Collection(helper.RESULTS)
	_, err = collection.InsertOne(context.TODO(), wh)
	if err != nil {
		return err
	}
	return nil

}

func SuccessResponse(message string, w http.ResponseWriter) {
	var response = models.SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    message,
	}
	msg, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write(msg)
}

func GetError(err error, w http.ResponseWriter) {
	var response = models.ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
