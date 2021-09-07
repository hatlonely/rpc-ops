import 'package:test/test.dart';
import 'package:api/api.dart';


/// tests for OpsServiceApi
void main() {
  final instance = Api().getOpsServiceApi();

  group(OpsServiceApi, () {
    //Future<JsonObject> opsServiceDelJob(String id) async
    test('test opsServiceDelJob', () async {
      // TODO
    });

    //Future<JsonObject> opsServiceDelRepository(String id) async
    test('test opsServiceDelRepository', () async {
      // TODO
    });

    //Future<JsonObject> opsServiceDelVariable(String id) async
    test('test opsServiceDelVariable', () async {
      // TODO
    });

    //Future<ApiPlaybook> opsServiceDescribeRepository(ApiDescribeRepositoryReq body) async
    test('test opsServiceDescribeRepository', () async {
      // TODO
    });

    //Future<ApiJob> opsServiceGetJob(String id) async
    test('test opsServiceGetJob', () async {
      // TODO
    });

    //Future<ApiRepository> opsServiceGetRepository(String id) async
    test('test opsServiceGetRepository', () async {
      // TODO
    });

    //Future<ApiVariable> opsServiceGetVariable(String id) async
    test('test opsServiceGetVariable', () async {
      // TODO
    });

    //Future<ApiListJobRes> opsServiceListJob({ String repositoryID, int offset, int limit }) async
    test('test opsServiceListJob', () async {
      // TODO
    });

    //Future<ApiListRepositoryRes> opsServiceListRepository({ int offset, int limit }) async
    test('test opsServiceListRepository', () async {
      // TODO
    });

    //Future<ApiListVariableRes> opsServiceListVariable({ int offset, int limit }) async
    test('test opsServiceListVariable', () async {
      // TODO
    });

    //Future<JsonObject> opsServicePing() async
    test('test opsServicePing', () async {
      // TODO
    });

    //Future<ApiRepositoryID> opsServicePutRepository(ApiRepository body) async
    test('test opsServicePutRepository', () async {
      // TODO
    });

    //Future<ApiVariableID> opsServicePutVariable(ApiVariable body) async
    test('test opsServicePutVariable', () async {
      // TODO
    });

    //Future<ApiRunOpsRes> opsServiceRunOps(ApiRunOpsReq body) async
    test('test opsServiceRunOps', () async {
      // TODO
    });

    //Future<JsonObject> opsServiceUpdateRepository(String id, InlineObject body) async
    test('test opsServiceUpdateRepository', () async {
      // TODO
    });

    //Future<JsonObject> opsServiceUpdateVariable(String id, InlineObject1 body) async
    test('test opsServiceUpdateVariable', () async {
      // TODO
    });

  });
}
