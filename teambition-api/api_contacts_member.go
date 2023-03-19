package teambitionapi

import (
	"fmt"

	"github.com/pkg/errors"
)

type Member struct {
	UserID         string `json:"userId"`         // 用户 ID
	MemberID       string `json:"memberId"`       // 企业成员 ID
	Name           string `json:"name"`           // 成员名称
	Pinyin         string `json:"pinyin"`         // 成员名称的拼音
	Py             string `json:"py"`             // 成员名称的拼音简写
	Title          string `json:"title"`          // 职位
	Role           int64  `json:"role"`           // 成员角色，取值为：-1: 外部成员, 0: 成员, 1: 管理员, 2: 拥有者
	StaffType      string `json:"staffType"`      // 员工类型
	EmployeeNumber string `json:"employeeNumber"` // 员工工号
	Email          string `json:"email"`          // 电子邮箱
	AvatarURL      string `json:"avatarUrl"`      // 用户头像
	Birthday       string `json:"birthday"`       // 出生日期
	City           string `json:"city"`           // 工作地点（城市）
	Country        string `json:"country"`        // 工作地点（国家）
	EntryTime      string `json:"entryTime"`      // 入职时间
	IsDisabled     int64  `json:"isDisabled"`     // 成员账号是否被停用：0 启用，1 停用
	Phone          string `json:"phone"`          // 联系电话
	Province       string `json:"province"`       // 工作地点（省份）
}

// 获取企业的管理员（含拥有者）
// GET https://open.teambition.com/api/org/admins
// 接口地址： https://open.teambition.com/docs/apis/6321c6ce912d20d3b5a48974
// @param orgId 企业ID 必填
func (c *Client) GetOrgAdmins(orgId string) (resp *Response[[]Member], err error) {
	resp, err = Get[[]Member](c, fmt.Sprintf("/org/admins?orgId=%s", orgId), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 获取企业拥有者
//  GET https://open.teambition.com/api/org/owners
//  接口地址：https://open.teambition.com/docs/apis/6321c6ce912d20d3b5a489be
//  @param orgId 企业ID 必填
func (c *Client) GetOrgOwners(orgId string) (resp *Response[[]Member], err error) {
	resp, err = Get[[]Member](c, fmt.Sprintf("/org/owners?orgId=%s", orgId), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 获取企业成员列表
//   GET https://open.teambition.com/api/org/member/list
//   接口地址：https://open.teambition.com/docs/apis/6321c6ce912d20d3b5a489ee
//   @param orgId 企业ID 必填
//   @param pageToken 分页标记，第一次请求不填或者空字符串，表示从头开始遍历；分页查询结果还有更多项时返回 nextPageToken，下次遍历可采用该 nextPageToken 入参 pageToken 获取查询结果，示例值："cfcb90voe9jct71bqkfg"
//   @param pageSize 分页大小,默认值 10
//   @param filter 该参数用于指定查询范围，initiate：新加入的成员,disabled：已禁用的成员,enabled：未禁用的成员,external：外部成员
func (c *Client) GetOrgMembers(orgId string, pageToken string, pageSize int, filter string) (resp *Response[[]Member], err error) {
	resp, err = Get[[]Member](c, fmt.Sprintf("/org/member/list?orgId=%s&pageToken=%s&pageSize=%d&filter=%s", orgId, pageToken, pageSize, filter), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
