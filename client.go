package ks_shop_go_sdk

type KsShopClient struct {
	AppId     string
	AppSecret string
	Env       string
}

func NewKsShopClient(appId, appSecret string) *KsShopClient {
	return &KsShopClient{
		AppId:     appId,
		AppSecret: appSecret,
		Env:       OnlineEnv,
	}
}

func (k *KsShopClient) SetEnv(env string) {
	k.Env = env
}

const (
	OnlineEnv       = "https://openapi.kwaixiaodian.com"
	OnlineEnvBatest = "https://open.kwaixiaodian.com"
)

const (
	// Oauth2AccessTokenApi
	// 用授权码code换取长时令牌refreshToken以及访问令牌accessToken
	Oauth2AccessTokenApi = "/oauth2/access_token"
)
