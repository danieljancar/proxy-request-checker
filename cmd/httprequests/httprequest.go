package httprequests

import (
	"bufio"
	"fmt"
	"github.com/danieljancar/go-proxy-request-checker/cmd/reportgenerator"
	"log"
	"net/http"
	"os"
	"strings"
)

type ProxyObject struct {
	URL              string `json:"url"`
	ExpectedResponse int    `json:"expectedResponse"`
}

type ProxyObjects []ProxyObject

func (objects ProxyObjects) Init(report *reportgenerator.Report) {
	reader := bufio.NewReader(os.Stdin)

	for _, object := range objects {
		log.Printf("Requesting %s\n", object.URL)
		object.Request(report)
	}

	fmt.Print("Do you wish to export the results? (y/n): ")
	result, _, err := reader.ReadRune()
	if err != nil {
		log.Printf("Error reading input: %s\n", err)
		return
	}

	if strings.ToLower(string(result)) == "y" {
		filePath := "report.json"
		fmt.Print("Enter file path to save the report (default: report.json): ")

		reader.ReadString('\n')

		input, err := reader.ReadString('\n')
		if err == nil && strings.TrimSpace(input) != "" {
			filePath = strings.TrimSpace(input)
		}

		err = report.SaveToFile(filePath)
		if err != nil {
			log.Printf("Error saving report: %s\n", err)
		}
	}
}

func (object ProxyObject) Request(report *reportgenerator.Report) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", object.URL, nil)
	if err != nil {
		log.Printf("Error creating request: %s\n", err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error sending request: %s\n", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == object.ExpectedResponse {
		log.Printf("Request to %s successful with expected status code %d\n", object.URL, object.ExpectedResponse)
	} else {
		log.Printf("Unexpected status code %d for %s\n", response.StatusCode, object.URL)
	}

	report.AddRequest(object.URL, object.ExpectedResponse, response.StatusCode)
}
