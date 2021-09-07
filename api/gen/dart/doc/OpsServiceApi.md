# openapi.api.OpsServiceApi

## Load the API package
```dart
import 'package:openapi/api.dart';
```

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**opsServiceDelJob**](OpsServiceApi.md#opsservicedeljob) | **DELETE** /v1/job/{id} | 
[**opsServiceDelRepository**](OpsServiceApi.md#opsservicedelrepository) | **DELETE** /v1/repository/{id} | 
[**opsServiceDelVariable**](OpsServiceApi.md#opsservicedelvariable) | **DELETE** /v1/variable/{id} | 
[**opsServiceDescribeRepository**](OpsServiceApi.md#opsservicedescriberepository) | **POST** /v1/describeRepository | 
[**opsServiceGetJob**](OpsServiceApi.md#opsservicegetjob) | **GET** /v1/job/{id} | 
[**opsServiceGetRepository**](OpsServiceApi.md#opsservicegetrepository) | **GET** /v1/repository/{id} | 
[**opsServiceGetVariable**](OpsServiceApi.md#opsservicegetvariable) | **GET** /v1/variable/{id} | 
[**opsServiceListJob**](OpsServiceApi.md#opsservicelistjob) | **GET** /v1/job | 
[**opsServiceListRepository**](OpsServiceApi.md#opsservicelistrepository) | **GET** /v1/repository | 
[**opsServiceListVariable**](OpsServiceApi.md#opsservicelistvariable) | **GET** /v1/variable | 
[**opsServicePing**](OpsServiceApi.md#opsserviceping) | **GET** /ping | 
[**opsServicePutRepository**](OpsServiceApi.md#opsserviceputrepository) | **POST** /v1/repository | 
[**opsServicePutVariable**](OpsServiceApi.md#opsserviceputvariable) | **POST** /v1/variable | 
[**opsServiceRunOps**](OpsServiceApi.md#opsservicerunops) | **POST** /v1/runOps | 
[**opsServiceUpdateRepository**](OpsServiceApi.md#opsserviceupdaterepository) | **PUT** /v1/repository/{id} | 
[**opsServiceUpdateVariable**](OpsServiceApi.md#opsserviceupdatevariable) | **PUT** /v1/variable/{id} | 


# **opsServiceDelJob**
> Object opsServiceDelJob(id)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.opsServiceDelJob(id);
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

final api_instance = OpsServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.opsServiceDelRepository(id);
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

final api_instance = OpsServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.opsServiceDelVariable(id);
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

final api_instance = OpsServiceApi();
final body = ApiDescribeRepositoryReq(); // ApiDescribeRepositoryReq | 

try { 
    final result = api_instance.opsServiceDescribeRepository(body);
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
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.opsServiceGetJob(id);
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
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.opsServiceGetRepository(id);
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
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final id = id_example; // String | 

try { 
    final result = api_instance.opsServiceGetVariable(id);
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
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final repositoryID = repositoryID_example; // String | 
final offset = 56; // int | 
final limit = 56; // int | 

try { 
    final result = api_instance.opsServiceListJob(repositoryID, offset, limit);
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
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final offset = 56; // int | 
final limit = 56; // int | 

try { 
    final result = api_instance.opsServiceListRepository(offset, limit);
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
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final offset = 56; // int | 
final limit = 56; // int | 

try { 
    final result = api_instance.opsServiceListVariable(offset, limit);
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

# **opsServicePing**
> Object opsServicePing()



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();

try { 
    final result = api_instance.opsServicePing();
    print(result);
} catch (e) {
    print('Exception when calling OpsServiceApi->opsServicePing: $e\n');
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Object**](Object.md)

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

final api_instance = OpsServiceApi();
final body = ApiRepository(); // ApiRepository | 

try { 
    final result = api_instance.opsServicePutRepository(body);
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
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final body = ApiVariable(); // ApiVariable | 

try { 
    final result = api_instance.opsServicePutVariable(body);
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
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final body = ApiRunOpsReq(); // ApiRunOpsReq | 

try { 
    final result = api_instance.opsServiceRunOps(body);
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
> Object opsServiceUpdateRepository(id, body)



### Example 
```dart
import 'package:openapi/api.dart';

final api_instance = OpsServiceApi();
final id = id_example; // String | 
final body = InlineObject(); // InlineObject | 

try { 
    final result = api_instance.opsServiceUpdateRepository(id, body);
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

final api_instance = OpsServiceApi();
final id = id_example; // String | 
final body = InlineObject1(); // InlineObject1 | 

try { 
    final result = api_instance.opsServiceUpdateVariable(id, body);
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

[**Object**](Object.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

