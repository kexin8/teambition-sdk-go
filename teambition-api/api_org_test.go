package teambitionapi_test

import (
	"testing"

	teambition "github.com/kexin8/teambition-sdk-go"
	example "github.com/kexin8/teambition-sdk-go/example"
	teambitionapi "github.com/kexin8/teambition-sdk-go/teambition-api"
)

func TestGetOrgInfo(t *testing.T) {

	var (
		orgId     = example.OrgId
		appId     = example.AppId
		appSecret = example.AppSecret
	)

	teambitionapi.NewClient(teambition.NewOptions(orgId, appId, appSecret))

	orgInfo, err := teambitionapi.GetOrgInfo(orgId)
	if err != nil {
		t.Error(err)
	}
	t.Log(orgInfo)
}
