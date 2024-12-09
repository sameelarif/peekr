package main

import (
	"fmt"
)

func main() {
	url := "https://www.zillow.com/"

	services, err := ParseWebsite(url) 
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Detected Bot Mitigation Services:")
	for _, service := range services {
		fmt.Println(service)
	}
}