# ks-shop-go-sdk

快手小店 Go 语言 SDK。

## 功能

提供快手小店 OpenAPI 的 Go 语言封装，方便开发者快速接入。

## 快速开始

### 安装

```bash
go get github.com/ealink1/ks-shop-go-sdk
```
> 注意：请根据实际仓库地址调整上述 import 路径。

### 使用示例

```go
package main

import (
	"context"
	"fmt"
	"ks_shop_go_sdk"
)

func main() {
	// 初始化客户端
	appId := "your_app_id"
	appSecret := "your_app_secret"
	client := ks_shop_go_sdk.NewKsShopClient(appId, appSecret)

	// 调用接口，例如获取 Access Token
	ctx := context.Background()
	code := "your_auth_code"
	
	// 注意：当前实现可能还在开发中，具体返回值需参考源码
	client.Oauth2AccessToken(ctx, code)
}
```

## 目录结构

- `client.go`: 客户端结构定义及初始化
- `shop_api.go`: API 接口实现

## License

MIT
