//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiJob {
  /// Returns a new [ApiJob] instance.
  ApiJob({
    this.id,
    this.seq,
    this.state,
    this.repositoryID,
    this.variableID,
    this.version,
    this.createAt,
    this.updateAt,
    this.scheduleAt,
    this.elapseSeconds,
    this.output,
  });

  String id;

  int seq;

  String state;

  String repositoryID;

  String variableID;

  String version;

  String createAt;

  String updateAt;

  String scheduleAt;

  int elapseSeconds;

  String output;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiJob &&
     other.id == id &&
     other.seq == seq &&
     other.state == state &&
     other.repositoryID == repositoryID &&
     other.variableID == variableID &&
     other.version == version &&
     other.createAt == createAt &&
     other.updateAt == updateAt &&
     other.scheduleAt == scheduleAt &&
     other.elapseSeconds == elapseSeconds &&
     other.output == output;

  @override
  int get hashCode =>
    (id == null ? 0 : id.hashCode) +
    (seq == null ? 0 : seq.hashCode) +
    (state == null ? 0 : state.hashCode) +
    (repositoryID == null ? 0 : repositoryID.hashCode) +
    (variableID == null ? 0 : variableID.hashCode) +
    (version == null ? 0 : version.hashCode) +
    (createAt == null ? 0 : createAt.hashCode) +
    (updateAt == null ? 0 : updateAt.hashCode) +
    (scheduleAt == null ? 0 : scheduleAt.hashCode) +
    (elapseSeconds == null ? 0 : elapseSeconds.hashCode) +
    (output == null ? 0 : output.hashCode);

  @override
  String toString() => 'ApiJob[id=$id, seq=$seq, state=$state, repositoryID=$repositoryID, variableID=$variableID, version=$version, createAt=$createAt, updateAt=$updateAt, scheduleAt=$scheduleAt, elapseSeconds=$elapseSeconds, output=$output]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (id != null) {
      json[r'id'] = id;
    }
    if (seq != null) {
      json[r'seq'] = seq;
    }
    if (state != null) {
      json[r'state'] = state;
    }
    if (repositoryID != null) {
      json[r'repositoryID'] = repositoryID;
    }
    if (variableID != null) {
      json[r'variableID'] = variableID;
    }
    if (version != null) {
      json[r'version'] = version;
    }
    if (createAt != null) {
      json[r'createAt'] = createAt;
    }
    if (updateAt != null) {
      json[r'updateAt'] = updateAt;
    }
    if (scheduleAt != null) {
      json[r'scheduleAt'] = scheduleAt;
    }
    if (elapseSeconds != null) {
      json[r'elapseSeconds'] = elapseSeconds;
    }
    if (output != null) {
      json[r'output'] = output;
    }
    return json;
  }

  /// Returns a new [ApiJob] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiJob fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiJob(
        id: json[r'id'],
        seq: json[r'seq'],
        state: json[r'state'],
        repositoryID: json[r'repositoryID'],
        variableID: json[r'variableID'],
        version: json[r'version'],
        createAt: json[r'createAt'],
        updateAt: json[r'updateAt'],
        scheduleAt: json[r'scheduleAt'],
        elapseSeconds: json[r'elapseSeconds'],
        output: json[r'output'],
    );

  static List<ApiJob> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiJob>[]
      : json.map((v) => ApiJob.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiJob> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiJob>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiJob.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiJob-objects as value to a dart map
  static Map<String, List<ApiJob>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiJob>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiJob.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

