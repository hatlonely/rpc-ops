//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_describe_repository_req.g.dart';



abstract class ApiDescribeRepositoryReq implements Built<ApiDescribeRepositoryReq, ApiDescribeRepositoryReqBuilder> {
    @BuiltValueField(wireName: r'id')
    String? get id;

    @BuiltValueField(wireName: r'version')
    String? get version;

    ApiDescribeRepositoryReq._();

    static void _initializeBuilder(ApiDescribeRepositoryReqBuilder b) => b;

    factory ApiDescribeRepositoryReq([void updates(ApiDescribeRepositoryReqBuilder b)]) = _$ApiDescribeRepositoryReq;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiDescribeRepositoryReq> get serializer => _$ApiDescribeRepositoryReqSerializer();
}

class _$ApiDescribeRepositoryReqSerializer implements StructuredSerializer<ApiDescribeRepositoryReq> {
    @override
    final Iterable<Type> types = const [ApiDescribeRepositoryReq, _$ApiDescribeRepositoryReq];

    @override
    final String wireName = r'ApiDescribeRepositoryReq';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiDescribeRepositoryReq object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.id != null) {
            result
                ..add(r'id')
                ..add(serializers.serialize(object.id,
                    specifiedType: const FullType(String)));
        }
        if (object.version != null) {
            result
                ..add(r'version')
                ..add(serializers.serialize(object.version,
                    specifiedType: const FullType(String)));
        }
        return result;
    }

    @override
    ApiDescribeRepositoryReq deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiDescribeRepositoryReqBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'id':
                    result.id = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'version':
                    result.version = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
            }
        }
        return result.build();
    }
}

