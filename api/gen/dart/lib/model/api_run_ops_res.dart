part of openapi.api;

class ApiRunOpsRes {
  
  String jobID = null;
  ApiRunOpsRes();

  @override
  String toString() {
    return 'ApiRunOpsRes[jobID=$jobID, ]';
  }

  ApiRunOpsRes.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    jobID = json['jobID'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (jobID != null)
      json['jobID'] = jobID;
    return json;
  }

  static List<ApiRunOpsRes> listFromJson(List<dynamic> json) {
    return json == null ? List<ApiRunOpsRes>() : json.map((value) => ApiRunOpsRes.fromJson(value)).toList();
  }

  static Map<String, ApiRunOpsRes> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ApiRunOpsRes>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ApiRunOpsRes.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ApiRunOpsRes-objects as value to a dart map
  static Map<String, List<ApiRunOpsRes>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ApiRunOpsRes>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ApiRunOpsRes.listFromJson(value);
       });
     }
     return map;
  }
}

