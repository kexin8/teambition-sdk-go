package teambitionapi_test

import "testing"

func TestCreateTask(t *testing.T) {

	resp, err := tbapi.CreateTask("项目ID", "测试任务", "创建人ID")
	if err != nil {
		t.Errorf("%v", err)
	}
	// t.Log(resp)
	t.Log(resp.Result)
	t.Log(resp)
	t.Logf("%#v", resp)
}
