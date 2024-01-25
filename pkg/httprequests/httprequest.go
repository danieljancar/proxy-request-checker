package httprequests

import (
	"github.com/danieljancar/go-proxy-request-checker/pkg/proxyresponse"
	"github.com/danieljancar/go-proxy-request-checker/pkg/reportgenerator"
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
		statusCode, htmlValues, err := proxy.SendRequest()
		if err != nil {
			log.Printf("Error sending request: %s\n", err)
			continue
		}
		proxyresponse.IsProxyHandlingValid(statusCode, proxy.URL)
		report.AddRequest(proxy.URL, proxy.ExpectedResponse, statusCode, htmlValues)
	}
}

func (p ProxyObject) SendRequest() (int, []string, error) {
	response, err := http.Get(p.URL)
	if err != nil {
		return 0, nil, err
	}
	defer response.Body.Close()

	depth, tag := proxyresponse.GetHTMLParseConfig()

	htmlValues := proxyresponse.ParseHTML(response.Body, depth, tag)

	return response.StatusCode, htmlValues, nil
}
