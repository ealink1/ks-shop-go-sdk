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

// OpenShopInfoGet 获取店铺信息
func (k *KsShopClient) OpenShopInfoGet(ctx context.Context, reqData *OpenShopInfoGetRequest) (*OpenShopInfoGetResponse, error) {
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
		"method":       k.FormatApi(OpenShopInfoGetApi),
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
	values.Set("method", k.FormatApi(OpenShopInfoGetApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

	// 发起请求
	endpoint := k.Env + OpenShopInfoGetApi + "?" + values.Encode()
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
		return nil, fmt.Errorf("open_shop_info_get status=%d body=%s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result OpenShopInfoGetResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_shop_info_get json_parse failed: %w", err)
	}

	// 业务校验
	//if result.Result != 1 {
	//	codeText := strconv.Itoa(result.Result)
	//	if result.Code != "" {
	//		codeText = result.Code
	//	}
	//	return &result, fmt.Errorf("open_shop_info_get failed: code=%s msg=%s", codeText, result.ErrorMsg)
	//}

	return &result, nil
}

// OpenShopInfoGetRequest 获取店铺信息请求参数
type OpenShopInfoGetRequest struct {
}

// OpenShopInfoGetResponse 获取店铺信息响应
type OpenShopInfoGetResponse struct {
	Result    int                 `json:"result"`    // 结果码
	Msg       string              `json:"msg"`       // 结果描述
	ErrorMsg  string              `json:"error_msg"` // 错误信息
	Code      string              `json:"code"`      // 业务码
	Data      OpenShopInfoGetData `json:"data"`      // 店铺信息
	RequestId string              `json:"requestId"` // 请求 ID
	SubMsg    string              `json:"sub_msg"`   // 子错误信息
	SubCode   string              `json:"sub_code"`  // 子错误码
}

// OpenShopInfoGetData 店铺信息
type OpenShopInfoGetData struct {
	ShopName      string                   `json:"shopName"`      // 店铺名称
	ShopType      int                      `json:"shopType"`      // 店铺类型
	ShopScoreInfo OpenShopInfoGetScoreInfo `json:"shopScoreInfo"` // 店铺评分
}

// OpenShopInfoGetScoreInfo 店铺评分信息
type OpenShopInfoGetScoreInfo struct {
	ContentQualifyScoreStr    string `json:"contentQualifyScoreStr"`    // 内容资质得分
	AfterSalesServiceScoreStr string `json:"afterSalesServiceScoreStr"` // 售后服务得分
	LogisticsServiceScoreStr  string `json:"logisticsServiceScoreStr"`  // 物流服务得分
	ShopExpScoreStr           string `json:"shopExpScoreStr"`           // 购物体验得分
	ProductQualityScoreStr    string `json:"productQualityScoreStr"`    // 商品质量得分
	CustomerServiceScoreStr   string `json:"customerServiceScoreStr"`   // 客服服务得分
}
