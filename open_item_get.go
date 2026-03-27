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

func (k *KsShopClient) OpenItemGet(ctx context.Context, reqData *OpenItemGetRequest) (*OpenItemGetResponse, error) {
	paramBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	values := url.Values{}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign, err := k.Sign(map[string]string{
		"access_token": k.AccToken,
		"method":       k.FormatApi(OpenItemGetApi),
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
	values.Set("method", k.FormatApi(OpenItemGetApi))
	values.Set("param", string(paramBytes))
	values.Set("sign", sign)
	values.Set("appkey", k.AppId)
	values.Set("version", k.Version)
	values.Set("signMethod", k.SignMethod)
	values.Set("timestamp", timestamp)

	endpoint := k.Env + OpenItemGetApi + "?" + values.Encode()
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
		return nil, fmt.Errorf("open_item_get status=%d body=%s", resp.StatusCode, string(body))
	}

	var result OpenItemGetResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("open_item_get json_parse failed: %w", err)
	}

	return &result, nil
}

// OpenItemGetRequest 获取商品详情请求参数
type OpenItemGetRequest struct {
	KwaiItemId           int64 `json:"kwaiItemId"`           // 商品 ID
	SupportNegativeStock bool  `json:"supportNegativeStock"` // 是否查询负库存
}

// OpenItemGetResponse 获取商品详情响应
type OpenItemGetResponse struct {
	Result    int             `json:"result"`    // 结果码
	Msg       string          `json:"msg"`       // 结果描述
	ErrorMsg  string          `json:"error_msg"` // 错误信息
	Code      string          `json:"code"`      // 主返回码
	Data      OpenItemGetData `json:"data"`      // 商品详情数据
	RequestId string          `json:"requestId"` // 请求 ID
	SubMsg    string          `json:"sub_msg"`   // 子错误信息
	SubCode   string          `json:"sub_code"`  // 子返回码
}

// OpenItemGetData 商品详情数据
type OpenItemGetData struct {
	ItemId             int64            `json:"itemId"`             // 商品 ID
	RelItemId          int64            `json:"relItemId"`          // 外部商品 ID
	CategoryId         int64            `json:"categoryId"`         // 类目 ID
	CategoryName       string           `json:"categoryName"`       // 类目名称
	ParentCategoryId   int64            `json:"parentCategoryId"`   // 父类目 ID
	ParentCategoryName string           `json:"parentCategoryName"` // 父类目名称
	RootCategoryId     int64            `json:"rootCategoryId"`     // 根类目 ID
	RootCategoryName   string           `json:"rootCategoryName"`   // 根类目名称
	Title              string           `json:"title"`              // 商品标题
	Details            string           `json:"details"`            // 商品详情描述
	ImageUrls          []string         `json:"imageUrls"`          // 商品主图列表
	DetailImageUrls    []string         `json:"detailImageUrls"`    // 详情图列表
	SkuInfos           []OpenItemGetSku `json:"skuInfos"`           // SKU 列表
	Price              int              `json:"price"`              // 商品价格（分）
	Volume             int              `json:"volume"`             // 销量
	Status             int              `json:"status"`             // 商品状态
	ItemStatus         int              `json:"itemStatus"`         // 商品审核状态
	ShelfStatus        int              `json:"shelfStatus"`        // 上下架状态
	MainImageUrl       string           `json:"mainImageUrl"`       // 商品主图
	LinkUrl            string           `json:"linkUrl"`            // 商品链接
	ExpressTemplateId  int64            `json:"expressTemplateId"`  // 运费模板 ID
	ItemType           int              `json:"itemType"`           // 商品类型
	CreateTime         int64            `json:"createTime"`         // 创建时间戳（毫秒）
	UpdateTime         int64            `json:"updateTime"`         // 更新时间戳（毫秒）
}

