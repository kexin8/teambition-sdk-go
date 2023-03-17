package teambitionapi

import "fmt"

// OrgDomain 企业信息
type OrgDomain struct {
	OrgID       string `json:"orgId"`       // 企业 ID
	Name        string `json:"name"`        // 企业名称
	Pinyin      string `json:"pinyin"`      // 企业名称的拼音
	Py          string `json:"py"`          // 企业名称的拼音缩写
	Logo        string `json:"logo"`        // 企业 Logo 的 URL
	IsPublic    int64  `json:"isPublic"`    // 是否公开
	Description string `json:"description"` // 企业简介
	Created     string `json:"created"`     // 企业的创建时间
}

//获取企业信息
// GET https://open.teambition.com/api/org/info
func GetOrgInfo(orgId string) (orgDomain OrgDomain, err error) {
	orgDomain, err = Get[OrgDomain](fmt.Sprintf("/org/info?orgId=%s", orgId), nil)
	return
}
