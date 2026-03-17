package ks_shop_go_sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

var (
	appId      = "ks23423423423423423434"
	appSecret  = "safsfsfsdfsdafsdf"
	signSecret = "sfsdafsdfsdfsdafsdfsdfsdf"
	accToken   = "sadfsadfsdfsdfsdfasfsdaf-FBaVwIfY_fu0ysCJOvAG4Lpy9bxumlR397jYE0qpCAOkzn-4_rsIY6bXN80pygDlL9ZHD1Z3QmHp0I53fcUPW_P27B60S-asfdasdfsdf-61BUXjkUvKAUwAQ"
)

func TestKsShopClient_OpenUserInfoGet(t *testing.T) {
	k := NewKsShopClient(appId, appSecret, signSecret, accToken)
	info, err := k.OpenServiceMarketBuyerServiceInfo(context.Background(), &OpenServiceMarketBuyerServiceInfoRequest{
		BuyerOpenId: "f1b68e8f2a4995efsdfsadfsdf8d14",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}

func TestKsShopClient_OpenShopInfoGet(t *testing.T) {
	k := NewKsShopClient(appId, appSecret, signSecret, accToken)
	info, err := k.OpenShopInfoGet(context.Background(), &OpenShopInfoGetRequest{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(must2Json(info))
}

func TestKsShopClient_OpenDistributionSellerActivityOpenInfo(t *testing.T) {
	k := NewKsShopClient(appId, appSecret, signSecret, accToken)
	info, err := k.OpenDistributionSellerActivityOpenInfo(context.Background(), &OpenDistributionSellerActivityOpenInfoRequest{
		ActivityId: 10064904230,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(must2Json(info))
}

func must2Json(i any) string {
	bytes, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
