//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_variable.g.dart';



abstract class ApiVariable implements Built<ApiVariable, ApiVariableBuilder> {
    @BuiltValueField(wireName: r'id')
    String? get id;

    @BuiltValueField(wireName: r'name')
    String? get name;

    @BuiltValueField(wireName: r'body')
    String? get body;

    @BuiltValueField(wireName: r'createAt')
    String? get createAt;

    @BuiltValueField(wireName: r'updateAt')
    String? get updateAt;

    ApiVariable._();

    static void _initializeBuilder(ApiVariableBuilder b) => b;

    factory ApiVariable([void updates(ApiVariableBuilder b)]) = _$ApiVariable;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiVariable> get serializer => _$ApiVariableSerializer();
}

class _$ApiVariableSerializer implements StructuredSerializer<ApiVariable> {
    @override
    final Iterable<Type> types = const [ApiVariable, _$ApiVariable];

    @override
    final String wireName = r'ApiVariable';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiVariable object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.id != null) {
            result
                ..add(r'id')
                ..add(serializers.serialize(object.id,
                    specifiedType: const FullType(String)));
        }
        if (object.name != null) {
            result
                ..add(r'name')
                ..add(serializers.serialize(object.name,
                    specifiedType: const FullType(String)));
        }
        if (object.body != null) {
            result
                ..add(r'body')
                ..add(serializers.serialize(object.body,
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
    ApiVariable deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiVariableBuilder();

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
                case r'name':
                    result.name = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'body':
                    result.body = serializers.deserialize(value,
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

