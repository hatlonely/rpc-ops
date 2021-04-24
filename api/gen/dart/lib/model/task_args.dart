//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class TaskArgs {
  /// Returns a new [TaskArgs] instance.
  TaskArgs({
    this.type,
    this.default_,
    this.validation,
  });

  String type;

  String default_;

  String validation;

  @override
  bool operator ==(Object other) => identical(this, other) || other is TaskArgs &&
     other.type == type &&
     other.default_ == default_ &&
     other.validation == validation;

  @override
  int get hashCode =>
    (type == null ? 0 : type.hashCode) +
    (default_ == null ? 0 : default_.hashCode) +
    (validation == null ? 0 : validation.hashCode);

  @override
  String toString() => 'TaskArgs[type=$type, default_=$default_, validation=$validation]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (type != null) {
      json[r'type'] = type;
    }
    if (default_ != null) {
      json[r'default'] = default_;
    }
    if (validation != null) {
      json[r'validation'] = validation;
    }
    return json;
  }

  /// Returns a new [TaskArgs] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static TaskArgs fromJson(Map<String, dynamic> json) => json == null
    ? null
    : TaskArgs(
        type: json[r'type'],
        default_: json[r'default'],
        validation: json[r'validation'],
    );

  static List<TaskArgs> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <TaskArgs>[]
      : json.map((v) => TaskArgs.fromJson(v)).toList(growable: true == growable);

  static Map<String, TaskArgs> mapFromJson(Map<String, dynamic> json) {
    final map = <String, TaskArgs>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = TaskArgs.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of TaskArgs-objects as value to a dart map
  static Map<String, List<TaskArgs>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<TaskArgs>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = TaskArgs.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

