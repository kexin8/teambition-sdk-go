package teambitionapi

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type Client struct {
	*Options

	request *resty.Client
}

func NewClient(options *Options) (client *Client) {
	c := resty.New()
	c.SetBaseURL(options.GetBaseUrl())
	c.SetHeader("X-Tenant-Id", options.GetOrgId())
	c.SetHeader("X-Tenant-Type", "organization")

	client = &Client{
		Options: options,
		request: c,
	}

	return
}

// GetToken 获取token
func GetToken(appid, appSecret string, expoesIn time.Duration) (token string, exp time.Time, err error) {
	t := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	exp = time.Now().Add(expoesIn)
	claims["exp"] = exp
	claims["iat"] = time.Now().Unix()
	claims["_appId"] = appid
	t.Claims = claims

	tokenStr, err := t.SignedString([]byte(appSecret))
	if err != nil {
		return "", time.Time{}, errors.WithStack(err)
	}

	return "Bearer " + tokenStr, exp, nil
}

// RefreshToken 刷新token
func (c *Client) RefreshToken() error {
	//是否需要从缓存中获取token
	if !c.isCacheToken {
		//不需要从缓存中获取token，直接刷新token
		token, _, err := GetToken(c.GetAppId(), c.GetAppSecret(), c.GetTokenExpies())
		if err != nil {
			return errors.WithStack(err)
		}
		c.request.SetHeader("Authorization", token)
		return nil
	}

	//判断token是否过期
	token, ok := c.tokenCacheExecutor.GetToken(c.GetOrgId())
	if ok {
		//token未过期，直接返回
		c.request.SetHeader("Authorization", token)
		return nil
	}
	//刷新token
	token, exp, err := GetToken(c.GetAppId(), c.GetAppSecret(), c.GetTokenExpies())
	if err != nil {
		return errors.WithStack(err)
	}
	c.request.SetHeader("Authorization", token)
	//将token存入缓存
	c.tokenCacheExecutor.SetToken(c.GetOrgId(), token, exp.Unix())
	return nil
}
