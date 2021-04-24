//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

library openapi.api;

import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart';
import 'package:intl/intl.dart';

import 'package:meta/meta.dart';

part 'api_client.dart';
part 'api_helper.dart';
part 'api_exception.dart';
part 'auth/authentication.dart';
part 'auth/api_key_auth.dart';
part 'auth/oauth.dart';
part 'auth/http_basic_auth.dart';
part 'auth/http_bearer_auth.dart';

part 'api/ops_service_api.dart';

part 'model/api_describe_repository_req.dart';
part 'model/api_job.dart';
part 'model/api_list_job_res.dart';
part 'model/api_list_repository_res.dart';
part 'model/api_list_variable_res.dart';
part 'model/api_playbook.dart';
part 'model/api_repository.dart';
part 'model/api_repository_id.dart';
part 'model/api_run_ops_req.dart';
part 'model/api_run_ops_res.dart';
part 'model/api_variable.dart';
part 'model/api_variable_id.dart';
part 'model/inline_object.dart';
part 'model/inline_object1.dart';
part 'model/playbook_env.dart';
part 'model/playbook_task.dart';
part 'model/protobuf_any.dart';
part 'model/rpc_status.dart';
part 'model/task_args.dart';


const _delimiters = {'csv': ',', 'ssv': ' ', 'tsv': '\t', 'pipes': '|'};
const _dateEpochMarker = 'epoch';
final _dateFormatter = DateFormat('yyyy-MM-dd');
final _regList = RegExp(r'^List<(.*)>$');
final _regSet = RegExp(r'^Set<(.*)>$');
final _regMap = RegExp(r'^Map<String,(.*)>$');

ApiClient defaultApiClient = ApiClient();
