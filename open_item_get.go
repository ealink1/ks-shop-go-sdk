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
	ItemId                int64                          `json:"itemId"`                // 商品 ID
	RelItemId             int64                          `json:"relItemId"`             // 外部商品 ID
	CategoryId            int64                          `json:"categoryId"`            // 类目 ID
	CategoryName          string                         `json:"categoryName"`          // 类目名称
	ParentCategoryId      int64                          `json:"parentCategoryId"`      // 父类目 ID
	ParentCategoryName    string                         `json:"parentCategoryName"`    // 父类目名称
	RootCategoryId        int64                          `json:"rootCategoryId"`        // 根类目 ID
	RootCategoryName      string                         `json:"rootCategoryName"`      // 根类目名称
	Title                 string                         `json:"title"`                 // 商品标题
	Details               string                         `json:"details"`               // 商品详情描述
	ImageUrls             []string                       `json:"imageUrls"`             // 商品主图列表
	DetailImageUrls       []string                       `json:"detailImageUrls"`       // 详情图列表
	SkuInfos              []OpenItemGetSku               `json:"skuInfos"`              // SKU 列表
	ServiceRule           OpenItemGetServiceRule         `json:"serviceRule"`           // 服务规则
	Price                 int                            `json:"price"`                 // 商品价格（分）
	Volume                int                            `json:"volume"`                // 销量
	Status                int                            `json:"status"`                // 商品状态
	ItemStatus            int                            `json:"itemStatus"`            // 商品审核状态
	ShelfStatus           int                            `json:"shelfStatus"`           // 上下架状态
	AuditStatus           int                            `json:"auditStatus"`           // 审核状态
	AuditReason           string                         `json:"auditReason"`           // 审核原因
	OnOfflineStatus       int                            `json:"onOfflineStatus"`       // 上下线状态
	DuplicationStatus     int                            `json:"duplicationStatus"`     // 重复铺货状态
	DuplicationReason     string                         `json:"duplicationReason"`     // 重复铺货原因
	MultipleStock         bool                           `json:"multipleStock"`         // 是否多仓库存
	MainImageUrl          string                         `json:"mainImageUrl"`          // 商品主图
	LinkUrl               string                         `json:"linkUrl"`               // 商品链接
	TimeOfSale            int64                          `json:"timeOfSale"`            // 开售时间戳
	ItemRemark            string                         `json:"itemRemark"`            // 商品备注
	Instructions          string                         `json:"instructions"`          // 使用说明
	SellingPoint          string                         `json:"sellingPoint"`          // 卖点
	ShortTitle            string                         `json:"shortTitle"`            // 短标题
	ContractPhone         bool                           `json:"contractPhone"`         // 是否支持联系手机号
	OfflineReason         string                         `json:"offlineReason"`         // 下架原因
	ExpressTemplateId     int64                          `json:"expressTemplateId"`     // 运费模板 ID
	ItemType              int                            `json:"itemType"`              // 商品类型
	PurchaseLimit         bool                           `json:"purchaseLimit"`         // 是否限购
	LimitCount            int                            `json:"limitCount"`            // 限购数量
	PoiIds                []int64                        `json:"poiIds"`                // 适用门店 ID 列表
	ParentCategoryList    []OpenItemGetCategory          `json:"parentCategoryList"`    // 父类目信息列表
	ItemPropValues        []OpenItemGetItemPropValue     `json:"itemPropValues"`        // 商品属性值列表
	ItemQualification     []OpenItemGetItemQualification `json:"itemQualification"`     // 商品资质列表
	ThreeQuartersImageUrl []string                       `json:"threeQuartersImageUrl"` // 三比四主图列表
	WhiteBaseImageUrl     string                         `json:"whiteBaseImageUrl"`     // 白底图地址
	TransparentImageUrl   string                         `json:"transparentImageUrl"`   // 透明底图地址
	ItemVideo             OpenItemGetItemVideo           `json:"itemVideo"`             // 商品视频信息
	SizeChart             OpenItemGetSizeChart           `json:"sizeChart"`             // 尺码表信息
	SpuId                 int64                          `json:"spuId"`                 // SPU ID
	CreatedTime           int64                          `json:"createdTime"`           // 创建时间戳（毫秒）
	CreateTime            int64                          `json:"createTime"`            // 创建时间戳（毫秒）
	UpdateTime            int64                          `json:"updateTime"`            // 更新时间戳（毫秒）
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

