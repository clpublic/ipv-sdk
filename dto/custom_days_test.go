package dto

import (
	"encoding/json"
	"testing"
)

func TestCustomDaysOpenRequestsMarshal(t *testing.T) {
	one, fifteen, zero, negative, sixteen := int32(1), int32(15), int32(0), int32(-1), int32(16)
	for _, tc := range []struct {
		name string
		days *int32
	}{
		{"omitted", nil},
		{"one", &one},
		{"fifteen", &fifteen},
		{"zero", &zero},
		{"negative", &negative},
		{"sixteen", &sixteen},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, req := range []struct {
				name   string
				value  any
				nested bool
			}{
				{"normal", AppInstanceOpenReq{Params: []OpenParam{{ProductNo: "product_1", Days: tc.days, Unit: 3, Duration: 2, CycleTimes: 4}}}, true},
				{"assign_ip", AppAssignIpInstanceOpenReq{ProductNo: "product_1", AssignIp: "192.0.2.1", Days: tc.days, CycleTimes: 4}, false},
			} {
				t.Run(req.name, func(t *testing.T) {
					data, err := json.Marshal(req.value)
					if err != nil {
						t.Fatal(err)
					}
					var fields map[string]json.RawMessage
					if err := json.Unmarshal(data, &fields); err != nil {
						t.Fatal(err)
					}
					if req.nested {
						var params []map[string]json.RawMessage
						if err := json.Unmarshal(fields["params"], &params); err != nil {
							t.Fatal(err)
						}
						if len(params) != 1 {
							t.Fatalf("unexpected params length: %d", len(params))
						}
						fields = params[0]
						if string(fields["unit"]) != "3" || string(fields["duration"]) != "2" {
							t.Fatalf("legacy unit/duration changed: %s", data)
						}
					}
					if string(fields["cycleTimes"]) != "4" {
						t.Fatalf("legacy cycleTimes changed: %s", data)
					}
					days, present := fields["days"]
					if tc.days == nil {
						if present {
							t.Fatalf("unset days must be omitted: %s", data)
						}
						return
					}
					if !present {
						t.Fatalf("explicit days must be present: %s", data)
					}
					var got int32
					if err := json.Unmarshal(days, &got); err != nil {
						t.Fatal(err)
					}
					if got != *tc.days {
						t.Fatalf("days = %d, want %d", got, *tc.days)
					}
				})
			}
		})
	}
}

func TestAppProductSyncRespUnmarshalAllowCustomDays(t *testing.T) {
	for _, tc := range []struct {
		name string
		data string
		want int
	}{
		{"allowed", `{"allowCustomDays":1}`, 1},
		{"disallowed", `{"allowCustomDays":-1}`, -1},
		{"legacy", `{"productNo":"product_1"}`, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var resp AppProductSyncResp
			if err := json.Unmarshal([]byte(tc.data), &resp); err != nil {
				t.Fatal(err)
			}
			if resp.AllowCustomDays != tc.want {
				t.Fatalf("allowCustomDays = %d, want %d", resp.AllowCustomDays, tc.want)
			}
		})
	}
}
