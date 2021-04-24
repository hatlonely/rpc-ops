//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_repository.g.dart';



abstract class ApiRepository implements Built<ApiRepository, ApiRepositoryBuilder> {
    @BuiltValueField(wireName: r'id')
    String? get id;

    @BuiltValueField(wireName: r'username')
    String? get username;

    @BuiltValueField(wireName: r'password')
    String? get password;

    @BuiltValueField(wireName: r'endpoint')
    String? get endpoint;

    @BuiltValueField(wireName: r'name')
    String? get name;

    @BuiltValueField(wireName: r'playbook')
    String? get playbook;

    @BuiltValueField(wireName: r'createAt')
    String? get createAt;

    @BuiltValueField(wireName: r'updateAt')
    String? get updateAt;

    ApiRepository._();

    static void _initializeBuilder(ApiRepositoryBuilder b) => b;

    factory ApiRepository([void updates(ApiRepositoryBuilder b)]) = _$ApiRepository;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiRepository> get serializer => _$ApiRepositorySerializer();
}

class _$ApiRepositorySerializer implements StructuredSerializer<ApiRepository> {
    @override
    final Iterable<Type> types = const [ApiRepository, _$ApiRepository];

    @override
    final String wireName = r'ApiRepository';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiRepository object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.id != null) {
            result
                ..add(r'id')
                ..add(serializers.serialize(object.id,
                    specifiedType: const FullType(String)));
        }
        if (object.username != null) {
            result
                ..add(r'username')
                ..add(serializers.serialize(object.username,
                    specifiedType: const FullType(String)));
        }
        if (object.password != null) {
            result
                ..add(r'password')
                ..add(serializers.serialize(object.password,
                    specifiedType: const FullType(String)));
        }
        if (object.endpoint != null) {
            result
                ..add(r'endpoint')
                ..add(serializers.serialize(object.endpoint,
                    specifiedType: const FullType(String)));
        }
        if (object.name != null) {
            result
                ..add(r'name')
                ..add(serializers.serialize(object.name,
                    specifiedType: const FullType(String)));
        }
        if (object.playbook != null) {
            result
                ..add(r'playbook')
                ..add(serializers.serialize(object.playbook,
                    specifiedType: const FullType(String)));
        }
        if (object.createAt != null) {
            result
                ..add(r'createAt')
                ..add(serializers.serialize(object.createAt,
                    specifiedType: const FullType(String)));
        }
        if (object.updateAt != null) {
            result
                ..add(r'updateAt')
                ..add(serializers.serialize(object.updateAt,
                    specifiedType: const FullType(String)));
        }
        return result;
    }

    @override
    ApiRepository deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiRepositoryBuilder();

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
                case r'username':
                    result.username = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'password':
                    result.password = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'endpoint':
                    result.endpoint = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'name':
                    result.name = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'playbook':
                    result.playbook = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'createAt':
                    result.createAt = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'updateAt':
                    result.updateAt = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
            }
        }
        return result.build();
    }
}

