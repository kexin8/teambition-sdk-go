package teambitionapi

import (
	"fmt"

	"github.com/pkg/errors"
)

// 任务详情
type Task struct {
	AccomplishTime string        `json:"accomplishTime"` // 任务完成时间(UTC)
	AncestorIDS    []string      `json:"ancestorIds"`    // 祖先任务ID列表
	Content        string        `json:"content"`        // 任务标题
	Created        string        `json:"created"`        // 创建时间(UTC)
	CreatorID      string        `json:"creatorId"`      // 创建人ID
	Customfields   []Customfield `json:"customfields"`   // 自定义字段值集合
	DueDate        string        `json:"dueDate"`        // 任务截止时间(UTC)
	ExecutorID     string        `json:"executorId"`     // 执行人ID
	ID             string        `json:"id"`             // 任务ID
	InvolveMembers []string      `json:"involveMembers"` // 参与者ID集合
	IsArchived     bool          `json:"isArchived"`     // 是否任务放入回收站
	IsDone         bool          `json:"isDone"`         // 是否任务已完成
	Note           string        `json:"note"`           // 任务备注
	ParentTaskID   string        `json:"parentTaskId"`   // 父任务ID
	Priority       int64         `json:"priority"`       // 任务优先级
	ProjectID      string        `json:"projectId"`      // 项目ID
	Recurrence     []string      `json:"recurrence"`     // 重复规则列表
	SfcID          string        `json:"sfcId"`          // 任务类型ID
	SprintID       string        `json:"sprintId"`       // 迭代ID
	StageID        string        `json:"stageId"`        // 任务列ID
	StartDate      string        `json:"startDate"`      // 任务开始时间(UTC)
	StoryPoint     string        `json:"storyPoint"`     // StoryPoint
	TagIDS         []string      `json:"tagIds"`         // 标签ID集合
	TasklistID     string        `json:"tasklistId"`     // 任务分组ID
	TfsID          string        `json:"tfsId"`          // 任务状态ID
	UniqueID       string        `json:"uniqueId"`       // 任务数字ID
	Updated        string        `json:"updated"`        // 更新时间(UTC)
	Visible        string        `json:"visible"`        // 任务隐私性，'involves'表达仅参与者可见; 'members'表达项目成员可见
}

// 查询项目任务
// GET /v3/project/{projectId}/task/query
// @param projectId 项目ID
// @param q	查询语句 参考[TQL查询文档](doc.fullPath=/tql-doc)
// @param pageToken	分页标
// @param pageSize	每页任务数量，默认为10
func (c *Client) GetProjectTasks(projectId, q, pageToken string, pageSize int) (resp *Response[[]Task], err error) {
	resp, err = Get[[]Task](c, fmt.Sprintf("/v3/project/%s/task/query?q=%s&pageToken=%s&pageSize=%d", projectId, q, pageToken, pageSize), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

//		查询任务详情
//	  GET /v3/task/query
//	  @param taskId 任务ID集合,使用逗号分隔,和parentTaskId冲突(选其一)
//	  @param shortIds 任务短ID集合,使用逗号分隔
//	  @param parentTaskId 父任务ID,和taskIds冲突(选其一)
func (c *Client) GetTaskDetail(taskId, shortIds, parentTaskId string) (resp *Response[[]Task], err error) {
	resp, err = Get[[]Task](c, fmt.Sprintf("/v3/task/query?taskId=%s&shortIds=%s&parentTaskId=%s", taskId, shortIds, parentTaskId), nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}
