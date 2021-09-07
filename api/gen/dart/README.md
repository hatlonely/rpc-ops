# openapi
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

This Dart package is automatically generated by the [OpenAPI Generator](https://openapi-generator.tech) project:

- API version: version not set
- Build package: org.openapitools.codegen.languages.DartClientCodegen

## Requirements

Dart 2.0 or later

## Installation & Usage

### Github
If this Dart package is published to Github, add the following dependency to your pubspec.yaml
```
dependencies:
  openapi:
    git: https://github.com/GIT_USER_ID/GIT_REPO_ID.git
```

### Local
To use the package in your local drive, add the following dependency to your pubspec.yaml
```
dependencies:
  openapi:
    path: /path/to/openapi
```

## Tests

TODO

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

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

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OpsServiceApi* | [**opsServiceDelJob**](doc//OpsServiceApi.md#opsservicedeljob) | **DELETE** /v1/job/{id} | 
*OpsServiceApi* | [**opsServiceDelRepository**](doc//OpsServiceApi.md#opsservicedelrepository) | **DELETE** /v1/repository/{id} | 
*OpsServiceApi* | [**opsServiceDelVariable**](doc//OpsServiceApi.md#opsservicedelvariable) | **DELETE** /v1/variable/{id} | 
*OpsServiceApi* | [**opsServiceDescribeRepository**](doc//OpsServiceApi.md#opsservicedescriberepository) | **POST** /v1/describeRepository | 
*OpsServiceApi* | [**opsServiceGetJob**](doc//OpsServiceApi.md#opsservicegetjob) | **GET** /v1/job/{id} | 
*OpsServiceApi* | [**opsServiceGetRepository**](doc//OpsServiceApi.md#opsservicegetrepository) | **GET** /v1/repository/{id} | 
*OpsServiceApi* | [**opsServiceGetVariable**](doc//OpsServiceApi.md#opsservicegetvariable) | **GET** /v1/variable/{id} | 
*OpsServiceApi* | [**opsServiceListJob**](doc//OpsServiceApi.md#opsservicelistjob) | **GET** /v1/job | 
*OpsServiceApi* | [**opsServiceListRepository**](doc//OpsServiceApi.md#opsservicelistrepository) | **GET** /v1/repository | 
*OpsServiceApi* | [**opsServiceListVariable**](doc//OpsServiceApi.md#opsservicelistvariable) | **GET** /v1/variable | 
*OpsServiceApi* | [**opsServicePutRepository**](doc//OpsServiceApi.md#opsserviceputrepository) | **POST** /v1/repository | 
*OpsServiceApi* | [**opsServicePutVariable**](doc//OpsServiceApi.md#opsserviceputvariable) | **POST** /v1/variable | 
*OpsServiceApi* | [**opsServiceRunOps**](doc//OpsServiceApi.md#opsservicerunops) | **POST** /v1/runOps | 
*OpsServiceApi* | [**opsServiceUpdateRepository**](doc//OpsServiceApi.md#opsserviceupdaterepository) | **PUT** /v1/repository/{id} | 
*OpsServiceApi* | [**opsServiceUpdateVariable**](doc//OpsServiceApi.md#opsserviceupdatevariable) | **PUT** /v1/variable/{id} | 


## Documentation For Models

 - [ApiDescribeRepositoryReq](doc//ApiDescribeRepositoryReq.md)
 - [ApiJob](doc//ApiJob.md)
 - [ApiListJobRes](doc//ApiListJobRes.md)
 - [ApiListRepositoryRes](doc//ApiListRepositoryRes.md)
 - [ApiListVariableRes](doc//ApiListVariableRes.md)
 - [ApiPlaybook](doc//ApiPlaybook.md)
 - [ApiRepository](doc//ApiRepository.md)
 - [ApiRepositoryID](doc//ApiRepositoryID.md)
 - [ApiRunOpsReq](doc//ApiRunOpsReq.md)
 - [ApiRunOpsRes](doc//ApiRunOpsRes.md)
 - [ApiVariable](doc//ApiVariable.md)
 - [ApiVariableID](doc//ApiVariableID.md)
 - [InlineObject](doc//InlineObject.md)
 - [InlineObject1](doc//InlineObject1.md)
 - [PlaybookEnv](doc//PlaybookEnv.md)
 - [PlaybookTask](doc//PlaybookTask.md)
 - [ProtobufAny](doc//ProtobufAny.md)
 - [RpcStatus](doc//RpcStatus.md)
 - [TaskArgs](doc//TaskArgs.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author



