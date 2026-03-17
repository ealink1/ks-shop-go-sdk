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

// OpenDistributionSellerActivityOpenInfo 获取团长活动信息
func (k *KsShopClient) OpenDistributionSellerActivityOpenInfo(ctx context.Context, reqData *OpenDistributionSellerActivityOpenInfoRequest) (*OpenDistributionSellerActivityOpenInfoResponse, error) {
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
		"method":       k.FormatApi(OpenDistributionSellerActivityOpenInfoApi),
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
	values.Set("method", k.FormatApi(OpenDistributionSellerActivityOpenInfoApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

	// 发起请求
	endpoint := k.Env + OpenDistributionSellerActivityOpenInfoApi + "?" + values.Encode()
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
		return nil, fmt.Errorf("open_distribution_seller_activity_open_info status=%d body=%s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result OpenDistributionSellerActivityOpenInfoResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_distribution_seller_activity_open_info json_parse failed: %w", err)
	}

	// 业务校验
	if result.Result != 1 {
		codeText := strconv.Itoa(result.Result)
		if result.Code != "" {
			codeText = result.Code
		}
		return &result, fmt.Errorf("open_distribution_seller_activity_open_info failed: code=%s msg=%s", codeText, result.ErrorMsg)
	}

	return &result, nil
}

// OpenDistributionSellerActivityOpenInfoRequest 获取团长活动信息请求参数
type OpenDistributionSellerActivityOpenInfoRequest struct {
	ActivityId int64 `json:"activityId"` // 活动 ID
}

// OpenDistributionSellerActivityOpenInfoResponse 获取团长活动信息响应
type OpenDistributionSellerActivityOpenInfoResponse struct {
	Result    int                                        `json:"result"`    // 结果码
	Msg       string                                     `json:"msg"`       // 结果描述
	ErrorMsg  string                                     `json:"error_msg"` // 错误信息
	Code      string                                     `json:"code"`      // 业务码
	Data      OpenDistributionSellerActivityOpenInfoData `json:"data"`      // 活动信息
	RequestId string                                     `json:"requestId"` // 请求 ID
	SubMsg    string                                     `json:"sub_msg"`   // 子错误信息
	SubCode   string                                     `json:"sub_code"`  // 子错误码
}

// OpenDistributionSellerActivityOpenInfoData 团长活动信息
type OpenDistributionSellerActivityOpenInfoData struct {
	ActivityUserId                         int64                                            `json:"activityUserId"`                         // 活动创建人 ID
	UserOpenInfo                           []OpenDistributionSellerActivityUserOpenInfo     `json:"userOpenInfo"`                           // 用户开放信息列表
	ActivityBeginTime                      int64                                            `json:"activityBeginTime"`                      // 活动开始时间（毫秒）
	UpdateTime                             int64                                            `json:"updateTime"`                             // 更新时间（毫秒）
	ActivityId                             int64                                            `json:"activityId"`                             // 活动 ID
	ActivityEndTime                        int64                                            `json:"activityEndTime"`                        // 活动结束时间（毫秒），0 表示无限制
	ActivityUserHistoryDataInfoViewOpenDto OpenDistributionSellerActivityHistoryDataInfo    `json:"activityUserHistoryDataInfoViewOpenDto"` // 活动历史数据信息
	ActivityRuleSetViewOpenDto             OpenDistributionSellerActivityRuleSetViewOpenDto `json:"activityRuleSetViewOpenDto"`             // 活动规则设置
	ActivityTitle                          string                                           `json:"activityTitle"`                          // 活动标题
	CreateTime                             int64                                            `json:"createTime"`                             // 创建时间（毫秒）
	ActivityStatus                         int                                              `json:"activityStatus"`                         // 活动状态
	ActivityUserNickName                   string                                           `json:"activityUserNickName"`                   // 活动创建人昵称
	ActivityType                           int                                              `json:"activityType"`                           // 活动类型
}

// OpenDistributionSellerActivityUserOpenInfo 用户开放信息
type OpenDistributionSellerActivityUserOpenInfo struct {
	UserId int64  `json:"userId"` // 用户 ID
	Name   string `json:"name"`   // 用户昵称
}

// OpenDistributionSellerActivityHistoryDataInfo 活动历史数据信息
type OpenDistributionSellerActivityHistoryDataInfo struct {
	ActivityCount              int64 `json:"activityCount"`              // 活动数量
	DistributeTradeSellerCount int64 `json:"distributeTradeSellerCount"` // 分销商家数量
	DistributeTradeItemCount   int64 `json:"distributeTradeItemCount"`   // 分销商品数量
	DistributeItemOrderCount   int64 `json:"distributeItemOrderCount"`   // 分销订单数量
	CustomerPrice              int64 `json:"customerPrice"`              // 客户价格（分）
}

// OpenDistributionSellerActivityRuleSetViewOpenDto 活动规则设置
type OpenDistributionSellerActivityRuleSetViewOpenDto struct {
	PromotionActivityMarketingRule OpenDistributionSellerMarketingRule `json:"promotionActivityMarketingRule"` // 营销活动规则
}

// OpenDistributionSellerMarketingRule 营销活动规则
type OpenDistributionSellerMarketingRule struct {
	CategoryCommissionRate         []OpenDistributionSellerCategoryCommissionRate `json:"categoryCommissionRate"`         // 类目佣金比例
	MinInvestmentPromotionRate     int                                            `json:"minInvestmentPromotionRate"`     // 最低推广费率
	MinItemCommissionRate          int                                            `json:"minItemCommissionRate"`          // 最低商品佣金率
	RuleType                       int                                            `json:"ruleType"`                       // 规则类型
	GlobalMinCommissionRate        int                                            `json:"globalMinCommissionRate"`        // 全局最低佣金率
	MinInvestmentPromotionRateView string                                         `json:"minInvestmentPromotionRateView"` // 最低推广费率（展示用）
	MinItemCommissionRateView      string                                         `json:"minItemCommissionRateView"`      // 最低商品佣金率（展示用）
	GlobalMinCommissionRateView    string                                         `json:"globalMinCommissionRateView"`    // 全局最低佣金率（展示用）
}

// OpenDistributionSellerCategoryCommissionRate 类目佣金比例
type OpenDistributionSellerCategoryCommissionRate struct {
	CategoryId                int64                                `json:"categoryId"`                // 类目 ID
	MinCommissionRate         int                                  `json:"minCommissionRate"`         // 最低佣金率
	CategorySortedByHierarchy []OpenDistributionSellerCategoryInfo `json:"categorySortedByHierarchy"` // 层级排序的类目列表
	MinCommissionRateView     string                               `json:"minCommissionRateView"`     // 最低佣金率（展示用）
}

// OpenDistributionSellerCategoryInfo 类目信息
type OpenDistributionSellerCategoryInfo struct {
	CategoryId   int64  `json:"categoryId"`   // 类目 ID
	CategoryName string `json:"categoryName"` // 类目名称
}
