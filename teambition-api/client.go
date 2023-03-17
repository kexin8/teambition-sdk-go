package teambitionapi

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-resty/resty/v2"
	teambition "github.com/kexin8/teambition-sdk-go"
	"github.com/pkg/errors"
)

var ApiClient *Client

type Client struct {
	options *teambition.Options

	token       string    //token
	tokenExpies time.Time //token过期时间

	*resty.Client
}

func NewClient(options *teambition.Options) {
	c := resty.New()
	c.SetBaseURL(options.GetBaseUrl())
	c.SetHeader("X-Tenant-Id", options.GetOrgId())
	c.SetHeader("X-Tenant-Type", "organization")

	token, exp, err := GetToken(options.GetAppId(), options.GetAppSecret(), options.GetTokenExpies())
	if err != nil {
		panic(err)
	}

	c.SetHeader("Authorization", token)

	ApiClient = &Client{
		options:     options,
		token:       token,
		tokenExpies: exp,
		Client:      c,
	}
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

// ValidToken 校验token是否过期
func (c *Client) ValidToken() bool {
	if c.token == "" {
		return false
	}

	return time.Now().Before(c.tokenExpies)
}

// RefreshToken 刷新token
func (c *Client) RefreshToken() error {
	token, exp, err := GetToken(c.options.GetAppId(), c.options.GetAppSecret(), c.options.GetTokenExpies())
	if err != nil {
		return errors.WithStack(err)
	}

	c.token = token
	c.tokenExpies = exp
	c.SetHeader("Authorization", token)

	return nil
}
