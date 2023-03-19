package teambitionapi_test

import (
	"encoding/json"
	"os"

	teambitionapi "github.com/kexin8/teambition-sdk-go/teambition-api"
)

var (
	orgId     = "..."
	appId     = "..."
	appSecret = "..."

	tbapi = teambitionapi.NewClient(teambitionapi.NewOptions(orgId, appId, appSecret))
)

func init() {

	const confile = "../config.json"

	//判断文件是否存在
	if _, err := os.Stat(confile); err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}

	bs, err := os.ReadFile(confile)
	if err != nil {
		panic(err)
	}

	var conf map[string]string
	if err = json.Unmarshal(bs, &conf); err != nil {
		panic(err)
	}

	orgId = conf["orgId"]
	appId = conf["appId"]
	appSecret = conf["appSecret"]

	tbapi = teambitionapi.NewClient(teambitionapi.NewOptions(orgId, appId, appSecret))
}
