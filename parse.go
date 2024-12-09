package main

import (
	"fmt"
	"net/http"
)

func ParseWebsite(url string) ([]BotMitigationService, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-encoding", "gzip, deflate, br, zstd")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("if-none-match", "\"s5tu3vn58o6bzh\"")
	req.Header.Set("priority", "u=0, i")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"")
	req.Header.Set("sec-ch-ua-full-version-list", "\"Google Chrome\";v=\"131.0.6778.109\", \"Chromium\";v=\"131.0.6778.109\", \"Not_A Brand\";v=\"24.0.0.0\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-model", "\"\"")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-ch-ua-platform-version", "\"19.0.0\"")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
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

	return services, nil
}
