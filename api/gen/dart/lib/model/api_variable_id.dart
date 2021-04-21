part of openapi.api;

class ApiVariableID {
  
  String id = null;
  ApiVariableID();

  @override
  String toString() {
    return 'ApiVariableID[id=$id, ]';
  }

  ApiVariableID.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    id = json['id'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (id != null)
      json['id'] = id;
    return json;
  }

  static List<ApiVariableID> listFromJson(List<dynamic> json) {
    return json == null ? List<ApiVariableID>() : json.map((value) => ApiVariableID.fromJson(value)).toList();
  }

  static Map<String, ApiVariableID> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ApiVariableID>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ApiVariableID.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ApiVariableID-objects as value to a dart map
  static Map<String, List<ApiVariableID>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ApiVariableID>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ApiVariableID.listFromJson(value);
       });
     }
     return map;
  }
}

