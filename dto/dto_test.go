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
