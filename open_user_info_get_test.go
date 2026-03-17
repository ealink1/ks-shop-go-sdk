package ks_shop_go_sdk

import (
	"context"
	"testing"
)

func TestKsShopClient_OpenUserInfoGet(t *testing.T) {
	accToken := "ChFvYXV0aC5hY2Nlc3NUb2tlbhJgT3lr43C6MC0drrLbM6y-rSS9c8ccuHBpc4SlFLQEvuNV5vvY9wR4aG6tJe6EhaPppPMFzx5VqgbEQmecNILFx2EqAU9xZI6WWzP4DlWKD-7jQoBzfQvkJZVbXRLZIfgHGhJ0lNpl9NtFt7jM9LzyFTTPFEciIDPqePkz_53U1h5ZbQ--wszCdZfYDOJ0p_TcJnXu_9zCKAUwAQ"
	k := NewKsShopClient("ks699183844582124027", "LzOy5TKflTe14M9pw4GZxg", "b312340697d26ebb2cca5bfbe8be45fb", accToken)
	info, err := k.OpenServiceMarketBuyerServiceInfo(context.Background(), &OpenServiceMarketBuyerServiceInfoRequest{
		BuyerOpenId: "f1b68e8f2a4995ed14bf5725f63e8d14",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}
