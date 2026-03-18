package ks_shop_go_sdk

const (
	KwaishopOrderDeliverySuccess = "kwaishop_order_deliverySuccess" // 订单已收货消息
	KwaishopOrderDelivering      = "kwaishop_order_delivering"      // 订单已发货消息
	KwaishopOrderPaySuccess      = "kwaishop_order_paySuccess"      // 订单已支付消息
	KwaishopOrderTotalFeeChange  = "kwaishop_order_totalFeeChange"  // 订单费用变更消息
	KwaishopOrderOrderSuccess    = "kwaishop_order_orderSuccess"    // 订单交易成功消息
	KwaishopOrderAddOrder        = "kwaishop_order_addOrder"        // 订单新增消息
	KwaishopOrderOrderFail       = "kwaishop_order_orderFail"       // 订单交易失败消息
)

// OrderMessage 订单消息
// {"eventId":"243_4027_2375715766426","msgId":"84_2375715766425","bizId":2607602240240301,"userId":2161524184,"openId":"f1b68e8f2a4995ed14bf57245903fd9a","appKey":"ks699183844582124027","event":"kwaishop_order_addOrder","info":"{\"oid\":2607602240240301,\"sellerId\":2161524184,\"createTime\":1773750942425,\"customizedInfo\":{\"customizedItemType\":0,\"customizedUrl\":\"\",\"customizedCode\":\"\"},\"openId\":\"Nb8oPhmFTcwPpYcJdfRnTr70nlt1u2eTFhvicWXG13P\",\"buyerOpenId\":\"f1b68e8f2a4995ed14bf572415fe3aff\"}","status":2,"createTime":1773750943589,"updateTime":1773750945776,"operator":"2323245113"}
type OrderMessage struct {
	EventId    string `json:"eventId"`    // 事件ID
	MsgId      string `json:"msgId"`      // 消息ID
	BizId      int64  `json:"bizId"`      // 业务ID
	UserId     int64  `json:"userId"`     // 用户ID
	OpenId     string `json:"openId"`     // 开放ID
	AppKey     string `json:"appKey"`     // 应用Key
	Event      string `json:"event"`      // 事件类型
	Info       string `json:"info"`       // 消息内容 (JSON 字符串)
	Status     int    `json:"status"`     // 状态
	CreateTime int64  `json:"createTime"` // 创建时间
	UpdateTime int64  `json:"updateTime"` // 更新时间
	Operator   string `json:"operator"`   // 操作者
}

// OrderMessageInfo 订单消息详情
type OrderAddOrder struct {
	Oid            int64                       `json:"oid"`            // 订单ID
	SellerId       int64                       `json:"sellerId"`       // 卖家ID
	CreateTime     int64                       `json:"createTime"`     // 创建时间
	CustomizedInfo OrderAddOrderCustomizedInfo `json:"customizedInfo"` // 定制信息
	OpenId         string                      `json:"openId"`         // 卖家开放ID
	BuyerOpenId    string                      `json:"buyerOpenId"`    // 买家开放ID
}

// OrderAddOrderCustomizedInfo 订单定制信息
type OrderAddOrderCustomizedInfo struct {
	CustomizedItemType int    `json:"customizedItemType"` // 定制商品类型
	CustomizedUrl      string `json:"customizedUrl"`      // 定制URL
	CustomizedCode     string `json:"customizedCode"`     // 定制编码
}

// 订单已发货消息
type OrderMessageStatusChange struct {
	Oid        int64 `json:"oid"`        // 订单ID
	SellerId   int64 `json:"sellerId"`   // 商家ID
	Status     int   `json:"status"`     // 订单状态：[0, "未知状态"], [10, "待付款"], [30, "已付款"], [40, "已发货"], [50, "已签收"], [70, "订单成功"], [80, "订单失败"]; 订单取消后会转为“订单失败”状态
	UpdateTime int64 `json:"updateTime"` // 业务变更时间
}

type OrderMessageFeeChange struct {
	Oid              int    `json:"oid"`
	SellerId         int    `json:"sellerId"`
	TotalFee         int    `json:"totalFee"`
	BeforeTotalFee   int    `json:"beforeTotalFee"`
	ExpressFee       int    `json:"expressFee"`
	BeforeExpressFee int    `json:"beforeExpressFee"`
	OpenId           string `json:"openId"`
	UpdateTime       int    `json:"updateTime"`
	BuyerOpenId      string `json:"buyerOpenId"`
}
