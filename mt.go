package main

import (
	"fmt"
	"github.com/usualoma/mt-data-api-sdk-go"
)

type Entry struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Excerpt string `json:"excerpt"`
}

type EntriesResult struct {
	dataapi.Result
	TotalResults int     `json:"totalResults"`
	Items        []Entry `json:"items"`
}

func main() {
	client := dataapi.NewClient(dataapi.ClientOptionsStruct{
		OptBaseUrl:    "https://www.example.com/mt/mt-data-api.cgi",
		OptApiVersion: "3",
		OptClientId:   "go",
	})

	result := EntriesResult{}
	err := client.SendRequest(
		"GET",
		"/sites/1/entries",
		&dataapi.RequestParameters{
			"limit": "100",
		},
		&result)

	if err != nil {
		panic(err)
	}

	if result.Error != nil {
		panic(result.Error.Message)
	}

	fmt.Printf("total: %d\n", result.TotalResults)
	for _, e := range result.Items {
		fmt.Printf("%d: %s\n excerpt:%s\n", e.Id, e.Title, e.Excerpt)
	}
}
