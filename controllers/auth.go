package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Auth endpoint
func Auth(w http.ResponseWriter, request *http.Request) {

	response, err := http.Post(os.Getenv("URL_AUTH"), "application/json", request.Body)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {
		data, _ := ioutil.ReadAll(response.Body)
		// fmt.Println(string(data))
		w.WriteHeader(response.StatusCode)
		w.Write(data)
	}

}

// Hello endpoint
func Hello(w http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(w, "hello world")

}
