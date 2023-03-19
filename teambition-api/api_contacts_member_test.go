package teambitionapi_test

import (
	"testing"
)

func TestGetOrgAdmins(t *testing.T) {
	members, err := tbapi.GetOrgAdmins(orgId)
	if err != nil {
		t.Error(err)
	}
	t.Log(members)
}
