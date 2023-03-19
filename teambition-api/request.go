package teambitionapi

import (
	"github.com/pkg/errors"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

// Request 请求
func Request[T any](c *Client, method, url string, body any, headers map[string]string) (resp *Response[T], err error) {

	err = c.RefreshToken()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp = &Response[T]{}
	response, err := c.request.R().
		SetBody(body).
		SetResult(resp).
		SetHeaders(headers).
		Execute(method, url)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	if response.IsError() {
		return nil, errors.Errorf("请求失败: %s", response.String())
	}

	if resp.Code != 200 {
		return nil, errors.Errorf("请求失败: %#v", resp)
	}

	return resp, nil
}

// Get 请求
func Get[T any](c *Client, url string, headers map[string]string) (resp *Response[T], err error) {
	return Request[T](c, GET, url, nil, headers)
}

// Post 请求
func Post[T any](c *Client, url string, body any, headers map[string]string) (resp *Response[T], err error) {
	return Request[T](c, POST, url, body, headers)
}

// PostJson 请求
func PostJson[T any](c *Client, url string, body any, headers map[string]string) (resp *Response[T], err error) {
	headers["Content-Type"] = "application/json"
	return Request[T](c, POST, url, body, headers)
}

// Put 请求
func Put[T any](c *Client, url string, body any, headers map[string]string) (resp *Response[T], err error) {
	return Request[T](c, PUT, url, body, headers)
}

// Delete 请求
func Delete[T any](c *Client, url string, body any, headers map[string]string) (resp *Response[T], err error) {
	return Request[T](c, DELETE, url, body, headers)
}
