//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiRepositoryID {
  /// Returns a new [ApiRepositoryID] instance.
  ApiRepositoryID({
    this.id,
  });

  String id;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiRepositoryID &&
     other.id == id;

  @override
  int get hashCode =>
    (id == null ? 0 : id.hashCode);

  @override
  String toString() => 'ApiRepositoryID[id=$id]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (id != null) {
      json[r'id'] = id;
    }
    return json;
  }

  /// Returns a new [ApiRepositoryID] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiRepositoryID fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiRepositoryID(
        id: json[r'id'],
    );

  static List<ApiRepositoryID> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiRepositoryID>[]
      : json.map((v) => ApiRepositoryID.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiRepositoryID> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiRepositoryID>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiRepositoryID.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiRepositoryID-objects as value to a dart map
  static Map<String, List<ApiRepositoryID>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiRepositoryID>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiRepositoryID.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

