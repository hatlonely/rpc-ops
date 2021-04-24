//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_repository_id.g.dart';



abstract class ApiRepositoryID implements Built<ApiRepositoryID, ApiRepositoryIDBuilder> {
    @BuiltValueField(wireName: r'id')
    String? get id;

    ApiRepositoryID._();

    static void _initializeBuilder(ApiRepositoryIDBuilder b) => b;

    factory ApiRepositoryID([void updates(ApiRepositoryIDBuilder b)]) = _$ApiRepositoryID;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiRepositoryID> get serializer => _$ApiRepositoryIDSerializer();
}

class _$ApiRepositoryIDSerializer implements StructuredSerializer<ApiRepositoryID> {
    @override
    final Iterable<Type> types = const [ApiRepositoryID, _$ApiRepositoryID];

    @override
    final String wireName = r'ApiRepositoryID';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiRepositoryID object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.id != null) {
            result
                ..add(r'id')
                ..add(serializers.serialize(object.id,
                    specifiedType: const FullType(String)));
        }
        return result;
    }

    @override
    ApiRepositoryID deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiRepositoryIDBuilder();

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
            }
        }
        return result.build();
    }
}

