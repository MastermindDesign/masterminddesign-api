package utils

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

var domain = "https://search-mastermind-nmwc6cisf7mszachqwe7rb34hi.us-east-2.es.amazonaws.com"
var region = "us-east-2" // e.g. us-east-1
var service = "es"

var awsCredentials *credentials.Credentials
var awsSigner *v4.Signer

func init() {
	awsCredentials = credentials.NewEnvCredentials()
	awsSigner = v4.NewSigner(awsCredentials)
}

func LogToElasticsearch(index string, jsonString string) {
	client := &http.Client{}
	body := strings.NewReader(jsonString)
	endpoint := domain + "/" + index + "/" + "_doc"

	req, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		fmt.Print(err)
	}

	req.Header.Add("Content-Type", "application/json")

	awsSigner.Sign(req, body, service, region, time.Now())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(resp)
	}
}

type HTTPReqInfo struct {
	Method    string        `json:"method"`
	Uri       string        `json:"uri"`
	Referer   string        `json:"referer"`
	Ipaddr    string        `json:"ip"`
	Code      int           `json:"code"`
	Size      int64         `json:"size"`
	Duration  time.Duration `json:"duration"`
	UserAgent string        `json:"userAgent"`
}
