//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_collection/built_collection.dart';
import 'package:openapi/src/model/api_repository.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_list_repository_res.g.dart';



abstract class ApiListRepositoryRes implements Built<ApiListRepositoryRes, ApiListRepositoryResBuilder> {
    @BuiltValueField(wireName: r'repositories')
    BuiltList<ApiRepository>? get repositories;

    ApiListRepositoryRes._();

    static void _initializeBuilder(ApiListRepositoryResBuilder b) => b;

    factory ApiListRepositoryRes([void updates(ApiListRepositoryResBuilder b)]) = _$ApiListRepositoryRes;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiListRepositoryRes> get serializer => _$ApiListRepositoryResSerializer();
}

class _$ApiListRepositoryResSerializer implements StructuredSerializer<ApiListRepositoryRes> {
    @override
    final Iterable<Type> types = const [ApiListRepositoryRes, _$ApiListRepositoryRes];

    @override
    final String wireName = r'ApiListRepositoryRes';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiListRepositoryRes object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.repositories != null) {
            result
                ..add(r'repositories')
                ..add(serializers.serialize(object.repositories,
                    specifiedType: const FullType(BuiltList, [FullType(ApiRepository)])));
        }
        return result;
    }

    @override
    ApiListRepositoryRes deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiListRepositoryResBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'repositories':
                    result.repositories.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltList, [FullType(ApiRepository)])) as BuiltList<ApiRepository>);
                    break;
            }
        }
        return result.build();
    }
}

