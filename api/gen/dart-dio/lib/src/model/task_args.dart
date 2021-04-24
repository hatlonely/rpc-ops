//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'task_args.g.dart';



abstract class TaskArgs implements Built<TaskArgs, TaskArgsBuilder> {
    @BuiltValueField(wireName: r'type')
    String? get type;

    @BuiltValueField(wireName: r'default')
    String? get default_;

    @BuiltValueField(wireName: r'validation')
    String? get validation;

    TaskArgs._();

    static void _initializeBuilder(TaskArgsBuilder b) => b;

    factory TaskArgs([void updates(TaskArgsBuilder b)]) = _$TaskArgs;

    @BuiltValueSerializer(custom: true)
    static Serializer<TaskArgs> get serializer => _$TaskArgsSerializer();
}

class _$TaskArgsSerializer implements StructuredSerializer<TaskArgs> {
    @override
    final Iterable<Type> types = const [TaskArgs, _$TaskArgs];

    @override
    final String wireName = r'TaskArgs';

    @override
    Iterable<Object?> serialize(Serializers serializers, TaskArgs object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.type != null) {
            result
                ..add(r'type')
                ..add(serializers.serialize(object.type,
                    specifiedType: const FullType(String)));
        }
        if (object.default_ != null) {
            result
                ..add(r'default')
                ..add(serializers.serialize(object.default_,
                    specifiedType: const FullType(String)));
        }
        if (object.validation != null) {
            result
                ..add(r'validation')
                ..add(serializers.serialize(object.validation,
                    specifiedType: const FullType(String)));
        }
        return result;
    }

    @override
    TaskArgs deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = TaskArgsBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'type':
                    result.type = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'default':
                    result.default_ = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'validation':
                    result.validation = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
            }
        }
        return result.build();
    }
}

