package teambitionapi_test

import (
	"testing"

	teambition "github.com/kexin8/teambition-sdk-go"
	example "github.com/kexin8/teambition-sdk-go/example"
	teambitionapi "github.com/kexin8/teambition-sdk-go/teambition-api"
)

var (
	orgId     = example.OrgId
	appId     = example.AppId
	appSecret = example.AppSecret
)

func init() {
	teambitionapi.NewClient(teambition.NewOptions(orgId, appId, appSecret))
}

func TestGetOrgAdmins(t *testing.T) {
	members, err := teambitionapi.GetOrgAdmins(orgId)
	if err != nil {
		t.Error(err)
	}
	t.Log(members)
}
