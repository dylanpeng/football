package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Client struct {
	client *http.Client
}

func (c *Client) Get(url string, header map[string][]string, body []byte) (rspCode int, rspBody []byte, err error) {
	return c.Do(http.MethodGet, url, header, body)
}

func (c *Client) Post(url string, header map[string][]string, body []byte) (rspCode int, rspBody []byte, err error) {
	return c.Do(http.MethodPost, url, header, body)
}

func (c *Client) Do(method, url string, header map[string][]string, body []byte) (rspCode int, rspBody []byte, err error) {
	var req *http.Request

	if len(body) > 0 {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return
	}

	if len(header) > 0 {
		req.Header = header
	}

	response, err := c.client.Do(req)

	if err != nil {
		return
	}

	rspCode = response.StatusCode
	rspBody, err = io.ReadAll(response.Body)

	if err != nil {
		return
	}

	if rspCode != http.StatusOK {
		err = fmt.Errorf("error http code %d", rspCode)
		return
	}

	return
}

func NewClient(timeout time.Duration) *Client {
	cookie, _ := cookiejar.New(nil)
	return &Client{client: &http.Client{Jar: cookie, Timeout: timeout}}
}
