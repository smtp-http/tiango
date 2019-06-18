# http api
API基于HTTP协议JSON RPC，其请求方法必须是POST。 它的URL是：/，Content-Type是：application / json

## api 概述

#### 请求

method: method，String
params: parameters，Json Array or object
id: Request id, Integer


#### 应答

result: null for failure
error: Json object，null for success，non-null for failure
code: error code
message: error message
id: Request id, Integer


常见 error code:

1: invalid argument
2: internal error
3: service unavailable
4: method not found
5: service timeout



## 实时数据（real time data）API

#### 产品信息
method: ProductInformation

http请求：
post  /api/v1/productinformation

example:
请求body:
```json
{
  "req_id":1,
  "data":{
      "DomSupplier":"公司A",
      "DpSupplier":"公司B",
      "ProductCn":"2395739",
      "LRstationDifference":"wfa1w35",
      "A_B":3.1242,
      "B_D":31.315,
      "E_F":8.2243,
      "G_H":9.29843,
      "Result":true,
      "Angle":32.0,
      "SizeA":54.2,
      "SizeB":253.5,
      "SizeC":87.2,
      "SizeD":67.5,
      "SizeE":65.3,
      "SizeF":36.8,
      "SizeG":65.8,
      "SizeH":53.0
  }
}
```
应答body:
```json
{
  "req_id":1,
  "rescode":0,
  "result":null
}
```


#### DP尺寸


#### Dom尺寸




## 参数更新（Param update）API

