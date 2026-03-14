package ks_shop_go_sdk

type KsShopClient struct {
	AppId      string
	AppSecret  string
	SignSecret string
	Env        string
}

func NewKsShopClient(appId, appSecret, signSecret string) *KsShopClient {
	return &KsShopClient{
		AppId:      appId,
		AppSecret:  appSecret,
		SignSecret: signSecret,
		Env:        OnlineEnv,
	}
}

func (k *KsShopClient) SetEnv(env string) {
	k.Env = env
}

const (
	OnlineEnv        = "https://openapi.kwaixiaodian.com"
	OnlineRefreshEnv = "https://open.kuaishou.com"
	OnlineEnvBatest  = "https://open.kwaixiaodian.com"
)

const (
	// Oauth2AccessTokenApi
	// 用授权码code换取长时令牌refreshToken以及访问令牌accessToken
	Oauth2AccessTokenApi = "/oauth2/access_token"
	// Oauth2RefreshTokenApi 用刷新令牌refreshToken换取新的访问令牌accessToken
	Oauth2RefreshTokenApi = "/oauth2/refresh_token"

	// OpenOrderCursorListApi 获取订单列表
	OpenOrderCursorListApi = "/open/order/cursor/list"
)