// OpenItemGetServiceRule 商品服务规则
type OpenItemGetServiceRule struct {
	RefundRule               string                         `json:"refundRule"`               // 退货规则
	TheDayOfDeliverGoodsTime int                            `json:"theDayOfDeliverGoodsTime"` // 当日发货截止时间
	PromiseDeliveryTime      int                            `json:"promiseDeliveryTime"`      // 承诺发货时效（秒）
	ServicePromise           OpenItemGetServicePromise      `json:"servicePromise"`           // 服务承诺
	UnavailableTimeRule      OpenItemGetUnavailableTimeRule `json:"unavailableTimeRule"`      // 不可售时间规则
	CertMerchantCode         string                         `json:"certMerchantCode"`         // 凭证码商家编码
	DeliveryMethod           string                         `json:"deliveryMethod"`           // 发货方式
	CertExpireType           int                            `json:"certExpireType"`           // 凭证过期类型
	CertStartTime            int64                          `json:"certStartTime"`            // 凭证生效时间戳（毫秒）
	CertEndTime              int64                          `json:"certEndTime"`              // 凭证失效时间戳（毫秒）
	CertExpDays              int64                          `json:"certExpDays"`              // 凭证有效天数
	OrderPurchaseLimitType   int                            `json:"orderPurchaseLimitType"`   // 每单限购类型
	MinOrderCount            int                            `json:"minOrderCount"`            // 每单最小购买数量
	MaxOrderCount            int                            `json:"maxOrderCount"`            // 每单最大购买数量
	CustomerInfo             OpenItemGetCustomerInfo        `json:"customerInfo"`             // 用户信息收集规则
	PriceProtectDays         int                            `json:"priceProtectDays"`         // 价格保护天数
	DeliveryTimeMode         string                         `json:"deliveryTimeMode"`         // 发货时间模式
}

// OpenItemGetServicePromise 服务承诺
type OpenItemGetServicePromise struct {
	FreshRotRefund  bool `json:"freshRotRefund"`  // 坏果包赔
	BrokenRefund    bool `json:"brokenRefund"`    // 破损包赔
	AllergyRefund   bool `json:"allergyRefund"`   // 过敏包赔
	CrabRefund      bool `json:"crabRefund"`      // 缺蟹包赔
	WeightGuarantee bool `json:"weightGuarantee"` // 足斤足两保障
}

// OpenItemGetUnavailableTimeRule 不可售时间规则
type OpenItemGetUnavailableTimeRule struct {
	Weeks      []int                         `json:"weeks"`      // 不可售周几
	Holidays   []int                         `json:"holidays"`   // 不可售节假日类型
	TimeRanges []OpenItemGetUnavailableRange `json:"timeRanges"` // 不可售时间段
}

// OpenItemGetUnavailableRange 不可售时间段
type OpenItemGetUnavailableRange struct {
	StartTime int64 `json:"startTime"` // 开始时间戳（毫秒）
	EndTime   int64 `json:"endTime"`   // 结束时间戳（毫秒）
}

// OpenItemGetCustomerInfo 用户信息收集规则
type OpenItemGetCustomerInfo struct {
	CustomerInfoType        string `json:"customerInfoType"`        // 用户信息收集类型
	CustomerCertificateType []int  `json:"customerCertificateType"` // 证件类型列表
}

// OpenItemGetCategory 父类目信息
type OpenItemGetCategory struct {
	CategoryId   int64  `json:"categoryId"`   // 类目 ID
	CategoryName string `json:"categoryName"` // 类目名称
	CategoryPid  int64  `json:"categoryPid"`  // 父类目 ID
}

