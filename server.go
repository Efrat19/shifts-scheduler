package main

import (
	//"encoding/json"
	"fmt"
	//"log"
	//"math"
	"net/http"
	"os"

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


func slashCommandHandler(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)

	if err != nil {
		fmt.Printf("[ERROR] on parsing: %v",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(os.Getenv("SLACK_VERIFICATION_TOKEN")) {
		fmt.Printf("[ERROR] invalid token")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	//case "/weather":
		//params := &slack.Msg{Text: s.Text}
		//zipCode := params.Text
		//
		//// Generates API request URL with zip code and API key
		//url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?zip=%v&APPID=%v&units=imperial", zipCode, os.Getenv("API_KEY"))
		//
		//req, err := http.NewRequest("GET", url, nil)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//client := &http.Client{}
		//resp, err := client.Do(req)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//defer resp.Body.Close()
		//
		//if resp.StatusCode == http.StatusOK {
		//	apiResponse := &APIResponse{}
		//
		//	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//
		//	// Round temp to nearest integer
		//	roundedTemp := math.Round(apiResponse.Weather.Temp)
		//
		//	response := fmt.Sprintf("The weather in %v is %v. The temperature is %v\u00B0 F.", apiResponse.Location, apiResponse.Summary[0].Description, roundedTemp)
		//	w.Write([]byte(response))
		//}
	case "/devops-on-duty":
		fmt.Println("[INFO] Receiving /devops-on-duty request")
		err,onDuty := whoIsOnDutyNow()
		var response string
		if err != nil {
			response = fmt.Sprintf("%s is on duty today",onDuty)
		} else {
			response = "Error finding DevOps on duty today"
		}
		w.Write([]byte(response))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}