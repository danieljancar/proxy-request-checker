package proxyresponse

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	HTMLParsing struct {
		Depth    int    `json:"depth"`
		Tag      string `json:"tag"`
		Contains string `json:"contains"`
	} `json:"html_parsing"`
}

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
	data, err := ioutil.ReadFile("config/proxy/proxy-config.json")
	if err != nil {
		log.Printf("Error reading configuration file: %s", err)
		return
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Printf("Error parsing configuration file: %s", err)
		return
	}

	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error making GET request: %s", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return
	}
	log.Print("Response body: ", string(body))

	result := parseHTML(response.Body, config.HTMLParsing.Depth, config.HTMLParsing.Tag, config.HTMLParsing.Contains)
	log.Print("Parsed HTML result: ", result)
}

func parseHTML(body io.ReadCloser, depth int, tag string, contains string) []string {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Printf("Error reading HTML document: %s", err)
		return nil
	}

	var results []string
	tag = strings.ToLower(tag)

	doc.Find(tag).Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		if contains == "" || strings.Contains(strings.ToLower(content), strings.ToLower(contains)) {
			results = append(results, content)
		}
	})

	return results
}
