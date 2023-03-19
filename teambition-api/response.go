package teambitionapi

import "fmt"

type Response[T any] struct {
	Code         int    `json:"code"`         //返回码，200 表示成功
	ErrorMessage string `json:"errorMessage"` // 调用失败时的错误信息
	RequestId    string `json:"requestId"`    // 请求 ID，用于排查问题

	NextPageToken string `json:"nextPageToken,omitempty"`
	TotalSize     int64  `json:"totalSize,omitempty"`

	Result T `json:"result"`
}

// impl Response GoStringer
func (r *Response[T]) GoString() string {
	return fmt.Sprintf("requestId: %s, code: %d, errorMessage: %s", r.RequestId, r.Code, r.ErrorMessage)
}
