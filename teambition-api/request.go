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
func Request[T any](method, url string, body any, headers map[string]string) (result T, resp Response[T], err error) {

	if ApiClient.ValidToken() {
		if err := ApiClient.RefreshToken(); err != nil {
			return result, resp, errors.WithStack(err)
		}
	}

	response, err := ApiClient.R().
		SetBody(body).
		SetResult(new(Response[T])).
		SetHeaders(headers).
		Execute(method, url)

	if err != nil {
		return result, resp, errors.WithStack(err)
	}

	if !response.IsSuccess() {
		return result, resp, errors.Errorf("请求失败: %s", response.String())
	}

	respBody := response.Result().(*Response[T])

	if respBody.Code != 200 {
		return result, resp, errors.Errorf("请求失败: %s", respBody.ErrorMessage)
	}

	return respBody.Result, *respBody, nil
}

// Get 请求
func Get[T any](url string, headers map[string]string) (result T, resp Response[T], err error) {
	return Request[T](GET, url, nil, headers)
}

// Post 请求
func Post[T any](url string, body any, headers map[string]string) (result T, resp Response[T], err error) {
	return Request[T](POST, url, body, headers)
}

// PostJson 请求
func PostJson[T any](url string, body any, headers map[string]string) (result T, resp Response[T], err error) {
	headers["Content-Type"] = "application/json"
	return Request[T](POST, url, body, headers)
}

// Put 请求
func Put[T any](url string, body any, headers map[string]string) (result T, resp Response[T], err error) {
	return Request[T](PUT, url, body, headers)
}

// Delete 请求
func Delete[T any](url string, body any, headers map[string]string) (result T, resp Response[T], err error) {
	return Request[T](DELETE, url, body, headers)
}
