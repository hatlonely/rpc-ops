//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class PlaybookTask {
  /// Returns a new [PlaybookTask] instance.
  PlaybookTask({
    this.args = const {},
    this.const_ = const {},
    this.step = const [],
  });

  Map<String, TaskArgs> args;

  Map<String, String> const_;

  List<String> step;

  @override
  bool operator ==(Object other) => identical(this, other) || other is PlaybookTask &&
     other.args == args &&
     other.const_ == const_ &&
     other.step == step;

  @override
  int get hashCode =>
    (args == null ? 0 : args.hashCode) +
    (const_ == null ? 0 : const_.hashCode) +
    (step == null ? 0 : step.hashCode);

  @override
  String toString() => 'PlaybookTask[args=$args, const_=$const_, step=$step]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (args != null) {
      json[r'args'] = args;
    }
    if (const_ != null) {
      json[r'const'] = const_;
    }
    if (step != null) {
      json[r'step'] = step;
    }
    return json;
  }

  /// Returns a new [PlaybookTask] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static PlaybookTask fromJson(Map<String, dynamic> json) => json == null
    ? null
    : PlaybookTask(
        args: json[r'args'] == null
          ? null
          : TaskArgs.mapFromJson(json[r'args']),
        const_: json[r'const'] == null ?
          null :
          (json[r'const'] as Map).cast<String, String>(),
        step: json[r'step'] == null
          ? null
          : (json[r'step'] as List).cast<String>(),
    );

  static List<PlaybookTask> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <PlaybookTask>[]
      : json.map((v) => PlaybookTask.fromJson(v)).toList(growable: true == growable);

  static Map<String, PlaybookTask> mapFromJson(Map<String, dynamic> json) {
    final map = <String, PlaybookTask>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = PlaybookTask.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of PlaybookTask-objects as value to a dart map
  static Map<String, List<PlaybookTask>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<PlaybookTask>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = PlaybookTask.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

