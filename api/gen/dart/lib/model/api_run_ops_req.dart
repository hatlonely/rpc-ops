//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiRunOpsReq {
  /// Returns a new [ApiRunOpsReq] instance.
  ApiRunOpsReq({
    this.repositoryID,
    this.variableID,
    this.version,
    this.environment,
    this.task,
    this.args,
  });

  String repositoryID;

  String variableID;

  String version;

  String environment;

  String task;

  String args;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiRunOpsReq &&
     other.repositoryID == repositoryID &&
     other.variableID == variableID &&
     other.version == version &&
     other.environment == environment &&
     other.task == task &&
     other.args == args;

  @override
  int get hashCode =>
    (repositoryID == null ? 0 : repositoryID.hashCode) +
    (variableID == null ? 0 : variableID.hashCode) +
    (version == null ? 0 : version.hashCode) +
    (environment == null ? 0 : environment.hashCode) +
    (task == null ? 0 : task.hashCode) +
    (args == null ? 0 : args.hashCode);

  @override
  String toString() => 'ApiRunOpsReq[repositoryID=$repositoryID, variableID=$variableID, version=$version, environment=$environment, task=$task, args=$args]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (repositoryID != null) {
      json[r'repositoryID'] = repositoryID;
    }
    if (variableID != null) {
      json[r'variableID'] = variableID;
    }
    if (version != null) {
      json[r'version'] = version;
    }
    if (environment != null) {
      json[r'environment'] = environment;
    }
    if (task != null) {
      json[r'task'] = task;
    }
    if (args != null) {
      json[r'args'] = args;
    }
    return json;
  }

  /// Returns a new [ApiRunOpsReq] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiRunOpsReq fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiRunOpsReq(
        repositoryID: json[r'repositoryID'],
        variableID: json[r'variableID'],
        version: json[r'version'],
        environment: json[r'environment'],
        task: json[r'task'],
        args: json[r'args'],
    );

  static List<ApiRunOpsReq> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiRunOpsReq>[]
      : json.map((v) => ApiRunOpsReq.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiRunOpsReq> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiRunOpsReq>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiRunOpsReq.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiRunOpsReq-objects as value to a dart map
  static Map<String, List<ApiRunOpsReq>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiRunOpsReq>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiRunOpsReq.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

