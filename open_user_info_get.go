package ks_shop_go_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (k *KsShopClient) OpenUserInfoGet(ctx context.Context, reqData *OpenUserInfoGetRequest) (*OpenUserInfoGetResponse, error) {

	paramBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign := ""
	sign, err = k.Sign(map[string]string{
		"access_token": k.AccToken,
		"method":       k.FormatApi(OpenUserInfoGetApi),
		"param":        string(paramBytes),
		"appkey":       k.AppId,
		"version":      k.Version,
		"signMethod":   k.SignMethod,
		"timestamp":    timestamp,
	})
	if err != nil {
		return nil, err
	}
	values.Set("access_token", k.AccToken)
	values.Set("method", k.FormatApi(OpenUserInfoGetApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

	endpoint := k.Env + OpenUserInfoGetApi + "?" + values.Encode()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("open_user_info_get status=%d body=%s", resp.StatusCode, string(body))
	}

	var result OpenUserInfoGetResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_user_info_get json_parse failed: %w", err)
	}

	//if result.Result != 1 {
	//	codeText := strconv.Itoa(result.Result)
	//	if result.Code != "" {
	//		codeText = result.Code
	//	}
	//	return &result, fmt.Errorf("open_user_info_get failed: code=%s msg=%s", codeText, result.ErrorMsg)
	//}

	return &result, nil
}

type OpenUserInfoGetRequest struct {
}

type OpenUserInfoGetResponse struct {
	Result    int                 `json:"result"`
	Msg       string              `json:"msg"`
	ErrorMsg  string              `json:"error_msg"`
	Code      string              `json:"code"`
	Data      OpenUserInfoGetData `json:"data"`
	RequestId string              `json:"requestId"`
	SubMsg    string              `json:"sub_msg"`
	SubCode   string              `json:"sub_code"`
}

type OpenUserInfoGetData struct {
	Head    string `json:"head"`
	OpenId  string `json:"openId"`
	Sex     string `json:"sex"`
	Name    string `json:"name"`
	BigHead string `json:"bigHead"`
}
