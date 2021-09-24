package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/cmwylie19/sonar-webhook-operator/webhook/helper"
	"github.com/cmwylie19/sonar-webhook-operator/webhook/models"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

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
