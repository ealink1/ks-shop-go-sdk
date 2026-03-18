package ks_shop_go_sdk

import (
	"encoding/json"
	"fmt"
)

// DecodeAndUnmarshalOrderAddOrder 解码并反序列化订单新增消息
// content: 加密的消息内容
// signKey: 签名密钥 (推送消息解密 Key)
func DecodeAndUnmarshalOrderAddOrder(content, signKey string) (*OrderMessage, *OrderAddOrder, error) {
	// 1. 解码加密的消息
	decoded := MessageDecode(content, signKey)
	if decoded == nil {
		return nil, nil, fmt.Errorf("message decode failed")
	}

	// 2. 解析外层 OrderMessage 结构
	res := &OrderMessage{}
	if err := json.Unmarshal(decoded, res); err != nil {
		return nil, nil, fmt.Errorf("unmarshal OrderMessage failed: %w", err)
	}

	// 3. 解析内层 Info 字段 (OrderAddOrder)
	info := &OrderAddOrder{}
	if res.Info != "" {
		if err := json.Unmarshal([]byte(res.Info), info); err != nil {
			return res, nil, fmt.Errorf("unmarshal OrderAddOrder failed: %w", err)
		}
	}

	return res, info, nil
}

func DecodeAndUnmarshalOrderStatusChange(content, signKey string) (*OrderMessage, *OrderMessageStatusChange, error) {
	// 1. 解码加密的消息
	decoded := MessageDecode(content, signKey)
	if decoded == nil {
		return nil, nil, fmt.Errorf("message decode failed")
	}

	// 2. 解析外层 OrderMessage 结构
	res := &OrderMessage{}
	if err := json.Unmarshal(decoded, res); err != nil {
		return nil, nil, fmt.Errorf("unmarshal OrderMessage failed: %w", err)
	}

	// 3. 解析内层 Info 字段 (OrderAddOrder)
	info := &OrderMessageStatusChange{}
	if res.Info != "" {
		if err := json.Unmarshal([]byte(res.Info), info); err != nil {
			return res, nil, fmt.Errorf("unmarshal OrderAddOrder failed: %w", err)
		}
	}

	return res, info, nil
}

// DecodeAndUnmarshalOrderFeeChange 订单状态变更消息解码方法
func DecodeAndUnmarshalOrderFeeChange(content, signKey string) (*OrderMessage, *OrderMessageFeeChange, error) {
	// 1. 解码加密的消息
	decoded := MessageDecode(content, signKey)
	if decoded == nil {
		return nil, nil, fmt.Errorf("message decode failed")
	}

	// 2. 解析外层 OrderMessage 结构
	res := &OrderMessage{}
	if err := json.Unmarshal(decoded, res); err != nil {
		return nil, nil, fmt.Errorf("unmarshal OrderMessage failed: %w", err)
	}

	// 3. 解析内层 Info 字段 (OrderAddOrder)
	info := &OrderMessageFeeChange{}
	if res.Info != "" {
		if err := json.Unmarshal([]byte(res.Info), info); err != nil {
			return res, nil, fmt.Errorf("unmarshal OrderAddOrder failed: %w", err)
		}
	}

	return res, info, nil
}

// DecodeOrderMessage 基础消息解码
// 仅解码消息并解析出外层结构，方便根据 Event 类型后续自行解析 Info
func DecodeOrderMessage(content, signKey string) (*OrderMessage, error) {
	decoded := MessageDecode(content, signKey)
	if decoded == nil {
		return nil, fmt.Errorf("message decode failed")
	}

	res := &OrderMessage{}
	if err := json.Unmarshal(decoded, res); err != nil {
		return nil, fmt.Errorf("unmarshal OrderMessage failed: %w", err)
	}

	return res, nil
}
