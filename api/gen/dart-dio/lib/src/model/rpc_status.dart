//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_collection/built_collection.dart';
import 'package:openapi/src/model/protobuf_any.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'rpc_status.g.dart';



abstract class RpcStatus implements Built<RpcStatus, RpcStatusBuilder> {
    @BuiltValueField(wireName: r'code')
    int? get code;

    @BuiltValueField(wireName: r'message')
    String? get message;

    @BuiltValueField(wireName: r'details')
    BuiltList<ProtobufAny>? get details;

    RpcStatus._();

    static void _initializeBuilder(RpcStatusBuilder b) => b;

    factory RpcStatus([void updates(RpcStatusBuilder b)]) = _$RpcStatus;

    @BuiltValueSerializer(custom: true)
    static Serializer<RpcStatus> get serializer => _$RpcStatusSerializer();
}

class _$RpcStatusSerializer implements StructuredSerializer<RpcStatus> {
    @override
    final Iterable<Type> types = const [RpcStatus, _$RpcStatus];

    @override
    final String wireName = r'RpcStatus';

    @override
    Iterable<Object?> serialize(Serializers serializers, RpcStatus object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.code != null) {
            result
                ..add(r'code')
                ..add(serializers.serialize(object.code,
                    specifiedType: const FullType(int)));
        }
        if (object.message != null) {
            result
                ..add(r'message')
                ..add(serializers.serialize(object.message,
                    specifiedType: const FullType(String)));
        }
        if (object.details != null) {
            result
                ..add(r'details')
                ..add(serializers.serialize(object.details,
                    specifiedType: const FullType(BuiltList, [FullType(ProtobufAny)])));
        }
        return result;
    }

    @override
    RpcStatus deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = RpcStatusBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'code':
                    result.code = serializers.deserialize(value,
                        specifiedType: const FullType(int)) as int;
                    break;
                case r'message':
                    result.message = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'details':
                    result.details.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltList, [FullType(ProtobufAny)])) as BuiltList<ProtobufAny>);
                    break;
            }
        }
        return result.build();
    }
}

