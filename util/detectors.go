package util

import (
	"compress/gzip"
	"io"
	"net/http"
	"regexp"
	"slices"
	"strings"
)

// AKA_A2 - Akamai cache
// bm_* - Antibot
// ak_bmsc - Antibot
// _abck (deprecated) - Antibot
func DetectAkamai(resp *http.Response) bool {
	cookies := resp.Cookies()
	akamaiCookies := []string{"_abck", "bm_sv", "ak_bmsc", "akavpau_p2", "bm_"}

	for _, cookie := range cookies {
		if slices.Contains(akamaiCookies, cookie.Name) {
			return true
		}
	}

	return false
}

// <noscript>
// <img src="https://www.vhnissanstevenspoint.com/akam/13/pixel_2d0a6cd4?a=dD1kMjQ2Y2IyNDg1MDc1NzVhNGRkNDNlNzg5NmNlZjQ1Mjk2ZjEyYjkxJmpzPW9mZg==" style="visibility: hidden; position: absolute; left: -999px; top: -999px;"/>
// </noscript>
func DetectAkamaiPixel(resp *http.Response) bool {
	return false
}

func DetectCloudflare(resp *http.Response) bool {
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if strings.HasPrefix(cookie.Name, "__cf_") {
			return true
		}
	}
	return false
}

func DetectImperva(resp *http.Response) bool {
	return false
}

// _pxvid - Browser detection
// _px* - Session
// _pxff_* - Fingerprint
// _pxhd - Server-side detection
// _pxde - Data enrichment
// window._pxAppId - Session
func DetectPerimeterX(resp *http.Response) bool {
	cookies := resp.Cookies()
	perimeterXCookies := []string{"_pxvid", "_px", "_pxff_", "_pxhd", "_pxde"}

	for _, cookie := range cookies {
		for _, pxCookie := range perimeterXCookies {
			if strings.HasPrefix(cookie.Name, pxCookie) {
				return true
			}
		}
	}

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

	return strings.Contains(html, "_pxAppId")
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

// KPSDK.configure 
// <script src="{uuid}/{uuid}/p.js"></script>
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

	if strings.Contains(html, "KPSDK.configure") {
		return true
	}

	re := regexp.MustCompile(`<script src="/[0-9a-fA-F-]+/[0-9a-fA-F-]+/p.js"></script>`)
	
	return re.MatchString(html)
}