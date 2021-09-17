package lib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Curl interface {
	Call() ResponseCurl
}

type Header struct {
	Key   string
	Value string
}

type curl struct {
	Url    string
	Method string
	Body   string
	Header []Header
}

func GetInstanceCurl(url string, method string, body string, header []Header) Curl {
	return &curl{
		Url:    url,
		Method: method,
		Body:   body,
		Header: header,
	}
}

type ResponseCurl struct {
	Status int
	Data   interface{}
}

func (c *curl) Call() ResponseCurl {
	client := &http.Client{}

	payload := strings.NewReader(c.Body)
	req, err := http.NewRequest(c.Method, c.Url, payload)

	if err != nil {
		panic(err)
	}
	if len(c.Header) > 0 {
		for _, v := range c.Header {
			req.Header.Add(v.Key, v.Value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var response ResponseCurl

	json.Unmarshal([]byte(string(body)), &response)
	return response
}
