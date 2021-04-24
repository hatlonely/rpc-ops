//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: lines_longer_than_80_chars

part of openapi.api;

class ApiRepository {
  /// Returns a new [ApiRepository] instance.
  ApiRepository({
    this.id,
    this.username,
    this.password,
    this.endpoint,
    this.name,
    this.playbook,
    this.createAt,
    this.updateAt,
  });

  String id;

  String username;

  String password;

  String endpoint;

  String name;

  String playbook;

  String createAt;

  String updateAt;

  @override
  bool operator ==(Object other) => identical(this, other) || other is ApiRepository &&
     other.id == id &&
     other.username == username &&
     other.password == password &&
     other.endpoint == endpoint &&
     other.name == name &&
     other.playbook == playbook &&
     other.createAt == createAt &&
     other.updateAt == updateAt;

  @override
  int get hashCode =>
    (id == null ? 0 : id.hashCode) +
    (username == null ? 0 : username.hashCode) +
    (password == null ? 0 : password.hashCode) +
    (endpoint == null ? 0 : endpoint.hashCode) +
    (name == null ? 0 : name.hashCode) +
    (playbook == null ? 0 : playbook.hashCode) +
    (createAt == null ? 0 : createAt.hashCode) +
    (updateAt == null ? 0 : updateAt.hashCode);

  @override
  String toString() => 'ApiRepository[id=$id, username=$username, password=$password, endpoint=$endpoint, name=$name, playbook=$playbook, createAt=$createAt, updateAt=$updateAt]';

  Map<String, dynamic> toJson() {
    final json = <String, dynamic>{};
    if (id != null) {
      json[r'id'] = id;
    }
    if (username != null) {
      json[r'username'] = username;
    }
    if (password != null) {
      json[r'password'] = password;
    }
    if (endpoint != null) {
      json[r'endpoint'] = endpoint;
    }
    if (name != null) {
      json[r'name'] = name;
    }
    if (playbook != null) {
      json[r'playbook'] = playbook;
    }
    if (createAt != null) {
      json[r'createAt'] = createAt;
    }
    if (updateAt != null) {
      json[r'updateAt'] = updateAt;
    }
    return json;
  }

  /// Returns a new [ApiRepository] instance and imports its values from
  /// [json] if it's non-null, null if [json] is null.
  static ApiRepository fromJson(Map<String, dynamic> json) => json == null
    ? null
    : ApiRepository(
        id: json[r'id'],
        username: json[r'username'],
        password: json[r'password'],
        endpoint: json[r'endpoint'],
        name: json[r'name'],
        playbook: json[r'playbook'],
        createAt: json[r'createAt'],
        updateAt: json[r'updateAt'],
    );

  static List<ApiRepository> listFromJson(List<dynamic> json, {bool emptyIsNull, bool growable,}) =>
    json == null || json.isEmpty
      ? true == emptyIsNull ? null : <ApiRepository>[]
      : json.map((v) => ApiRepository.fromJson(v)).toList(growable: true == growable);

  static Map<String, ApiRepository> mapFromJson(Map<String, dynamic> json) {
    final map = <String, ApiRepository>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) => map[key] = ApiRepository.fromJson(v));
    }
    return map;
  }

  // maps a json object with a list of ApiRepository-objects as value to a dart map
  static Map<String, List<ApiRepository>> mapListFromJson(Map<String, dynamic> json, {bool emptyIsNull, bool growable,}) {
    final map = <String, List<ApiRepository>>{};
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic v) {
        map[key] = ApiRepository.listFromJson(v, emptyIsNull: emptyIsNull, growable: growable);
      });
    }
    return map;
  }
}

