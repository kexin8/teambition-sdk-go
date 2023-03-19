package teambitionapi_test

import (
	"testing"
)

func TestGetOrgInfo(t *testing.T) {

	orgInfo, err := tbapi.GetOrgInfo(orgId)
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Log(orgInfo)
}
