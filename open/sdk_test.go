package sdk

import (
	"testing"
	"time"

	"github.com/clpublic/ipv-sdk/dto"
)

func getClient() *IpvClient {
	//根据实际需要 以下参数需要修改配置
	//BaseUrl   string = "https://api.ipipv.com"

	Endpoint := "http://192.168.0.78:10888"
	appid := "AppId"
	appKey := `xxxxx`
	client, err := NewClient(Endpoint, appid, appKey, Encrypt_AES)
	if err != nil {
		panic(err)
	}
	return client
}

func TestGetProductStock(t *testing.T) {
	//ps, err := GetProductStock(dto.AppProductSyncReq{ProxyType: 1})
	ps, err := getClient().GetProductStock(dto.AppProductSyncReq{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestCreateUser(t *testing.T) {
	ps, err := getClient().CreateUser(dto.AppUserReq{AuthName: "", No: "", AppUsername: "ad67a4bb7e1e4486bba8bc77027295b8", Password: ""})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestCreateUser2(t *testing.T) {
	ps, err := getClient().CreateUser(dto.AppUserReq{AuthName: "", No: "", AppUsername: "aaaabbb", Password: "bbbbbb"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestUserAuth(t *testing.T) {
	ps, err := getClient().UserAuth(dto.AppAuthUserReq{AuthName: "aaaa", No: "ccc", Username: "aaaabbb"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestGetOrder(t *testing.T) {
	orders := []string{"C20240510134623046203"}
	ps, err := getClient().GetOrder(dto.AppGetOrderReq{orders})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestGetInstance(t *testing.T) {
	instances := []string{"aaaa", "c_gzmtux84eu8ipkq"}
	ps, err := getClient().GetInstance(dto.AppGetInstanceReq{instances})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestGetArea(t *testing.T) {
	ps, err := getClient().GetArea(dto.AppGetAreaReq{})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func getOrderNo() string {
	return "TEST" + time.Now().Format("20060102150405")
}

func TestInstanceOpen(t *testing.T) {
	params := []dto.OpenParam{dto.OpenParam{
		ProductNo:    "aws_light_206", // tx_166 ip90_1189 aws_light_205  jg_226=随机地区  jg_227=安徽省芜湖市 ipidea_590=纽约 ipidea_d_283=美国
		OrderNo:      getOrderNo(),
		Count:        1,
		Duration:     1,
		Unit:         3,
		Renew:        false,
		Username:     "",
		CountryCode:  "",
		CityCode:     "",
		ProxyType:    0,
		SupplierCode: "",
		IspType:      0,
		ExtBandWidth: 0,
		AppUsername:  "",
	},
	}
	o := dto.AppInstanceOpenReq{
		Params: params,
	}
	ps, err := getClient().InstanceOpen(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestInstanceRenew(t *testing.T) {
	// ipidea => c_gzm9tradpjmqwg4
	o := dto.AppInstanceRenewReq{
		OrderNo:   getOrderNo(),
		Instances: []dto.Instance{dto.Instance{Id: "c_gzqk6ykpa2qb59o"}}, // tx=>c_gzmtux84eu8ipkq  ip90=>c_gzmk5maz55cv536  aws=>c_gzmk7asab4wy6vp jg => c_gzms93ztk8jnh0r
	}
	ps, err := getClient().InstanceRenew(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestInstanceRelease(t *testing.T) {
	o := dto.AppInstanceReleaseReq{
		OrderNo:   getOrderNo(),
		Instances: []string{"c_gzraaycyuhps5pg"}, // c_gzmycfycqejct4m  ip90=>c_gzmk5maz55cv536 aws=>c_gzmk7asab4wy6vp
	}
	ps, err := getClient().InstanceRelease(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ps)
}

func TestDrawByPwd(t *testing.T) {
	o := dto.AppDrawByPwdReq{
		AppUsername:  "app6_user1",
		AddressCode:  "USA",
		SessTime:     "",
		Num:          1,
		ProxyType:    104,
		MaxFlowLimit: 10,
	}
	ps, err := getClient().DrawByPwd(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}

func TestProxyInfo(t *testing.T) {
	o := dto.AppProxyInfoReq{
		Username:  "app6_user",
		ProxyType: 104,
	}
	ps, err := getClient().ProxyInfo(o)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", ps)
}
