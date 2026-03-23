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

func (k *KsShopClient) OpenSellerOrderCpsDetail(ctx context.Context, reqData *OpenSellerOrderCpsDetailRequest) (*OpenSellerOrderCpsDetailResponse, error) {
	paramBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign, err := k.Sign(map[string]string{
		"access_token": k.AccToken,
		"method":       k.FormatApi(OpenSellerOrderCpsDetailApi),
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
	values.Set("method", k.FormatApi(OpenSellerOrderCpsDetailApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

	endpoint := k.Env + OpenSellerOrderCpsDetailApi + "?" + values.Encode()
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
		return nil, fmt.Errorf("open_seller_order_cps_detail status=%d body=%s", resp.StatusCode, string(body))
	}

	var result OpenSellerOrderCpsDetailResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_seller_order_cps_detail json_parse failed: %w", err)
	}

	//if result.Result != 1 {
	//	codeText := strconv.Itoa(result.Result)
	//	if result.Code != "" {
	//		codeText = result.Code
	//	}
	//	return &result, fmt.Errorf("open_seller_order_cps_detail failed: code=%s msg=%s", codeText, result.ErrorMsg)
	//}

	return &result, nil
}

type OpenSellerOrderCpsDetailRequest struct {
	DistributorId int64 `json:"distributorId"`
	OrderId       int64 `json:"orderId"`
}

type OpenSellerOrderCpsDetailResponse struct {
	Result    int                          `json:"result"`
	Msg       string                       `json:"msg"`
	ErrorMsg  string                       `json:"error_msg"`
	Code      string                       `json:"code"`
	Data      OpenSellerOrderCpsDetailData `json:"data"`
	RequestId string                       `json:"requestId"`
	SubMsg    string                       `json:"sub_msg"`
	SubCode   string                       `json:"sub_code"`
}

type OpenSellerOrderCpsDetailData struct {
	CommissionRate            int    `json:"commissionRate"`
	PayTime                   int64  `json:"payTime"`
	InvestmentPromotionAmount int    `json:"investmentPromotionAmount"`
	PlatformDpRate            int    `json:"platformDpRate"`
	Oid                       int64  `json:"oid"`
	ExpressFee                int    `json:"expressFee"`
	OrderChannel              string `json:"orderChannel"`
	SettlementBizType         int    `json:"settlementBizType"`
	ActivityId                int64  `json:"activityId"`
	DistributorName           string `json:"distributorName"`
	SellerId                  int64  `json:"sellerId"`
	SettlementSuccessTime     int64  `json:"settlementSuccessTime"`
	ActivityUserId            int64  `json:"activityUserId"`
	InvestmentPromotionRate   int    `json:"investmentPromotionRate"`
	BuyerOpenId               string `json:"buyerOpenId"`
	DistributorId             int64  `json:"distributorId"`
	RefundTime                int64  `json:"refundTime"`
	ActivityUserNickname      string `json:"activityUserNickname"`
	UpdateTime                int64  `json:"updateTime"`
	ItemId                    int64  `json:"itemId"`
	CreateTime                int64  `json:"createTime"`
	SettlementTime            int64  `json:"settlementTime"`
	TotalFee                  int    `json:"totalFee"`
	EstimatedIncome           int    `json:"estimatedIncome"`
	Status                    int    `json:"status"`
}
