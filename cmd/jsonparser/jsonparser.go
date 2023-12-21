package jsonparser

import (
	"encoding/json"
	"github.com/danieljancar/go-proxy-request-checker/cmd/httprequests"
	"io/ioutil"
	"log"
)

func ParseFromJsonFile(filePath string, target *httprequests.ProxyObjects) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Using default values...")
		*target = httprequests.ProxyObjects{
			{
				URL:              "https://www.google.com",
				ExpectedResponse: 200,
			},
			{
				URL:              "https://www.slack.com",
				ExpectedResponse: 200,
			},
		}
		return nil
	}

	err = json.Unmarshal(file, target)
	return err
}
