package dto

import (
	"encoding/json"
	"testing"
)

func TestAppChangeProxyReqMarshal(t *testing.T) {
	req := AppChangeProxyReq{
		InstanceNo:      "c_test",
		AppOrderNo:      "change_order_1",
		TargetProductNo: "product_1",
		TargetCIDRBlocks: []CIDRBlock{
			{
				CIDR:  "192.0.2.0/24",
				Count: 1,
			},
		},
		Reason: "replace",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"instanceNo":"c_test","appOrderNo":"change_order_1","targetProductNo":"product_1","targetCidrBlocks":[{"cidr":"192.0.2.0/24","count":1,"asn":"","isp":"","netType":0,"asnType":0,"costPrice":null,"riskLevel":"","fraudScore":null,"totalInventory":null,"projectList":null}],"reason":"replace"}`
	if string(data) != want {
		t.Fatalf("unexpected json: %s", data)
	}
}

func TestAppInstanceOpenReqMarshalExtraParams(t *testing.T) {
	req := AppInstanceOpenReq{
		AppOrderNo:         "order_1",
		UserId:             "user_1",
		ExcludeUserHistory: 1,
		ExcludeIp:          "1.2.3.4,5.6.7.8",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"appOrderNo":"order_1","userId":"user_1","excludeUserHistory":1,"excludeIp":"1.2.3.4,5.6.7.8","params":null}`
	if string(data) != want {
		t.Fatalf("unexpected json: %s", data)
	}
}

func TestAppProductSyncRespUnmarshalAsnType(t *testing.T) {
	var resp AppProductSyncResp
	if err := json.Unmarshal([]byte(`{"productNo":"product_1","asnType":3}`), &resp); err != nil {
		t.Fatal(err)
	}

	if resp.AsnType != 3 {
		t.Fatalf("unexpected asnType: %d", resp.AsnType)
	}
}

func TestAppProductSyncRespUnmarshalCIDRBlockFields(t *testing.T) {
	var resp AppProductSyncResp
	err := json.Unmarshal([]byte(`{
		"productNo": "product_1",
		"totalInventory": 100,
		"cidrBlocks": [{
			"cidr": "192.0.2.0/24",
			"netType": 1,
			"asnType": 2,
			"costPrice": "12.30",
			"riskLevel": "low",
			"fraudScore": 5,
			"totalInventory": 256
		}, {
			"cidr": "198.51.100.0/24",
			"costPrice": null,
			"fraudScore": null,
			"totalInventory": null
		}]
	}`), &resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.TotalInventory == nil || *resp.TotalInventory != 100 {
		t.Fatalf("unexpected totalInventory: %v", resp.TotalInventory)
	}
	if len(resp.CIDRBlocks) != 2 {
		t.Fatalf("unexpected cidrBlocks length: %d", len(resp.CIDRBlocks))
	}
	item := resp.CIDRBlocks[0]
	if item.NetType != 1 || item.AsnType != 2 || item.RiskLevel != "low" {
		t.Fatalf("unexpected cidr block fields: %+v", item)
	}
	if item.CostPrice == nil || *item.CostPrice != "12.30" {
		t.Fatalf("unexpected costPrice: %v", item.CostPrice)
	}
	if item.FraudScore == nil || *item.FraudScore != 5 {
		t.Fatalf("unexpected fraudScore: %v", item.FraudScore)
	}
	if item.TotalInventory == nil || *item.TotalInventory != 256 {
		t.Fatalf("unexpected cidr totalInventory: %v", item.TotalInventory)
	}

	unknown := resp.CIDRBlocks[1]
	if unknown.CostPrice != nil || unknown.FraudScore != nil || unknown.TotalInventory != nil {
		t.Fatalf("unknown values should stay nil: %+v", unknown)
	}
}
