package teambitionapi

type Response[T any] struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`

	NextPageToken string `json:"nextPageToken,omitempty"`
	TotalSize     int    `json:"totalSize,omitempty"`

	Result T `json:"result"`
}
