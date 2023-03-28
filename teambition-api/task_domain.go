package teambitionapi

type CreateTaskRequest struct {
	ProjectID             string        `json:"projectId"`             // 项目id, 必填
	Content               string        `json:"content"`               // 任务标题, 必填
	ExecutorID            string        `json:"executorId"`            // 执行者id, 必填
	Customfields          []Customfield `json:"customfields"`          // 自定义字段值列表
	DueDate               string        `json:"dueDate"`               // 截止日期
	InvolveMembers        []string      `json:"involveMembers"`        // 参与者用户ID列表
	Note                  string        `json:"note"`                  // 任务备注
	ParentTaskID          string        `json:"parentTaskId"`          // 父任务id
	Priority              int64         `json:"priority"`              // 执行优先级
	ScenariofieldconfigID string        `json:"scenariofieldconfigId"` // 任务类型id
	SprintID              string        `json:"sprintId"`              // 迭代id
	StageID               string        `json:"stageId"`               // 任务列id
	StartDate             string        `json:"startDate"`             // 开始日期
	StoryPoint            string        `json:"storyPoint"`
	TagIDS                []string      `json:"tagIds"`           // 标签id列表
	TaskflowstatusID      string        `json:"taskflowstatusId"` // 任务状态id
	TasklistID            string        `json:"tasklistId"`       // 任务组id
	Visible               string        `json:"visible"`          // 任务的可见性规则 involves | members
}
