package ks_shop_go_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (k *KsShopClient) Oauth2AccessToken(ctx context.Context, code string) (*Oauth2AccessTokenResponse, error) {
	values := url.Values{}
	values.Set("app_id", k.AppId)
	values.Set("grant_type", "code")
	values.Set("code", code)
	values.Set("app_secret", k.AppSecret)

	endpoint := k.Env + Oauth2AccessTokenApi + "?" + values.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("oauth2_access_token status=%d body=%s", resp.StatusCode, string(body))
	}

	var result Oauth2AccessTokenResponse
	if err = json.Unmarshal(body, &result); err != nil {
		err = fmt.Errorf("oauth2_access_token json_parse failed: %w", err)
		return nil, err
	}

	//if result.Result != 1 {
	//	codeText := strconv.Itoa(result.Result)
	//	if result.Error != "" {
	//		codeText = result.Error
	//	}
	//	return &result, fmt.Errorf("oauth2_access_token failed: code=%s msg=%s", codeText, result.ErrorMsg)
	//}

	return &result, nil
}

type Oauth2AccessTokenResponse struct {
	Result                int      `json:"result"`
	AccessToken           string   `json:"access_token"`
	OpenId                string   `json:"open_id"`
	ExpiresIn             int      `json:"expires_in"`
	TokenType             string   `json:"token_type"`
	RefreshToken          string   `json:"refresh_token"`
	RefreshTokenExpiresIn int      `json:"refresh_token_expires_in"`
	Scopes                []string `json:"scopes"`
	Error                 string   `json:"error"`
	ErrorMsg              string   `json:"error_msg"`
}
