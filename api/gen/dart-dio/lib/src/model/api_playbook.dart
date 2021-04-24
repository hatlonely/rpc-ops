//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:api/src/model/playbook_task.dart';
import 'package:built_collection/built_collection.dart';
import 'package:api/src/model/playbook_env.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_playbook.g.dart';



abstract class ApiPlaybook implements Built<ApiPlaybook, ApiPlaybookBuilder> {
    @BuiltValueField(wireName: r'name')
    String? get name;

    @BuiltValueField(wireName: r'envs')
    BuiltList<PlaybookEnv>? get envs;

    @BuiltValueField(wireName: r'tasks')
    BuiltMap<String, PlaybookTask>? get tasks;

    ApiPlaybook._();

    static void _initializeBuilder(ApiPlaybookBuilder b) => b;

    factory ApiPlaybook([void updates(ApiPlaybookBuilder b)]) = _$ApiPlaybook;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiPlaybook> get serializer => _$ApiPlaybookSerializer();
}

class _$ApiPlaybookSerializer implements StructuredSerializer<ApiPlaybook> {
    @override
    final Iterable<Type> types = const [ApiPlaybook, _$ApiPlaybook];

    @override
    final String wireName = r'ApiPlaybook';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiPlaybook object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.name != null) {
            result
                ..add(r'name')
                ..add(serializers.serialize(object.name,
                    specifiedType: const FullType(String)));
        }
        if (object.envs != null) {
            result
                ..add(r'envs')
                ..add(serializers.serialize(object.envs,
                    specifiedType: const FullType(BuiltList, [FullType(PlaybookEnv)])));
        }
        if (object.tasks != null) {
            result
                ..add(r'tasks')
                ..add(serializers.serialize(object.tasks,
                    specifiedType: const FullType(BuiltMap, [FullType(String), FullType(PlaybookTask)])));
        }
        return result;
    }

    @override
    ApiPlaybook deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiPlaybookBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'name':
                    result.name = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'envs':
                    result.envs.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltList, [FullType(PlaybookEnv)])) as BuiltList<PlaybookEnv>);
                    break;
                case r'tasks':
                    result.tasks.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltMap, [FullType(String), FullType(PlaybookTask)])) as BuiltMap<String, PlaybookTask>);
                    break;
            }
        }
        return result.build();
    }
}