// OpenItemGetSku 商品 SKU 数据
type OpenItemGetSku struct {
	KwaiSkuId     int64                 `json:"kwaiSkuId"`     // 快手 SKU ID
	RelSkuId      int64                 `json:"relSkuId"`      // 外部 SKU ID
	SkuStock      int                   `json:"skuStock"`      // SKU 库存
	ImageUrl      string                `json:"imageUrl"`      // SKU 图片
	SkuSalePrice  int                   `json:"skuSalePrice"`  // SKU 售价（分）
	Volume        int                   `json:"volume"`        // SKU 销量
	IsValid       int                   `json:"isValid"`       // 是否有效
	CreateTime    int64                 `json:"createTime"`    // 创建时间戳（毫秒）
	UpdateTime    int64                 `json:"updateTime"`    // 更新时间戳（毫秒）
	Specification string                `json:"specification"` // SKU 规格描述
	AppKey        string                `json:"appkey"`        // 应用标识
	SkuNick       string                `json:"skuNick"`       // SKU 商家编码
	SkuProp       []OpenItemGetSkuProp  `json:"skuProp"`       // SKU 属性列表
	KwaiItemId    int64                 `json:"kwaiItemId"`    // 关联商品 ID
	GtinCode      string                `json:"gtinCode"`      // 商品条形码
	MealDetail    OpenItemGetMealDetail `json:"mealDetail"`    // 套餐信息
	ReserveStock  int                   `json:"reserveStock"`  // 预留库存
	PackageCode   string                `json:"packageCode"`   // 包装编码
	GoodsId       string                `json:"goodsId"`       // 货品 ID
	GoodsCode     []string              `json:"goodsCode"`     // 货品编码列表
	NegativeStock int                   `json:"negativeStock"` // 负库存数量
}

// OpenItemGetSkuProp SKU 属性数据
type OpenItemGetSkuProp struct {
	SkuPropId        int64                  `json:"skuPropId"`        // SKU 属性 ID
	PropId           int64                  `json:"propId"`           // 属性 ID
	PropName         string                 `json:"propName"`         // 属性名称
	PropValueId      int64                  `json:"propValueId"`      // 属性值 ID
	PropValueName    string                 `json:"propValueName"`    // 属性值名称
	ItemId           int64                  `json:"itemId"`           // 商品 ID
	SkuId            int64                  `json:"skuId"`            // SKU ID
	ImageUrl         string                 `json:"imageUrl"`         // 属性图片
	IsValid          int                    `json:"isValid"`          // 是否有效
	CreateTime       int64                  `json:"createTime"`       // 创建时间戳（毫秒）
	UpdateTime       int64                  `json:"updateTime"`       // 更新时间戳（毫秒）
	PropSortNum      int                    `json:"propSortNum"`      // 属性排序
	PropValueSortNum int                    `json:"propValueSortNum"` // 属性值排序
	PropValueRemarks string                 `json:"propValueRemarks"` // 属性值备注
	PropValueGroupId int64                  `json:"propValueGroupId"` // 属性值组 ID
	IsMainProp       int                    `json:"isMainProp"`       // 是否主属性
	PropVersion      int                    `json:"propVersion"`      // 属性版本
	MeasureInfo      OpenItemGetMeasureInfo `json:"measureInfo"`      // 计量信息
}

// OpenItemGetMeasureInfo 计量信息
type OpenItemGetMeasureInfo struct {
	TemplateId int64                     `json:"templateId"` // 模板 ID
	Value      []OpenItemGetMeasureValue `json:"value"`      // 计量值列表
}

// OpenItemGetMeasureValue 计量值
type OpenItemGetMeasureValue struct {
	Type          string `json:"type"`          // 值类型
	Value         string `json:"value"`         // 值内容
	UnitValueId   int64  `json:"unitValueId"`   // 单位值 ID
	UnitValueName string `json:"unitValueName"` // 单位名称
}

// OpenItemGetMealDetail 套餐详情
type OpenItemGetMealDetail struct {
	MealGroupDTOList []OpenItemGetMealGroupDTO `json:"mealGroupDTOList"` // 套餐分组列表
	LowestPeopleNum  int                       `json:"lowestPeopleNum"`  // 最少用餐人数
	HighestPeopleNum int                       `json:"highestPeopleNum"` // 最多用餐人数
	Remark           string                    `json:"remark"`           // 套餐备注
}

// OpenItemGetMealGroupDTO 套餐分组
type OpenItemGetMealGroupDTO struct {
	Title              string                      `json:"title"`              // 分组标题
	MealContentDTOList []OpenItemGetMealContentDTO `json:"mealContentDTOList"` // 分组内容列表
	FromNum            int                         `json:"fromNum"`            // 起选数量
	SelectNum          int                         `json:"selectNum"`          // 可选数量
}

// OpenItemGetMealContentDTO 套餐内容项
type OpenItemGetMealContentDTO struct {
	Title string `json:"title"` // 内容标题
	Count int    `json:"count"` // 数量
	Price int    `json:"price"` // 价格（分）
}
