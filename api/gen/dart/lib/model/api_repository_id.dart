part of openapi.api;

class ApiRepositoryID {
  
  String id = null;
  ApiRepositoryID();

  @override
  String toString() {
    return 'ApiRepositoryID[id=$id, ]';
  }

  ApiRepositoryID.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    id = json['id'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (id != null)
      json['id'] = id;
    return json;
  }

  static List<ApiRepositoryID> listFromJson(List<dynamic> json) {
    return json == null ? List<ApiRepositoryID>() : json.map((value) => ApiRepositoryID.fromJson(value)).toList();
  }

  static Map<String, ApiRepositoryID> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ApiRepositoryID>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ApiRepositoryID.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ApiRepositoryID-objects as value to a dart map
  static Map<String, List<ApiRepositoryID>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ApiRepositoryID>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ApiRepositoryID.listFromJson(value);
       });
     }
     return map;
  }
}

