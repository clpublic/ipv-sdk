package dto

import (
	"time"
)

// 请求统一为密文传输
type AppOpenReq struct {
	ReqId   string `json:"reqId"`   // 请求id，每次生成，如果失败重复请求，保持不变
	Version string `json:"version"` // 版本 v2
	Encrypt string `json:"encrypt"` // 加密方式 aes,rsa（版本2以上提供，默认aes,老接口rsa）aes cbc模式
	AppKey  string `json:"appKey"`  // appKey 渠道商号
	Params  string `json:"params"`  // 根据加密方式密文 转base64
}

// 统一返回
type Res struct {
	ReqId string `json:"reqId"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  string `json:"data"` //密文转base64
}

// 获取产品列表请求
type AppProductSyncReq struct {
	ProxyType    []int  `json:"proxyType"`    //代理类型 可选  101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	ProductNo    string `json:"productNo"`    //产品编号 可选
	CountryCode  string `json:"countryCode"`  //国家代码 可选
	CityCode     string `json:"cityCode"`     //城市代码 可选
	SupplierCode string `json:"supplierCode"` //供应商代码 可选
	Unit         int8   `json:"unit"`         //时长单位 可选 1=天 2=周(7天) 3=月(自然月) 4=年(自然年365，366) 10=无限制
	IspType      int    `json:"ispType"`      //isp类型 可选 1=单isp 2=双isp
	NetType      int    `json:"netType"`      //网络类型 1=原生 2=广播
	Duration     int32  `json:"duration"`     //相对于时长单位的最小购买时长 可选
}

type AppInfoResp struct {
	AppName     string `json:"appName"`          //应用名
	Coin        string `json:"coin"`             //余额
	UseBridge   int8   `json:"useBridge"`        //使用桥 1 不使用 2使用
	CallbackUrl string `json:"callbackUrl"`      //回调地址
	Credit      string `json:"credit,omitempty"` //授信额度
	Status      int    `json:"status"`           // 1正常 -1禁用
}

// 获取产品列表返回 （明文返回）
type AppProductSyncResp struct {
	ProductNo            string             `json:"productNo"`            //必要，产品Id 保持唯一
	ProductName          string             `json:"productName"`          //可选,商品名 后续该字段逐步废弃
	ProxyType            int16              `json:"proxyType"`            //必要, 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	UseType              string             `json:"useType"`              //必要, ,分割  1=账密 2=白名单 3=uuid
	Protocol             string             `json:"protocol"`             //必要,1=socks5 2=http 3=https 4=ssh
	UseLimit             int8               `json:"useLimit"`             //必要,1=出口ip国外 2=出口ip国内 3=无限制
	SellLimit            int8               `json:"sellLimit"`            //必要,1=大陆可售 2=海外可售 3=无限制
	AreaCode             string             `json:"areaCode"`             //区域code
	CountryCode          string             `json:"countryCode"`          //必要,国家代码 3位 iso标准
	StateCode            string             `json:"stateCode"`            //必要,州省代码 6位
	CityCode             string             `json:"cityCode"`             //必要,城市代码 向我方提取
	Detail               string             `json:"detail"`               //商品描述
	CostPrice            string             `json:"costPrice"`            //必要 价格
	Inventory            int                `json:"inventory"`            //必要 库存
	IpType               int                `json:"ipType"`               //ip类型 1=ipv4 2=ipv6 3=随机 默认1
	IspType              int                `json:"ispType"`              //ispType 1=单isp 2=双isp
	NetType              int                `json:"netType"`              //网络类型 1=原生 2=广播
	Duration             int                `json:"duration"`             //必要 时长 0无限制
	Unit                 int                `json:"unit"`                 //单位 1=天 2=周(7天) 3=月(自然月) 4=年(自然年365，366)
	BandWidth            int                `json:"bandWidth"`            //带宽|流量时必要 单位 MB
	BandWidthType        int                `json:"bandWidthType"`        //带宽类型 1=独享 2=共享 3=流量包
	BandWidthPrice       string             `json:"bandWidthPrice"`       //额外带宽价格
	MaxBandWidth         int                `json:"maxBandWidth"`         //可设置最大带宽
	Flow                 int                `json:"flow"`                 //流量包
	Cpu                  int                `json:"cpu"`                  //cpu数
	Memory               float64            `json:"memory"`               //内存容量
	Enable               int8               `json:"enable"`               //是否可以购买 1可以
	SupplierCode         string             `json:"supplierCode"`         //供应商代码 后续该字段逐步废弃
	CIDRBlocks           []CIDRBlock        `json:"cidrBlocks"`           //支持网段及数量 新增 2024/06/27
	DrawType             int                `json:"drawType"`             //代理提取方式 0=不需要提取(静态代理) 1=白名单提取(api) 2=账密提取 3=都支持 默认为0 新增于2024/08/12
	RefundDuration       int                `json:"refundDuration"`       //退款时效 单位秒 0=不支持退款 大于0表示从创建订单之后多少秒内可以退款 默认为0 新增于2024/08/12
	IpCount              int                `json:"ipCount"`              //ip数量 动态代理按照ip数量购买 该字段大于0 默认为0 新增于2024/08/26
	IpDuration           int                `json:"ipDuration"`           //ip时长 动态代理按照ip数量购买 单位分钟 该字段大于0 默认为0 新增于2024/08/26
	AssignIp             int                `json:"assignIp"`             //是否支持指定ip购买 1=是 -1=否 默认为-1 新增于2024/10/10
	ParentNo             string             `json:"parentNo"`             //父产品编号 新增于2024/10/15
	CIDRStatus           int                `json:"cidrStatus"`           //ip段支持状态 1=支持 -1=不支持 默认为-1 新增于2024/10/15
	OneDay               int                `json:"oneDay"`               //是否支持1天的测试 1=是 -1=否 新增于2024/10/17
	OfflineCIDRBlocks    []OfflineCIDRBlock `json:"offlineCidrBlocks"`    //最近1个月下架的网段 新增 2025/07/07
	ProxyEverytimeChange int                `json:"proxyEverytimeChange"` //动态代理账密提取 是否支持每次更换代理  1=是 -1=否 默认为否 新增于2025/08/14
	ProxyGlobalRandom    int                `json:"proxyGlobalRandom"`    //动态代理提取 是否支持全球混播 1=是 -1=否 默认为否 新增于2025/08/14
	ApiDrawGlobalRandom  int                `json:"apiDrawGlobalRandom"`  //动态代理Api提取是否支持全球混播 1=是 -1=否 默认为否 新增于2025/09/5
	IpWhiteList          int                `json:"ipWhiteList"`          //动态代理是否支持IP白名单功能 1=是 -1=否 默认为否 新增于2025/09/5
	PwdDrawProxyUser     int                `json:"pwdDrawProxyUser"`     //动态代理账密提取是否支持子账号 1=是 -1=否 默认为否 新增于2025/09/5
	ProxyUserFlowLimit   int                `json:"proxyUserFlowLimit"`   //动态代理子账号是否支持流量上限管理 1=是 -1=否 默认为否 新增于2025/09/5
	FlowUseLog           int                `json:"flowUseLog"`           //动态代理是否支持流量明细查询 1=是 -1=否 默认为否 新增于2025/09/5
	PwdDrawSessionRange  string             `json:"pwdDrawSessionRange"`  //动态代理账密流量提取持续时间范围 单位分钟 新增于2025/09/5
	FlowConversionBase   int                `json:"flowConversionBase"`   //动态代理流量单位转化基准 1000 或者 1024 0表示未知或不支持  新增于2025/09/5
	ProjectList          []ProjectItem      `json:"projectList"`          // ProjectList
	ProductType          int                `json:"productType"`          // 2=share
	ResetPassword        int                `json:"resetPassword"`        // 代理是否支持重置密码 1=是 -1=否 默认为否
}

// 下架网段 新增 2025/0707
type OfflineCIDRBlock struct {
	CIDR        string `json:"cidr"`        // 网段 192.168.0.0/24 172.16.0.0/16 10.0.0.0/8
	OfflineTime string `json:"offlineTime"` // 网段下架时间 格式 2024-06-27 12:00:00
}

// 新增 2024/06/27
// 网段
type CIDRBlock struct {
	CIDR        string        `json:"cidr"`        // 网段 192.168.0.0/24 172.16.0.0/16 10.0.0.0/8
	Count       int           `json:"count"`       // 网段ip数量 购买代理的时候如果传了网段 该字段必传
	Asn         string        `json:"asn"`         // 该网段属于哪个asn 购买代理的时候可选 新增 2024/09/21
	Isp         string        `json:"isp"`         // 该网段属于哪个网络提供商 购买代理的时候可选 新增 2024/10/10
	ProjectList []ProjectItem `json:"projectList"` // projectList
}

// 创建或修改主账号请求
type AppUserReq struct {
	AppUsername string `json:"appUsername"` //渠道商主账号 该渠道商唯一 不支持修改
	Password    string `json:"password"`    //主账号密码(不传随机生成)
	Phone       string `json:"phone"`       //主账号手机号
	Email       string `json:"email"`       //主账号邮箱
	AuthType    int8   `json:"authType"`    //认证类型 1=未实名 2=个人实名 3=企业实名
	AuthName    string `json:"authName"`    //主账号实名认证的真实名字或者企业名
	No          string `json:"no"`          //主账号实名认证的实名证件号码或者企业营业执照号码
	VSP         uint8  `json:"vsp"`         //vsp
	Status      int8   `json:"status"`      //状态 1=正常 2=禁用
}

// 创建用户返回
type AppCreateUserResp struct {
	AppUsername string `json:"appUsername"` //渠道商主账号
	Username    string `json:"username"`    //平台主账号
	Password    string `json:"password"`    //主账号密码
	Status      int8   `json:"status"`      //用户状态 1=正常 2=禁用
	AuthStatus  int8   `json:"authStatus"`  //认证状态 1=未实名 2=个人实名 3=企业实名
}

// 创建或修改代理用户请求
type AppProxyUserReq struct {
	AppUsername     string `json:"appUsername"`     //渠道商子账号 该渠道商唯一 (不传随机生成) 不支持修改
	Password        string `json:"password"`        //密码(不传随机生成)
	LimitFlow       int    `json:"limitFlow"`       //动态流量上限
	MainUsername    string `json:"mainUsername"`    //平台主账号  选填 平台主账号和渠道商主账号两个必填一个
	AppMainUsername string `json:"appMainUsername"` //渠道商主账号 选填 平台主账号和渠道商主账号两个必填一个
	Remark          string `json:"remark"`          //备注
	Status          int8   `json:"status"`          //状态 1=正常 2=禁用
}

// 创建或修改代理用户返回
type AppProxyUserResp struct {
	AppUsername string `json:"appUsername"` //渠道商子账号
	Username    string `json:"username"`    //平台子账号
	Password    string `json:"password"`    //子账号密码
	Status      int8   `json:"status"`      //用户状态 1=正常 2=禁用
	AuthStatus  int8   `json:"authStatus"`  //认证状态 1=未实名 2=个人实名 3=企业实名
}

// 同步实名请求
type AppAuthUserReq struct {
	Username    string `json:"username"`    //平台主账号 选填 平台主账号和渠道商主账号两个必填一个
	AppUsername string `json:"appUsername"` //渠道商主账号 选填 平台主账号和渠道商主账号两个必填一个
	AuthType    int8   `json:"authType"`    //认证类型 1 未实名 2 个人实名 3 企业实名
	AuthName    string `json:"authName"`    //真实姓名或者企业名
	No          string `json:"no"`          //实名证件号码或者企业营业执照号码
	VSP         string `json:"vsp"`         //vsp
}

// 同步实名返回
type AppAuthUserResp struct {
	Username   string `json:"username"`   //平台账号
	AuthStatus int    `json:"authStatus"` //认证状态 1=未实名 2=个人实名 3=企业实名
}

// 获取订单列表请求
type AppGetOrderReq struct {
	OrderNo    string `json:"orderNo"`    // 平台订单编号
	AppOrderNo string `json:"appOrderNo"` // 渠道商（购买方）订单编号
	Page       int    `json:"page"`       // 页码  默认1
	PageSize   int    `json:"pageSize"`   // 每页显示数量  默认10 最大100
}

// 订单返回
type AppOrderResp struct {
	OrderNo    string            `json:"orderNo"`    //平台订单号
	AppOrderNo string            `json:"appOrderNo"` //渠道商（购买方）订单编号
	Type       int8              `json:"type"`       //订单类型 1=新建 2=续费 3=释放
	Status     int8              `json:"status"`     //订单状态 1=待处理 2=处理中 3=处理成功 4=处理失败 5=部分完成
	Count      int               `json:"count"`      //购买数量
	Amount     string            `json:"amount"`     //总价
	Refund     int               `json:"refund"`     //是否有退费 1存在退费
	Page       int               `json:"page"`       //页码 原样返回
	PageSize   int               `json:"pageSize"`   //每页显示数量  原样返回
	Total      int64             `json:"total"`      //订单对应实例总数量
	Instances  []AppInstanceResp `json:"instances"`  //订单对应实例列表
}

// 获取实列列表请求
type AppGetInstanceReq struct {
	Instances []string `json:"Instances"` //平台实例编号
}

// 实例返回
type AppInstanceResp struct {
	InstanceNo  string    `json:"instanceNo"`  //平台实例编号（渠道商续费和释放操作使用该编号）
	ProxyType   uint      `json:"proxyType"`   //代理类型 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	Protocol    string    `json:"protocol" `   //协议类型 多个用英文逗号分隔 1=socks5 2=http 3=https 4=ssh
	Ip          string    `json:"ip"`          //代理地址 用户实际代理访问使用
	Port        uint      `json:"port"`        //代理端口
	RegionId    string    `json:"regionId"`    //区域地址
	CountryCode string    `json:"countryCode"` //国家代码
	CityCode    string    `json:"cityCode"`    //城市代码
	UseType     string    `json:"useType"`     //使用方式 多个用英文逗号分隔  1=账密 2=ip白名单 3=uuid（uuid写password内）
	Username    string    `json:"username"`    //账户名或uuid 动态为平台主账号
	Pwd         string    `json:"pwd"`         //密码
	OrderNo     string    `json:"orderNo"`     //创建该实例的平台订单号
	UserExpired int64     `json:"userExpired"` //到期时间
	FlowTotal   float64   `json:"flowTotal"`   //总流量
	FlowBalance float64   `json:"flowBalance"` //剩余流量
	Status      int8      `json:"status"`      //1=待创建 2=创建中 3=运行中 6=已停止 10=关闭 11=释放
	Renew       int8      `json:"renew"`       //1 自动续费
	Bridges     []string  `json:"bridges"`     //桥地址列表
	OpenAt      time.Time `json:"openAt"`      //开通时间
	RenewAt     time.Time `json:"renewAt"`     //最后成功续费时间
	ReleaseAt   time.Time `json:"releaseAt"`   //释放成功时间
	ProductNo   string    `json:"productNo"`   //产品编号
	ExtendIp    string    `json:"extendIp"`    //扩展地址 仅供展示，部分产品该字段有值 2024-11-20 新增
}

// 同步地域请求
type AppGetAreaReq struct {
	Codes []string `json:"codes"` //获取地域代码对应列表，为null获取全部
}

// 同步地域返回
type AppAreaResp struct {
	Code     string        `json:"code"`               //地域代码
	Name     string        `json:"name"`               //地域英文名称
	Cname    string        `json:"cname"`              //地域中文名
	Children []AppAreaResp `json:"children,omitempty"` //下级地域
}

// 开通代理资源请求
type AppInstanceOpenReq struct {
	AppOrderNo string      `json:"appOrderNo"` //购买者订单号(渠道商订单号)
	Params     []OpenParam `json:"params"`     //购买代理产品列表
}

type OpenParam struct {
	ProductNo    string      `json:"productNo"`    //产品编号（如果存在，后面7项无意义）
	ProxyType    uint16      `json:"proxyType"`    //代理类型（匹配产品用，如果产品编号有值忽略该项） 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	CountryCode  string      `json:"countryCode"`  //国家代码（匹配产品用，如果产品编号有值忽略该项）
	CityCode     string      `json:"cityCode"`     //城市代码（匹配产品用，如果产品编号有值忽略该项）
	SupplierCode string      `json:"supplierCode"` //供应商代码（匹配产品用，如果产品编号有值忽略该项）（可为null,随机分配）
	Unit         int8        `json:"unit"`         //单位（匹配产品用，如果产品编号有值忽略该项） 1=天 2=周(7天) 3=月(自然月) 4=年(自然年365，366) 10=无限制
	IspType      int         `json:"ispType"`      //isp类型（匹配产品用，如果产品编号有值忽略该项） 1=单isp 2=双isp
	Count        int         `json:"count"`        //购买数量 （实例个数）静态必填 默认1 一次最大20
	Duration     int32       `json:"duration"`     //相对时长 必要 默认1 为Unit的时长 此处不是指的绝对时长 指的是相对x个时间单位(unit)的时长 2024-9-21改为产品定义的时长单位 （匹配产品用，如果产品编号有值忽略该项）  之前对接的定义不变
	Renew        bool        `json:"renew"`        //是否续费 1续费 默认0
	ExtBandWidth int32       `json:"extBandWidth"` //额外增加带宽 单位Mbps
	AppUsername  string      `json:"appUsername"`  //渠道商主账号，开通动态代理的时候必填(必须在平台上注册过)
	Flow         int         `json:"flow"`         //动态流量 最大102400MB 动态流量必填 单位MB
	UseBridge    uint8       `json:"useBridge"`    //是否使用桥 0=随app设置 1=不使用桥 2=使用桥 默认0
	CIDRBlocks   []CIDRBlock `json:"cidrBlocks"`   //静态购买所在网段及数量（产品有的才支持） 2024/06/27新增
	ProjectId    string      `json:"projectId"`    //购买项目id,保留字段，后续支持
	CycleTimes   int32       `json:"cycleTimes"`   //购买时长周期数，此字段对有时长的产品有意义，默认1 表示产品的unit * duration 如果该字段大于0 duration字段不作为购买时长使用 2024/09/21新增
}

// 开通代理资源返回
type AppInstanceOpenResp struct {
	OrderNo    string `json:"orderNo"`    //平台订单号
	AppOrderNo string `json:"appOrderNo"` //购买者订单号(渠道商订单号) 原样返回
	Amount     string `json:"amount"`     //花费金额
}

// 续费代理资源请求
type AppInstanceRenewReq struct {
	AppOrderNo string     `json:"appOrderNo"` //购买者订单号(渠道商订单号)
	Instances  []Instance `json:"instances"`  //实例列表
}

type Instance struct {
	InstanceNo string `json:"instanceNo"` //平台实例编号
	Duration   int32  `json:"duration"`   //可选 时长 续费的时候默认是购买的时候的时长
	CycleTimes int32  `json:"cycleTimes"` //可选 购买时长周期数，此字段对有时长的产品有意义，默认1 表示cycleTimes个产品的最低单位时长 有这个优先使用 2024/10/25新增
}

// 续费代理资源返回
type AppInstanceRenewResp struct {
	OrderNo    string `json:"orderNo"`    //平台订单号
	AppOrderNo string `json:"appOrderNo"` //购买者订单号(渠道商订单号) 原样返回
	Amount     string `json:"amount"`     //花费金额
}

// 释放代理资源请求
type AppInstanceReleaseReq struct {
	OrderNo   string   `json:"orderNo"`   //购买者订单号(渠道商订单号)
	Instances []string `json:"instances"` //平台实例编号
}

// 释放代理资源返回
type AppInstanceReleaseResp struct {
	OrderNo    string `json:"orderNo"`    //平台订单号
	AppOrderNo string `json:"appOrderNo"` //购买者订单号(渠道商订单号) 原样返回
	Amount     string `json:"amount"`     //花费金额
}

// 账密提取请求
type AppDrawByPwdReq struct {
	AppUsername  string `json:"appUsername"`  //必要（渠道商子账号名）
	AddressCode  string `json:"addressCode"`  //地址代码 可以传 areaCode countryCode stateCode cityCode 四种之一
	SessTime     string `json:"sessTime"`     //有效时间 1-120分钟 默认5分钟
	Num          int    `json:"num"`          //数量 默认1
	ProxyType    uint16 `json:"proxyType"`    //代理类型 104=动态国外 105=动态国内
	MaxFlowLimit int    `json:"maxFlowLimit"` //子账号最大流量限制 可选 大于0的时候生效
	ProductNo    string `json:"productNo"`    //产品编号 必要
}

// 账密提取返回
type AppDrawByPwdResp struct {
	List []AppDrawByPwdItem `json:"list"`
}

type AppDrawByPwdItem struct {
	ProxyUrl string   `json:"proxyUrl"` //代理地址
	List     []string `json:"list"`     //
}

// 动态代理余额信息请求
type AppProxyInfoReq struct {
	Username    string `json:"username"`    //平台主账号，选填 平台主账号和渠道商主账号两个必填一个
	AppUsername string `json:"appUsername"` //渠道商主账号，选填 平台主账号和渠道商主账号两个必填一个
	ProxyType   uint16 `json:"proxyType"`   //代理类型 必填 104=动态国外 105=动态国内
	ProductNo   string `json:"productNo"`   //产品编号 必填
}

// 动态代理余额信息返回
type AppProxyInfoResp struct {
	// Used        string                `json:"used"`        //已使用 单位M
	// Total       string                `json:"total"`       //购买总数 单位M
	// Balance     string                `json:"balance"`     //剩余 单位M
	// IpWhiteList []string              `json:"ipWhiteList"` //ip白名单
	List []AppProxyInfoProduct `json:"list"` //产品列表数据
}
type AppProxyInfoProduct struct {
	Used        string   `json:"used"`        //已使用 单位M
	Total       string   `json:"total"`       //总数   单位M
	Balance     string   `json:"balance"`     //剩余   单位M
	IpWhiteList []string `json:"ipWhiteList"` //ip白名单
	ProductNo   string   `json:"productNo"`   //产品编号
	IpUsed      int      `json:"ipUsed"`      //已使用ip 单位个
	IpTotal     int      `json:"ipTotal"`     //总数ip   单位个
}

// 动态产品区域列表请求
type AppProductAreaReq struct {
	ProductNo string `json:"productNo" `             //平台产品编号
	ProxyType int16  `json:"proxyType" label:"代理类型"` //代理类型 104=动态国外 105=动态国内
}

// 动态产品区域列表返回
type AppProductAreaResp struct {
	ProductNo    string `json:"productNo"`    //平台产品编号
	ProxyType    int16  `json:"proxyType"`    //代理类型
	AreaCode     string `json:"areaCode"`     //区域代码（洲）
	CountryCode  string `json:"countryCode"`  //国家代码
	StateCode    string `json:"stateCode"`    //州省代码
	CityCode     string `json:"cityCode"`     //城市代码
	Status       int8   `json:"status"`       //状态 1=上架 -1=下架
	Region       string `json:"region"`       //上游供应商区域
	SupplierCode string `json:"supplierCode"` //供应商code
}

// 异步订单，开通成功或者失败后，进行回调，回调地址为app配置的回调地址
// 回调参数为
// type=order|instance
// no=订单号（请求方）|实列id
// eg：
// callbackUrl?type=order&no=apporder123456
// 订单回调app端返回的结果，Code 为success 表示成功,不再重复回调
type NotifyRes struct {
	Code string //成功success 多次回调上一次已经成功，第二次还是返回success
	Msg  string
}

// 添加ip白名单
type AppAddIpWhiteListReq struct {
	AppUsername string `json:"appUsername"` //渠道商主账号 必要
	Ip          string `json:"ip"`          //ip地址 必要
	ProxyType   uint16 `json:"proxyType"`   //代理类型 可选 默认104 104=动态国外 105=动态国内
	ProductNo   string `json:"productNo"`   //产品编号
}

// 添加ip白名单返回
type AppAddIpWhiteListResp struct {
	IpWhiteList []string `json:"ipWhiteList"` //ip白名单
}

// 删除ip白名单
type AppDelIpWhiteListReq struct {
	AppUsername string `json:"appUsername"` //渠道商主账号 必要
	Ip          string `json:"ip"`          //ip地址 必要
	ProxyType   uint16 `json:"proxyType"`   //代理类型 可选 默认104 104=动态国外 105=动态国内
	ProductNo   string `json:"productNo"`   //产品编号
}

// 删除ip白名单返回
type AppDelIpWhiteListResp struct {
	IpWhiteList []string `json:"ipWhiteList"` //ip白名单
}

// Api提取代理请求
type AppDrawByApiReq struct {
	AppUsername  string `json:"appUsername"`  //渠道商主账号 必要
	ProxyType    uint16 `json:"proxyType"`    //代理类型 必要 104=动态国外 105=动态国内
	Num          int    `json:"num"`          //提取ip数量 可选 默认1
	AddressCode  string `json:"addressCode"`  //地址代码 可选  取值 areaCode countryCode stateCode cityCode 四种之一
	Protocol     string `json:"protocol"`     //协议 可选 默认socks5  取值 socks5 http 之一
	ReturnType   string `json:"returnType"`   //数据格式 可选 默认txt  取值 txt json 之一
	Delimiter    int    `json:"delimiter"`    //分隔符 可选 只有数据格式是txt的时候生效 默认1 (1=\r\n 2=/br 3=\r 4=\n 5=\t)
	MaxFlowLimit int    `json:"maxFlowLimit"` //最大流量限制 可选 大于0的时候生效
	ProductNo    string `json:"productNo"`    //产品编号 必要
}

// Api提取代理返回
type AppDrawByApiResp struct {
	List []AppDrawByApiItem `json:"list"`
}

type AppDrawByApiItem struct {
	ProxyUrl string `json:"proxyUrl"` //提取代理Api地址
}

// 流量使用记录
type AppFlowUseLogReq struct {
	AppUsername string `json:"appUsername"` //渠道商主账号 必要
	StartTime   string `json:"startTime"`   //开始时间 可选 默认7天前 格式 2021-01-01 00:00:00
	EndTime     string `json:"endTime"`     //结束时间 可选当天 格式 2021-01-01 00:00:00
	ProductNo   string `json:"productNo"`   //产品编号 必须
	Page        int    `json:"page"`        //页码 可选 默认1
	PageSize    int    `json:"pageSize"`    //每页数量 可选 默认10 最大100
}

// 流量使用记录返回
type AppFlowUseLogResp struct {
	List    []AppFlowUseLogItem `json:"list"`    // list
	Total   int                 `json:"total"`   // 总数量
	CurPage int                 `json:"curPage"` // 当前页
}
type AppFlowUseLogItem struct {
	Used      int64  `json:"used"`      //已使用流量 B
	Total     int64  `json:"total"`     //总流量 B
	Balance   int64  `json:"balance"`   //剩余流量 B
	UsedTime  uint64 `json:"usedTime"`  //使用时间 单位秒
	ProductNo string `json:"productNo"` //产品编号
}

// 同步城市列表请求
type AppCityListReq struct {
	Codes []string `json:"codes" form:"codes"` //城市代码列表，为null获取全部
}

// 同步城市列表返回
type AppCityListResp struct {
	CityCode      string `json:"cityCode"`      //城市代码
	CityName      string `json:"cityName"`      //城市中文名称
	CityEnName    string `json:"cityEnName"`    //城市英文名称
	StateCode     string `json:"stateCode"`     //州、省代码
	StateName     string `json:"stateName"`     //州、省中文名称
	StateEnName   string `json:"stateEnName"`   //州、省英文名称
	CountryCode   string `json:"countryCode"`   //国家代码
	CountryName   string `json:"countryName"`   //国家中文名称
	CountryEnName string `json:"countryEnName"` //国家英文名称
	AreaCode      string `json:"areaCode"`      //洲代码
	AreaName      string `json:"areaName"`      //洲中文名称
	AreaEnName    string `json:"areaEnName"`    //洲英文名称
	Status        int    `json:"status"`        //状态 1=上架 -1=下架
}

// 动态代理回收请求
type AppProxyReturnReq struct {
	AppUsername string `json:"appUsername"` //渠道商主账号，必填
	ProxyType   uint16 `json:"proxyType"`   //代理类型 必填 104=动态国外 105=动态国内
	ProductNo   string `json:"productNo"`   //产品编号 必填
	IpNum       int64  `json:"ipNum"`       //回收ip数量  单位个  产品如果是按照ip数量购买 使用该字段
	FlowNum     int64  `json:"flowNum"`     //回收流量数量  单位M  产品如果是按照流量购买 使用该字段
	Remark      string `json:"remark"`      //备注 最多250个字符
}

// 动态代理回收返回
type AppProxyReturnResp struct {
	ReturnAmount float64 `json:"returnAmount"` //回收代理退还金额 单位元
}

// 指定ip开通代理资源请求
type AppAssignIpInstanceOpenReq struct {
	AppOrderNo   string `json:"appOrderNo"`   //购买者订单号(渠道商订单号) 必须
	ProductNo    string `json:"productNo"`    //产品编号 必须 但是指定购买成功的ip不一定和产品的所属区域一致
	Renew        bool   `json:"renew"`        //是否自动续费 1自动续费 默认0
	ExtBandWidth int32  `json:"extBandWidth"` //额外增加带宽 单位Mbps
	UseBridge    uint8  `json:"useBridge"`    //是否使用桥 0=随app设置 1=不使用桥 2=使用桥 默认0
	AssignIp     string `json:"assignIp"`     //指定ip购买 必须
	CycleTimes   int32  `json:"cycleTimes"`   //购买时长周期数，此字段对有时长的产品有意义，默认1 表示cycleTimes个产品的最低单位时长  必须
}

// 获取指定ip是否可以创建代理
type AppGetAssignIpInfoReq struct {
	Ip string `json:"ip"` //必要 ip
}

type AppGetAssignIpInfoResp struct {
	Ip           string `json:"ip"`           //必要 ip
	CanBuyStatus bool   `json:"canBuyStatus"` //必要 指定ip购买状态  true=可以购买 false=不能购买 默认不能购买
}

// 获取历史订单列表请求 最近1年
type AppGetOrderListReq struct {
	PageSize  int    `json:"pageSize"`  // 每页显示数量 默认20 最大100
	StartTime string `json:"startTime"` //开始时间 可选 默认365天前 格式 2021-01-01 00:00:00
	EndTime   string `json:"endTime"`   //结束时间 可选当天 格式 2021-01-01 00:00:00
	Page      int    `json:"page"`      //页码 可选 默认1
}

// 历史订单返回
type AppOrderListResp struct {
	Page     int            `json:"page"`     // 页码 原样返回或者返回默认
	PageSize int            `json:"pageSize"` // 每页显示数量  原样返回或者返回默认
	Total    int64          `json:"total"`    // 总数量
	CurPage  int            `json:"curPage"`  // 当前页
	List     []AppOrderItem `json:"list"`     // 订单列表
}

type AppOrderItem struct {
	OrderNo      string    `json:"orderNo"`      //平台订单编号
	AppOrderNo   string    `json:"appOrderNo"`   //渠道商（购买方）订单编号
	CityCode     string    `json:"cityCode"`     //城市编号
	InstanceNo   string    `json:"instanceNo"`   //代理实例编号  新建购买代理订单该字段为空
	Type         int8      `json:"type"`         //订单类型 1=新建购买代理 2=续费代理 3=释放代理
	Status       int8      `json:"status"`       //订单状态 1=待处理 2=处理中 3=处理成功 4=处理失败 5=部分完成
	Count        int       `json:"count"`        //购买数量
	SuccessCount int       `json:"successCount"` //成功数量
	Amount       string    `json:"amount"`       //金额
	ExtBandWidth int32     `json:"extBandWidth"` //额外带宽 单位MB
	Flow         int       `json:"flow"`         //总流量MB或者总ip个数 动态代理订单才有值
	CreatedAt    time.Time `json:"createdAt"`    //下单时间
}

// 获取实列历史列表请求
type AppGetInstanceListReq struct {
	PageSize  int    `json:"pageSize"`  // 每页显示数量 默认20 最大100
	StartTime string `json:"startTime"` //开始时间 可选 默认365天前 格式 2021-01-01 00:00:00
	EndTime   string `json:"endTime"`   //结束时间 可选当天 格式 2021-01-01 00:00:00
	Page      int    `json:"page"`      //页码 可选 默认1
}

// 实例返回
type AppInstanceListResp struct {
	Page     int               `json:"page"`     // 页码 原样返回或者返回默认
	PageSize int               `json:"pageSize"` // 每页显示数量  原样返回或者返回默认
	Total    int64             `json:"total"`    // 总数量
	CurPage  int               `json:"curPage"`  // 当前页
	List     []AppInstanceItem `json:"list"`     // 实例列表
}

type AppInstanceItem struct {
	InstanceNo  string    `json:"instanceNo"`  //平台实例编号（渠道商续费和释放操作使用该编号）
	ProxyType   uint      `json:"proxyType"`   //代理类型 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	Protocol    string    `json:"protocol" `   //协议类型 多个用英文逗号分隔 1=socks5 2=http 3=https 4=ssh
	Ip          string    `json:"ip"`          //代理地址 用户实际代理访问使用
	Port        uint      `json:"port"`        //代理端口
	RegionId    string    `json:"regionId"`    //区域地址
	CountryCode string    `json:"countryCode"` //国家代码
	CityCode    string    `json:"cityCode"`    //城市代码
	UseType     string    `json:"useType"`     //使用方式 多个用英文逗号分隔  1=账密 2=ip白名单 3=uuid（uuid写password内）
	Username    string    `json:"username"`    //账户名或uuid 动态为平台主账号
	Pwd         string    `json:"pwd"`         //密码
	OrderNo     string    `json:"orderNo"`     //创建该代理实例的平台订单号
	UserExpired int64     `json:"userExpired"` //到期时间
	FlowTotal   float64   `json:"flowTotal"`   //总流量
	FlowBalance float64   `json:"flowBalance"` //剩余流量
	Status      int8      `json:"status"`      //1=待创建 2=创建中 3=运行中 6=已停止 10=关闭 11=释放
	Bridges     []string  `json:"bridges"`     //桥地址列表
	OpenAt      time.Time `json:"openAt"`      //开通时间
	RenewAt     time.Time `json:"renewAt"`     //最后成功续费时间
	ReleaseAt   time.Time `json:"releaseAt"`   //释放成功时间
	ProductNo   string    `json:"productNo"`   //产品编号
	ExtendIp    string    `json:"extendIp"`    //扩展地址 仅供展示，部分产品该字段有值 2024-11-20 新增
}

// 设置子账户流量上限请求
type AppSetProxyUserFlowLimitReq struct {
	ProductNo   string `json:"productNo"`   //平台产品编号 必填
	AppUsername string `json:"appUsername"` //渠道商子账号 必填
	LimitFlow   int64  `json:"limitFlow"`   //动态流量上限 必填 -1表示不限制  单位B
	Remark      string `json:"remark"`      //备注
}

// 获取子账户信息请求
type AppProxyUserInfoReq struct {
	ProductNo   string `json:"productNo"`   //平台产品编号 必填
	AppUsername string `json:"appUsername"` //渠道商子账号 必填
}

// 获取子账户信息返回
type AppProxyUserInfoResp struct {
	ProductNo   string `json:"productNo"`   //平台产品编号
	AppUsername string `json:"appUsername"` //渠道商子账号名称
	Username    string `json:"username"`    //平台子账号名称
	LimitFlow   int64  `json:"limitFlow"`   //动态流量上限 -1表示不限制  单位B
	UseFlow     int64  `json:"useFlow"`     //使用流量  单位B
	Password    string `json:"password"`    //子账号密码
	Remark      string `json:"remark"`      //备注
	Status      int8   `json:"status"`      //子账号状态 1=正常 2=禁用
}

type AppInstanceAfterSaleReleaseReq struct {
	OrderNo          string         `json:"orderNo"`          //订单号(渠道商订单号)
	Instances        []string       `json:"instances"`        //平台实例编号
	Remark           string         `json:"remark"`           //备注
	Reason           string         `json:"reason"`           //原因
	DingTalkUserList []DingTalkUser `json:"dingTalkUserList"` //用户列表
}

type DingTalkUser struct {
	Uid   string `json:"uid"`   //
	Phone string `json:"phone"` //
}

type AppInstanceAfterSaleReleaseResp struct {
	OrderNo    string `json:"orderNo"`    //平台订单号
	AppOrderNo string `json:"appOrderNo"` //购买者订单号(渠道商订单号) 原样返回
	Amount     string `json:"amount"`     //金额
}

type AppProjectListReq struct {
	Codes []string `json:"codes" form:"codes"` //项目代码列表，为null获取全部
}

type AppProjectListResp struct {
	Code   string `json:"code"`   //项目代码
	Name   string `json:"name"`   //项目名称
	Status int    `json:"status"` //状态 1=上架 -1=下架
}

type ProjectItem struct {
	Code      string `json:"code"`      // id
	Inventory int    `json:"inventory"` // 库存
}

// 获取产品信息请求
type AppProductInfoReq struct {
	ProductNo string `json:"productNo"` //产品编号 必填
}

// 重置实例密码请求
type AppResetProxyPasswordReq struct {
	InstanceNoList []string `json:"instanceNoList"` //实例编号数组 必填
	ResetNo        string   `json:"resetNo"`        //重置单号 必填 平台根据该单号做幂等判断，同样的单号重复请求只会处理第一次的请求，之后的请求只会返回第一次的结果 最长32位，只能包含字母、数字、下划线
}

// 重置实例密码返回
type AppResetProxyPasswordResp struct {
	List    []AppResetProxyPasswordResult `json:"list"`    //重置结果列表
	ResetNo string                        `json:"resetNo"` //重置单号 平台自动生成或者渠道商传递
}

type AppResetProxyPasswordResult struct {
	InstanceNo  string `json:"instanceNo"`  //实例编号
	NewPassword string `json:"newPassword"` //新密码 如果为空 说明还在处理中或者处理失败
	Status      int8   `json:"status"`      //状态 1待处理 2处理中 3成功 4失败
}
