---
description: 
---
1、每个接口、方法、常量、函数、结构体、结构体字段都要写好注释，包括方法内的具体逻辑
 示例：
```go
// OpenItemListGetApi 获取商品列表
	OpenItemListGetApi = "/open/item/list/get"
```


2、字段类型不要使用json.RawMessage，我要知道每个字段的详细信息。
3、结构体放在文件后面。
    示例：
```go
func (k *KsShopClient) OpenItemListGet(ctx context.Context, reqData OpenItemListGetRequest) (*OpenItemListGetResponse, error) {
    method := reqData.Method
    // 其他逻辑
}

// OpenItemListGetParam 获取商品列表参数
type OpenItemListGetParam struct {
    AppKey     string   `json:"appkey"`
    GoodsCode  []string `json:"goodsCode"`
}
、、、

4、接口完成后需要补充到README.md中。