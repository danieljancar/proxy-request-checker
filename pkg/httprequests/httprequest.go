package httprequests

import (
	"github.com/danieljancar/go-proxy-request-checker/pkg/proxyresponse"
	"github.com/danieljancar/go-proxy-request-checker/pkg/reportgenerator"
	"io"
	"log"
	"net/http"
)

type ProxyObject struct {
	URL              string `json:"url"`
	ExpectedResponse int    `json:"expectedResponse"`
}

type ProxyObjects []ProxyObject

func ProcessProxies(proxies *ProxyObjects, report *reportgenerator.Report) {
	for _, proxy := range *proxies {
		log.Printf("Requesting %s\n", proxy.URL)
		statusCode, err := proxy.SendRequest()
		if err != nil {
			log.Printf("Error sending request: %s\n", err)
			continue
		}
		proxyresponse.IsProxyHandlingValid(statusCode, proxy.URL)
		report.AddRequest(proxy.URL, proxy.ExpectedResponse, statusCode)
	}
}

func (p ProxyObject) SendRequest() (int, error) {
	response, err := http.Get(p.URL)
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %s\n", err)
		}
	}(response.Body)

	return response.StatusCode, nil
}
