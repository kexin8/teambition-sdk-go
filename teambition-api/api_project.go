package teambitionapi

import "fmt"

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

	Customfields []struct {
		CustomfieldId string `json:"customfieldId"` //自定义字段ID
		Typ           string `json:"type"`          //类型
		Value         []struct {
			Id         string `json:"id"`         //自定义字段值ID
			Title      string `json:"title"`      //自定义字段值
			MetaString string `json:"metaString"` //自定义字段值
		} `json:"value"` //自定义字段值
	} `json:"customfields"` //自定义字段
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

// 查询项目应用列表
// GET https://open.teambition.com/api/v3/project/{projectId}/application/list
// 接口地址：https://open.teambition.com/docs/apis/6321c6d0912d20d3b5a496c1
// @param projectId 项目ID
// @param appIds 应用ID集合，逗号分隔，如果传递该参数仅查询指定 应用ID
// @param scope installed 或者 all，默认查询已经安装的应用
func GetProjectApplications(projectId string, appIds string, scope string) ([]ProjectApplication, error) {
	result, _, err := Get[[]ProjectApplication](fmt.Sprintf("/project/%s/application/list?appIds=%s&scope=%s", projectId, appIds, scope), nil)
	return result, err
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
func GetProjectMembers(projectId string, userIds string, projectRoleId string, limit int, skip int, pageSize int, pageToken string) (Response[[]ProjectMember], error) {
	_, resp, err := Get[[]ProjectMember](fmt.Sprintf("/v3/project/%s/member?userIds=%s&projectRoleId=%s&limit=%d&skip=%d&pageSize=%d&pageToken=%s", projectId, userIds, projectRoleId, limit, skip, pageSize, pageToken), nil)
	return resp, err
}

// 查询项目
//	GET https://open.teambition.com/api/v3/project/query
// 接口地址： https://open.teambition.com/docs/apis/6321c6d0912d20d3b5a49aa7
// @projectIds 项目ID列表,多个ID用逗号分隔
// @name 项目名称
// @sourceId 原始项目ID
// @pageToken 分页标识
// @pageSize 分页大小
func QueryProject(projectIds string, name string, sourceId string, pageToken string, pageSize int) (resp Response[[]Project], err error) {
	_, resp, err = Get[[]Project](fmt.Sprintf("/v3/project/query?projectIds=%s&name=%s&sourceId=%s&pageToken=%s&pageSize=%d", projectIds, name, sourceId, pageToken, pageSize), nil)
	return resp, err
}
