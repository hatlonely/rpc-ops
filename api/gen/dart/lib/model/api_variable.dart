//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiVariable {
  /// Returns a new [ApiVariable] instance.
  ApiVariable({
    this.id,
    this.name,
    this.body,
    this.createAt,
    this.updateAt,
  });

  String id;

  String name;

  String body;

  String createAt;

  String updateAt;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiVariable &&
     other.id == id &&
     other.name == name &&
     other.body == body &&
     other.createAt == createAt &&
     other.updateAt == updateAt;

  @override
  int get hashCode =>
    (id == null ? 0 : id.hashCode) +
    (name == null ? 0 : name.hashCode) +
    (body == null ? 0 : body.hashCode) +
    (createAt == null ? 0 : createAt.hashCode) +
    (updateAt == null ? 0 : updateAt.hashCode);

  @override
  String toString() => 'ApiVariable[id=$id, name=$name, body=$body, createAt=$createAt, updateAt=$updateAt]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (id != null) {
      json[r'id'] = id;
    }
    if (name != null) {
      json[r'name'] = name;
    }
    if (body != null) {
      json[r'body'] = body;
    }
    if (createAt != null) {
      json[r'createAt'] = createAt;
    }
    if (updateAt != null) {
      json[r'updateAt'] = updateAt;
    }
    return json;
  }

  /// Returns a new [ApiVariable] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiVariable fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiVariable(
        id: json[r'id'],
        name: json[r'name'],
        body: json[r'body'],
        createAt: json[r'createAt'],
        updateAt: json[r'updateAt'],
    );

  static List<ApiVariable> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiVariable>[]
      : json.map((v) => ApiVariable.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiVariable> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiVariable>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiVariable.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiVariable-objects as value to a dart map
  static Map<String, List<ApiVariable>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiVariable>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiVariable.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

