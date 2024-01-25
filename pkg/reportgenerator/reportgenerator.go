package reportgenerator

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Report struct {
	Name     string          `json:"name"`
	Date     string          `json:"date"`
	Status   string          `json:"status"`
	Total    int             `json:"total"`
	Success  int             `json:"success"`
	Fail     int             `json:"fail"`
	Requests []RequestReport `json:"requests"`
}

type RequestReport struct {
	URL              string   `json:"url"`
	ExpectedResponse int      `json:"expectedResponse"`
	ActualResponse   int      `json:"actualResponse"`
	HtmlValues       []string `json:"htmlValues"`
}

func NewReport() *Report {
	return &Report{
		Name:     "Report",
		Date:     time.Now().Format(time.RFC3339),
		Status:   "Success",
		Total:    0,
		Success:  0,
		Fail:     0,
		Requests: []RequestReport{},
	}
}

func (r *Report) Analyze() {
	for _, request := range r.Requests {
		r.Total++
		if request.ExpectedResponse == request.ActualResponse {
			r.Success++
		} else {
			r.Fail++
			r.Status = "Fail"
		}
	}
}

func (r *Report) AddRequest(url string, expectedResponse, actualResponse int, htmlValues []string) {
	r.Requests = append(r.Requests, RequestReport{
		URL:              url,
		ExpectedResponse: expectedResponse,
		ActualResponse:   actualResponse,
		HtmlValues:       htmlValues,
	})
}

func (r *Report) SaveToFile(filePath string) error {
	log.Printf("Saving report to: %s", filePath)
	r.Analyze()
	data, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
