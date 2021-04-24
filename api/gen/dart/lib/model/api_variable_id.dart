//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiVariableID {
  /// Returns a new [ApiVariableID] instance.
  ApiVariableID({
    this.id,
  });

  String id;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiVariableID &&
     other.id == id;

  @override
  int get hashCode =>
    (id == null ? 0 : id.hashCode);

  @override
  String toString() => 'ApiVariableID[id=$id]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (id != null) {
      json[r'id'] = id;
    }
    return json;
  }

  /// Returns a new [ApiVariableID] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiVariableID fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiVariableID(
        id: json[r'id'],
    );

  static List<ApiVariableID> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiVariableID>[]
      : json.map((v) => ApiVariableID.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiVariableID> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiVariableID>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiVariableID.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiVariableID-objects as value to a dart map
  static Map<String, List<ApiVariableID>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiVariableID>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiVariableID.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

