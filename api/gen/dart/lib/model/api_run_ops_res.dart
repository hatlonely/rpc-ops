//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiRunOpsRes {
  /// Returns a new [ApiRunOpsRes] instance.
  ApiRunOpsRes({
    this.jobID,
  });

  String jobID;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiRunOpsRes &&
     other.jobID == jobID;

  @override
  int get hashCode =>
    (jobID == null ? 0 : jobID.hashCode);

  @override
  String toString() => 'ApiRunOpsRes[jobID=$jobID]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (jobID != null) {
      json[r'jobID'] = jobID;
    }
    return json;
  }

  /// Returns a new [ApiRunOpsRes] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiRunOpsRes fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiRunOpsRes(
        jobID: json[r'jobID'],
    );

  static List<ApiRunOpsRes> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiRunOpsRes>[]
      : json.map((v) => ApiRunOpsRes.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiRunOpsRes> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiRunOpsRes>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiRunOpsRes.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiRunOpsRes-objects as value to a dart map
  static Map<String, List<ApiRunOpsRes>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiRunOpsRes>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiRunOpsRes.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

