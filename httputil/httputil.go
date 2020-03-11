package httputil

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

const HTTP_CONTENT_TYPE_JSON = "application/json;charset=utf-8"

// Get request
func DoGet(url string) interface{} {
	// GET
	log.Info("-------- Get request start -------")
	log.Infof("url : %v\n", url)
	resp, err := http.Get(url)
	r := handleResponse(resp, err)
	log.Info("-------- Get request end -------")
	return r
}

// Post request
func DoPost(url string, contentType string, requestBody string) interface{} {
	// POST
	log.Info("-------- Post request start -------")
	log.Infof("url : %v\n", url)
	log.Infof("contentType : %v\n", contentType)
	log.Infof("requestBody : %v\n", requestBody)
	resp, err := http.Post(url, contentType, strings.NewReader(requestBody))
	r := handleResponse(resp, err)
	log.Info("-------- Post request end -------")
	return r
}

func handleResponse(resp *http.Response, err error) interface{} {
	if err != nil {
		log.Error(err)
		return nil
	}
	result, err := readBody(resp)
	if err != nil {
		log.Error(err)
		return nil
	}
	return result
}

func readBody(resp *http.Response) (interface{}, error) {
	// Response handling
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result interface{}
	json.Unmarshal(body, &result)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Panic("resp.Body.Close() failed.")
		}
	}()
	return result, nil
}
