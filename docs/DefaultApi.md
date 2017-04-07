# contract_center api

## 目录

Method | HTTP request | Description
------------- | ------------- | -------------
[**1.1 开通电子签章服务**](DefaultApi.md#CompanyTsignCreatePost) | **Post** /contractcenter/company/tsign_create | 


## API

### <a id="CompanyTsignCreatePost">1.1 开通电子签章服务</a>

#### Request

> **Post** /contractcenter/company/tsign_create

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

