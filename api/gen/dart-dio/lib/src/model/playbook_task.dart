//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_collection/built_collection.dart';
import 'package:api/src/model/task_args.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'playbook_task.g.dart';



abstract class PlaybookTask implements Built<PlaybookTask, PlaybookTaskBuilder> {
    @BuiltValueField(wireName: r'args')
    BuiltMap<String, TaskArgs>? get args;

    @BuiltValueField(wireName: r'const')
    BuiltMap<String, String>? get const_;

    @BuiltValueField(wireName: r'step')
    BuiltList<String>? get step;

    PlaybookTask._();

    static void _initializeBuilder(PlaybookTaskBuilder b) => b;

    factory PlaybookTask([void updates(PlaybookTaskBuilder b)]) = _$PlaybookTask;

    @BuiltValueSerializer(custom: true)
    static Serializer<PlaybookTask> get serializer => _$PlaybookTaskSerializer();
}

class _$PlaybookTaskSerializer implements StructuredSerializer<PlaybookTask> {
    @override
    final Iterable<Type> types = const [PlaybookTask, _$PlaybookTask];

    @override
    final String wireName = r'PlaybookTask';

    @override
    Iterable<Object?> serialize(Serializers serializers, PlaybookTask object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.args != null) {
            result
                ..add(r'args')
                ..add(serializers.serialize(object.args,
                    specifiedType: const FullType(BuiltMap, [FullType(String), FullType(TaskArgs)])));
        }
        if (object.const_ != null) {
            result
                ..add(r'const')
                ..add(serializers.serialize(object.const_,
                    specifiedType: const FullType(BuiltMap, [FullType(String), FullType(String)])));
        }
        if (object.step != null) {
            result
                ..add(r'step')
                ..add(serializers.serialize(object.step,
                    specifiedType: const FullType(BuiltList, [FullType(String)])));
        }
        return result;
    }

    @override
    PlaybookTask deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = PlaybookTaskBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'args':
                    result.args.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltMap, [FullType(String), FullType(TaskArgs)])) as BuiltMap<String, TaskArgs>);
                    break;
                case r'const':
                    result.const_.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltMap, [FullType(String), FullType(String)])) as BuiltMap<String, String>);
                    break;
                case r'step':
                    result.step.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltList, [FullType(String)])) as BuiltList<String>);
                    break;
            }
        }
        return result.build();
    }
}

