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

	want := `{"instanceNo":"c_test","appOrderNo":"change_order_1","targetProductNo":"product_1","targetCidrBlocks":[{"cidr":"192.0.2.0/24","count":1,"asn":"","isp":"","projectList":null}],"reason":"replace"}`
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
