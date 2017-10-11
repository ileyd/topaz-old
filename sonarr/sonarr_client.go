package sonarr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/Sirupsen/logrus"
)

type SonarrClient struct {
	address    *url.URL
	apiKey     string
	HttpClient *http.Client
}

func NewSonarrClient(address string, apiKey string) (*SonarrClient, error) {

	if address == "" {
		return nil, errors.New("No address specified")
	}

	addressUrl, err := url.Parse(address)

	path := addressUrl.Path
	//
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	if !strings.HasSuffix(path, "api/") {
		path += "api/"
	}

	addressUrl.Path = path

	if err != nil {
		return nil, err
	}

	return &SonarrClient{
		address:    addressUrl,
		apiKey:     apiKey,
		HttpClient: http.DefaultClient,
	}, nil
}

func (sc *SonarrClient) DoRequest(action, path string, params map[string]string, reqData, resData interface{}) error {
	lookupUrl := *sc.address

	parameters := url.Values{}

	if params != nil {
		for k, v := range params {
			parameters.Add(k, v)
		}
	}

	lookupUrl.RawQuery = parameters.Encode()
	lookupUrl.Path += path

	jsonValue, err := json.Marshal(reqData)

	if err != nil {
		return err
	}

	logrus.Debugf("Calling Sonarr at %v", lookupUrl.String())

	req, err := http.NewRequest(action, lookupUrl.String(), bytes.NewBuffer(jsonValue))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-Api-Key", sc.apiKey)

	response, err := sc.HttpClient.Do(req)

	if err != nil {
		return err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(response.Body)

		if err == nil {
			logrus.Debugf("Failing (%v) call returned:\n%v", response.StatusCode, string(bodyBytes))
		}

		return errors.New(fmt.Sprintf("Status code %v", response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(resData)

	if err != nil {
		return err
	}

	return nil
}
