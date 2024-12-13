package main

import (
	"fmt"
	"main/util"
	"testing"
)

func TestParseWebsite(t *testing.T) {
	// Define the URL to test
	url := "https://www.footlocker.com/" // Replace with the desired URL

	// Call the ParseWebsite function
	services, err := util.ParseWebsite(url)
	if err != nil {
		t.Fatalf("Error parsing website: %v", err)
	}

	// Print the detected services
	t.Log("Detected Bot Mitigation Services:")
	for _, service := range services {
		fmt.Println(service)
	}
}

