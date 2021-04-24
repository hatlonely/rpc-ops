//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiPlaybook {
  /// Returns a new [ApiPlaybook] instance.
  ApiPlaybook({
    this.name,
    this.envs = const [],
    this.tasks = const {},
  });

  String name;

  List<PlaybookEnv> envs;

  Map<String, PlaybookTask> tasks;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiPlaybook &&
     other.name == name &&
     other.envs == envs &&
     other.tasks == tasks;

  @override
  int get hashCode =>
    (name == null ? 0 : name.hashCode) +
    (envs == null ? 0 : envs.hashCode) +
    (tasks == null ? 0 : tasks.hashCode);

  @override
  String toString() => 'ApiPlaybook[name=$name, envs=$envs, tasks=$tasks]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (name != null) {
      json[r'name'] = name;
    }
    if (envs != null) {
      json[r'envs'] = envs;
    }
    if (tasks != null) {
      json[r'tasks'] = tasks;
    }
    return json;
  }

  /// Returns a new [ApiPlaybook] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiPlaybook fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiPlaybook(
        name: json[r'name'],
        envs: PlaybookEnv.listFromJson(json[r'envs']),
        tasks: json[r'tasks'] == null
          ? null
          : PlaybookTask.mapFromJson(json[r'tasks']),
    );

  static List<ApiPlaybook> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiPlaybook>[]
      : json.map((v) => ApiPlaybook.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiPlaybook> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiPlaybook>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiPlaybook.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiPlaybook-objects as value to a dart map
  static Map<String, List<ApiPlaybook>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiPlaybook>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiPlaybook.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

