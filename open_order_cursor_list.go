package ks_shop_go_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type OpenOrderCursorListParam struct {
	OrderViewStatus int    `json:"orderViewStatus"`
	PageSize        int    `json:"pageSize"`
	Sort            int    `json:"sort"`
	QueryType       int    `json:"queryType"`
	BeginTime       int64  `json:"beginTime"`
	EndTime         int64  `json:"endTime"`
	CpsType         int    `json:"cpsType"`
	Cursor          string `json:"cursor"`
}

type OpenOrderCursorListRequest struct {
	AccessToken string
	Sign        string
	Timestamp   int64
	AppKey      string
	Version     string
	SignMethod  string
	Method      string
	Param       OpenOrderCursorListParam
}

type OpenOrderCursorListResponse struct {
	Result    int                     `json:"result"`
	Msg       string                  `json:"msg"`
	ErrorMsg  string                  `json:"error_msg"`
	Code      string                  `json:"code"`
	Data      OpenOrderCursorListData `json:"data"`
	RequestId string                  `json:"requestId"`
	SubMsg    string                  `json:"sub_msg"`
	SubCode   string                  `json:"sub_code"`
}

type OpenOrderCursorListData struct {
	Cursor    string            `json:"cursor"`
	OrderList []json.RawMessage `json:"orderList"`
	PageSize  int               `json:"pageSize"`
	BeginTime int64             `json:"beginTime"`
	EndTime   int64             `json:"endTime"`
}

// OpenOrderCursorList 获取订单列表
func (k *KsShopClient) OpenOrderCursorList(ctx context.Context, reqData OpenOrderCursorListRequest) (*OpenOrderCursorListResponse, error) {
	method := reqData.Method
	if method == "" {
		method = "open.order.cursor.list"
	}

	appKey := reqData.AppKey
	if appKey == "" {
		appKey = k.AppId
	}

	version := reqData.Version
	if version == "" {
		version = "1"
	}

	signMethod := reqData.SignMethod
	if signMethod == "" {
		signMethod = "MD5"
	}

	paramBytes, err := json.Marshal(reqData.Param)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	values.Set("access_token", reqData.AccessToken)
	values.Set("method", method)
	values.Set("param", string(paramBytes))
	values.Set("sign", reqData.Sign)
	values.Set("appkey", appKey)
	values.Set("version", version)
	values.Set("signMethod", signMethod)
	values.Set("timestamp", strconv.FormatInt(reqData.Timestamp, 10))

	endpoint := k.Env + OpenOrderCursorListApi + "?" + values.Encode()
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
		return nil, fmt.Errorf("open_order_cursor_list status=%d body=%s", resp.StatusCode, string(body))
	}

	var result OpenOrderCursorListResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_order_cursor_list json_parse failed: %w", err)
	}

	if result.Result != 1 {
		codeText := strconv.Itoa(result.Result)
		if result.Code != "" {
			codeText = result.Code
		}
		return &result, fmt.Errorf("open_order_cursor_list failed: code=%s msg=%s", codeText, result.ErrorMsg)
	}

	return &result, nil
}
