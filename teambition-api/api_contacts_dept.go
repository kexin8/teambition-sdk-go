package teambitionapi

import (
	"fmt"

	"github.com/pkg/errors"
)

type Dept struct {
	AncestorIDS []string `json:"ancestorIds"` // 所有祖先部门 ID 列表，ancestorIds[0] 为直接父部门，ancestorIds[1] 为父部门的父部门，以此类推
	Created     string   `json:"created"`     // 创建时间
	CreatorID   string   `json:"creatorId"`   // 创建者 ID
	DeptID      string   `json:"deptId"`      // 部门 ID
	LeaderID    string   `json:"leaderId"`    // 部门负责人 ID
	Name        string   `json:"name"`        // 部门名称
	OrgID       string   `json:"orgId"`       // 企业 ID
	ParentID    string   `json:"parentId"`    // 父部门 ID （没有为空）
}

// 获取企业部门列表
//
//	GET https://open.teambition.com/api/org/departments
//	接口地址： https://open.teambition.com/docs/apis/6321c6ce912d20d3b5a48bbc
//	@param orgId 企业ID
//	@param pageToken 分页标识
//	@param pageSize 分页大小
func (c *Client) GetOrgDepartments(orgId string, pageToken string, pageSize int) (resp *Response[[]Dept], err error) {
	resp, err = Get[[]Dept](c, fmt.Sprintf("/org/departments?orgId=%s&pageToken=%s&pageSize=%d", orgId, pageToken, pageSize), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 获取部门成员列表
// GET https://open.teambition.com/api/departments/members
// 接口地址： https://open.teambition.com/docs/apis/6321c6ce912d20d3b5a48bfc
// @param orgId 企业ID
// @param deptId 部门ID
// @param omitSubDepartment 是否忽略子孙部门成员
// @param pageToken 分页标识
// @param pageSize 分页大小
func (c *Client) GetDeptMembers(orgId string, deptId string, omitSubDepartment bool, pageToken string, pageSize int) (resp *Response[[]Member], err error) {
	resp, err = Get[[]Member](c, fmt.Sprintf("/departments/members?orgId=%s&deptId=%s&omitSubDepartment=%t&pageToken=%s&pageSize=%d", orgId, deptId, omitSubDepartment, pageToken, pageSize), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 获取用户加入的企业部门列表
//
//	GET https://open.teambition.com/api/user/joinedDepartments
//	接口地址： https://open.teambition.com/docs/apis/6321c6ce912d20d3b5a48c30
//	@param orgId 企业ID
//	@param userId 用户ID
//	@param pageToken 分页标识
//	@param pageSize 分页大小
func (c *Client) GetUserJoinedDepartments(orgId string, userId string, pageToken string, pageSize int) (resp *Response[[]Dept], err error) {
	resp, err = Get[[]Dept](c, fmt.Sprintf("/user/joinedDepartments?orgId=%s&userId=%s&pageToken=%s&pageSize=%d", orgId, userId, pageToken, pageSize), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
