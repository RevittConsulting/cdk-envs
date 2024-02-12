package jsonrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) *Client {
	return &Client{
		url: url,
	}
}

func (c *Client) DoRequest(method string, params ...interface{}) (*Response, error) {
	par, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req := &Request{
		JsonRpc: "2.0",
		Id:      1,
		Method:  method,
		Params:  par,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpRes, err := http.Post(c.url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", httpRes.StatusCode)
	}

	resBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	var res *Response
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
