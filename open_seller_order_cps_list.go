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

func (k *KsShopClient) OpenSellerOrderCpsList(ctx context.Context, reqData *OpenSellerOrderCpsListRequest) (*OpenSellerOrderCpsListResponse, error) {
	paramBytes, err := json.Marshal(reqData.Param)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign, err := k.Sign(map[string]string{
		"access_token": k.AccToken,
		"method":       k.FormatApi(OpenSellerOrderCpsListApi),
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
	values.Set("method", k.FormatApi(OpenSellerOrderCpsListApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

	endpoint := k.Env + OpenSellerOrderCpsListApi + "?" + values.Encode()
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
		return nil, fmt.Errorf("open_seller_order_cps_list status=%d body=%s", resp.StatusCode, string(body))
	}

	var result OpenSellerOrderCpsListResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_seller_order_cps_list json_parse failed: %w", err)
	}

	if result.Result != 1 {
		codeText := strconv.Itoa(result.Result)
		if result.Code != "" {
			codeText = result.Code
		}
		return &result, fmt.Errorf("open_seller_order_cps_list failed: code=%s msg=%s", codeText, result.ErrorMsg)
	}

	return &result, nil
}

type OpenSellerOrderCpsListRequest struct {
	AccessToken string
	Sign        string
	Timestamp   int64
	AppKey      string
	Version     string
	SignMethod  string
	Method      string
	Param       OpenSellerOrderCpsListParam
}

type OpenSellerOrderCpsListParam struct {
	CurrentPage int    `json:"currentPage"`
	PageSize    int    `json:"pageSize"`
	Sort        int    `json:"sort"`
	QueryType   int    `json:"queryType"`
	Type        int    `json:"type"`
	Pcursor     string `json:"pcursor"`
	BeginTime   int64  `json:"beginTime"`
	EndTime     int64  `json:"endTime"`
}

type OpenSellerOrderCpsListResponse struct {
	Result    int                        `json:"result"`
	Msg       string                     `json:"msg"`
	ErrorMsg  string                     `json:"error_msg"`
	Code      string                     `json:"code"`
	Data      OpenSellerOrderCpsListData `json:"data"`
	RequestId string                     `json:"requestId"`
	SubMsg    string                     `json:"sub_msg"`
	SubCode   string                     `json:"sub_code"`
}

type OpenSellerOrderCpsListData struct {
	CpsOrderList []OpenSellerOrderCpsListItem `json:"cpsOrderList"`
	Pcursor      string                       `json:"pcursor"`
}

type OpenSellerOrderCpsListItem struct {
	CommissionRate            int    `json:"commissionRate"`
	PayTime                   int64  `json:"payTime"`
	InvestmentPromotionAmount int    `json:"investmentPromotionAmount"`
	PlatformDpRate            int    `json:"platformDpRate"`
	Oid                       int64  `json:"oid"`
	BuyerId                   int64  `json:"buyerId"`
	ExpressFee                int    `json:"expressFee"`
	OrderChannel              string `json:"orderChannel"`
	SettlementBizType         int    `json:"settlementBizType"`
	ActivityId                int64  `json:"activityId"`
	DistributorName           string `json:"distributorName"`
	SellerId                  int64  `json:"sellerId"`
	ServiceIncome             int    `json:"serviceIncome"`
	OrderTradeAmount          int    `json:"orderTradeAmount"`
	OrderUpdateTime           int64  `json:"orderUpdateTime"`
	SettlementSuccessTime     int64  `json:"settlementSuccessTime"`
	ActivityUserId            int64  `json:"activityUserId"`
	InvestmentPromotionRate   int    `json:"investmentPromotionRate"`
	BuyerOpenId               string `json:"buyerOpenId"`
	DistributorId             int64  `json:"distributorId"`
	RefundTime                int64  `json:"refundTime"`
	PromoterNickname          string `json:"promoterNickname"`
	OrderCreateTime           int64  `json:"orderCreateTime"`
	ActivityUserNickname      string `json:"activityUserNickname"`
	AttributionCpsType        int    `json:"attributionCpsType"`
	UpdateTime                int64  `json:"updateTime"`
	ShareRate                 int    `json:"shareRate"`
	PromoterId                int64  `json:"promoterId"`
	CpsOrderStatus            int    `json:"cpsOrderStatus"`
	BaseAmount                int    `json:"baseAmount"`
	ItemId                    int64  `json:"itemId"`
	CreateTime                int64  `json:"createTime"`
	SettlementTime            int64  `json:"settlementTime"`
	TotalFee                  int    `json:"totalFee"`
	PlatformRate              int    `json:"platformRate"`
	SettlementBizTypeDesc     string `json:"settlementBizTypeDesc"`
	SecondActivityUserId      int64  `json:"secondActivityUserId"`
	ExcitationInCome          int    `json:"excitationInCome"`
	KwaimoneyUserId           int64  `json:"kwaimoneyUserId"`
	EstimatedIncome           int    `json:"estimatedIncome"`
	Status                    int    `json:"status"`
}
