package main

import (
	"compress/gzip"
	"io"
	"net/http"
	"slices"
	"strings"
)

// AKA_A2 - Akamai cache
// bm_sv - Antibot
// ak_bmsc - Antibot
// _abck (deprecated) - Antibot
func DetectAkamai(resp *http.Response) bool {
	cookies := resp.Cookies()
	akamaiCookies := []string{"_abck", "bm_sv", "ak_bmsc"}

	for _, cookie := range cookies {
		if slices.Contains(akamaiCookies, cookie.Name) {
			return true
		}
	}

	return false
}

func DetectAkamaiPixel(resp *http.Response) bool {
	return false
}

func DetectCloudflare(resp *http.Response) bool {
	return false
}

func DetectImperva(resp *http.Response) bool {
	return false
}

func DetectPerimeterX(resp *http.Response) bool {
	return false
}

func DetectReblaze(resp *http.Response) bool {
	return false
}

func DetectRadware(resp *http.Response) bool {
	return false
}

// datadome - DataDome cookie
func DetectDataDome(resp *http.Response) bool {
	cookies := resp.Cookies()
	dataDomeCookies := []string{"datadome"}

	for _, cookie := range cookies {
		if slices.Contains(dataDomeCookies, cookie.Name) {
			return true
		}
	}

	return false
}

func DetectKasada(resp *http.Response) bool {
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return false
	}
	defer reader.Close()

	bytes, err := io.ReadAll(reader)
	if err != nil {
		return false
	}

	html := string(bytes)

	return strings.Contains(html, "KPSDK.configure")
} 