//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiDescribeRepositoryReq {
  /// Returns a new [ApiDescribeRepositoryReq] instance.
  ApiDescribeRepositoryReq({
    this.id,
    this.version,
  });

  String id;

  String version;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiDescribeRepositoryReq &&
     other.id == id &&
     other.version == version;

  @override
  int get hashCode =>
    (id == null ? 0 : id.hashCode) +
    (version == null ? 0 : version.hashCode);

  @override
  String toString() => 'ApiDescribeRepositoryReq[id=$id, version=$version]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (id != null) {
      json[r'id'] = id;
    }
    if (version != null) {
      json[r'version'] = version;
    }
    return json;
  }

  /// Returns a new [ApiDescribeRepositoryReq] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiDescribeRepositoryReq fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiDescribeRepositoryReq(
        id: json[r'id'],
        version: json[r'version'],
    );

  static List<ApiDescribeRepositoryReq> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiDescribeRepositoryReq>[]
      : json.map((v) => ApiDescribeRepositoryReq.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiDescribeRepositoryReq> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiDescribeRepositoryReq>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiDescribeRepositoryReq.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiDescribeRepositoryReq-objects as value to a dart map
  static Map<String, List<ApiDescribeRepositoryReq>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiDescribeRepositoryReq>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiDescribeRepositoryReq.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

