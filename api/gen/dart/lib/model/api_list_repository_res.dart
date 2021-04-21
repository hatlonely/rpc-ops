part of openapi.api;

class ApiListRepositoryRes {
  
  List<ApiRepository> repositories = [];
  ApiListRepositoryRes();

  @override
  String toString() {
    return 'ApiListRepositoryRes[repositories=$repositories, ]';
  }

  ApiListRepositoryRes.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    repositories = (json['repositories'] == null) ?
      null :
      ApiRepository.listFromJson(json['repositories']);
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (repositories != null)
      json['repositories'] = repositories;
    return json;
  }

  static List<ApiListRepositoryRes> listFromJson(List<dynamic> json) {
    return json == null ? List<ApiListRepositoryRes>() : json.map((value) => ApiListRepositoryRes.fromJson(value)).toList();
  }

  static Map<String, ApiListRepositoryRes> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ApiListRepositoryRes>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ApiListRepositoryRes.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ApiListRepositoryRes-objects as value to a dart map
  static Map<String, List<ApiListRepositoryRes>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ApiListRepositoryRes>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ApiListRepositoryRes.listFromJson(value);
       });
     }
     return map;
  }
}

