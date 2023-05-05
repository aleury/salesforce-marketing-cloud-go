package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"salesforce-marketing-cloud-go/datafolders"
	"salesforce-marketing-cloud-go/soap"
)

func main() {
	client := soap.NewClient(os.Getenv("SALESFORCE_MARKETING_CLOUD_SOAP_URL"))

	// Create a datafolder
	params := datafolders.CreateParams{
		Name:        "Adams Data Extensions Tests",
		Description: "Testing creating data folders",
		ContentType: datafolders.DataExtension,
	}
	folderId, err := datafolders.Create(context.Background(), client, params)
	if err != nil {
		log.Fatalln("failed to create data folder")
	}
	fmt.Println("Created a new data folder with id:", folderId)

	// Fetch a data folder by its id
	df, err := datafolders.GetById(context.Background(), client, folderId)
	if err != nil {
		log.Fatalln("failed to fetch data folder", err)
	}
	fmt.Printf("Fetched data folder id: %+v\n", df)
}
