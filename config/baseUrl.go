package config

import "os"

var (
	// BaseURL is the base url of the server
	BaseURL = os.Getenv("BASE_URL")
)

func SetUrl(url string) string {
	if BaseURL == "" {
		BaseURL = "https://fcp-webapp-sikm-ruangguru-production.up.railway.app/"
	}

	return BaseURL + url
}
