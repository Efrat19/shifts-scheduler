package main

import (
	//"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	//"log"
	//"math"
	"net/http"
	//"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

// APIResponse maps to the JSON response from the Open Weather Map API
type APIResponse struct {
	Summary []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Weather struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
		Pressure int     `json:"pressure"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Location string `json:"name"`
}

func main() {
	// Load environment variables
	//err := godotenv.Load("environment.env")
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	http.HandleFunc("/devops-on-duty", slashCommandHandler)
	http.HandleFunc("/healthz", healthzHandler)

	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":8080", nil)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO] Receiving /healthz request")
	w.WriteHeader(http.StatusOK)
	return
}

func logRequest(r *http.Request) {
	fmt.Println("[INFO] Logging /devops-on-duty request")
	fmt.Printf("%v just triggered a /devops-on-duty request",r.Body)

}

func slashCommandHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO] Receiving /devops-on-duty request")
	logRequest(r)
	signingSecret := getEnv("SLACK_SIGNING_SECRET","")
	verifier, err := slack.NewSecretsVerifier(r.Header, signingSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.Body = ioutil.NopCloser(io.TeeReader(r.Body, &verifier))
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		fmt.Printf("[ERROR] on parsing: %v",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = verifier.Ensure(); err != nil {
		fmt.Printf("[ERROR] invalid verfier")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/devops-on-duty":
		err,onDuty := whoIsOnDutyNow()
		var response string
		if err != nil {
			fmt.Printf("[ERROR] Error finding DevOps on duty today %v",err)
			response = "Error finding DevOps on duty today"
		} else {
			response = fmt.Sprintf("%s is on duty today",onDuty)
		}
		w.Write([]byte(response))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}