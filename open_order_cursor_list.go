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

const (
	PayTypeUnknown      = 0  // 未知
	PayTypeWechat       = 1  // 微信
	PayTypeAlipay       = 2  // 支付宝
	PayTypePingan       = 3  // 平安
	PayTypeBankTransfer = 99 // 银行转账
	PayTypeAlipayCredit = 88 // 支付宝先用后付

	// 订单状态常量
	OrderStatusUnknown        = 0  // 未知状态
	OrderStatusPendingPayment = 10 // 待付款
	OrderStatusPaid           = 30 // 已付款/待发货
	OrderStatusShipped        = 40 // 已发货
	OrderStatusSigned         = 50 // 已签收
	OrderStatusSuccess        = 70 // 订单成功
	OrderStatusFailed         = 80 // 订单失败/订单关闭（订单取消会转为该状态）

	// 订单载体
	OrderCarrierLive       = "live"       // 直播
	OrderCarrierItemCard   = "itemCard"   // 商品卡
	OrderCarrierShortVideo = "shortVideo" // 短视频
	OrderCarrierOther      = "other"      // 其他
)

// OpenOrderCursorList 获取订单列表
func (k *KsShopClient) OpenOrderCursorList(ctx context.Context, reqData *OpenOrderCursorListRequest) (*OpenOrderCursorListResponse, error) {
	paramBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign, err := k.Sign(map[string]string{
		"access_token": k.AccToken,
		"method":       k.FormatApi(OpenOrderCursorListApi),
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
	values.Set("method", k.FormatApi(OpenOrderCursorListApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

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

	//if result.Result != 1 {
	//	codeText := strconv.Itoa(result.Result)
	//	if result.Code != "" {
	//		codeText = result.Code
	//	}
	//	return &result, fmt.Errorf("open_order_cursor_list failed: code=%s msg=%s", codeText, result.ErrorMsg)
	//}

	return &result, nil
}

type OpenOrderCursorListRequest struct {
	OrderViewStatus int    `json:"orderViewStatus"`
	PageSize        int    `json:"pageSize"`
	Sort            int    `json:"sort"`
	QueryType       int    `json:"queryType"`
	BeginTime       int64  `json:"beginTime"`
	EndTime         int64  `json:"endTime"`
	CpsType         int    `json:"cpsType"`
	Cursor          string `json:"cursor"`
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
	Cursor    string                     `json:"cursor"`
	OrderList []OpenOrderCursorListOrder `json:"orderList"`
	PageSize  int                        `json:"pageSize"`
	BeginTime int64                      `json:"beginTime"`
	EndTime   int64                      `json:"endTime"`
}

type OpenOrderCursorListOrder struct {
	OrderBaseInfo      OpenOrderCursorListOrderBaseInfo     `json:"orderBaseInfo"`
	OrderRefundList    []OpenOrderCursorListOrderRefund     `json:"orderRefundList"`
	OrderAddress       OpenOrderCursorListOrderAddress      `json:"orderAddress"`
	OrderLogisticsInfo []OpenOrderCursorListLogisticsInfo   `json:"orderLogisticsInfo"`
	OrderItemInfo      OpenOrderCursorListOrderItemInfo     `json:"orderItemInfo"`
	OrderDeliveryInfo  OpenOrderCursorListOrderDeliveryInfo `json:"orderDeliveryInfo"`
	SubOrderInfo       []OpenOrderCursorListSubOrderInfo    `json:"subOrderInfo"`
	OrderCpsInfo       OrderCpsInfo                         `json:"orderCpsInfo"`
}

type OpenOrderCursorListOrderBaseInfo struct {
	DiscountFee                   int                                     `json:"discountFee"`
	BuyerNick                     string                                  `json:"buyerNick"`
	PayTime                       int64                                   `json:"payTime"`
	OrderLabels                   any                                     `json:"orderLabels"`
	Channel                       string                                  `json:"channel"`
	Remark                        string                                  `json:"remark"`
	RemindShipmentSign            int                                     `json:"remindShipmentSign"`
	Oid                           int64                                   `json:"oid"`
	SellerOpenId                  string                                  `json:"sellerOpenId"`
	ExpressFee                    int                                     `json:"expressFee"`
	OrderSellerRoleInfo           OpenOrderCursorListOrderSellerRoleInfo  `json:"orderSellerRoleInfo"`
	BuyerImage                    string                                  `json:"buyerImage"`
	PayType                       int                                     `json:"payType"`
	MultiplePiecesNo              int                                     `json:"multiplePiecesNo"`
	OrderDomainCode               string                                  `json:"orderDomainCode"`
	ExpressCode                   string                                  `json:"expressCode"`
	EnableSplitDeliveryOrder      bool                                    `json:"enableSplitDeliveryOrder"`
	ValidPromiseShipmentTimeStamp int64                                   `json:"validPromiseShipmentTimeStamp"`
	GovernmentDiscount            int                                     `json:"governmentDiscount"`
	SellerNick                    string                                  `json:"sellerNick"`
	DisableDeliveryReasonCode     []int                                   `json:"disableDeliveryReasonCode"`
	RecvTime                      int64                                   `json:"recvTime"`
	BuyerOpenId                   string                                  `json:"buyerOpenId"`
	CpsType                       int                                     `json:"cpsType"`
	PromiseTimeStampOfDelivery    int64                                   `json:"promiseTimeStampOfDelivery"`
	RefundTime                    int64                                   `json:"refundTime"`
	RiskCode                      int                                     `json:"riskCode"`
	UpdateTime                    int64                                   `json:"updateTime"`
	OrderRiskEventInfo            []OpenOrderCursorListOrderRiskEventInfo `json:"orderRiskEventInfo"`
	TheDayOfDeliverGoodsTime      int                                     `json:"theDayOfDeliverGoodsTime"`
	CommentStatus                 int                                     `json:"commentStatus"`
	SendTime                      int64                                   `json:"sendTime"`
	TradeInPayAfterPromoAmount    int                                     `json:"tradeInPayAfterPromoAmount"`
	PreSale                       int                                     `json:"preSale"`
	CoType                        int                                     `json:"coType"`
	CreateTime                    int64                                   `json:"createTime"`
	TotalFee                      int                                     `json:"totalFee"`
	AllActivityType               []int                                   `json:"allActivityType"`
	SellerDelayPromiseTimeStamp   int64                                   `json:"sellerDelayPromiseTimeStamp"`
	PayChannel                    string                                  `json:"payChannel"`
	RemindShipmentTime            int64                                   `json:"remindShipmentTime"`
	ActivityType                  int                                     `json:"activityType"`
	AllowanceExpressFee           int                                     `json:"allowanceExpressFee"`
	PriorityDelivery              bool                                    `json:"priorityDelivery"`
	PayChannelDiscount            int                                     `json:"payChannelDiscount"`
	Status                        int                                     `json:"status"`
}

type OpenOrderCursorListOrderSellerRoleInfo struct {
	RoleId   int64  `json:"roleId"`
	RoleName string `json:"roleName"`
	RoleType int    `json:"roleType"`
}

type OpenOrderCursorListOrderRiskEventInfo struct {
	RiskType int    `json:"riskType"`
	RiskMsg  string `json:"riskMsg"`
}

type OpenOrderCursorListOrderRefund struct {
	RefundId     int64  `json:"refundId"`
	RefundStatus int    `json:"refundStatus"`
	RefundReason string `json:"refundReason"`
}

type OpenOrderCursorListOrderAddress struct {
	DistrictCode         int64  `json:"districtCode"`
	Town                 string `json:"town"`
	City                 string `json:"city"`
	TownCode             int64  `json:"townCode"`
	CityCode             int64  `json:"cityCode"`
	ProvinceCode         int64  `json:"provinceCode"`
	EncryptedMobile      string `json:"encryptedMobile"`
	EncryptedConsignee   string `json:"encryptedConsignee"`
	DesensitiseConsignee string `json:"desensitiseConsignee"`
	EncryptedAddress     string `json:"encryptedAddress"`
	Province             string `json:"province"`
	District             string `json:"district"`
	DesensitiseMobile    string `json:"desensitiseMobile"`
	DesensitiseAddress   string `json:"desensitiseAddress"`
}

type OpenOrderCursorListLogisticsInfo struct {
	LogisticsId int64  `json:"logisticsId"`
	ExpressNo   string `json:"expressNo"`
	ExpressCode int    `json:"expressCode"`
}

type OpenOrderCursorListOrderItemInfo struct {
	ItemPicUrl    string                            `json:"itemPicUrl"`
	ItemType      int                               `json:"itemType"`
	DiscountFee   int                               `json:"discountFee"`
	OriginalPrice int                               `json:"originalPrice"`
	ItemTitle     string                            `json:"itemTitle"`
	OrderItemId   int64                             `json:"orderItemId"`
	Num           int                               `json:"num"`
	ItemExtra     OpenOrderCursorListOrderItemExtra `json:"itemExtra"`
	WarehouseCode string                            `json:"warehouseCode"`
	ItemId        int64                             `json:"itemId"`
	RelItemId     int64                             `json:"relItemId"`
	RelSkuId      int64                             `json:"relSkuId"`
	Price         int                               `json:"price"`
	ItemLinkUrl   string                            `json:"itemLinkUrl"`
	SkuNick       string                            `json:"skuNick"`
	SkuDesc       string                            `json:"skuDesc"`
	GoodsCode     string                            `json:"goodsCode"`
	SkuId         int64                             `json:"skuId"`
}

type OpenOrderCursorListOrderItemExtra struct {
	BrandName    string                                   `json:"brandName"`
	EnergyLevel  string                                   `json:"energyLevel"`
	CategoryInfo OpenOrderCursorListOrderItemCategoryInfo `json:"categoryInfo"`
	ProductNo    string                                   `json:"productNo"`
}

type OpenOrderCursorListOrderItemCategoryInfo struct {
	GovCategory     string `json:"govCategory"`
	ItemCid         int64  `json:"itemCid"`
	GovCategoryCode string `json:"govCategoryCode"`
	CategoryName    string `json:"categoryName"`
}

type OpenOrderCursorListOrderDeliveryInfo struct {
	SplitDeliveryOrder  bool   `json:"splitDeliveryOrder"`
	MergeDeliveryType   int    `json:"mergeDeliveryType"`
	DeliveryNum         int    `json:"deliveryNum"`
	EnableAppendPackage bool   `json:"enableAppendPackage"`
	TotalPackageNum     int    `json:"totalPackageNum"`
	OpenAddressId       string `json:"openAddressId"`
	MaxPackageNum       int    `json:"maxPackageNum"`
	DeliveryStatus      int    `json:"deliveryStatus"`
	PackageNum          int    `json:"packageNum"`
}

type OpenOrderCursorListSubOrderInfo struct {
	SubOrderId int64 `json:"subOrderId"`
}

type OrderCpsInfo struct {
	DistributorId         int64  `json:"distributorId"`
	DistributorName       string `json:"distributorName"`
	ActivityUserId        int64  `json:"activityUserId"`
	ActivityUserNickName  string `json:"activityUserNickName"`
	KwaiMoneyUserId       int64  `json:"kwaiMoneyUserId"`
	KwaiMoneyUserNickName string `json:"kwaiMoneyUserNickName"`
}
