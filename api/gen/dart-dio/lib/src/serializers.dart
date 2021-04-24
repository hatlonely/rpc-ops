//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_collection/built_collection.dart';
import 'package:built_value/json_object.dart';
import 'package:built_value/serializer.dart';
import 'package:built_value/standard_json_plugin.dart';
import 'package:built_value/iso_8601_date_time_serializer.dart';

import 'package:openapi/src/model/api_describe_repository_req.dart';
import 'package:openapi/src/model/api_job.dart';
import 'package:openapi/src/model/api_list_job_res.dart';
import 'package:openapi/src/model/api_list_repository_res.dart';
import 'package:openapi/src/model/api_list_variable_res.dart';
import 'package:openapi/src/model/api_playbook.dart';
import 'package:openapi/src/model/api_repository.dart';
import 'package:openapi/src/model/api_repository_id.dart';
import 'package:openapi/src/model/api_run_ops_req.dart';
import 'package:openapi/src/model/api_run_ops_res.dart';
import 'package:openapi/src/model/api_variable.dart';
import 'package:openapi/src/model/api_variable_id.dart';
import 'package:openapi/src/model/inline_object.dart';
import 'package:openapi/src/model/inline_object1.dart';
import 'package:openapi/src/model/playbook_env.dart';
import 'package:openapi/src/model/playbook_task.dart';
import 'package:openapi/src/model/protobuf_any.dart';
import 'package:openapi/src/model/rpc_status.dart';
import 'package:openapi/src/model/task_args.dart';

part 'serializers.g.dart';

@SerializersFor([
  ApiDescribeRepositoryReq,
  ApiJob,
  ApiListJobRes,
  ApiListRepositoryRes,
  ApiListVariableRes,
  ApiPlaybook,
  ApiRepository,
  ApiRepositoryID,
  ApiRunOpsReq,
  ApiRunOpsRes,
  ApiVariable,
  ApiVariableID,
  InlineObject,
  InlineObject1,
  PlaybookEnv,
  PlaybookTask,
  ProtobufAny,
  RpcStatus,
  TaskArgs,
])
Serializers serializers = (_$serializers.toBuilder()
      ..add(Iso8601DateTimeSerializer()))
    .build();

Serializers standardSerializers =
    (serializers.toBuilder()..addPlugin(StandardJsonPlugin())).build();
