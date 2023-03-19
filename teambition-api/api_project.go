package teambitionapi

import (
	"fmt"

	"github.com/pkg/errors"
)

type Project struct {
	Id             string `json:"id"`             //项目ID
	Name           string `json:"name"`           //项目名称
	Logo           string `json:"logo"`           //项目LOGO
	Description    string `json:"description"`    //项目描述
	OrganizationId string `json:"organizationId"` //企业ID
	Visibility     string `json:"visibility"`     //可见性，project | organization
	IsTemplate     bool   `json:"isTemplate"`     //是模版项目
	CreatorId      string `json:"creatorId"`      //创建人ID
	IsArchived     bool   `json:"isArchived"`     //是否放入回收站
	IsSuspended    bool   `json:"isSuspended"`    //是否归档
	UniqueIdPrefix string `json:"uniqueIdPrefix"` //任务ID前缀
	Created        string `json:"created"`        //创建时间
	Updated        string `json:"updated"`        //更新时间
	StartDate      string `json:"startDate"`      //项目开始时间
	EndDate        string `json:"endDate"`        //项目结束时间

	Customfields []Customfield `json:"customfields"` //自定义字段
}

type ProjectApplication struct {
	Id      string `json:"id"`      // 应用ID
	AppId   string `json:"appId"`   // 应用ID
	Version string `json:"version"` // 应用版本
	Enabled bool   `json:"enabled"` // 是否启用
}

type ProjectMember struct {
	Id      string `json:"id"`      // 项目成员ID
	UserID  string `json:"userId"`  // 用户ID
	Role    string `json:"role"`    // 项目角色，0=成员；1=管理员；2=拥有者
	RoleIds string `json:"roleIds"` // 成员角色ID
}

// 标签
type ProjectTag struct {
	Color          string   `json:"color"`          // 标签颜色
	Created        string   `json:"created"`        // 创建时间
	CreatorID      string   `json:"creatorId"`      // 创建人ID
	ID             string   `json:"id"`             // 标签ID
	IsArchived     bool     `json:"isArchived"`     // 是否归档
	Name           string   `json:"name"`           // 标签名
	OrganizationID string   `json:"organizationId"` // 企业ID
	ProjectID      string   `json:"projectId"`      // 项目ID
	TagcategoryIDS []string `json:"tagcategoryIds"`
	Updated        string   `json:"updated"` // 更新时间
}

// 查询项目应用列表
// GET https://open.teambition.com/api/v3/project/{projectId}/application/list
// 接口地址：https://open.teambition.com/docs/apis/6321c6d0912d20d3b5a496c1
// @param projectId 项目ID
// @param appIds 应用ID集合，逗号分隔，如果传递该参数仅查询指定 应用ID
// @param scope installed 或者 all，默认查询已经安装的应用
func (c *Client) GetProjectApplications(projectId string, appIds string, scope string) (resp *Response[[]ProjectApplication], err error) {
	resp, err = Get[[]ProjectApplication](c, fmt.Sprintf("/project/%s/application/list?appIds=%s&scope=%s", projectId, appIds, scope), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 查询项目成员列表
// GET https://open.teambition.com/api/v3/project/{projectId}/member
// 接口地址：https://open.teambition.com/docs/apis/6321c6d0912d20d3b5a49906
// @param projectId 项目ID
// @param userIds 如果传递，仅查询这些用户ID， 用逗号组合
// @param projectRoleId 项目角色ID，仅查询拥有该角色的成员，并且仅支持单个角色查询
// @param limit 每一页数量
// @param skip 分页
// @param pageSize 分页大小
// @param pageToken 分页标识
func (c *Client) GetProjectMembers(projectId string, userIds string, projectRoleId string, limit int, skip int, pageSize int, pageToken string) (resp *Response[[]ProjectMember], err error) {
	resp, err = Get[[]ProjectMember](c, fmt.Sprintf("/v3/project/%s/member?userIds=%s&projectRoleId=%s&limit=%d&skip=%d&pageSize=%d&pageToken=%s", projectId, userIds, projectRoleId, limit, skip, pageSize, pageToken), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 查询项目
//
//	GET https://open.teambition.com/api/v3/project/query
//
// 接口地址： https://open.teambition.com/docs/apis/6321c6d0912d20d3b5a49aa7
// @param projectIds 项目ID列表,多个ID用逗号分隔
// @param name 项目名称
// @param sourceId 原始项目ID
// @param pageToken 分页标识
// @param pageSize 分页大小
func (c *Client) QueryProject(projectIds string, name string, sourceId string, pageToken string, pageSize int) (resp *Response[[]Project], err error) {
	resp, err = Get[[]Project](c, fmt.Sprintf("/v3/project/query?projectIds=%s&name=%s&sourceId=%s&pageToken=%s&pageSize=%d", projectIds, name, sourceId, pageToken, pageSize), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 搜索项目标签
//
//	GET /v3/project/{projectId}/tag/search
//
// @param projectId 项目ID
// @param tagIds 标签ID集合，逗号组合
// @param q 模糊查询标签名字
// @param pageSize 每页长度
// @param pageToken 分页标
func (c *Client) SearchProjectTags(projectId string, tagIds string, q string, pageSize int, pageToken string) (resp *Response[[]ProjectTag], err error) {
	resp, err = Get[[]ProjectTag](c, fmt.Sprintf("/v3/project/%s/tag/search?tagIds=%s&q=%s&pageSize=%d&pageToken=%s", projectId, tagIds, q, pageSize, pageToken), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
