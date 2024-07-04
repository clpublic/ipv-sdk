package dto

import "time"

// 请求统一为密文传输
type AppOpenReq struct {
	ReqId   string `json:"reqId"`   // 请求id，每次生成，如果失败重复请求，保持不变
	Version string `json:"version"` // 版本 v2
	Encrypt string `json:"encrypt"` // 加密方式 aes,rsa（版本2以上提供，默认aes,老接口rsa）aes cbc模式
	AppKey  string `json:"appKey"`  // appKey 渠道商号
	Params  string `json:"params"`  // 根据加密方式密文 转base64
}

type Res struct {
	ReqId string `json:"reqId"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  string `json:"data"` //密文转base64
}

// 获取产品列表请求
type AppProductSyncReq struct {
	ProxyType []int `json:"proxyType" form:"proxyType"` //代理类型  101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
}

// 获取产品列表返回 （明文返回）
type AppProductSyncResp struct {
	ProductNo      string      `json:"productNo" `     //必要，产品Id 保持唯一
	ProductName    string      `json:"productName"`    //必要,商品名
	ProxyType      int16       `json:"proxyType"`      //必要, 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	UseType        string      `json:"useType"`        //必要, ,分割  1=账密 2=白名单 3=uuid
	Protocol       string      `json:"protocol"`       //必要,1=socks5 2=http 3=https 4=ssh
	UseLimit       int8        `json:"useLimit"`       //必要,1=出口ip国外 2=出口ip国内 3=无限制
	SellLimit      int8        `json:"sellLimit"`      //必要,1=大陆可售 2=海外可售 3=无限制
	AreaCode       string      `json:"areaCode"`       //区域code
	CountryCode    string      `json:"countryCode"`    //必要,国家代码 3位 iso标准
	CityCode       string      `json:"cityCode"`       //必要,城市代码 向我方提取
	Detail         string      `json:"detail"`         //商品描述
	CostPrice      string      `json:"costPrice"`      //必要 价格
	Inventory      int         `json:"inventory"`      //必要 库存
	IpType         int         `json:"ipType"`         //ip类型 1=ipv4 2=ipv6 3=随机 默认1
	IspType        int         `json:"ispType"`        //ispType 1=单isp 2=双isp
	Duration       int         `json:"duration"`       //必要 时长 0无限制
	Unit           int         `json:"unit"`           //单位 1=天 2=周(7天) 3=月(自然月) 4=年(自然年365，366)
	BandWidth      int         `json:"bandWidth"`      //带宽|流量时必要 单位 MB
	BandWidthType  int         `json:"bandWidthType"`  //带宽类型 1=独享 2=共享 3=流量包
	BandWidthPrice string      `json:"bandWidthPrice"` //额外带宽价格
	MaxBandWidth   int         `json:"maxBandWidth"`   //可设置最大带宽
	Flow           int         `json:"flow"`           //流量包
	Cpu            int         `json:"cpu"`            //cpu数
	Memory         float64     `json:"memory"`         //内存容量
	Enable         int8        `json:"enable"`         //是否可以购买 1可以
	SupplierCode   string      `json:"supplierCode"`   //供应商代码
	CIDRBlocks     []CIDRBlock `json:"cidrBlocks"`     //支持网段及数量 新增 2024/06/27
}

// 新增 2024/06/27
// 网段
type CIDRBlock struct {
	CIDR  string `json:"cidr"` //网段 192.168.0.0/24 172.16.0.0/16 10.0.0.0/8
	Count int    `json:"count"`
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
	OrderNo  string `json:"orderNo"`  // 平台订单编号
	Page     int    `json:"page"`     // 页码  默认1
	PageSize int    `json:"pageSize"` // 每页显示数量  默认10 最大100
}

// 订单返回
type AppOrderResp struct {
	OrderNo    string            `json:"orderNo"`    //平台订单号
	AppOrderNo string            `json:"appOrderNo"` //渠道商（购买订单）订单号
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
	Ip          string    `json:"ip"`          //代理ip地址
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
}

// 同步地域请求
type AppGetAreaReq struct {
	Codes []string `json:"codes"` //获取地域代码对应表，为null获取全部
}

// 同步地域返回
type AppAreaResp struct {
	Code     string        `json:"code"`               //地域代码
	Name     string        `json:"name"`               //地域名称
	Cname    string        `json:"cname"`              //地域中文名
	Children []AppAreaResp `json:"children,omitempty"` //下级地域
}

// 开通代理资源请求
type AppInstanceOpenReq struct {
	AppOrderNo string      `json:"appOrderNo"` //购买者订单号(渠道商订单号)
	Params     []OpenParam `json:"params"`     //购买代理产品列表
}

type OpenParam struct {
	ProductNo    string      `json:"productNo"`    //商品编号（如果存在，后面6项无意义）
	ProxyType    uint16      `json:"proxyType"`    //代理类型 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	CountryCode  string      `json:"countryCode"`  //国家代码
	CityCode     string      `json:"cityCode"`     //城市代码
	SupplierCode string      `json:"supplierCode"` //供应商代码（可为null,随机分配）
	Unit         int8        `json:"unit"`         //单位 1=天 2=周(7天) 3=月(自然月) 4=年(自然年365，366) 10=无限制
	IspType      int         `json:"ispType"`      //isp类型 1=单isp 2=双isp
	Count        int         `json:"count"`        //购买数量 （实例个数）静态必填 默认1 一次最大20
	Duration     int32       `json:"duration"`     //必要 时长 默认1 为Unit的时长
	Renew        bool        `json:"renew"`        //是否续费 1续费 默认0
	ExtBandWidth int32       `json:"extBandWidth"` //额外增加带宽 单位Mbps
	AppUsername  string      `json:"appUsername"`  //渠道商主账号，开通动态代理的时候必填(必须在平台上注册过)
	Flow         int         `json:"flow"`         //动态流量 最大102400MB 动态必填 单位MB
	UseBridge    uint8       `json:"useBridge"`    //1=使用桥 2=不使用桥 3=随app设置 默认3
	CIDRBlocks   []CIDRBlock `json:"cidrBlocks"`   //静态购买所在网段及数量（产品有的才支持） 2024/06/27新增
	ProjectId    string      `json:"projectId"`    //购买项目id,保留字段，后续会支持
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
	Duration   int32  `json:"duration"`   //可选 时长 默认1
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
}

// 动态代理余额信息返回
type AppProxyInfoResp struct {
	Used        string   `json:"used"`        //已使用
	Total       string   `json:"total"`       //总数
	Balance     string   `json:"balance"`     //剩余
	IpWhiteList []string `json:"ipWhiteList"` //ip白名单
}

// 动态产品区域列表请求
type AppProductAreaReq struct {
	ProductNo string `json:"productNo" `             //平台产品编号
	ProxyType int16  `json:"proxyType" label:"代理类型"` //代理类型 104=动态国外 105=动态国内
}

// 动态产品区域列表返回
type AppProductAreaResp struct {
	ProductNo   string `json:"productNo"`   //平台产品编号
	ProxyType   int16  `json:"proxyType"`   //代理类型
	AreaCode    string `json:"areaCode"`    //区域代码（洲）
	CountryCode string `json:"countryCode"` //国家代码
	StateCode   string `json:"stateCode"`   //州省代码
	CityCode    string `json:"cityCode"`    //城市代码
	Status      int8   `json:"status"`      //状态 1=上架 -1=下架
	Region      string `json:"region"`      //上游供应商区域
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
}

// Api提取代理返回
type AppDrawByApiResp struct {
	List []AppDrawByApiItem `json:"list"`
}

type AppDrawByApiItem struct {
	ProxyUrl string `json:"proxyUrl"` //提取代理Api地址
}
