//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiListVariableRes {
  /// Returns a new [ApiListVariableRes] instance.
  ApiListVariableRes({
    this.variables = const [],
  });

  List<ApiVariable> variables;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiListVariableRes &&
     other.variables == variables;

  @override
  int get hashCode =>
    (variables == null ? 0 : variables.hashCode);

  @override
  String toString() => 'ApiListVariableRes[variables=$variables]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (variables != null) {
      json[r'variables'] = variables;
    }
    return json;
  }

  /// Returns a new [ApiListVariableRes] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiListVariableRes fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiListVariableRes(
        variables: ApiVariable.listFromJson(json[r'variables']),
    );

  static List<ApiListVariableRes> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiListVariableRes>[]
      : json.map((v) => ApiListVariableRes.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiListVariableRes> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiListVariableRes>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiListVariableRes.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiListVariableRes-objects as value to a dart map
  static Map<String, List<ApiListVariableRes>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiListVariableRes>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiListVariableRes.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

