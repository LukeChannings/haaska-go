package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

var hassURL = os.Getenv("HOME_ASSISTANT_URL")
var hassToken = os.Getenv("HOME_ASSISTANT_TOKEN")
var skipVerify = os.Getenv("SKIP_CERTIFICATE_VERIFICATION") == "true"

// ProxyHomeAssistant - proxies an Alexa Smart Home event to Home Assistant
func ProxyHomeAssistant(event json.RawMessage) (json.RawMessage, error) {

	req, err := http.NewRequest("POST", hassURL+"/api/alexa/smart_home", bytes.NewBuffer(event))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+hassToken)

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
	}}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func main() {
	lambda.Start(ProxyHomeAssistant)
}
