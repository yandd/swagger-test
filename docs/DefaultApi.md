# contract_center api

## 目录

Method | HTTP request | Description
------------- | ------------- | -------------
[**1.2 开通电子签章服务**](DefaultApi.md#CompanyTsignCreate1Post) | **Post** /contractcenter/company/tsign_create1 | 
[**1.3 开通电子签章服务**](DefaultApi.md#CompanyTsignCreate2Post) | **Post** /contractcenter/company/tsign_create2 | 
[**1.4 开通电子签章服务**](DefaultApi.md#CompanyTsignCreate3Post) | **Post** /contractcenter/company/tsign_create3 | 
[**1.1 开通电子签章服务**](DefaultApi.md#CompanyTsignCreatePost) | **Post** /contractcenter/company/tsign_create | 


## API

### <a id="CompanyTsignCreate1Post">1.2 开通电子签章服务</a>

#### Request

```http
**Post** /contractcenter/company/tsign_create1
```

#### Parameters

Name|Type|Must|Position|Description
----|----|----|--------|--------
 **authToken** | **string**| 是 | formData | token
 **mobile** | **string**| 是 | formData | 预留号码

#### Response
```json
{
  "data" : {
    "account_id" : "aeiou",
    "mobile" : "aeiou"
  },
  "code" : 0,
  "msg" : "aeiou"
}
```

### <a id="CompanyTsignCreate2Post">1.3 开通电子签章服务</a>

#### Request

```http
**Post** /contractcenter/company/tsign_create2
```

#### Parameters

Name|Type|Must|Position|Description
----|----|----|--------|--------
 **authToken** | **string**| 是 | formData | token
 **mobile** | **string**| 是 | formData | 预留号码

#### Response
```json
{
  "data" : {
    "account_id" : "aeiou",
    "mobile" : "aeiou"
  },
  "code" : 0,
  "msg" : "aeiou"
}
```

### <a id="CompanyTsignCreate3Post">1.4 开通电子签章服务</a>

#### Request

```http
**Post** /contractcenter/company/tsign_create3
```

#### Parameters

Name|Type|Must|Position|Description
----|----|----|--------|--------
 **authToken** | **string**| 是 | formData | token
 **mobile** | **string**| 是 | formData | 预留号码

#### Response
```json
{
  "data" : {
    "account_id" : "aeiou",
    "mobile" : "aeiou"
  },
  "code" : 0,
  "msg" : "aeiou"
}
```

### <a id="CompanyTsignCreatePost">1.1 开通电子签章服务</a>

#### Request

```http
**Post** /contractcenter/company/tsign_create
```

#### Parameters

Name|Type|Must|Position|Description
----|----|----|--------|--------
 **authToken** | **string**| 是 | formData | token
 **mobile** | **string**| 是 | formData | 预留号码

#### Response
```json
{
  "data" : {
    "account_id" : "aeiou",
    "mobile" : "aeiou"
  },
  "code" : 0,
  "msg" : "aeiou"
}
```

