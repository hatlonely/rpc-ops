part of openapi.api;

class PlaybookEnv {
  
  String key = null;
  
  Map<String, String> val = {};
  PlaybookEnv();

  @override
  String toString() {
    return 'PlaybookEnv[key=$key, val=$val, ]';
  }

  PlaybookEnv.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    key = json['key'];
    val = (json['val'] == null) ?
      null :
      (json['val'] as Map).cast<String, String>();
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (key != null)
      json['key'] = key;
    if (val != null)
      json['val'] = val;
    return json;
  }

  static List<PlaybookEnv> listFromJson(List<dynamic> json) {
    return json == null ? List<PlaybookEnv>() : json.map((value) => PlaybookEnv.fromJson(value)).toList();
  }

  static Map<String, PlaybookEnv> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, PlaybookEnv>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = PlaybookEnv.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of PlaybookEnv-objects as value to a dart map
  static Map<String, List<PlaybookEnv>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<PlaybookEnv>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = PlaybookEnv.listFromJson(value);
       });
     }
     return map;
  }
}

