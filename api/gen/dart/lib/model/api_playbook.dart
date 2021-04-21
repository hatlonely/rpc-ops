part of openapi.api;

class ApiPlaybook {
  
  String name = null;
  
  List<PlaybookEnv> envs = [];
  
  Map<String, PlaybookTask> tasks = {};
  ApiPlaybook();

  @override
  String toString() {
    return 'ApiPlaybook[name=$name, envs=$envs, tasks=$tasks, ]';
  }

  ApiPlaybook.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    name = json['name'];
    envs = (json['envs'] == null) ?
      null :
      PlaybookEnv.listFromJson(json['envs']);
    tasks = (json['tasks'] == null) ?
      null :
      PlaybookTask.mapFromJson(json['tasks']);
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (name != null)
      json['name'] = name;
    if (envs != null)
      json['envs'] = envs;
    if (tasks != null)
      json['tasks'] = tasks;
    return json;
  }

  static List<ApiPlaybook> listFromJson(List<dynamic> json) {
    return json == null ? List<ApiPlaybook>() : json.map((value) => ApiPlaybook.fromJson(value)).toList();
  }

  static Map<String, ApiPlaybook> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ApiPlaybook>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ApiPlaybook.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ApiPlaybook-objects as value to a dart map
  static Map<String, List<ApiPlaybook>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ApiPlaybook>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ApiPlaybook.listFromJson(value);
       });
     }
     return map;
  }
}

