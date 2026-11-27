# ipv-sdk

ipipv app开通实例 golang sdk

文档详见 [https://app.doc.ipipv.com](https://app.doc.ipipv.com)

## V2 自建资源指定天数首购

`GetProductStock` 和 `GetProductInfo` 的产品返回数据新增 `AllowCustomDays`（JSON 字段 `allowCustomDays`）：只有值为 `1` 时才允许指定天数首购，`-1` 表示不允许；旧服务端未返回时为 `0`，也应视为未开启。该能力要求产品属于自建 `ipv` / `ipv_cn` 静态资源，且供应商、渠道商两个开关均开启；产品周期须为正整数个月或 30 天的正整数倍。

普通开通在 `OpenParam` 中设置 `Days`，指定 IP 开通在 `AppAssignIpInstanceOpenReq` 中设置 `Days`。当前支持 **1～15 天**，必须指定 `ProductNo`。使用已初始化的 `client`（参数类型来自 `github.com/clpublic/ipv-sdk/dto`）：

```go
days := int32(7)
resp, err := client.InstanceOpen(dto.AppInstanceOpenReq{
	AppOrderNo: "custom-days-order-001",
	Params: []dto.OpenParam{
		{
			ProductNo: "ipv-product-001",
			Days:      &days,
			Count:     1,
		},
	},
})
```

指定 IP 开通示例：

```go
days := int32(7)
resp, err := client.InstanceOpenAssignIp(dto.AppAssignIpInstanceOpenReq{
	AppOrderNo: "custom-days-assign-ip-001",
	ProductNo:  "ipv-product-001",
	AssignIp:   "192.0.2.1", // 替换为可购买的 IP
	Days:       &days,
})
```

`Days` 为 `*int32`：不设置（`nil`）时 JSON 省略 `days`，保持原有一天测试或整周期开通逻辑；显式传入 `0`、负数或大于 15 的值会由服务端拒绝。传入 `days` 时，普通开通忽略 `Unit`、`Duration`、`CycleTimes`，指定 IP 开通忽略 `CycleTimes`；指定 IP 开通未传 `days` 时，`CycleTimes` 仍须大于 0。

指定天数只影响首购，价格按完整产品周期价格乘以 `days / 周期天数` 折算，月周期按每月 30 天计算。后续调用 `InstanceRenew` 按产品完整周期续费，推荐明确设置 `CycleTimes`：

```go
resp, err := client.InstanceRenew(dto.AppInstanceRenewReq{
	AppOrderNo: "renew-order-001",
	Instances: []dto.Instance{
		{InstanceNo: "c_instance_001", CycleTimes: 1},
	},
})
```

例如，按月产品首购 7 天后，`CycleTimes: 1` 续费一个完整月周期（30 天）；不传续费时长也默认一个完整产品周期，不会沿用首购的 7 天。

## 渠道商下架/即将下架代理 IP 列表

`GetSupplierOfflineList` 对应 `POST /api/open/app/instance/offline/list/v2`，只返回当前渠道商自己的、运行中的天翼代理：

```go
status := dto.SupplierOfflineStatusPending // 只看即将下架；传 nil 返回两类
resp, err := client.GetSupplierOfflineList(dto.AppSupplierOfflineListReq{
	OfflineStatus: &status,
	Page:          1,
	PageSize:      100,
})
if err != nil {
	// handle error
}
for _, item := range resp.List {
	// item.Ip、item.OfflineStatus、item.OfflineTime
}
```

字段说明：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `ip` | string | 代理 IP |
| `offlineStatus` | int8 | `-1`=已下架，`2`=即将下架 |
| `offlineTime` | string/null | 已下架为平台首次检测时间；即将下架为 `null` |

状态常量：

- `dto.SupplierOfflineStatusOffline`（`-1`）：已下架；
- `dto.SupplierOfflineStatusPending`（`2`）：即将下架。

说明：

- 天翼不提供真实下架时间；已下架的 `offlineTime` 是平台首次检测时间，即将下架固定为 `null`。
- 接口只覆盖天翼；已停止/已关闭/已释放的实例不返回。
- 不回填历史数据，接口从功能上线时刻开始覆盖。
- 建议渠道每 30~60 分钟拉取一次并整体替换本地异常 IP 集合；待下架被上游撤销后，下次查询不再返回。
