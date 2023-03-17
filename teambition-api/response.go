package teambitionapi

type Response[T any] struct {
	Code         int    `json:"code"`         //返回码，200 表示成功
	ErrorMessage string `json:"errorMessage"` // 调用失败时的错误信息

	NextPageToken string `json:"nextPageToken,omitempty"`
	TotalSize     int    `json:"totalSize,omitempty"`

	Result T `json:"result"`
}
