package handlers // GetPost Handler
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cmwylie19/sonar-webhook-operator/webhook/controllers"
	"github.com/cmwylie19/sonar-webhook-operator/webhook/models"
)

// ViewHandler for the UI
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1 style=background-color:blue;text-align:center;color:red;>%s</h1><div style=text-align:center>%s</div>", "WebHook Operator", "Accepts POST requests from Sonarqube")
}

// Authenticate Endpoint
func Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var creds models.User
	_ = json.NewDecoder(r.Body).Decode(&creds)
	user, err := controllers.GetUserByEmail(creds.Email)
	if err != nil {
		controllers.GetError(err, w)
		return
	}

	match := controllers.CheckPasswordHash(creds.Password, user.Password)
	if match {
		// getToken()
		fmt.Println("Correct Credentials")
		controllers.SuccessResponse("Correct creds", w)
	} else {
		fmt.Println("failure")
		err := fmt.Errorf("credential incorrect")
		controllers.GetError(err, w)
		return
	}

	result, _ := json.Marshal(user)
	w.Write(result)
}

// Create User Handler
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	err := controllers.CreateUser(user)
	if err != nil {
		controllers.GetError(err, w)
		return
	}

	// Generate a token
	validToken, err := controllers.GenerateJWT(user.Email)
	if err != nil {
		fmt.Println("Failed to generate token: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Token = validToken
	user.Password = ""
	result, _ := json.Marshal(user)
	w.Write(result)
}

// Checktoken Endpoint
func CheckToken(w http.ResponseWriter, r *http.Request) {}

// Receive the data- dont store it
func SonarEndpoint(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type to Application/JSON
	w.Header().Set("Content-Type", "application/json")

	var webhook models.SonarResult
	unmarshalError := json.NewDecoder(r.Body).Decode(&webhook)
	if unmarshalError != nil {
		log.Fatalf("Could not unmarshal post data: %v", unmarshalError)
	}

	// fmt.Println(webhook.Project.URL)
	fmt.Println(webhook.Status)

	w.Write([]byte(webhook.Status))

}

// GetRemote  Handler used for HealthCheck
func GetResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	results, err := controllers.ReadAllResults()
	if err != nil {
		controllers.GetError(err, w)
		return
	}

	result, _ := json.Marshal(results)
	w.Write(result)
}

// GetRemote  Handler used for HealthCheck
func GetRemote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	remote := models.Remote{XFF: r.Header.Get("x-forwarded-for")}
	result, err := json.Marshal(remote)
	if err != nil {
		controllers.GetError(err, w)
	}
	w.Write(result)
}

func GetPostTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var wh models.WebHook

	err := json.NewDecoder(r.Body).Decode(&wh)
	if err != nil {
		fmt.Println("Error:Unmarshal ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	saveError := controllers.StoreSonarResults(wh)
	if saveError != nil {
		controllers.GetError(saveError, w)
		return
	}
	fmt.Fprintf(w, "%+v", "saved")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// unmarshal
	var result models.SonarResult
	//var res map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error:Unmarshal ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%+v", result)

}
