//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiListRepositoryRes {
  /// Returns a new [ApiListRepositoryRes] instance.
  ApiListRepositoryRes({
    this.repositories = const [],
  });

  List<ApiRepository> repositories;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiListRepositoryRes &&
     other.repositories == repositories;

  @override
  int get hashCode =>
    (repositories == null ? 0 : repositories.hashCode);

  @override
  String toString() => 'ApiListRepositoryRes[repositories=$repositories]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (repositories != null) {
      json[r'repositories'] = repositories;
    }
    return json;
  }

  /// Returns a new [ApiListRepositoryRes] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiListRepositoryRes fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiListRepositoryRes(
        repositories: ApiRepository.listFromJson(json[r'repositories']),
    );

  static List<ApiListRepositoryRes> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiListRepositoryRes>[]
      : json.map((v) => ApiListRepositoryRes.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiListRepositoryRes> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiListRepositoryRes>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiListRepositoryRes.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiListRepositoryRes-objects as value to a dart map
  static Map<String, List<ApiListRepositoryRes>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiListRepositoryRes>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiListRepositoryRes.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

