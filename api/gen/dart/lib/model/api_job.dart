part of openapi.api;

class ApiJob {
  
  String id = null;
  
  int seq = null;
  
  String state = null;
  
  String repositoryID = null;
  
  String variableID = null;
  
  String version = null;
  
  String createAt = null;
  
  String updateAt = null;
  
  String scheduleAt = null;
  
  int elapseSeconds = null;
  
  String output = null;
  ApiJob();

  @override
  String toString() {
    return 'ApiJob[id=$id, seq=$seq, state=$state, repositoryID=$repositoryID, variableID=$variableID, version=$version, createAt=$createAt, updateAt=$updateAt, scheduleAt=$scheduleAt, elapseSeconds=$elapseSeconds, output=$output, ]';
  }

  ApiJob.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    id = json['id'];
    seq = json['seq'];
    state = json['state'];
    repositoryID = json['repositoryID'];
    variableID = json['variableID'];
    version = json['version'];
    createAt = json['createAt'];
    updateAt = json['updateAt'];
    scheduleAt = json['scheduleAt'];
    elapseSeconds = json['elapseSeconds'];
    output = json['output'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (id != null)
      json['id'] = id;
    if (seq != null)
      json['seq'] = seq;
    if (state != null)
      json['state'] = state;
    if (repositoryID != null)
      json['repositoryID'] = repositoryID;
    if (variableID != null)
      json['variableID'] = variableID;
    if (version != null)
      json['version'] = version;
    if (createAt != null)
      json['createAt'] = createAt;
    if (updateAt != null)
      json['updateAt'] = updateAt;
    if (scheduleAt != null)
      json['scheduleAt'] = scheduleAt;
    if (elapseSeconds != null)
      json['elapseSeconds'] = elapseSeconds;
    if (output != null)
      json['output'] = output;
    return json;
  }

  static List<ApiJob> listFromJson(List<dynamic> json) {
    return json == null ? List<ApiJob>() : json.map((value) => ApiJob.fromJson(value)).toList();
  }

  static Map<String, ApiJob> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, ApiJob>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = ApiJob.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of ApiJob-objects as value to a dart map
  static Map<String, List<ApiJob>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<ApiJob>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = ApiJob.listFromJson(value);
       });
     }
     return map;
  }
}

