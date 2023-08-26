package system

import "github.com/go-resty/resty/v2"

func Get(url string, headers map[string]string) string {
	client := resty.New().R()

	for key, value := range headers {
		client.SetHeader(key, value)
	}

	response, err := client.Get(url)

	if err != nil {
		println("failed to call =>", url)
		return ""
	}

	return response.String()
}

func Post(url string, payload map[string]string, headers map[string]string) string {
	return ""
}
