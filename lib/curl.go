package lib

import (
	"io/ioutil"
	"net/http"
	"strings"

	"fast.bibabo.vn/mongo_models"
)

type Curl interface {
	Call() []byte
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

type ResponsePromotionCurl struct {
	Success bool
	Data    map[string]mongo_models.Promotion
}

func (c *curl) Call() []byte {
	client := &http.Client{}

	payload := strings.NewReader(c.Body)
	req, err := http.NewRequest(c.Method, c.Url, payload)

	if err != nil {
		panic(err)
	}

	var headers []Header

	headers = append(headers, Header{
		Key:   "Content-Type",
		Value: "application/json",
	})
	if c.Header != nil {
		for _, h := range c.Header {
			headers = append(headers, h)
		}
	}
	if len(headers) > 0 {
		for _, v := range headers {
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
	return body
}
