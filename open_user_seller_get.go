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

// OpenUserSellerGet 获取商家信息
func (k *KsShopClient) OpenUserSellerGet(ctx context.Context, reqData *OpenUserSellerGetRequest) (*OpenUserSellerGetResponse, error) {
	// 序列化业务参数
	paramBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	// 生成时间戳
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	// 计算签名
	sign, err := k.Sign(map[string]string{
		"access_token": k.AccToken,
		"method":       k.FormatApi(OpenUserSellerGetApi),
		"param":        string(paramBytes),
		"appkey":       k.AppId,
		"version":      k.Version,
		"signMethod":   k.SignMethod,
		"timestamp":    timestamp,
	})
	if err != nil {
		return nil, err
	}

	// 组装请求参数
	values := url.Values{}
	values.Set("access_token", k.AccToken)
	values.Set("method", k.FormatApi(OpenUserSellerGetApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

	// 发起请求
	endpoint := k.Env + OpenUserSellerGetApi + "?" + values.Encode()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 校验 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("open_user_seller_get status=%d body=%s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result OpenUserSellerGetResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_user_seller_get json_parse failed: %w", err)
	}

	return &result, nil
}

// OpenUserSellerGetRequest 获取商家信息请求参数
type OpenUserSellerGetRequest struct {
}

// OpenUserSellerGetResponse 获取商家信息响应
type OpenUserSellerGetResponse struct {
	Result    int                    `json:"result"`    // 结果码
	Msg       string                 `json:"msg"`       // 结果描述
	ErrorMsg  string                 `json:"error_msg"` // 错误信息
	Code      string                 `json:"code"`      // 业务码
	Data      OpenUserSellerGetData  `json:"data"`      // 商家信息
	RequestId string                 `json:"requestId"` // 请求 ID
	SubMsg    string                 `json:"sub_msg"`   // 子错误信息
	SubCode   string                 `json:"sub_code"`  // 子错误码
}

// OpenUserSellerGetData 商家信息
type OpenUserSellerGetData struct {
	Name     string                  `json:"name"`     // 商家名称
	Sex      string                  `json:"sex"`      // 商家性别
	Head     string                  `json:"head"`     // 头像
	BigHead  string                  `json:"bigHead"`  // 高清头像
	SellerId int64                   `json:"sellerId"` // 商家id
	OpenId   string                  `json:"openId"`   // 商家在快手的唯一标识
	StaffInfo OpenUserSellerStaffInfo `json:"staffInfo"` // 子账号信息（子账号授权的token调用时才有该信息）
}

// OpenUserSellerStaffInfo 子账号信息
type OpenUserSellerStaffInfo struct {
	StaffId int64  `json:"staffId"` // 子账号id
	OpenId  string `json:"openId"`  // 子账号唯一id
	Name    string `json:"name"`    // 子账号昵称
}
