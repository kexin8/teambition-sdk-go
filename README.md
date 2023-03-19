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
    log.Info(resp)
    log.Info(resp.Result)
}
```

## API

## Maintainers

[@kexin8](https://github.com/kexin8)

## Contributing

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2023 kexin8
