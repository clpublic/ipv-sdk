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
	fmt.Println(string(data))
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	return
}

// 获取开通
func (c *IpvClient) InstanceOpen(params dto.AppInstanceOpenReq) (resp *dto.AppOrderResp, err error) {
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
func (c *IpvClient) InstanceRenew(params dto.AppInstanceRenewReq) (resp *dto.AppOrderResp, err error) {
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
		fmt.Println(c.Endpoint, uri, string(reqData))
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

	ap, err := json.Marshal(aoReq)
	if err != nil {
		slog.Error("ipipv_sdk", "json marshal error", err)
		return nil, err
	}
	req, err := http.NewRequest("POST", c.Endpoint+uri, bytes.NewBuffer(ap))
	if err != nil {
		slog.Error("ipipv_sdk", "Error request:", err)
		return
	}

	// 设置必要的Headers
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
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
	var res dto.Res
	err = json.Unmarshal(data, &res)
	if err != nil {
		slog.Error("ipipv_sdk", "Error unmarshaling response body:", err)
		return
	}
	if res.Code != 200 {
		slog.Error("ipipv_sdk", "Error response:", res)
		return nil, errors.New(res.Msg)
	}

	encrypted, err := base64.StdEncoding.DecodeString(res.Data)
	if err != nil {
		return nil, err
	}
	return cryptos.AesDecryptCBC(encrypted, c.AppKey)

	//return json.Marshal(res.Data)
}
