package teambitionapi

import (
	"fmt"

	"github.com/pkg/errors"
)

// Org 企业信息
type Org struct {
	OrgID       string `json:"orgId"`       // 企业 ID
	Name        string `json:"name"`        // 企业名称
	Pinyin      string `json:"pinyin"`      // 企业名称的拼音
	Py          string `json:"py"`          // 企业名称的拼音缩写
	Logo        string `json:"logo"`        // 企业 Logo 的 URL
	IsPublic    int64  `json:"isPublic"`    // 是否公开
	Description string `json:"description"` // 企业简介
	Created     string `json:"created"`     // 企业的创建时间
}

type CustomFieldCategory struct {
	Id             string `json:"id"`             //任务类型ID
	Name           string `json:"name"`           //名称
	OrganizationId string `json:"organizationId"` //企业ID
	CreatorId      string `json:"creatorId"`      //创建者ID
	Created        string `json:"created"`        //创建时间
	Updated        string `json:"updated"`        //更新时间
}

type Category struct {
	CategoryId string `json:"categoryId"` //任务类型ID
	Count      int64  `json:"count"`      //分类下自定义字段总数
}

type Scenariofieldconfig struct {
	BoundToObjectID   string          `json:"boundToObjectId"`   // 所属对象ID
	BoundToObjectType string          `json:"boundToObjectType"` // 所属对象类型
	Created           string          `json:"created"`           // 创建时间
	CreatorID         string          `json:"creatorId"`         // 创建人ID
	ID                string          `json:"id"`                // 任务类型ID
	IsArchived        bool            `json:"isArchived"`        // 是否归档
	Name              string          `json:"name"`              // 名称
	Scenariofields    []Scenariofield `json:"scenariofields"`
	Source            string          `json:"source"`     // 是否来自安装应用,可能有 application.risk | application.story | application.bug
	TaskflowID        string          `json:"taskflowId"` // 工作流ID
	Type              string          `json:"type"`       // 类型
}

type Scenariofield struct {
	CustomfieldID string `json:"customfieldId"` // 自定义字段ID
	FieldType     string `json:"fieldType"`     // 场景类型
	Required      bool   `json:"required"`      // 是否必填
}

// 获取企业信息
// GET https://open.teambition.com/api/org/info
// 接口地址： https://open.teambition.com/docs/apis/6321c6ce912d20d3b5a488f4
// @param orgId 企业ID
func (c *Client) GetOrgInfo(orgId string) (resp *Response[Org], err error) {
	resp, err = Get[Org](c, fmt.Sprintf("/org/info?orgId=%s", orgId), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 搜索企业自定义字段分类
// GET https://open.teambition.com/api/v3/customfield-category/search
// 接口地址： https://open.teambition.com/docs/apis/63ee3ea2912d20d3b543ee57
// @param q 搜索关键字
// @param categoryIds 自定义字段分类ID集合，逗号组合
// @param pageToken 分页标识
// @param pageSize 分页大小(默认50)
func (c *Client) V3CustomFieldCategorySearch(q, categoryIds, pageToken string, pageSize int) (resp *Response[CustomFieldCategory], err error) {
	resp, err = Get[CustomFieldCategory](c, fmt.Sprintf("/v3/customfield-category/search?q=%s&categoryIds=%s&pageToken=%s&pageSize=%d", q, categoryIds, pageToken, pageSize), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 根据自定义字段分类统计自定义字段数
// GET https://open.teambition.com/api/v3/customfield/count-by-category
// 接口地址： https://open.teambition.com/docs/apis/63ee3ea2912d20d3b543ee9c
// @param categoryIds 自定义字段分类ID集合，逗号组合，未分类可使用 uncategory
func (c *Client) V3CustomFieldCountByCategory(categoryIds string) (resp *Response[Category], err error) {
	resp, err = Get[Category](c, fmt.Sprintf("/v3/customfield/count-by-category?categoryIds=%s", categoryIds), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 搜索企业任务类型
// GET https://open.teambition.com/api/v3/scenariofieldconfig/search
// 接口地址： https://open.teambition.com/docs/apis/63ee3ea2912d20d3b543f1ae
// @param q 模糊查询任务类型名称
// @param sfcIds 任务类型ID集合，逗号组合
// @param pageToken 分页标识
// @param pageSize 分页大小(默认50)
func (c *Client) V3ScenarioFieldConfigSearch(q, sfcIds string, pageToken string, pageSize int) (resp *Response[Scenariofieldconfig], err error) {
	resp, err = Get[Scenariofieldconfig](c, fmt.Sprintf("/v3/scenariofieldconfig/search?q=%s&sfcIds=%s&pageToken=%s&pageSize=%d", q, sfcIds, pageToken, pageSize), nil)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return
}
