package domain

import (
	"time"

	"github.com/go-resty/resty/v2"
	teambitionapi "github.com/kexin8/teambition-sdk-go/teambition-api"
	"github.com/pkg/errors"
)

type Node struct {
	ID          string        `json:"_id"`
	AncestorIds []interface{} `json:"ancestorIds"`
	Name        string        `json:"name"`
	ObjectType  string        `json:"objectType"`
	ObjectID    string        `json:"objectId"`
	ParentID    string        `json:"parentId"`
	Order       int           `json:"order"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	V           int           `json:"__v"`
	Data        Data          `json:"data"`
	Children    []Node        `json:"children"`
}

type Data struct {
	ID            string `json:"_id"`
	Title         string `json:"title"`
	PermissionIds []any  `json:"permissionIds"`
	Status        string `json:"status"`
	Type          string `json:"type"`
	Deprecated    bool   `json:"deprecated"`
	BasePath      string `json:"basePath"`
	Method        string `json:"method"`
	Path          string `json:"path"`
	OperationID   string `json:"operationId"`
	Summary       string `json:"summary"`
	DocURL        string `json:"docUrl"`
	DebugURL      string `json:"debugUrl"`
	FullPath      string `json:"fullPath"`
}

func DocNode() (node []Node, err error) {
	r := new(teambitionapi.Response[[]Node])

	c := resty.New()
	resp, err := c.R().SetResult(r).Get("https://open.teambition.com/docs/api/doc/node")
	if err != nil {
		return node, errors.WithStack(err)
	}
	if resp.IsError() {
		return node, errors.New(resp.String())
	}

	return r.Result, nil
}
