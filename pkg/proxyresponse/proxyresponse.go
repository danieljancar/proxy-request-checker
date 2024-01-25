package proxyresponse

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	HTMLParsing struct {
		Depth int    `json:"depth"`
		Tag   string `json:"tag"`
	} `json:"html_parsing"`
}

var config Config

func init() {
	data, err := os.ReadFile("config/proxy/proxy-config.json")
	if err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
		return
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing configuration file: %s", err)
		return
	}
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
			}
		}
	}
}

func GetHTMLParseConfig() (int, string) {
	return config.HTMLParsing.Depth, config.HTMLParsing.Tag
}

func ParseHTML(body io.ReadCloser, depth int, tag string) []string {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Printf("Error reading HTML document: %s", err)
		return nil
	}

	var results []string
	tag = strings.ToLower(tag)

	doc.Find(tag).Each(func(i int, s *goquery.Selection) {
		if i < depth {
			content := s.Text()
			results = append(results, content)
		}
	})

	return results
}
