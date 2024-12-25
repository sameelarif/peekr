package util

import (
	"fmt"
	"main/models"

	http "github.com/useflyent/fhttp"
)

func ParseWebsite(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header = http.Header{
		"Accept":                      {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"Accept-Encoding":             {"gzip, deflate, br, zstd"},
		"Accept-Language":             {"en-US,en;q=0.9"},
		"If-None-Match":               {"\"s5tu3vn58o6bzh\""},
		"Priority":                    {"u=0, i"},
		"Sec-CH-UA":                   {"\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\""},
		"Sec-CH-UA-Full-Version-List": {"\"Google Chrome\";v=\"131.0.6778.109\", \"Chromium\";v=\"131.0.6778.109\", \"Not_A Brand\";v=\"24.0.0.0\""},
		"Sec-CH-UA-Mobile":            {"?0"},
		"Sec-CH-UA-Model":             {"\"\""},
		"Sec-CH-UA-Platform":          {"\"Windows\""},
		"Sec-CH-UA-Platform-Version":  {"\"19.0.0\""},
		"Sec-Fetch-Dest":              {"document"},
		"Sec-Fetch-Mode":              {"navigate"},
		"Sec-Fetch-Site":              {"none"},
		"Sec-Fetch-User":              {"?1"},
		"Upgrade-Insecure-Requests":   {"1"},
		"User-Agent":                  {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"},
		http.HeaderOrderKey: {
			"accept",
			"accept-encoding",
			"accept-language",
			"if-none-match",
			"priority",
			"sec-ch-ua",
			"sec-ch-ua-full-version-list",
			"sec-ch-ua-mobile",
			"sec-ch-ua-model",
			"sec-ch-ua-platform",
			"sec-ch-ua-platform-version",
			"sec-fetch-dest",
			"sec-fetch-mode",
			"sec-fetch-site",
			"sec-fetch-user",
			"upgrade-insecure-requests",
			"user-agent",
		},
		http.PHeaderOrderKey: {
			":method",
			":path",
			":authority",
			":scheme",
		},
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the website: %v", err)
	}
	defer resp.Body.Close()

	services := []string{}

	if DetectAkamai(resp) {
		services = append(services, models.Akamai.String())
	}
	if DetectAkamaiPixel(resp) {
		services = append(services, models.AkamaiPixel.String())
	}
	if DetectCloudflare(resp) {
		services = append(services, models.Cloudflare.String())
	}
	if DetectImperva(resp) {
		services = append(services, models.Imperva.String())
	}
	if DetectPerimeterX(resp) {
		services = append(services, models.PerimeterX.String())
	}
	if DetectReblaze(resp) {
		services = append(services, models.Reblaze.String())
	}
	if DetectRadware(resp) {
		services = append(services, models.Radware.String())
	}
	if DetectDataDome(resp) {
		services = append(services, models.DataDome.String())
	}
	if DetectKasada(resp) {
		services = append(services, models.Kasada.String())
	}

	return services, nil
}
