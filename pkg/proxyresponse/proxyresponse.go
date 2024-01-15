package proxyresponse

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func IsProxyHandlingValid(statusCode int, url string) {
	if os.Getenv("PROXY_HANDLING") == "1" {
		expectedResponses := strings.Split(os.Getenv("PROXY_EXPECTED_RESPONSE"), ",")
		for _, code := range expectedResponses {
			expectedCode, err := strconv.Atoi(code)
			if err != nil {
				continue
			}
			if expectedCode == statusCode {
				log.Printf("Proxy handling is needed for: %s", url)
				readRequestResponse(url)
			}
		}
	}
}

func readRequestResponse(url string) {
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}

	bodyString := string(bodyBytes)
	log.Print(bodyString)
}
