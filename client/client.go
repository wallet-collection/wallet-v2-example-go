package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	url    string
	client *http.Client
}

type Res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewClient 创建
func NewClient(url string, timeout int64) *Client {
	return &Client{
		url,
		&http.Client{
			Timeout: time.Millisecond * time.Duration(timeout),
		},
	}
}

// get 请求
func (w *Client) get(path string, header http.Header, params url.Values, res interface{}) error {

	return w.request(path, http.MethodGet, header, params, nil, res)

}

// post 请求
func (w *Client) post(path string, header http.Header, data map[string]interface{}, res interface{}) error {
	return w.request(path, http.MethodPost, header, nil, data, res)
}

// 请求
func (w *Client) request(path string, method string, header http.Header, params url.Values, data map[string]interface{}, res interface{}) error {
	var reqBody io.Reader
	if data != nil {

		bytesData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		reqBody = bytes.NewReader(bytesData)
	}

	urlStr := w.url + path

	if params != nil {
		Url, err := url.Parse(urlStr)
		if err != nil {
			return err
		}
		//如果参数中有中文参数,这个方法会进行URLEncode
		Url.RawQuery = params.Encode()
		urlStr = Url.String()
	}

	//fmt.Println(urlStr)

	req, err := http.NewRequest(method, urlStr, reqBody)
	if header != nil {
		req.Header = header
	}
	if err != nil {
		// handle error
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := w.client.Do(req)
	if err != nil {
		// handle error
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return err
	}

	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}

	return nil
}
