//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class PlaybookEnv {
  /// Returns a new [PlaybookEnv] instance.
  PlaybookEnv({
    this.key,
    this.val = const {},
  });

  String key;

  Map<String, String> val;

  @override
  bool operator ==(Object other) => identical(this, other) || other is PlaybookEnv &&
     other.key == key &&
     other.val == val;

  @override
  int get hashCode =>
    (key == null ? 0 : key.hashCode) +
    (val == null ? 0 : val.hashCode);

  @override
  String toString() => 'PlaybookEnv[key=$key, val=$val]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (key != null) {
      json[r'key'] = key;
    }
    if (val != null) {
      json[r'val'] = val;
    }
    return json;
  }

  /// Returns a new [PlaybookEnv] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static PlaybookEnv fromJson(Map<String, dynamic> json) => json == null
    ? null
    : PlaybookEnv(
        key: json[r'key'],
        val: json[r'val'] == null ?
          null :
          (json[r'val'] as Map).cast<String, String>(),
    );

  static List<PlaybookEnv> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <PlaybookEnv>[]
      : json.map((v) => PlaybookEnv.fromJson(v)).toList(growable: true == growable);

  static Map<String, PlaybookEnv> mapFromJson(Map<String, dynamic> json) {
    final map = <String, PlaybookEnv>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = PlaybookEnv.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of PlaybookEnv-objects as value to a dart map
  static Map<String, List<PlaybookEnv>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<PlaybookEnv>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = PlaybookEnv.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

