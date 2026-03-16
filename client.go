package ks_shop_go_sdk

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

type KsShopClient struct {
	AppId      string
	AppSecret  string
	SignSecret string
	AccToken   string
	Env        string
	SignMethod string // 签名方法
	Version    string
}

func NewKsShopClient(appId, appSecret, signSecret, AccToken string) *KsShopClient {
	return &KsShopClient{
		AppId:      appId,
		AppSecret:  appSecret,
		SignSecret: signSecret,
		AccToken:   AccToken,
		Env:        OnlineEnv,
		SignMethod: "MD5",
		Version:    "1",
	}
}

func (k *KsShopClient) SetEnv(env string) {
	k.Env = env
}

func (k *KsShopClient) SetSignMethod(signMethod string) {
	k.SignMethod = signMethod
}

func (k *KsShopClient) SetVersion(version string) {
	k.Version = version
}

func (k *KsShopClient) SetAccToken(accToken string) {
	k.AccToken = accToken
}

const (
	OnlineEnv = "https://openapi.kwaixiaodian.com"
	// OnlineRefreshEnv = "https://open.kuaishou.com"
	OnlineEnvBatest = "https://open.kwaixiaodian.com"
)

// 签名
func (k *KsShopClient) Sign(params map[string]string) (string, error) {
	if len(params) == 0 {
		return "", fmt.Errorf("sign params is empty")
	}

	if k.SignSecret == "" {
		return "", fmt.Errorf("sign secret is empty")
	}

	signMethod := strings.ToUpper(k.SignMethod)
	if signMethod == "" {
		signMethod = "MD5"
	}

	if signMethod != "MD5" {
		return "", fmt.Errorf("unsupported sign method: %s", signMethod)
	}

	keys := make([]string, 0, len(params))
	for key, value := range params {
		if key == "" || key == "sign" || value == "" {
			continue
		}
		keys = append(keys, key)
	}

	if len(keys) == 0 {
		return "", fmt.Errorf("sign params is empty")
	}

	sort.Strings(keys)

	builder := strings.Builder{}
	for i, key := range keys {
		if i > 0 {
			builder.WriteString("&")
		}
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(params[key])
	}
	builder.WriteString(k.SignSecret)

	signBytes := md5.Sum([]byte(builder.String()))
	return hex.EncodeToString(signBytes[:]), nil
}

func (k *KsShopClient) FormatApi(api string) string {
	if len(api) == 0 {
		return ""
	}
	api = strings.ReplaceAll(api, "/", ".")
	api = api[1:]
	return api
}

const (
	// Oauth2AccessTokenApi
	// 用授权码code换取长时令牌refreshToken以及访问令牌accessToken
	Oauth2AccessTokenApi = "/oauth2/access_token"
	// Oauth2RefreshTokenApi 用刷新令牌refreshToken换取新的访问令牌accessToken
	Oauth2RefreshTokenApi = "/oauth2/refresh_token"

	// OpenOrderCursorListApi 获取订单列表
	OpenOrderCursorListApi = "/open/order/cursor/list"
	// OpenUserInfoGetApi 获取授权用户信息
	OpenUserInfoGetApi = "/open/user/info/get"
	// OpenItemListGetApi 获取商品列表
	OpenItemListGetApi = "/open/item/list/get"
	// OpenSellerOrderCpsDetailApi 获取分销订单详情
	OpenSellerOrderCpsDetailApi = "/open/seller/order/cps/detail"
	// OpenSellerOrderCpsListApi 获取分销订单列表
	OpenSellerOrderCpsListApi = "/open/seller/order/cps/list"
	// OpenServiceMarketBuyerServiceInfoApi 获取买家服务市场授权信息
	OpenServiceMarketBuyerServiceInfoApi = "/open/service/market/buyer/service/info"
)
