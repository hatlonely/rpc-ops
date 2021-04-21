part of openapi.api;

class ApiRunOpsReq {
  
  String repositoryID = null;
  
  String variableID = null;
  
  String version = null;
  
  String environment = null;
  
  String task = null;
  
  String args = null;
  ApiRunOpsReq();

  @override
  String toString() {
    return 'ApiRunOpsReq[repositoryID=$repositoryID, variableID=$variableID, version=$version, environment=$environment, task=$task, args=$args, ]';
  }

  ApiRunOpsReq.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    repositoryID = json['repositoryID'];
    variableID = json['variableID'];
    version = json['version'];
    environment = json['environment'];
    task = json['task'];
    args = json['args'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (repositoryID != null)
      json['repositoryID'] = repositoryID;
    if (variableID != null)
      json['variableID'] = variableID;
    if (version != null)
      json['version'] = version;
    if (environment != null)
      json['environment'] = environment;
    if (task != null)
      json['task'] = task;
    if (args != null)
      json['args'] = args;
    return json;
  }

  static List<ApiRunOpsReq> listFromJson(List<dynamic> json) {
    return json == null ? List<ApiRunOpsReq>() : json.map((value) => ApiRunOpsReq.fromJson(value)).toList();
  }

  static Map<String, ApiRunOpsReq> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ApiRunOpsReq>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ApiRunOpsReq.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ApiRunOpsReq-objects as value to a dart map
  static Map<String, List<ApiRunOpsReq>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ApiRunOpsReq>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ApiRunOpsReq.listFromJson(value);
       });
     }
     return map;
  }
}

