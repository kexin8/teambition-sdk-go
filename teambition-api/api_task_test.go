package teambitionapi_test

import (
	"testing"

	teambitionapi "github.com/kexin8/teambition-sdk-go/teambition-api"
)

func TestCreateTask(t *testing.T) {
	task := teambitionapi.CreateTaskRequest{
		ProjectID: "",
		Content: "",
		ExecutorID: "",
	}

	resp, err := tbapi.CreateTask(task, "...")
	if err != nil {
		t.Errorf("%v", err)
	}
	// t.Log(resp)
	t.Log(resp.Result)
	t.Log(resp)
	t.Logf("%#v", resp)
}
