package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// inputs
	fmt.Print("Enter your zipcode: ")

	var zipcode string

	fmt.Scanln(&zipcode)

	fmt.Println("you searched for:" + zipcode)

	// http request

	webservice := "https://viacep.com.br/ws/" + zipcode + "/json"

	req, err := http.NewRequest("GET", webservice, nil)

	// req res
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))

}
