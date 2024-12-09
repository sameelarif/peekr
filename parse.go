package main

import (
	"fmt"
	"net/http"
)

func ParseWebsite(url string) ([]BotMitigationService, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the website: %v", err)
	}
	defer resp.Body.Close()


	services := []BotMitigationService{}
	
	if DetectAkamai(resp) {
		services = append(services, Akamai)
	}
	if DetectAkamaiPixel(resp) {
		services = append(services, AkamaiPixel)
	}
	if DetectCloudflare(resp) {
		services = append(services, Cloudflare)
	}
	if DetectImperva(resp) {
		services = append(services, Imperva)
	}
	if DetectPerimeterX(resp) {
		services = append(services, PerimeterX)
	}
	if DetectReblaze(resp) {
		services = append(services, Reblaze)
	}
	if DetectRadware(resp) {
		services = append(services, Radware)
	}
	if DetectDataDome(resp) {
		services = append(services, DataDome)
	}
	if DetectKasada(resp) {
		services = append(services, Kasada)
	}
}
