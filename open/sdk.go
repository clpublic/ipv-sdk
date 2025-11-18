package sdk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/clpublic/ipv-sdk/cryptos"
	"github.com/clpublic/ipv-sdk/dto"
)

const VERSION = "v2"

const (
	// 获取产品库存
	GetAppInfoUri = "/api/open/app/info/" + VERSION
	// 获取产品库存
	GetProductStockUri = "/api/open/app/product/query/" + VERSION
	//创建和修改主账户
	CreateUserUri = "/api/open/app/user/" + VERSION
	//同步实名
	UserAuthUri = "/api/open/app/userAuth/" + VERSION
	// 获取订单列表
	GetOrderUri = "/api/open/app/order/" + VERSION
	// 获取实列列表
	GetInstanceUri = "/api/open/app/instance/" + VERSION
	// 地域列表
	GetAreaUri = "/api/open/app/area/" + VERSION
	// 开通实例
	InstanceOpenUri = "/api/open/app/instance/open/" + VERSION
	// 续费实例
	InstanceRenewUri = "/api/open/app/instance/renew/" + VERSION
	// 释放实例
	InstanceReleaseUri = "/api/open/app/instance/release/" + VERSION
	//账密提取
	DrawByPwdUri = "/api/open/app/proxy/draw/pwd/" + VERSION
	//获取代理信息
	ProxyInfoUri = "/api/open/app/proxy/info/" + VERSION
	//创建和修改代理用户（子账号）
	CreateProxyUserUri = "/api/open/app/proxy/user/" + VERSION
	// 获取动态代理区域
	GetProductAreaListUri = "/api/open/app/product/area/" + VERSION
	// 添加ip白名单
	AddIpWhiteListUri = "/api/open/app/proxy/addIpWhiteList/" + VERSION
	// 删除ip白名单
	DelIpWhiteListUri = "/api/open/app/proxy/delIpWhiteList/" + VERSION
	//Api提取动态代理
	DrawByApiUri = "/api/open/app/proxy/draw/api/" + VERSION
	// 流量使用记录列表
	ProxyFlowUseLogUri = "/api/open/app/proxy/flow/use/log/" + VERSION
	// 城市列表
	GetCityListUri = "/api/open/app/city/list/" + VERSION

	// 动态代理回收
	DynamicProxyReturnUri = "/api/open/app/proxy/return/" + VERSION

	// 指定ip开通实例
	InstanceOpenAssignIpUri = "/api/open/app/instance/open/assign/ip/" + VERSION
	// 查询指定ip可用情况
	GetAssignIpInfoUri = "/api/open/app/assign/ip/info/" + VERSION

	// 获取订单列表
	GetOrderListUri = "/api/open/app/order/list/" + VERSION
	// 获取实列列表
	GetInstanceListUri       = "/api/open/app/instance/list/" + VERSION
	SetProxyUserFlowLimitUri = "/api/open/app/proxy/user/flow/limit/" + VERSION
	GetProxyUserInfoUri      = "/api/open/app/proxy/user/info/" + VERSION
	InstanceAfterSaleUri     = "/api/open/app/instance/aftersale/" + VERSION

	Encrypt_AES = "AES" //aes cbc模式
)

type IpvClient struct {
	Endpoint string // 请求地址
	AppId    string // 应用ID
	AppKey   []byte // 应用密钥
	Encrypt  string // 加密方式
}

func NewClient(endpoint, appId, appKey, encrypt string) (*IpvClient, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint is empty")
	}
	if appKey == "" {
		return nil, errors.New("appKey is empty")
	}
	if appId == "" {
		return nil, errors.New("appId is empty")
	}
	if encrypt == "" {
		encrypt = Encrypt_AES
	}
	endpoint = strings.TrimSuffix(endpoint, "/")
	return &IpvClient{
		Endpoint: endpoint,
		AppId:    appId,
		AppKey:   []byte(appKey),
		Encrypt:  encrypt,
	}, nil
}

