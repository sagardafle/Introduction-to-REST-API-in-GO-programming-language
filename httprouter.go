package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Input Struct of json string from Postman to localhost
type Input struct {
	Name string `json:"name"`
}

//Greet Struct of json string to be displayed on the POSTMAN
type Greet struct {
	Greeting string `json:"greeting"`
}

func sayHello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var t Input
	if err := json.NewDecoder(req.Body).Decode(&t); err != nil {
		fmt.Println(err)
	}
	output := fmt.Sprintf("Hello, %s!", t.Name)
	m := &Greet{
		Greeting: output,
	}
	js, err := json.Marshal(m)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

}

func main() {
	mux := httprouter.New()
	mux.POST("/sayhello", sayHello)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
