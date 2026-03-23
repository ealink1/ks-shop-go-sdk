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

func (k *KsShopClient) OpenItemListGet(ctx context.Context, reqData *OpenItemListGetRequest) (*OpenItemListGetResponse, error) {
	paramBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign, err := k.Sign(map[string]string{
		"access_token": k.AccToken,
		"method":       k.FormatApi(OpenItemListGetApi),
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
	values.Set("method", k.FormatApi(OpenItemListGetApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

	endpoint := k.Env + OpenItemListGetApi + "?" + values.Encode()
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
		return nil, fmt.Errorf("open_item_list_get status=%d body=%s", resp.StatusCode, string(body))
	}

	var result OpenItemListGetResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_item_list_get json_parse failed: %w", err)
	}

	//if result.Result != 1 {
	//	codeText := strconv.Itoa(result.Result)
	//	if result.Code != "" {
	//		codeText = result.Code
	//	}
	//	return &result, fmt.Errorf("open_item_list_get failed: code=%s msg=%s", codeText, result.ErrorMsg)
	//}

	return &result, nil
}

type OpenItemListGetRequest struct {
	ItemStatus           int  `json:"itemStatus"`
	ItemType             int  `json:"itemType"`
	PageNumber           int  `json:"pageNumber"`
	PageSize             int  `json:"pageSize"`
	OnOfflineStatus      int  `json:"onOfflineStatus"`
	SupportNegativeStock bool `json:"supportNegativeStock"`
}

type OpenItemListGetResponse struct {
	Result    int                 `json:"result"`
	Msg       string              `json:"msg"`
	ErrorMsg  string              `json:"error_msg"`
	Code      string              `json:"code"`
	Data      OpenItemListGetData `json:"data"`
	RequestId string              `json:"requestId"`
	SubMsg    string              `json:"sub_msg"`
	SubCode   string              `json:"sub_code"`
}

type OpenItemListGetData struct {
	TotalPage            int                   `json:"totalPage"`
	CurrentPageNumber    int                   `json:"currentPageNumber"`
	CurrentPageItemCount int                   `json:"currentPageItemCount"`
	Items                []OpenItemListGetItem `json:"items"`
	TotalItemCount       int                   `json:"totalItemCount"`
}

type OpenItemListGetItem struct {
	Instructions            string                     `json:"instructions"`
	ItemType                int                        `json:"itemType"`
	ShelfStatus             int                        `json:"shelfStatus"`
	SkuList                 []OpenItemListGetSku       `json:"skuList"`
	DuplicationStatus       int                        `json:"duplicationStatus"`
	Title                   string                     `json:"title"`
	CategoryName            string                     `json:"categoryName"`
	DuplicationReason       string                     `json:"duplicationReason"`
	RelItemId               int64                      `json:"relItemId"`
	FromType                int                        `json:"fromType"`
	Price                   int                        `json:"price"`
	ItemStatus              int                        `json:"itemStatus"`
	LinkUrl                 string                     `json:"linkUrl"`
	Details                 string                     `json:"details"`
	MainImageUrl            string                     `json:"mainImageUrl"`
	ServiceRule             OpenItemListGetServiceRule `json:"serviceRule"`
	KwaiItemId              int64                      `json:"kwaiItemId"`
	ExpressTemplateId       int64                      `json:"expressTemplateId"`
	UpdateTime              int64                      `json:"updateTime"`
	UserId                  int64                      `json:"userId"`
	Volume                  int                        `json:"volume"`
	MultipleStock           bool                       `json:"multipleStock"`
	AuditReason             string                     `json:"auditReason"`
	ShelfStatusUpdateReason string                     `json:"shelfStatusUpdateReason"`
	CreateTime              int64                      `json:"createTime"`
	ImageUrls               []string                   `json:"imageUrls"`
	AuditStatus             int                        `json:"auditStatus"`
	AppKey                  string                     `json:"appkey"`
	CategoryId              int64                      `json:"categoryId"`
	Status                  int                        `json:"status"`
}

type OpenItemListGetSku struct {
	KwaiItemId    int64                     `json:"kwaiItemId"`
	GoodsId       string                    `json:"goodsId"`
	IsValid       int                       `json:"isValid"`
	SkuSalePrice  int                       `json:"skuSalePrice"`
	SkuStock      int                       `json:"skuStock"`
	Specification string                    `json:"specification"`
	UpdateTime    int64                     `json:"updateTime"`
	Volume        int                       `json:"volume"`
	MealDetail    OpenItemListGetMealDetail `json:"mealDetail"`
	CreateTime    int64                     `json:"createTime"`
	RelSkuId      int64                     `json:"relSkuId"`
	ImageUrl      string                    `json:"imageUrl"`
	KwaiSkuId     int64                     `json:"kwaiSkuId"`
	NegativeStock int                       `json:"negativeStock"`
	SkuNick       string                    `json:"skuNick"`
	GtinCode      string                    `json:"gtinCode"`
	SkuProp       []OpenItemListGetSkuProp  `json:"skuProp"`
	AppKey        string                    `json:"appkey"`
	GoodsCode     []string                  `json:"goodsCode"`
}

type OpenItemListGetMealDetail struct {
	MealGroup        []OpenItemListGetMealGroup    `json:"mealGroup"`
	LowestPeopleNum  int                           `json:"lowestPeopleNum"`
	HighestPeopleNum int                           `json:"highestPeopleNum"`
	MealGroupDTOList []OpenItemListGetMealGroupDTO `json:"mealGroupDTOList"`
}

type OpenItemListGetMealGroup struct {
	MealId      int64   `json:"mealId"`
	MealName    string  `json:"mealName"`
	PeopleNum   int     `json:"peopleNum"`
	RelSkuIds   []int64 `json:"relSkuIds"`
	RelItemIds  []int64 `json:"relItemIds"`
	MealPrice   int     `json:"mealPrice"`
	MealDesc    string  `json:"mealDesc"`
	MealEnabled bool    `json:"mealEnabled"`
}

type OpenItemListGetMealGroupDTO struct {
	MealGroupId   int64   `json:"mealGroupId"`
	MealGroupName string  `json:"mealGroupName"`
	RelSkuIds     []int64 `json:"relSkuIds"`
	RelItemIds    []int64 `json:"relItemIds"`
	Required      bool    `json:"required"`
}

type OpenItemListGetSkuProp struct {
	PropValueGroupId int64  `json:"propValueGroupId"`
	IsValid          int    `json:"isValid"`
	PropId           int64  `json:"propId"`
	PropValueSortNum int    `json:"propValueSortNum"`
	SkuPropId        int64  `json:"skuPropId"`
	UpdateTime       int64  `json:"updateTime"`
	PropName         string `json:"propName"`
	ItemId           int64  `json:"itemId"`
	PropValueId      int64  `json:"propValueId"`
	IsMainProp       int    `json:"isMainProp"`
	CreateTime       int64  `json:"createTime"`
	PropVersion      int    `json:"propVersion"`
	ImageUrl         string `json:"imageUrl"`
	PropSortNum      int    `json:"propSortNum"`
	PropValueName    string `json:"propValueName"`
	PropValueRemarks string `json:"propValueRemarks"`
	SkuId            int64  `json:"skuId"`
}

type OpenItemListGetServiceRule struct {
	CertStartTime            int64                         `json:"certStartTime"`
	OrderPurchaseLimitType   int                           `json:"orderPurchaseLimitType"`
	DeliveryMethod           string                        `json:"deliveryMethod"`
	CertMerchantCode         string                        `json:"certMerchantCode"`
	CustomerInfo             OpenItemListGetCustomerInfo   `json:"customerInfo"`
	DeliveryTimeMode         string                        `json:"deliveryTimeMode"`
	MaxOrderCount            int                           `json:"maxOrderCount"`
	TheDayOfDeliverGoodsTime int                           `json:"theDayOfDeliverGoodsTime"`
	MinOrderCount            int                           `json:"minOrderCount"`
	PriceProtectDays         int                           `json:"priceProtectDays"`
	RefundRule               string                        `json:"refundRule"`
	ServicePromise           OpenItemListGetServicePromise `json:"servicePromise"`
	CertExpireType           int                           `json:"certExpireType"`
	CertExpDays              int                           `json:"certExpDays"`
	PromiseDeliveryTime      int                           `json:"promiseDeliveryTime"`
	CertEndTime              int64                         `json:"certEndTime"`
}

type OpenItemListGetCustomerInfo struct {
	CustomerInfoType        string   `json:"customerInfoType"`
	CustomerCertificateType []string `json:"customerCertificateType"`
}

type OpenItemListGetServicePromise struct {
	FreshRotRefund  bool `json:"freshRotRefund"`
	BrokenRefund    bool `json:"brokenRefund"`
	AllergyRefund   bool `json:"allergyRefund"`
	CrabRefund      bool `json:"crabRefund"`
	WeightGuarantee bool `json:"weightGuarantee"`
}