// OpenItemGetItemPropValue 商品属性值
type OpenItemGetItemPropValue struct {
	PropId                 int64                  `json:"propId"`                 // 属性 ID
	RadioPropValue         OpenItemGetPropValue   `json:"radioPropValue"`         // 单选属性值
	CheckBoxPropValuesList []OpenItemGetPropValue `json:"checkBoxPropValuesList"` // 多选属性值列表
	TextPropValue          string                 `json:"textPropValue"`          // 文本属性值
	DatetimeTimestamp      int64                  `json:"datetimeTimestamp"`      // 日期时间戳（毫秒）
	DateRange              OpenItemGetDateRange   `json:"dateRange"`              // 日期区间
	SortNum                int                    `json:"sortNum"`                // 排序号
	ImagePropValues        []OpenItemGetPropValue `json:"imagePropValues"`        // 图片属性值列表
	PropName               string                 `json:"propName"`               // 属性名称
	PropAlias              string                 `json:"propAlias"`              // 属性别名
	InputType              int                    `json:"inputType"`              // 输入类型
	PropType               int                    `json:"propType"`               // 属性类型
	UnitPropValueId        int64                  `json:"unitPropValueId"`        // 单位属性值 ID
	UnitPropValueName      string                 `json:"unitPropValueName"`      // 单位属性值名称
}

// OpenItemGetPropValue 属性值
type OpenItemGetPropValue struct {
	PropValueId int64  `json:"propValueId"` // 属性值 ID
	PropValue   string `json:"propValue"`   // 属性值名称
}

// OpenItemGetDateRange 日期区间
type OpenItemGetDateRange struct {
	StartTimeTimestamp int64 `json:"startTimeTimestamp"` // 开始时间戳（毫秒）
	EndTimeTimestamp   int64 `json:"endTimeTimestamp"`   // 结束时间戳（毫秒）
}

// OpenItemGetItemQualification 商品资质
type OpenItemGetItemQualification struct {
	QualificationDataId   int64    `json:"qualificationDataId"`   // 资质数据 ID
	Certification         string   `json:"certification"`         // 资质证书号
	QualificationFiles    []string `json:"qualificationFiles"`    // 资质文件列表
	ValidDateStart        string   `json:"validDateStart"`        // 有效期开始时间
	ValidDateEnd          string   `json:"validDateEnd"`          // 有效期结束时间
	AuditTime             string   `json:"auditTime"`             // 审核时间
	RelatedItemCount      int      `json:"relatedItemCount"`      // 关联商品数
	Remark                string   `json:"remark"`                // 资质备注
	DetailPageUrl         string   `json:"detailPageUrl"`         // 详情页链接
	QualificationMetaName string   `json:"qualificationMetaName"` // 资质模板名称
	QualificationMetaId   int64    `json:"qualificationMetaId"`   // 资质模板 ID
}

// OpenItemGetItemVideo 商品视频
type OpenItemGetItemVideo struct {
	VideoId   string `json:"videoId"`   // 视频 ID
	VideoType int    `json:"videoType"` // 视频类型
}

// OpenItemGetSizeChart 尺码表
type OpenItemGetSizeChart struct {
	SizeChartId             int64                     `json:"sizeChartId"`             // 尺码表 ID
	TemplateTypePropValueId int64                     `json:"templateTypePropValueId"` // 模板类型属性值 ID
	GroupId                 int64                     `json:"groupId"`                 // 分组 ID
	SizeChartNote           string                    `json:"sizeChartNote"`           // 尺码表备注
	SizeChartTable          OpenItemGetSizeChartTable `json:"sizeChartTable"`          // 尺码表内容
	ImageWidth              int                       `json:"imageWidth"`              // 图片宽度
	ImageHeight             int                       `json:"imageHeight"`             // 图片高度
	ImageUrl                string                    `json:"imageUrl"`                // 图片地址
}

// OpenItemGetSizeChartTable 尺码表内容
type OpenItemGetSizeChartTable struct {
	HeaderParam  []string                  `json:"headerParam"`  // 表头参数
	SizeChartRow []OpenItemGetSizeChartRow `json:"sizeChartRow"` // 表格行
}

// OpenItemGetSizeChartRow 尺码表行
type OpenItemGetSizeChartRow struct {
	SizeChartUnit []OpenItemGetSizeChartUnit `json:"sizeChartUnit"` // 行内单元
}

// OpenItemGetSizeChartUnit 尺码表单元
type OpenItemGetSizeChartUnit struct {
	ParamName      string   `json:"paramName"`      // 参数名称
	ParamUnit      string   `json:"paramUnit"`      // 参数单位
	ParamType      string   `json:"paramType"`      // 参数类型
	ParamTypeDesc  string   `json:"paramTypeDesc"`  // 参数类型描述
	ParamValueList []string `json:"paramValueList"` // 参数值列表
}
