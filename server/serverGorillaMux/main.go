package main

import (
	xy "server/modelsWithInterface"
	"server/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fetch the data
	res, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("We encounterd error ", err)
		return
	}
	// we need to Unmarshal it
	p1 := xy.Payload{}
	p1.CollectMetricsforPayload()
	err = json.Unmarshal(res, &p1)
	if err != nil {
		fmt.Println("We Encountered Error in Unmarshal", err)
		return
	}
	service.PrintMetrics(p1)

}

func main() {
	// initialize the port number
	port := "8000"
	// initialize the router
	mux := mux.NewRouter()

	// launch the handler func at path
	mux.HandleFunc("/", handler)

	// Listen and Serve at port number
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Println("We encountered the error in Listen And Serve errors:=", err)
		return
	}
}
