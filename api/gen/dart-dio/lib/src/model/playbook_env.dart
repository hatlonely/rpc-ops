//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_collection/built_collection.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'playbook_env.g.dart';



abstract class PlaybookEnv implements Built<PlaybookEnv, PlaybookEnvBuilder> {
    @BuiltValueField(wireName: r'key')
    String? get key;

    @BuiltValueField(wireName: r'val')
    BuiltMap<String, String>? get val;

    PlaybookEnv._();

    static void _initializeBuilder(PlaybookEnvBuilder b) => b;

    factory PlaybookEnv([void updates(PlaybookEnvBuilder b)]) = _$PlaybookEnv;

    @BuiltValueSerializer(custom: true)
    static Serializer<PlaybookEnv> get serializer => _$PlaybookEnvSerializer();
}

class _$PlaybookEnvSerializer implements StructuredSerializer<PlaybookEnv> {
    @override
    final Iterable<Type> types = const [PlaybookEnv, _$PlaybookEnv];

    @override
    final String wireName = r'PlaybookEnv';

    @override
    Iterable<Object?> serialize(Serializers serializers, PlaybookEnv object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.key != null) {
            result
                ..add(r'key')
                ..add(serializers.serialize(object.key,
                    specifiedType: const FullType(String)));
        }
        if (object.val != null) {
            result
                ..add(r'val')
                ..add(serializers.serialize(object.val,
                    specifiedType: const FullType(BuiltMap, [FullType(String), FullType(String)])));
        }
        return result;
    }

    @override
    PlaybookEnv deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = PlaybookEnvBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'key':
                    result.key = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'val':
                    result.val.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltMap, [FullType(String), FullType(String)])) as BuiltMap<String, String>);
                    break;
            }
        }
        return result.build();
    }
}

