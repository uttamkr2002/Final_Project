package controller

import (
	models "client/modelsWithInterface"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

// present data fetch kar ke db ko de do

func Handler(wg *sync.WaitGroup) {
	defer wg.Done()
	// baseUrl defiene
	baseUrl := "http://localhost:8080"

	// data will be in the form of struct
	// convert the struct into json := Marshal or Encode
	payload := models.Payload{}
	payload.CollectMetricsforPayload()
	reqByte, err := json.Marshal(payload) // cannot pass the struct create an obj then pas
	if err != nil {
		fmt.Println("We Encountered Error in Marshal", err)
		return
	}
	// convert the json data into json Reader
	jsonReader := strings.NewReader(string(reqByte)) //

	// http.Post(path, content-type, data in form of jsonReader)

	res, err := http.Post(baseUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("We Encountered the error in Posting the data", err)
		return
	}
	fmt.Println(res.StatusCode)
}
