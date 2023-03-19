package teambitionapi_test

import (
	"testing"
)

func TestGetOrgInfo(t *testing.T) {

	resp, err := tbapi.GetOrgInfo(orgId)
	if err != nil {
		t.Errorf("%v", err)
	}
	// t.Log(resp)
	t.Log(resp.Result)
	t.Log(resp)
	t.Logf("%#v", resp)
}
