package dto

import "time"

// AppOpen 请求统计为密文传输
type AppOpenReq struct {
	Version string `json:"version" form:"version"` //版本
	Encrypt string `json:"encrypt" form:"encrypt"` //加密方式 aes,rsa（版本2以上提供，默认aes,老接口rsa）aes cbc模式
	AppKey  string `json:"appKey" form:"appKey"`   //appKey
	Params  string `json:"params" form:"params"`   // 根据加密方式密文 转base64
}

type Res struct {
	ReqId string `json:"reqId"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  string `json:"data"` //密文转base64
}

// 获取产品列表请求
type AppProductSyncReq struct {
	ProxyType []int `json:"proxyType" form:"proxyType"` //代理类型 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
}

// 获取产品列表返回 （明文返回）
type AppProductSyncResp struct {
	ProductNo      string  `json:"productNo" form:"productNo"`     //必要，产品Id 保持唯一
	ProductName    string  `json:"productName" form:"productName"` //必要,商品名
	ProxyType      int16   `json:"proxyType" form:"proxyType"`     //必要, 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	UseType        string  `json:"useType" form:"useType"`         //必要, ,分割  1=账密 2=白名单 3=uuid
	Protocol       string  `json:"protocol" form:"protocol"`       //必要,1=socks5 2=http 3=https 4=ssh
	UseLimit       int8    `json:"useLimit" form:"useLimit"`       //必要,1=出口ip国外 2=出口ip国内 3=无限制
	SellLimit      int8    `json:"sellLimit" form:"sellLimit"`     //必要,1=大陆可售 2=海外可售 3=无限制
	AreaCode       string  `json:"areaCode" form:"areaCode"`       //区域code
	CountryCode    string  `json:"countryCode" form:"countryCode"` //必要,国家代码 3位 iso标准
	CityCode       string  `json:"cityCode" form:"cityCode"`       //必要,城市代码 向我方提取
	Detail         string  `json:"detail" form:"detail"`           //商品描述
	CostPrice      string  `json:"costPrice" `                     //必要 价格
	Inventory      int     `json:"inventory"`                      //必要 库存
	IpType         int     `json:"ipType" `                        //ip类型 1=ipv4 2=ipv6 3=随机 默认1
	IspType        int     `json:"ispType" `                       //ispType 1=单isp 2=双isp
	Duration       int     `json:"duration"`                       //必要 时长 0无限制
	Unit           int     `json:"unit" `                          //单位 1=天 2=周(7天) 3=月(自然月) 4=年(自然年365，366)
	BandWidth      int     `json:"bandWidth"`                      //带宽|流量时必要 单位 MB
	BandWidthType  int     `json:"bandWidthType" `                 //带宽类型 1=独享 2=共享 3=流量包
	BandWidthPrice string  `json:"bandWidthPrice" `                //额外带宽价格
	MaxBandWidth   int     `json:"maxBandWidth" `                  //可设置最大带宽
	Flow           int     `json:"flow"`                           //流量包
	Cpu            int     `json:"cpu"`                            //cpu数
	Memory         float64 `json:"memory"`                         //内存容量
	Enable         int8    `json:"enable"`                         //是否可以购买 1可以
	SupplierCode   string  `json:"supplierCode"`                   //供应商代码
	//新增 2024/06/27
	CIDRBlocks []CIDRBlock `json:"cidrBlocks"` //支持网段及数量
}

//新增 2024/06/27
//网段
type CIDRBlock struct {
	CIDR  string `json:"cidr"` //网段 192.168.0.0/24 172.16.0.0/16 10.0.0.0/8
	Count int    `json:"count"`
}

// 创建或修改主账号请求
type AppUserReq struct {
	AppUsername string `json:"appUsername" form:"appUsername"` //渠道商主账号 该渠道商唯一 不支持修改
	Password    string `json:"password" form:"password"`       //主账号密码(不传随机生成) 不支持修改
	Phone       string `json:"phone" form:"phone"`             //主账号手机号
	Email       string `json:"email" form:"email"`             //主账号邮箱
	AuthType    int8   `json:"authType" form:"authType"`       //认证类型
	AuthName    string `json:"authName" form:"authName"`       //用户名（非必要) 不支持修改
	No          string `json:"no" form:"no"`                   //实名证件号码(非必要) 不支持修改
	VSP         uint8  `json:"vsp" form:"vsp"`                 //vsp
	Status      int8   `json:"status" form:"status"`           // 状态 1=正常 2=禁用
}

// 创建用户返回
type AppCreateUserResp struct {
	AppUsername string `json:"appUsername" form:"appUsername"` //渠道商主账号
	Username    string `json:"username" form:"username"`       //平台主账号
	Password    string `json:"password" form:"password"`       //主账号密码
	Status      int8   `json:"status" form:"status"`           //用户状态 1=正常 2=禁用
	AuthStatus  int8   `json:"authStatus" form:"authStatus"`   //认证状态 1=未实名 2=个人实名 3=企业实名
}

// 创建或修改代理用户请求
type AppProxyUserReq struct {
	AppUsername     string `json:"appUsername" form:"appUsername"`         //渠道商子账号 该渠道商唯一  (不传随机生成) 不支持修改
	Password        string `json:"password" form:"password"`               //密码(不传随机生成) 不支持修改
	LimitFlow       int    `json:"limitFlow" form:"limitFlow"`             //动态流量上限
	MainUsername    string `json:"mainUsername" form:"mainUsername"`       //平台主账号 选填 平台主账号和渠道商主账号两个必填一个 不支持修改
	AppMainUsername string `json:"appMainUsername" form:"appMainUsername"` //渠道商主账号 选填 平台主账号和渠道商主账号两个必填一个 不支持修改
	Remark          string `json:"remark" form:"remark"`                   //备注
	Status          int8   `json:"status" form:"status"`                   //状态 1=正常 2=禁用
}

// AppProxyUserResp 创建或修改代理用户返回
type AppProxyUserResp struct {
	AppUsername string `json:"appUsername" form:"appUsername"` //渠道商子账号
	Username    string `json:"username" form:"username"`       //平台子账号
	Password    string `json:"password" form:"password"`       //子账号密码
	Status      int8   `json:"status" form:"status"`           //用户状态 1=正常 2=禁用
	AuthStatus  int8   `json:"authStatus" form:"authStatus"`   //认证状态 1=未实名 2=个人实名 3=企业实名
}

// 同步实名请求
type AppAuthUserReq struct {
	Username    string `json:"username" form:"username"`       //平台主账号 选填 平台主账号和渠道商主账号两个必填一个
	AppUsername string `json:"appUsername" form:"appUsername"` //渠道商主账号 选填 平台主账号和渠道商主账号两个必填一个
	AuthType    int8   `json:"authType" form:"authType"`       //认证类型
	AuthName    string `json:"authName" form:"authName"`       //用户名
	No          string `json:"no" form:"no"`                   //实名证件号码
	VSP         string `json:"vsp" form:"vsp"`                 //vsp
}

// 同步实名返回
type AppAuthUserResp struct {
	Username   string `json:"username" form:"username"`     //平台账号
	AuthStatus int    `json:"authStatus" form:"authStatus"` //认证状态
}

// 获取订单列表请求
type AppGetOrderReq struct {
	Orders []string `json:"orders" form:"orders"` //订单编号
}

// 订单返回
type AppOrderResp struct {
	No             string            `json:"no"`             //订单号（系统订单）
	Type           int8              `json:"type"`           //1=新建 2=续费 3=释放
	Status         int8              `json:"status"`         //状态 1=待处理 2=处理中 3=处理成功 4=处理失败 5=部分完成
	Count          int               `json:"count"`          //购买数量
	CostTotal      string            `json:"costTotal"`      //总成本
	Amount         string            `json:"amount"`         //总售价
	Discount       string            `json:"discount"`       //优惠券
	DiscountAmount string            `json:"discountAmount"` //优惠后价格
	ThirdNo        string            `json:"thirdNo"`        //三方（购买订单）订单
	CountryCode    string            `json:"countryCode"`    //国家代码
	CityCode       string            `json:"cityCode"`       //城市代码
	Refund         int               `json:"refund"`         //是否有退费 1存在退费
	UpdatedAt      time.Time         `json:"updatedAt"`      //更新时间
	Instance       []AppInstanceResp `json:"instances"`      //订单对应实列
}

// 获取实列列表请求
type AppGetInstanceReq struct {
	Instances []string `json:"Instances" form:"Instances"` //平台实例编号
}

// AppInstanceResp 实列返回
type AppInstanceResp struct {
	InstanceNo  string    `json:"instanceNo"`  //平台实例编号（续费释放使用该编号）
	ProxyType   uint      `json:"proxyType"`   //代理类型 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	Protocol    string    `json:"protocol" `   // 1=socks5 2=http 3=https 4=ssh
	Ip          string    `json:"ip"`          //代理ip地址
	Port        uint      `json:"port"`        //代理端口
	RegionId    string    `json:"regionId"`    //区域地址
	CountryCode string    `json:"countryCode"` //国家代码
	CityCode    string    `json:"cityCode"`    //城市代码
	UseType     string    `json:"useType"`     //使用方式  1=账密 2=ip白名单 3=uuid（uuid写password内）
	Username    string    `json:"username"`    //账户名或uuid 动态为平台主账号
	Pwd         string    `json:"pwd"`         //密码
	OrderNo     string    `json:"orderNo"`     //instance接口返回平台创建该实例的订单号 order接口返回操作该实例的订单号
	UserExpired int64     `json:"userExpired"` //到期时间
	FlowTotal   float64   `json:"flowTotal"`   //总流量
	FlowBalance float64   `json:"flowBalance"` //剩余流量
	Status      int8      `json:"status"`      //1=待创建 2=创建中 3=运行中 6=已停止 10=关闭 11=释放
	Renew       int8      `json:"renew"`       //1 自动续费
	Bridges     []string  `json:"bridges"`     //桥地址
	OpenAt      time.Time `json:"openAt"`      //开通时间
	RenewAt     time.Time `json:"renewAt"`     //最后一次续费时间
	ReleaseAt   time.Time `json:"releaseAt"`   //释放时间
}

// AppGetAreaReq 同步地域请求
type AppGetAreaReq struct {
	Codes []string `json:"codes" form:"codes"` //获取地域代码对应表，为null获取全部
}

// AppAreaResp 同步地域返回
type AppAreaResp struct {
	Code     string        `json:"code"`               //地域代码
	Name     string        `json:"name"`               //地域名称
	Cname    string        `json:"cname"`              //地域中文名
	Children []AppAreaResp `json:"children,omitempty"` //下级地域
}

// AppStaticProxyOpenReq 开通静态资源请求
type AppInstanceOpenReq struct {
	Params []OpenParam
}

type OpenParam struct {
	ProductNo    string `json:"productNo" form:"productNo"`       //商品spu（如果存在spu，后面6项无意义）
	ProxyType    uint16 `json:"proxyType" form:"proxyType"`       //代理类型 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	CountryCode  string `json:"countryCode" form:"countryCode"`   //国家代码
	CityCode     string `json:"cityCode" form:"cityCode"`         //城市代码
	SupplierCode string `json:"supplierCode" form:"supplierCode"` //供应商代码（可为null,随机分配）
	Unit         int8   `json:"unit" `                            //单位 1=天 2=周(7天) 3=月(自然月) 4=年(自然年365，366) 10=无限制
	IspType      int    `json:"ispType" `                         //isp类型 1=单isp 2=双isp
	OrderNo      string `json:"orderNo" form:"orderNo"`           //必要 购买发起方订单号必要，唯一
	Count        int    `json:"count" form:"count"`               //购买数量 （实例个数）静态必填 默认1 一次最大20
	Duration     int32  `json:"duration"`                         //必要 时长 默认1 为Unit的时长
	Renew        bool   `json:"renew" `                           //是否续费 1续费 默认0
	ExtBandWidth int32  `json:"extBandWidth"`                     //额外增加带宽 单位Mbps
	Username     string `json:"username" form:"username"`         //平台主账号，选填 开通动态代理的时候平台主账号和渠道商主账号两个必填一个
	AppUsername  string `json:"appUsername" form:"appUsername"`   //渠道商主账号，选填 开通动态代理的时候平台主账号和渠道商主账号两个必填一个
	Flow         int    `json:"flow" form:"flow"`                 //动态流量 最大102400MB 动态必填 单位MB
	UseBridge    uint8  `json:"useBridge"`                        //1=使用桥 2=不使用桥 3=随app设置 默认3
	//新增 2024/06/27
	CIDRBlocks []CIDRBlock `json:"cidrBlocks"`                 //静态购买所在网段及数量（产品有的才支持）
	ProjectId  string      `json:"projectId" form:"projectId"` //购买项目id,保留字段，后续会支持
}

// AppInstanceRenewReq 续费代理资源请求
type AppInstanceRenewReq struct {
	OrderNo  string `json:"orderNo" form:"orderNo"`   //购买者订单号(渠道商订单号)
	Instance string `json:"instance" form:"instance"` //平台实例编号
	Duration int32  `json:"duration" form:"duration"` //可选 时长 默认1
}

// AppInstanceReleaseReq 释放代理资源请求
type AppInstanceReleaseReq struct {
	OrderNo  string `json:"orderNo" form:"orderNo"`   //购买者订单号(渠道商订单号)
	Instance string `json:"instance" form:"instance"` //平台实例编号
}

// 释放代理资源返回
type AppInstanceReleaseResp struct {
	Status   int               `json:"status" form:"status"` //状态 1=待处理 2=处理中 3=处理成功 4=处理失败 5=部分完成
	No       string            `json:"no"`                   //订单号（系统订单）
	ThirdNo  string            `json:"thirdNo"`              //三方（购买订单）订单
	Instance []AppInstanceResp `json:"instances"`            //订单对应实例列表
}

// 购买国内动态套餐请求
type AppBuySetMealReq struct {
	OrderNo  string `json:"orderNo" form:"orderNo"`   //购买者订单号
	SKU      string `json:"sku" form:"sku"`           //商品skuid（如果存在sku，后面4项无意义）
	Username string `json:"username" form:"username"` //必要（主账号或者子账号名）
	Main     bool   `json:"main" form:"main"`         //true=主账号
	Count    int    `json:"count" form:"count"`       //数量
}

// 购买国外动态流量请求
type AppBuyFlowReq struct {
	OrderNo  string `json:"orderNo" form:"orderNo"`   //购买者订单号
	SKU      string `json:"sku" form:"sku"`           //商品skuid（如果存在sku，后面4项无意义）
	Username string `json:"username" form:"username"` //必要（主账号或者子账号名）
	Main     bool   `json:"main" form:"main"`         //true=主账号
	Amount   int    `json:"amount" form:"amount"`     //流量大小M
}

// AppDrawByPwdReq 账密提取请求
type AppDrawByPwdReq struct {
	//Main        bool   `json:"main" form:"main"`             //true=主账号 false=子账号  只有子账号提取
	AppUsername  string `json:"appUsername" form:"appUsername"`   //必要（渠道商子账号名）
	AddressCode  string `json:"addressCode" form:"addressCode"`   //地址代码 可以传 areaCode countryCode stateCode cityCode 四种之一
	SessTime     string `json:"sessTime" form:"sessTime"`         //有效时间
	Num          int    `json:"num" form:"num"`                   //数量
	ProxyType    uint16 `json:"proxyType" form:"proxyType"`       //代理类型 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
	MaxFlowLimit int    `json:"maxFlowLimit" form:"maxFlowLimit"` //子账号最大流量限制 可选 大于0的时候生效
}

// AppDrawByPwdResp 账密提取返回
type AppDrawByPwdResp struct {
	List []AppDrawByPwdItem `json:"list"`
}

type AppDrawByPwdItem struct {
	ProxyUrl string   `json:"proxyUrl"` //代理地址
	List     []string `json:"list"`     //
}

// 代理余额信息请求
type AppProxyInfoReq struct {
	Username    string `json:"username" form:"username"`       //平台主账号，选填 平台主账号和渠道商主账号两个必填一个
	AppUsername string `json:"appUsername" form:"appUsername"` //渠道商主账号，选填 平台主账号和渠道商主账号两个必填一个
	ProxyType   uint16 `json:"proxyType" form:"proxyType"`     //代理类型 101=静态云平台 102=静态国内家庭 103=静态国外家庭 104=动态国外 105=动态国内 201=whatsapp
}

// 代理余额信息返回
type AppProxyInfoResp struct {
	Used        string   `json:"used" form:"used"`               //已使用
	Total       string   `json:"total" form:"total"`             //总数
	Balance     string   `json:"balance" form:"balance"`         //剩余
	IpWhiteList []string `json:"ipWhiteList" form:"ipWhiteList"` //ip白名单
}

// AppProductAreaReq 动态产品区域列表请求
type AppProductAreaReq struct {
	ProductNo string `json:"productNo" form:"productNo"`              //平台产品编号
	ProxyType int16  `json:"proxyType" form:"proxyType" label:"代理类型"` //代理类型 104=动态国外 105=动态国内
}

// AppProductAreaResp 动态产品区域列表返回
type AppProductAreaResp struct {
	ProductNo   string `json:"productNo"`   //平台产品编号
	ProxyType   int16  `json:"proxyType"`   //代理类型
	AreaCode    string `json:"areaCode"`    //区域代码（洲）
	CountryCode string `json:"countryCode"` //国家代码
	StateCode   string `json:"stateCode"`   //州省代码
	CityCode    string `json:"cityCode"`    //城市代码
	Status      int8   `json:"status"`      //状态 1=上架 -1=下架
	RegionId    string `json:"regionId"`    //区域id
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

// AppAddIpWhiteListReq 添加ip白名单
type AppAddIpWhiteListReq struct {
	AppUsername string `json:"appUsername" form:"appUsername"` //渠道商主账号 必要
	Ip          string `json:"ip" form:"ip"`                   //ip地址 必要
	ProxyType   uint16 `json:"proxyType" form:"proxyType"`     //代理类型 可选 默认104 104=动态国外 105=动态国内
}

// AppAddIpWhiteListResp 添加ip白名单返回
type AppAddIpWhiteListResp struct {
	IpWhiteList []string `json:"ipWhiteList" form:"ipWhiteList"` //ip白名单
}

// AppDelIpWhiteListReq 删除ip白名单
type AppDelIpWhiteListReq struct {
	AppUsername string `json:"appUsername" form:"appUsername"` //渠道商主账号 必要
	Ip          string `json:"ip" form:"ip"`                   //ip地址 必要
	ProxyType   uint16 `json:"proxyType" form:"proxyType"`     //代理类型 可选 默认104 104=动态国外 105=动态国内
}

// AppDelIpWhiteListResp 删除ip白名单返回
type AppDelIpWhiteListResp struct {
	IpWhiteList []string `json:"ipWhiteList" form:"ipWhiteList"` //ip白名单
}

// AppDrawByApiReq Api提取代理请求
type AppDrawByApiReq struct {
	AppUsername  string `json:"appUsername" form:"appUsername"`   //渠道商主账号 必要
	ProxyType    uint16 `json:"proxyType" form:"proxyType"`       //代理类型 必要 104=动态国外 105=动态国内
	Num          int    `json:"num" form:"num"`                   //提取ip数量 可选 默认1
	AddressCode  string `json:"addressCode" form:"addressCode"`   //地址代码 可选  取值 areaCode countryCode stateCode cityCode 四种之一
	Protocol     string `json:"protocol" form:"protocol"`         //协议 可选 默认socks5  取值 socks5 http 之一
	ReturnType   string `json:"returnType" form:"returnType"`     //数据格式 可选 默认txt  取值 txt json 之一
	Delimiter    int    `json:"delimiter " form:"delimiter"`      //分隔符 可选 只有数据格式是txt的时候生效 默认1 (1=\r\n 2=/br 3=\r 4=\n 5=\t)
	MaxFlowLimit int    `json:"maxFlowLimit" form:"maxFlowLimit"` //最大流量限制 可选 大于0的时候生效
}

// AppDrawByApiResp Api提取代理返回
type AppDrawByApiResp struct {
	List []AppDrawByApiItem `json:"list"`
}

type AppDrawByApiItem struct {
	ProxyUrl string `json:"proxyUrl"` //提取代理Api地址
}
