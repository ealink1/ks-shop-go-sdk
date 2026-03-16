# ks-shop-go-sdk

快手小店 Go 语言 SDK。

## 功能

提供快手小店 OpenAPI 的 Go 语言封装，方便开发者快速接入。

## 快速开始

### 安装

```bash
go get github.com/ealink1/ks-shop-go-sdk
go get -u github.com/ealink1/ks-shop-go-sdk@latest
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

## 接口列表

| 方法名 | 接口路径 | 说明 |
| --- | --- | --- |
| `Oauth2AccessToken` | `/oauth2/access_token` | 用授权码 code 换取 access_token 与 refresh_token |
| `Oauth2RefreshToken` | `/oauth2/refresh_token` | 用 refresh_token 换取新的 access_token |
| `OpenUserInfoGet` | `/open/user/info/get` | 获取授权账号信息 |
| `OpenItemListGet` | `/open/item/list/get` | 获取商品列表 |
| `OpenOrderCursorList` | `/open/order/cursor/list` | 游标分页获取订单列表 |
| `OpenSellerOrderCpsDetail` | `/open/seller/order/cps/detail` | 获取分销订单详情 |
| `OpenSellerOrderCpsList` | `/open/seller/order/cps/list` | 获取分销订单列表 |
| `OpenServiceMarketBuyerServiceInfo` | `/open/service/market/buyer/service/info` | 获取买家服务市场授权信息 |

## License

MIT
