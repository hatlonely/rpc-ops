//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'protobuf_any.g.dart';



abstract class ProtobufAny implements Built<ProtobufAny, ProtobufAnyBuilder> {
    @BuiltValueField(wireName: r'typeUrl')
    String? get typeUrl;

    @BuiltValueField(wireName: r'value')
    String? get value;

    ProtobufAny._();

    static void _initializeBuilder(ProtobufAnyBuilder b) => b;

    factory ProtobufAny([void updates(ProtobufAnyBuilder b)]) = _$ProtobufAny;

    @BuiltValueSerializer(custom: true)
    static Serializer<ProtobufAny> get serializer => _$ProtobufAnySerializer();
}

class _$ProtobufAnySerializer implements StructuredSerializer<ProtobufAny> {
    @override
    final Iterable<Type> types = const [ProtobufAny, _$ProtobufAny];

    @override
    final String wireName = r'ProtobufAny';

    @override
    Iterable<Object?> serialize(Serializers serializers, ProtobufAny object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.typeUrl != null) {
            result
                ..add(r'typeUrl')
                ..add(serializers.serialize(object.typeUrl,
                    specifiedType: const FullType(String)));
        }
        if (object.value != null) {
            result
                ..add(r'value')
                ..add(serializers.serialize(object.value,
                    specifiedType: const FullType(String)));
        }
        return result;
    }

    @override
    ProtobufAny deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ProtobufAnyBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'typeUrl':
                    result.typeUrl = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'value':
                    result.value = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
            }
        }
        return result.build();
    }
}

