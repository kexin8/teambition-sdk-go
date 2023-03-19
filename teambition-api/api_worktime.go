package teambitionapi

import (
	"fmt"

	"github.com/pkg/errors"
)

type Worktime struct {
	CreatedAt   string `json:"createdAt"`   // 工时创建时间
	Date        string `json:"date"`        // 工时所属日期字符串
	ObjectID    string `json:"objectId"`    // 工时task ID
	ObjectType  string `json:"objectType"`  // 工时关联类型,默认task
	OrgID       string `json:"orgId"`       // 工时所属企业 ID
	SubmitterID string `json:"submitterId"` // 工时提交者 ID
	UpdatedAt   string `json:"updatedAt"`   // 工时更新时间
	UserID      string `json:"userId"`      // 工时执行者 ID
	Worktime    int64  `json:"worktime"`    // 工时时间(单位:毫秒)
	WorktimeID  string `json:"worktimeId"`  // 工时 ID
}

// 获取单个任务的实际工时列表
// GET /worktime/list/task/{taskId}
// @param taskId 任务ID
func (c *Client) GetWorktimeList(taskId string) (resp *Response[[]Worktime], err error) {
	resp, err = Get[[]Worktime](c, fmt.Sprintf("/worktime/list/task/%s", taskId), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

// 获取用户所有实际工时详情数据
// GET /worktime/query
// @param userId 用户ID
// @param startDate 查询起始时间，格式 “2022-08-01”
// @param endDate 查询结束时间，格式 “2022-08-01”，与起始时间间隔90天内
// @param pageToken	分页标
// @param pageSize	每页任务数量，默认为10
func (c *Client) GetWorktimeDetail(userId, startDate, endDate, pageToken string, pageSize int) (resp *Response[[]Worktime], err error) {
	resp, err = Get[[]Worktime](c, fmt.Sprintf("/worktime/query?userId=%s&startDate=%s&endDate=%s&pageToken=%s&pageSize=%d", userId, startDate, endDate, pageToken, pageSize), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
