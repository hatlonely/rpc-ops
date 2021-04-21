part of openapi.api;

class ApiRepository {
  
  String id = null;
  
  String username = null;
  
  String password = null;
  
  String endpoint = null;
  
  String name = null;
  
  String playbook = null;
  
  String createAt = null;
  
  String updateAt = null;
  ApiRepository();

  @override
  String toString() {
    return 'ApiRepository[id=$id, username=$username, password=$password, endpoint=$endpoint, name=$name, playbook=$playbook, createAt=$createAt, updateAt=$updateAt, ]';
  }

  ApiRepository.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    id = json['id'];
    username = json['username'];
    password = json['password'];
    endpoint = json['endpoint'];
    name = json['name'];
    playbook = json['playbook'];
    createAt = json['createAt'];
    updateAt = json['updateAt'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (id != null)
      json['id'] = id;
    if (username != null)
      json['username'] = username;
    if (password != null)
      json['password'] = password;
    if (endpoint != null)
      json['endpoint'] = endpoint;
    if (name != null)
      json['name'] = name;
    if (playbook != null)
      json['playbook'] = playbook;
    if (createAt != null)
      json['createAt'] = createAt;
    if (updateAt != null)
      json['updateAt'] = updateAt;
    return json;
  }

  static List<ApiRepository> listFromJson(List<dynamic> json) {
    return json == null ? List<ApiRepository>() : json.map((value) => ApiRepository.fromJson(value)).toList();
  }

  static Map<String, ApiRepository> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ApiRepository>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ApiRepository.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ApiRepository-objects as value to a dart map
  static Map<String, List<ApiRepository>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ApiRepository>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ApiRepository.listFromJson(value);
       });
     }
     return map;
  }
}

