part of openapi.api;

class PlaybookTask {
  
  Map<String, TaskArgs> args = {};
  
  Map<String, String> const_ = {};
  
  List<String> step = [];
  PlaybookTask();

  @override
  String toString() {
    return 'PlaybookTask[args=$args, const_=$const_, step=$step, ]';
  }

  PlaybookTask.fromJson(Map<String, dynamic> json) {
    if (json == null) return;
    args = (json['args'] == null) ?
      null :
      TaskArgs.mapFromJson(json['args']);
    const_ = (json['const'] == null) ?
      null :
      (json['const'] as Map).cast<String, String>();
    step = (json['step'] == null) ?
      null :
      (json['step'] as List).cast<String>();
  }

  Map<String, dynamic> toJson() {
    Map <String, dynamic> json = {};
    if (args != null)
      json['args'] = args;
    if (const_ != null)
      json['const'] = const_;
    if (step != null)
      json['step'] = step;
    return json;
  }

  static List<PlaybookTask> listFromJson(List<dynamic> json) {
    return json == null ? List<PlaybookTask>() : json.map((value) => PlaybookTask.fromJson(value)).toList();
  }

  static Map<String, PlaybookTask> mapFromJson(Map<String, dynamic> json) {
    var map = Map<String, PlaybookTask>();
    if (json != null && json.isNotEmpty) {
      json.forEach((String key, dynamic value) => map[key] = PlaybookTask.fromJson(value));
    }
    return map;
  }

  // maps a json object with a list of PlaybookTask-objects as value to a dart map
  static Map<String, List<PlaybookTask>> mapListFromJson(Map<String, dynamic> json) {
    var map = Map<String, List<PlaybookTask>>();
     if (json != null && json.isNotEmpty) {
       json.forEach((String key, dynamic value) {
         map[key] = PlaybookTask.listFromJson(value);
       });
     }
     return map;
  }
}

