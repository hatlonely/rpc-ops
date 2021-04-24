//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiListJobRes {
  /// Returns a new [ApiListJobRes] instance.
  ApiListJobRes({
    this.jobs = const [],
  });

  List<ApiJob> jobs;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiListJobRes &&
     other.jobs == jobs;

  @override
  int get hashCode =>
    (jobs == null ? 0 : jobs.hashCode);

  @override
  String toString() => 'ApiListJobRes[jobs=$jobs]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (jobs != null) {
      json[r'jobs'] = jobs;
    }
    return json;
  }

  /// Returns a new [ApiListJobRes] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiListJobRes fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiListJobRes(
        jobs: ApiJob.listFromJson(json[r'jobs']),
    );

  static List<ApiListJobRes> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiListJobRes>[]
      : json.map((v) => ApiListJobRes.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiListJobRes> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiListJobRes>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiListJobRes.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiListJobRes-objects as value to a dart map
  static Map<String, List<ApiListJobRes>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiListJobRes>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiListJobRes.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

