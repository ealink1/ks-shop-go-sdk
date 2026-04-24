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

// OpenDistributionInvestmentActivityOpenInfo 获取招商活动详情
func (k *KsShopClient) OpenDistributionInvestmentActivityOpenInfo(ctx context.Context, reqData *OpenDistributionInvestmentActivityOpenInfoRequest) (*OpenDistributionInvestmentActivityOpenInfoResponse, error) {
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
		"method":       k.FormatApi(OpenDistributionInvestmentActivityOpenInfoApi),
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
	values.Set("method", k.FormatApi(OpenDistributionInvestmentActivityOpenInfoApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

	// 发起请求
	endpoint := k.Env + OpenDistributionInvestmentActivityOpenInfoApi + "?" + values.Encode()
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
		return nil, fmt.Errorf("open_distribution_investment_activity_open_info status=%d body=%s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result OpenDistributionInvestmentActivityOpenInfoResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_distribution_investment_activity_open_info json_parse failed: %w", err)
	}

	return &result, nil
}

// OpenDistributionInvestmentActivityOpenInfoRequest 获取招商活动详情请求参数
type OpenDistributionInvestmentActivityOpenInfoRequest struct {
	ActivityId int64 `json:"activityId"` // 活动ID
}

// OpenDistributionInvestmentActivityOpenInfoResponse 获取招商活动详情响应
type OpenDistributionInvestmentActivityOpenInfoResponse struct {
	Result    int                                    `json:"result"`    // 结果码
	Msg       string                                 `json:"msg"`       // 结果描述
	ErrorMsg  string                                 `json:"error_msg"` // 错误信息
	Code      string                                 `json:"code"`      // 业务码
	Data      OpenDistributionInvestmentActivityData `json:"data"`      // 招商活动信息
	RequestId string                                 `json:"requestId"` // 请求 ID
	SubMsg    string                                 `json:"sub_msg"`   // 子错误信息
	SubCode   string                                 `json:"sub_code"`  // 子错误码
}

// OpenDistributionInvestmentActivityData 招商活动信息
type OpenDistributionInvestmentActivityData struct {
	ActivityId                             int64                                                `json:"activityId"`                             // 活动ID
	ActivityUserId                         int64                                                `json:"activityUserId"`                         // 活动团长ID
	ActivityUserNickName                   string                                               `json:"activityUserNickName"`                   // 团长昵称
	ActivityTitle                          string                                               `json:"activityTitle"`                          // 活动标题
	ActivityType                           int                                                  `json:"activityType"`                           // 活动类型
	ActivityBeginTime                      int64                                                `json:"activityBeginTime"`                      // 活动开始时间
	ActivityEndTime                        int64                                                `json:"activityEndTime"`                        // 活动结束时间
	ActivityStatus                         int                                                  `json:"activityStatus"`                         // 活动状态 (1-未发布 2-已发布 3-推广中 4-已失效 5-已删除)
	UserOpenInfo                           []OpenDistributionInvestmentActivityUserOpenInfo     `json:"userOpenInfo"`                           // 专属达人列表
	ActivityRuleSetViewOpenDto             OpenDistributionInvestmentActivityRuleSetViewOpenDto `json:"activityRuleSetViewOpenDto"`             // 活动规则信息
	ActivityUserHistoryDataInfoViewOpenDto OpenDistributionInvestmentActivityHistoryDataInfo    `json:"activityUserHistoryDataInfoViewOpenDto"` // 团长数据统计信息
	CreateTime                             int64                                                `json:"createTime"`                             // 创建时间
	UpdateTime                             int64                                                `json:"updateTime"`                             // 更新时间
	PreActivityUser                        []OpenDistributionInvestmentActivityUserOpenInfo     `json:"preActivityUser"`                        // 支持报名的专属团长列表
	PreExclusiveActivitySignType           int                                                  `json:"preExclusiveActivitySignType"`           // 该专属活动是否支持其余团长报名(1:不支持，2:全部，3:部分)
}

// OpenDistributionInvestmentActivityUserOpenInfo 用户开放信息
type OpenDistributionInvestmentActivityUserOpenInfo struct {
	UserId int64  `json:"userId"` // 用户ID
	Name   string `json:"name"`   // 用户昵称
}

// OpenDistributionInvestmentActivityRuleSetViewOpenDto 活动规则信息
type OpenDistributionInvestmentActivityRuleSetViewOpenDto struct {
	PromotionActivityMarketingRule OpenDistributionInvestmentActivityMarketingRule `json:"promotionActivityMarketingRule"` // 活动营销规则
}

// OpenDistributionInvestmentActivityMarketingRule 活动营销规则
type OpenDistributionInvestmentActivityMarketingRule struct {
	CategoryCommissionRate     []OpenDistributionInvestmentActivityCategoryCommissionRate `json:"categoryCommissionRate"`     // 分类佣金率
	MinItemCommissionRate      int                                                        `json:"minItemCommissionRate"`      // 最小商品佣金率
	MinInvestmentPromotionRate int                                                        `json:"minInvestmentPromotionRate"` // 最小投资推广率
}

// OpenDistributionInvestmentActivityCategoryCommissionRate 分类佣金率
type OpenDistributionInvestmentActivityCategoryCommissionRate struct {
	CategoryId                int64                                            `json:"categoryId"`                // 分类ID
	MinCommissionRate         int                                              `json:"minCommissionRate"`         // 最小佣金率
	CategorySortedByHierarchy []OpenDistributionInvestmentActivityCategoryInfo `json:"categorySortedByHierarchy"` // 按层级排序的分类
}

// OpenDistributionInvestmentActivityCategoryInfo 分类信息
type OpenDistributionInvestmentActivityCategoryInfo struct {
	CategoryName string `json:"categoryName"` // 分类名称
	CategoryId   int64  `json:"categoryId"`   // 分类ID
}

// OpenDistributionInvestmentActivityHistoryDataInfo 团长数据统计信息
type OpenDistributionInvestmentActivityHistoryDataInfo struct {
	CustomerPrice              int64 `json:"customerPrice"`              // 客户成交额
	ActivityCount              int64 `json:"activityCount"`              // 活动次数
	DistributeItemOrderCount   int64 `json:"distributeItemOrderCount"`   // 分销商品订单数
	DistributeTradeSellerCount int64 `json:"distributeTradeSellerCount"` // 分销交易卖家数
	DistributeTradeItemCount   int64 `json:"distributeTradeItemCount"`   // 分销交易商品数
}
