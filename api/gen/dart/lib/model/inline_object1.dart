//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class InlineObject1 {
  /// Returns a new [InlineObject1] instance.
  InlineObject1({
    this.name,
    this.body,
    this.createAt,
    this.updateAt,
  });

  String name;

  String body;

  String createAt;

  String updateAt;

  @override
  bool operator ==(Object other) => identical(this, other) || other is InlineObject1 &&
     other.name == name &&
     other.body == body &&
     other.createAt == createAt &&
     other.updateAt == updateAt;

  @override
  int get hashCode =>
    (name == null ? 0 : name.hashCode) +
    (body == null ? 0 : body.hashCode) +
    (createAt == null ? 0 : createAt.hashCode) +
    (updateAt == null ? 0 : updateAt.hashCode);

  @override
  String toString() => 'InlineObject1[name=$name, body=$body, createAt=$createAt, updateAt=$updateAt]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
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

  /// Returns a new [InlineObject1] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static InlineObject1 fromJson(Map<String, dynamic> json) => json == null
    ? null
    : InlineObject1(
        name: json[r'name'],
        body: json[r'body'],
        createAt: json[r'createAt'],
        updateAt: json[r'updateAt'],
    );

  static List<InlineObject1> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <InlineObject1>[]
      : json.map((v) => InlineObject1.fromJson(v)).toList(growable: true == growable);

  static Map<String, InlineObject1> mapFromJson(Map<String, dynamic> json) {
    final map = <String, InlineObject1>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = InlineObject1.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of InlineObject1-objects as value to a dart map
  static Map<String, List<InlineObject1>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<InlineObject1>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = InlineObject1.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

