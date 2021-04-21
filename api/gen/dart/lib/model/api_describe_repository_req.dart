part of openapi.api;

class ApiDescribeRepositoryReq {
  
  String id = null;
  
  String version = null;
  ApiDescribeRepositoryReq();

  @override
  String toString() {
    return 'ApiDescribeRepositoryReq[id=$id, version=$version, ]';
  }

  ApiDescribeRepositoryReq.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    id = json['id'];
    version = json['version'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (id != null)
      json['id'] = id;
    if (version != null)
      json['version'] = version;
    return json;
  }

  static List<ApiDescribeRepositoryReq> listFromJson(List<dynamic> json) {
    return json == null ? List<ApiDescribeRepositoryReq>() : json.map((value) => ApiDescribeRepositoryReq.fromJson(value)).toList();
  }

  static Map<String, ApiDescribeRepositoryReq> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ApiDescribeRepositoryReq>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ApiDescribeRepositoryReq.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ApiDescribeRepositoryReq-objects as value to a dart map
  static Map<String, List<ApiDescribeRepositoryReq>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ApiDescribeRepositoryReq>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ApiDescribeRepositoryReq.listFromJson(value);
       });
     }
     return map;
  }
}

