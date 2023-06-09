package teambitionapi

import "time"

type Options struct {
	baseUrl   string //https://open.teambition.com/api
	orgId     string //企业ID
	appID     string
	appSecret string

	tokenExpies        int64              //token过期时间
	isCacheToken       bool               //是否缓存token
	tokenCacheExecutor TokenCacheExecutor //token缓存执行器
}

func NewOptions(orgId, appId, appSecret string) *Options {
	return &Options{
		baseUrl:            "https://open.teambition.com/api",
		orgId:              orgId,
		appID:              appId,
		appSecret:          appSecret,
		tokenExpies:        1 * 3600 * 1000,
		isCacheToken:       true,
		tokenCacheExecutor: &DefaultTokenCacheExecutor{},
	}
}

func (c *Options) SetBaseUrl(baseUrl string) {
	c.baseUrl = baseUrl
}

func (c *Options) SetTokenExpies(tokenExpies int64) {
	c.tokenExpies = tokenExpies
}

func (c *Options) SetIsCacheToken(isCacheToken bool) {
	c.isCacheToken = isCacheToken
}

func (c *Options) SetTokenCacheExecutor(tokenCacheExecutor TokenCacheExecutor) {
	c.tokenCacheExecutor = tokenCacheExecutor
}

func (c *Options) GetBaseUrl() string {
	return c.baseUrl
}

func (c *Options) GetOrgId() string {
	return c.orgId
}

func (c *Options) GetAppId() string {
	return c.appID
}

func (c *Options) GetAppSecret() string {
	return c.appSecret
}

func (c *Options) GetTokenExpies() time.Duration {
	return time.Duration(c.tokenExpies)
}
