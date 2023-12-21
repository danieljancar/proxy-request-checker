package main

import (
	"github.com/danieljancar/go-proxy-request-checker/cmd/httprequests"
	"github.com/danieljancar/go-proxy-request-checker/cmd/jsonparser"
	"github.com/danieljancar/go-proxy-request-checker/cmd/reportgenerator"
	"log"
)

func main() {
	var proxies httprequests.ProxyObjects
	report := reportgenerator.NewReport()

	err := jsonparser.ParseFromJsonFile("config/links.json", &proxies)
	if err != nil {
		log.Fatalf("Error parsing json file: %s", err.Error())
		return
	}

	proxies.Init(report)
}
