package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("Subcriber server is running...")
	http.HandleFunc("/", sendEmail)
	http.ListenAndServe(":8080", nil)
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer r.Body.Close()
	log.Println(string(requestBody))

	// write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{message: "accepted"}`))
}
