package main

import (
	"fmt"
	"github.com/danieljancar/go-proxy-request-checker/pkg/httprequests"
	"github.com/danieljancar/go-proxy-request-checker/pkg/jsonparser"
	"github.com/danieljancar/go-proxy-request-checker/pkg/reportgenerator"
	"log"
)

func main() {
	var proxies httprequests.ProxyObjects
	report := reportgenerator.NewReport()

	err := jsonparser.ParseFromJsonFile("config/links.json", &proxies)
	if err != nil {
		log.Fatalf("Error processing json file: %s", err)
		return
	}

	httprequests.ProcessProxies(&proxies, report)

	filePath := fmt.Sprintf("reports/%s_report.json", report.Date)
	if err := report.SaveToFile(filePath); err != nil {
		log.Fatalf("Failed to save the report to file: %s", err)
		return
	}
}
