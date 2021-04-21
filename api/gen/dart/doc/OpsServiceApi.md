# openapi.api.OpsServiceApi

## Load the API package
```dart
import 'package:openapi/api.dart';
```

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**opsServiceDelJob**](OpsServiceApi.md#opsServiceDelJob) | **DELETE** /v1/job/{id} | 
[**opsServiceDelRepository**](OpsServiceApi.md#opsServiceDelRepository) | **DELETE** /v1/repository/{id} | 
[**opsServiceDelVariable**](OpsServiceApi.md#opsServiceDelVariable) | **DELETE** /v1/variable/{id} | 
[**opsServiceDescribeRepository**](OpsServiceApi.md#opsServiceDescribeRepository) | **POST** /v1/describeRepository | 
[**opsServiceGetJob**](OpsServiceApi.md#opsServiceGetJob) | **GET** /v1/job/{id} | 
[**opsServiceGetRepository**](OpsServiceApi.md#opsServiceGetRepository) | **GET** /v1/repository/{id} | 
[**opsServiceGetVariable**](OpsServiceApi.md#opsServiceGetVariable) | **GET** /v1/variable/{id} | 
[**opsServiceListJob**](OpsServiceApi.md#opsServiceListJob) | **GET** /v1/job | 
[**opsServiceListRepository**](OpsServiceApi.md#opsServiceListRepository) | **GET** /v1/repository | 
[**opsServiceListVariable**](OpsServiceApi.md#opsServiceListVariable) | **GET** /v1/variable | 
[**opsServicePutRepository**](OpsServiceApi.md#opsServicePutRepository) | **POST** /v1/repository | 
[**opsServicePutVariable**](OpsServiceApi.md#opsServicePutVariable) | **POST** /v1/variable | 
[**opsServiceRunOps**](OpsServiceApi.md#opsServiceRunOps) | **POST** /v1/runOps | 
[**opsServiceUpdateRepository**](OpsServiceApi.md#opsServiceUpdateRepository) | **PUT** /v1/repository/{id} | 
[**opsServiceUpdateVariable**](OpsServiceApi.md#opsServiceUpdateVariable) | **PUT** /v1/variable/{id} | 


# **opsServiceDelJob**
> Object opsServiceDelJob(id)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceDelJob(id);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceDelJob: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceDelRepository**
> Object opsServiceDelRepository(id)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceDelRepository(id);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceDelRepository: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceDelVariable**
> Object opsServiceDelVariable(id)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceDelVariable(id);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceDelVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceDescribeRepository**
> ApiPlaybook opsServiceDescribeRepository(body)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var body = ApiDescribeRepositoryReq(); // ApiDescribeRepositoryReq | 

try { 
    var result = api_instance.opsServiceDescribeRepository(body);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceDescribeRepository: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiDescribeRepositoryReq**](ApiDescribeRepositoryReq.md)|  | 

### Return type

[**ApiPlaybook**](ApiPlaybook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceGetJob**
> ApiJob opsServiceGetJob(id)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceGetJob(id);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceGetJob: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**ApiJob**](ApiJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceGetRepository**
> ApiRepository opsServiceGetRepository(id)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceGetRepository(id);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceGetRepository: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**ApiRepository**](ApiRepository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceGetVariable**
> ApiVariable opsServiceGetVariable(id)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceGetVariable(id);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceGetVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]

### Return type

[**ApiVariable**](ApiVariable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceListJob**
> ApiListJobRes opsServiceListJob(repositoryID, offset, limit)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var repositoryID = repositoryID_example; // String | 
var offset = 56; // int | 
var limit = 56; // int | 

try { 
    var result = api_instance.opsServiceListJob(repositoryID, offset, limit);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceListJob: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repositoryID** | **String**|  | [optional] [default to null]
 **offset** | **int**|  | [optional] [default to null]
 **limit** | **int**|  | [optional] [default to null]

### Return type

[**ApiListJobRes**](ApiListJobRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceListRepository**
> ApiListRepositoryRes opsServiceListRepository(offset, limit)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var offset = 56; // int | 
var limit = 56; // int | 

try { 
    var result = api_instance.opsServiceListRepository(offset, limit);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceListRepository: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int**|  | [optional] [default to null]
 **limit** | **int**|  | [optional] [default to null]

### Return type

[**ApiListRepositoryRes**](ApiListRepositoryRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceListVariable**
> ApiListVariableRes opsServiceListVariable(offset, limit)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var offset = 56; // int | 
var limit = 56; // int | 

try { 
    var result = api_instance.opsServiceListVariable(offset, limit);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceListVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int**|  | [optional] [default to null]
 **limit** | **int**|  | [optional] [default to null]

### Return type

[**ApiListVariableRes**](ApiListVariableRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServicePutRepository**
> ApiRepositoryID opsServicePutRepository(body)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var body = ApiRepository(); // ApiRepository | 

try { 
    var result = api_instance.opsServicePutRepository(body);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServicePutRepository: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiRepository**](ApiRepository.md)|  | 

### Return type

[**ApiRepositoryID**](ApiRepositoryID.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServicePutVariable**
> ApiVariableID opsServicePutVariable(body)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var body = ApiVariable(); // ApiVariable | 

try { 
    var result = api_instance.opsServicePutVariable(body);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServicePutVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiVariable**](ApiVariable.md)|  | 

### Return type

[**ApiVariableID**](ApiVariableID.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceRunOps**
> ApiRunOpsRes opsServiceRunOps(body)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var body = ApiRunOpsReq(); // ApiRunOpsReq | 

try { 
    var result = api_instance.opsServiceRunOps(body);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceRunOps: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiRunOpsReq**](ApiRunOpsReq.md)|  | 

### Return type

[**ApiRunOpsRes**](ApiRunOpsRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceUpdateRepository**
> Object opsServiceUpdateRepository(id, body)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var id = id_example; // String | 
var body = ApiRepository(); // ApiRepository | 

try { 
    var result = api_instance.opsServiceUpdateRepository(id, body);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceUpdateRepository: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]
 **body** | [**ApiRepository**](ApiRepository.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceUpdateVariable**
> Object opsServiceUpdateVariable(id, body)



### Example 
```dart
import 'package:openapi/api.dart';

var api_instance = OpsServiceApi();
var id = id_example; // String | 
var body = ApiVariable(); // ApiVariable | 

try { 
    var result = api_instance.opsServiceUpdateVariable(id, body);
    print(result);
} catch (e) {
    print("Exception when calling OpsServiceApi->opsServiceUpdateVariable: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | [default to null]
 **body** | [**ApiVariable**](ApiVariable.md)|  | 

### Return type

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

