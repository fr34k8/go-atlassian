package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
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

	options := &confluence.GetContentOptionsScheme{
		//ContextType: "page",
		SpaceKey:    "",
		Title:       "",
		Trigger:     "",
		OrderBy:     "",
		//Status:      []string{"any", "any"},
		//Expand:      []string{"childTypes.all", "space"},
		//PostingDay:  time.Now(),
	}

	page, response, err := instance.Content.Get(context.Background(), options, 0, 50)
	if err != nil {

		if response != nil {
			log.Fatal(response.API)
		}
		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	for _, content := range page.Results {
		log.Println(content)
	}
}
