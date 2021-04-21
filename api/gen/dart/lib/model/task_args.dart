part of openapi.api;

class TaskArgs {
  
  String type = null;
  
  String default_ = null;
  
  String validation = null;
  TaskArgs();

  @override
  String toString() {
    return 'TaskArgs[type=$type, default_=$default_, validation=$validation, ]';
  }

  TaskArgs.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    type = json['type'];
    default_ = json['default'];
    validation = json['validation'];
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (type != null)
      json['type'] = type;
    if (default_ != null)
      json['default'] = default_;
    if (validation != null)
      json['validation'] = validation;
    return json;
  }

  static List<TaskArgs> listFromJson(List<dynamic> json) {
    return json == null ? List<TaskArgs>() : json.map((value) => TaskArgs.fromJson(value)).toList();
  }

  static Map<String, TaskArgs> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, TaskArgs>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = TaskArgs.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of TaskArgs-objects as value to a dart map
  static Map<String, List<TaskArgs>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<TaskArgs>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = TaskArgs.listFromJson(value);
       });
     }
     return map;
  }
}

