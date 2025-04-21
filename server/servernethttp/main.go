package main

import (
	xy "server/modelsWithInterface"
	"server/service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// I have to recieve data from client

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We Are In Server Handler")
	// whether the client has requested post request
	fmt.Println(r.Method)
	if r.Method != http.MethodPost {
		fmt.Println("We are not recieving the Post Request from client")
		return
	}
	// recieve the json data
	res, _ := io.ReadAll(r.Body)
	// convert the json data into struct Unmarshal
	// create a obj of type payload
	p1 := xy.Payload{}
	json.Unmarshal(res, &p1)
	// Print the struct
	w.WriteHeader(http.StatusOK)
	service.PrintMetrics(p1)

}

func main() {
	// defiene the port number
	port := "8080"
	// create router
	router := http.NewServeMux()

	// launch the handler
	router.HandleFunc("/", handler)

	// listen and Serve at port number
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Println("We Encountered Error in Listen And Serve")
		return
	}
}
