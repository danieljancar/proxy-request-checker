package configparser

import (
	"encoding/json"
	"log"
	"os"

	"github.com/danieljancar/go-proxy-request-checker/pkg/httprequests"
)

func ParseFromJsonFile(filePath string, target *httprequests.ProxyObjects) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading json file: %s. Using default values...", err)
		*target = defaultProxyObjects()
		return nil
	}

	err = json.Unmarshal(fileData, target)
	if err != nil {
		log.Printf("Error unmarshalling json file: %s", err)
		*target = defaultProxyObjects()
		return err
	}

	return nil
}

func defaultProxyObjects() httprequests.ProxyObjects {
	return httprequests.ProxyObjects{
		{URL: "https://www.google.com", ExpectedResponse: 200},
	}
}
