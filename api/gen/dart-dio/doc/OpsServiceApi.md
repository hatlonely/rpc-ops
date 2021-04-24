# api.api.OpsServiceApi

## Load the API package
```dart
import 'package:api/api.dart';
```

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**opsServiceDelJob**](OpsServiceApi.md#opsservicedeljob) | **delete** /v1/job/{id} | 
[**opsServiceDelRepository**](OpsServiceApi.md#opsservicedelrepository) | **delete** /v1/repository/{id} | 
[**opsServiceDelVariable**](OpsServiceApi.md#opsservicedelvariable) | **delete** /v1/variable/{id} | 
[**opsServiceDescribeRepository**](OpsServiceApi.md#opsservicedescriberepository) | **post** /v1/describeRepository | 
[**opsServiceGetJob**](OpsServiceApi.md#opsservicegetjob) | **get** /v1/job/{id} | 
[**opsServiceGetRepository**](OpsServiceApi.md#opsservicegetrepository) | **get** /v1/repository/{id} | 
[**opsServiceGetVariable**](OpsServiceApi.md#opsservicegetvariable) | **get** /v1/variable/{id} | 
[**opsServiceListJob**](OpsServiceApi.md#opsservicelistjob) | **get** /v1/job | 
[**opsServiceListRepository**](OpsServiceApi.md#opsservicelistrepository) | **get** /v1/repository | 
[**opsServiceListVariable**](OpsServiceApi.md#opsservicelistvariable) | **get** /v1/variable | 
[**opsServicePutRepository**](OpsServiceApi.md#opsserviceputrepository) | **post** /v1/repository | 
[**opsServicePutVariable**](OpsServiceApi.md#opsserviceputvariable) | **post** /v1/variable | 
[**opsServiceRunOps**](OpsServiceApi.md#opsservicerunops) | **post** /v1/runOps | 
[**opsServiceUpdateRepository**](OpsServiceApi.md#opsserviceupdaterepository) | **put** /v1/repository/{id} | 
[**opsServiceUpdateVariable**](OpsServiceApi.md#opsserviceupdatevariable) | **put** /v1/variable/{id} | 


# **opsServiceDelJob**
> JsonObject opsServiceDelJob(id)



### Example 
```dart
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceDelJob(id);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceDelJob: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**JsonObject**](JsonObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceDelRepository**
> JsonObject opsServiceDelRepository(id)



### Example 
```dart
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceDelRepository(id);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceDelRepository: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**JsonObject**](JsonObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceDelVariable**
> JsonObject opsServiceDelVariable(id)



### Example 
```dart
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceDelVariable(id);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceDelVariable: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

### Return type

[**JsonObject**](JsonObject.md)

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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var body = new ApiDescribeRepositoryReq(); // ApiDescribeRepositoryReq | 

try { 
    var result = api_instance.opsServiceDescribeRepository(body);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceDescribeRepository: $e\n');
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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceGetJob(id);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceGetJob: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceGetRepository(id);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceGetRepository: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var id = id_example; // String | 

try { 
    var result = api_instance.opsServiceGetVariable(id);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceGetVariable: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 

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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var repositoryID = repositoryID_example; // String | 
var offset = 56; // int | 
var limit = 56; // int | 

try { 
    var result = api_instance.opsServiceListJob(repositoryID, offset, limit);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceListJob: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repositoryID** | **String**|  | [optional] 
 **offset** | **int**|  | [optional] 
 **limit** | **int**|  | [optional] 

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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var offset = 56; // int | 
var limit = 56; // int | 

try { 
    var result = api_instance.opsServiceListRepository(offset, limit);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceListRepository: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int**|  | [optional] 
 **limit** | **int**|  | [optional] 

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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var offset = 56; // int | 
var limit = 56; // int | 

try { 
    var result = api_instance.opsServiceListVariable(offset, limit);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceListVariable: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int**|  | [optional] 
 **limit** | **int**|  | [optional] 

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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var body = new ApiRepository(); // ApiRepository | 

try { 
    var result = api_instance.opsServicePutRepository(body);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServicePutRepository: $e\n');
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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var body = new ApiVariable(); // ApiVariable | 

try { 
    var result = api_instance.opsServicePutVariable(body);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServicePutVariable: $e\n');
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
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var body = new ApiRunOpsReq(); // ApiRunOpsReq | 

try { 
    var result = api_instance.opsServiceRunOps(body);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceRunOps: $e\n');
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
> JsonObject opsServiceUpdateRepository(id, body)



### Example 
```dart
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var id = id_example; // String | 
var body = new InlineObject(); // InlineObject | 

try { 
    var result = api_instance.opsServiceUpdateRepository(id, body);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceUpdateRepository: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 
 **body** | [**InlineObject**](InlineObject.md)|  | 

### Return type

[**JsonObject**](JsonObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opsServiceUpdateVariable**
> JsonObject opsServiceUpdateVariable(id, body)



### Example 
```dart
import 'package:api/api.dart';

var api_instance = new OpsServiceApi();
var id = id_example; // String | 
var body = new InlineObject1(); // InlineObject1 | 

try { 
    var result = api_instance.opsServiceUpdateVariable(id, body);
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServiceUpdateVariable: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**|  | 
 **body** | [**InlineObject1**](InlineObject1.md)|  | 

### Return type

[**JsonObject**](JsonObject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

