# teambition-sdk-go

![banner](./img/banner.png)

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/kexin8/teambition-sdk-go?filename=go.mod)
![teambition](https://img.shields.io/badge/teambition-API-brightgreen)


teambition open api for go

TODO: Fill out this long description.

## 目录

- [背景](#背景)
- [功能列表](./doc/API.md)
- [安装](#安装)
- [Usage](#usage)
- [API](#api)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)


## 背景
Go 实现的 [teambition](https://www.teambition.com/) Open Api 集合




## 安装

```
go get github.com/kexin8/teambition-sdk-go
```

## Usage

```go
package main

import (
    "log"
    teambitionapi "github.com/kexin8/teambition-sdk-go/teambition-api"
)

func main() {
    var (
        orgId     = "..."
        appId     = "..."
        appSecret = "..."

        tbapi = teambitionapi.NewClient(teambitionapi.NewOptions(orgId, appId, appSecret))
    )
    
    resp, err := tbapi.GetOrgInfo(orgId)
    if err != nil {
    	log.Errorf("%v", err)
    }
    log.Println(resp)
    log.Println(resp.Result)
}
```

### 配置

```go
tbapi = teambitionapi.NewClient(teambitionapi.NewOptions(orgId, appId, appSecret))
```

配置相关属性参考[config.go#Options](./teambition-api/config.go)

```go
//config.go

type Options struct {
	baseUrl   string //https://open.teambition.com/api
	orgId     string //企业ID
	appID     string
	appSecret string

	tokenExpies        int64              //token过期时间
	isCacheToken       bool               //是否缓存token
	tokenCacheExecutor TokenCacheExecutor //token缓存执行器
}
```

**配置属性说明：**

> orgId、appID、appSecret需要从teambition申请获取

tokenExpies: 请求`tb open api`时token的有效时间，单位ms

isCacheToken: 是否缓存token，默认`true`;如果为`false`,则每次请求都会重新生成token

tokenCacheExecutor: token缓存执行器，用于指定token缓存策略,默认为本地缓存;当`isCacheToken == true`时生效，可通过实现`TokenCacheExecutor`接口自定义缓存策略

#### 自定义token缓存策略

> 用户可根据业务需求自定义token缓存策略，例如redis缓存

`TokenCacheExecutor`接口

```go
type TokenCacheExecutor interface {
	SetToken(tenantId, token string, expireTime int64) //设置token
	GetToken(tenantId string) (token string, ok bool)  //获取token
}
```

1. 自定义缓存`CustomToeknCacheExecutor`

```go
type CustomToeknCacheExecutor struct {
}

func (d *CustomToeknCacheExecutor) SetToken(tenantId, token string, expireTime int64) {
	//...
}

func (d *CustomToeknCacheExecutor) GetToken(tenantId string) (token string, ok bool) {
	//...
}
```

2. 使用

```go
options := teambitionapi.NewOptions(orgId, appId, appSecret)
options.tokenCacheExecutor = &CustomToeknCacheExecutor{}

tbapi := teambitionapi.NewClient(options)
```



## API

## Maintainers

[@kexin8](https://github.com/kexin8)

## Contributing

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2023 kexin8
