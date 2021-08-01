package form3

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Options struct {
	Base    string
	Timeout int
}

type Client struct {
	BasePath   *url.URL
	httpClient *http.Client
}

type service struct {
	form3Client *Client
}

type payloadBody struct {
	Data interface{} `json:"data"`
}

type Response struct {
	*http.Response
}

type respError struct {
	ErrorMSG string `json:"error_message"`
}

func (c *Client) createRequest(method string, url string, payload interface{}) (*http.Request, error) {
	fmtEndPoint := fmt.Sprintf("%s/%s", c.BasePath.String(), url)
	var data io.Reader
	if payload != nil {
		body := payloadBody{Data: payload}
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		data = bytes.NewBuffer(jsonData)
	}
	reqest, err := http.NewRequest(method, fmtEndPoint, data)
	if err != nil {
		return nil, err
	}
	reqest.Header.Set("Content-Type", "application/json")
	return reqest, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, obj interface{}) (*Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	fmt.Println("???????????????????", err.Error())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("???????????????????", err.Error())
	response := &Response{Response: resp}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, nil
	}

	repErr := &respError{}
	err = json.Unmarshal(resData, repErr)
	if err != nil && repErr.ErrorMSG != "" {
		return response, errors.New(repErr.ErrorMSG)
	}

	if obj != nil {
		result := &payloadBody{}
		err := json.Unmarshal(resData, result)
		if err != nil {
			return response, err
		}

		encodeData, err := json.Marshal(result.Data)
		if err != nil {
			err = json.Unmarshal(encodeData, obj)
		}
	}
	return response, err
}

func (c *Client) GET(url string, reqBody interface{}) (*http.Request, error) {
	return c.createRequest("GET", url, reqBody)
}

func (c *Client) POST(url string, reqBody interface{}) (*http.Request, error) {
	return c.createRequest("POST", url, reqBody)
}

func (c *Client) DELETE(url string, reqBody interface{}) (*http.Request, error) {
	return c.createRequest("DELETE", url, reqBody)
}

func CreateClient(options *Options) *Client {
	if options == nil {
		options = &Options{
			Timeout: 3000000,
			Base:    "http://localhost:8080/v1",
		}
	}

	baseURL, _ := url.Parse(options.Base)

	return &Client{
		httpClient: &http.Client{Timeout: time.Duration(options.Timeout)},
		BasePath:   baseURL,
	}
}