// 获取产品库存
func (c *IpvClient) GetAppInfo() (resp dto.AppInfoResp, err error) {
	var data []byte
	data, err = c.postData(GetAppInfoUri, nil)
	if err != nil {
		return
	}
	//fmt.Println(string(data))
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取产品库存
func (c *IpvClient) GetProductStock(params dto.AppProductSyncReq) (resp []dto.AppProductSyncResp, err error) {
	data, err := c.postData(GetProductStockUri, params)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(data))
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 创建用户
func (c *IpvClient) CreateUser(params dto.AppUserReq) (resp *dto.AppCreateUserResp, err error) {
	data, err := c.postData(CreateUserUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 创建代理用户
func (c *IpvClient) CreateProxyUser(params dto.AppProxyUserReq) (resp *dto.AppProxyUserResp, err error) {
	data, err := c.postData(CreateProxyUserUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 同步实名
func (c *IpvClient) UserAuth(params dto.AppAuthUserReq) (resp *dto.AppAuthUserResp, err error) {
	data, err := c.postData(UserAuthUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 账密提取
func (c *IpvClient) DrawByPwd(params dto.AppDrawByPwdReq) (resp *dto.AppDrawByPwdResp, err error) {
	data, err := c.postData(DrawByPwdUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取订单信息
func (c *IpvClient) GetOrder(params dto.AppGetOrderReq) (resp *dto.AppOrderResp, err error) {
	data, err := c.postData(GetOrderUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		slog.Error("ipipv_sdk", "GetOrder-json.Unmarshal", err)
		return
	}
	return
}

// 获取实例信息
func (c *IpvClient) GetInstance(params dto.AppGetInstanceReq) (resp []dto.AppInstanceResp, err error) {
	data, err := c.postData(GetInstanceUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取地域信息
func (c *IpvClient) GetArea(params dto.AppGetAreaReq) (resp []dto.AppAreaResp, err error) {
	data, err := c.postData(GetAreaUri, params)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(data))
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取城市列表
func (c *IpvClient) GetCityList(params dto.AppCityListReq) (resp []dto.AppCityListResp, err error) {
	data, err := c.postData(GetCityListUri, params)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(data))
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取开通
func (c *IpvClient) InstanceOpen(params dto.AppInstanceOpenReq) (resp *dto.AppInstanceOpenResp, err error) {
	data, err := c.postData(InstanceOpenUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取续费
func (c *IpvClient) InstanceRenew(params dto.AppInstanceRenewReq) (resp *dto.AppInstanceRenewResp, err error) {
	data, err := c.postData(InstanceRenewUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取释放
func (c *IpvClient) InstanceRelease(params dto.AppInstanceReleaseReq) (resp *dto.AppInstanceReleaseResp, err error) {
	data, err := c.postData(InstanceReleaseUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取代理信息
func (c *IpvClient) ProxyInfo(params dto.AppProxyInfoReq) (resp *dto.AppProxyInfoResp, err error) {
	data, err := c.postData(ProxyInfoUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取动态产品地区列表
func (c *IpvClient) ProductAreaList(params dto.AppProductAreaReq) (resp *[]dto.AppProductAreaResp, err error) {
	data, err := c.postData(GetProductAreaListUri, params)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(data))
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 添加ip白名单
func (c *IpvClient) AddIpWhiteList(params dto.AppAddIpWhiteListReq) (resp *dto.AppAddIpWhiteListResp, err error) {
	data, err := c.postData(AddIpWhiteListUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 删除ip白名单
func (c *IpvClient) DelIpWhiteList(params dto.AppDelIpWhiteListReq) (resp *dto.AppDelIpWhiteListResp, err error) {
	data, err := c.postData(DelIpWhiteListUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// api提取动态代理
func (c *IpvClient) DrawByApi(params dto.AppDrawByApiReq) (resp *dto.AppDrawByApiResp, err error) {
	data, err := c.postData(DrawByApiUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 流量使用记录列表
func (c *IpvClient) ProxyFlowUseLog(params dto.AppFlowUseLogReq) (resp *dto.AppFlowUseLogResp, err error) {
	data, err := c.postData(ProxyFlowUseLogUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 动态代理回收
func (c *IpvClient) DynamicProxyReturn(params dto.AppProxyReturnReq) (resp *dto.AppProxyReturnResp, err error) {
	data, err := c.postData(DynamicProxyReturnUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 指定ip开通静态实例
func (c *IpvClient) InstanceOpenAssignIp(params dto.AppAssignIpInstanceOpenReq) (resp *dto.AppInstanceOpenResp, err error) {
	data, err := c.postData(InstanceOpenAssignIpUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

func (c *IpvClient) GetAssignIpInfo(params dto.AppGetAssignIpInfoReq) (resp *dto.AppGetAssignIpInfoResp, err error) {
	data, err := c.postData(GetAssignIpInfoUri, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取订单信息
func (c *IpvClient) GetOrderList(params dto.AppGetOrderListReq) (resp *dto.AppOrderListResp, err error) {
	data, err := c.postData(GetOrderListUri, params)
	if err != nil {
		return nil, fmt.Errorf("%s %w", GetOrderListUri, err)
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		slog.Error("ipipv_sdk", "GetOrderList-json.Unmarshal", err)
		return
	}
	return
}

// 获取实例信息
func (c *IpvClient) GetInstanceList(params dto.AppGetInstanceListReq) (resp *dto.AppInstanceListResp, err error) {
	data, err := c.postData(GetInstanceListUri, params)
	if err != nil {
		return resp, fmt.Errorf("%s %w", GetInstanceListUri, err)
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		slog.Error("ipipv_sdk", "GetInstanceList-json.Unmarshal", err)
		return
	}
	return
}

// 设置子账户流量上限
func (c *IpvClient) SetProxyUserFlowLimit(params dto.AppSetProxyUserFlowLimitReq) (err error) {
	_, err = c.postData(SetProxyUserFlowLimitUri, params)
	if err != nil {
		return fmt.Errorf("%s %w", SetProxyUserFlowLimitUri, err)
	}
	return
}

// 获取子账户信息
func (c *IpvClient) GetProxyUserInfo(params dto.AppProxyUserInfoReq) (resp *dto.AppProxyUserInfoResp, err error) {
	data, err := c.postData(GetProxyUserInfoUri, params)
	if err != nil {
		return resp, fmt.Errorf("%s %w", GetProxyUserInfoUri, err)
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		slog.Error("ipipv_sdk", "GetProxyUserInfo-json.Unmarshal", err)
		return
	}
	return
}

// InstanceAfterSale
func (c *IpvClient) InstanceAfterSale(params dto.AppInstanceAfterSaleReleaseReq) (resp *dto.AppInstanceAfterSaleReleaseResp, err error) {
	data, err := c.postData(InstanceAfterSaleUri, params)
	if err != nil {
		return resp, fmt.Errorf("%s %w", InstanceAfterSaleUri, err)
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		slog.Error("ipipv_sdk", "InstanceAfterSaleUri-json.Unmarshal", err, "resp", string(data))
		return
	}
	return
}

func (c *IpvClient) postData(uri string, params any) (resData []byte, err error) {
	aoReq := dto.AppOpenReq{
		Version: VERSION,
		Encrypt: c.Encrypt,
		AppKey:  c.AppId,
	}
	if params != nil {
		reqData, err := json.Marshal(params)
		if err != nil {
			slog.Error("ipipv_sdk", "json marshal error", err)
			return nil, err
		}
		slog.Debug("[ipipv_sdk]", "param", string(reqData))
		//fmt.Println(c.Endpoint, uri, string(reqData))
		if c.Encrypt == "" {
			c.Encrypt = Encrypt_AES
		}
		var ens []byte
		if c.Encrypt == Encrypt_AES {
			ens, err = cryptos.AesEncryptCBC(reqData, c.AppKey)
			if err != nil {
				slog.Error("ipipv_sdk", "AesEncryptCBC err", err)
				return nil, err
			}
		}
		aoReq.Params = base64.StdEncoding.EncodeToString(ens)
	}
	aoReq.ReqId = fmt.Sprintf("reqId_%d", time.Now().UnixNano())
	//fmt.Println("加密之后Params", aoReq.Params)

	ap, err := json.Marshal(aoReq)
	if err != nil {
		slog.Error("ipipv_sdk", "json marshal error", err)
		return nil, err
	}

	slog.Debug("[ipipv_sdk]", "req", string(ap))

	//fmt.Println(string(ap))
	req, err := http.NewRequest("POST", c.Endpoint+uri, bytes.NewBuffer(ap))
	if err != nil {
		slog.Error("ipipv_sdk", "Error request:", err)
		return
	}

	slog.Debug("[ipipv_sdk]", "endpoint", c.Endpoint+uri)
	// fmt.Println("url", c.Endpoint+uri)
	// fmt.Println("req", string(ap))
	// fmt.Println("err", err)
	// 设置必要的Headers
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("ipipv_sdk", "Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= http.StatusBadRequest {
		slog.Error("ipipv_sdk", "Error response:", resp)
		return
	}

	// 读取响应体
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("ipipv_sdk", "Error reading response body:", err)
		return
	}
	slog.Debug("[ipipv_sdk]", "result", string(data))
	//fmt.Println(string(data))
	var res dto.Res
	err = json.Unmarshal(data, &res)
	if err != nil {
		slog.Error("ipipv_sdk", "Error unmarshaling response body:", err)
		return
	}
	if res.Code != 200 {
		slog.Error("ipipv_sdk", "Error response:", res)
		return nil, fmt.Errorf("[code]=%d,[msg]=%s", res.Code, res.Msg)
	}

	if res.Data == "" {
		return nil, nil
	}
	encrypted, err := base64.StdEncoding.DecodeString(res.Data)
	if err != nil {
		return nil, err
	}
	return cryptos.AesDecryptCBC(encrypted, c.AppKey)

	//return json.Marshal(res.Data)
}

func (c *IpvClient) PostData(uri string, params any) (resData []byte, err error) {
	return c.postData(uri, params)
}
