package main

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"net/http"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	instance, err := confluence.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	instance.Auth.SetBasicAuth(mail, token)
	instance.Auth.SetUserAgent("curl/7.54.0")

	var payload = &confluence.ContentPropertyPayloadScheme{
		Key:   "key",
		Value: "value",
	}

	property, response, err := instance.Content.Property.Create(context.Background(), "80412692", payload)
	if err != nil {
		if response != nil {
			if response.Code == http.StatusBadRequest {
				log.Println(response.API)
			}
		}
		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)
	fmt.Println(property)

}