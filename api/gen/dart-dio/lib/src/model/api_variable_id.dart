//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_variable_id.g.dart';



abstract class ApiVariableID implements Built<ApiVariableID, ApiVariableIDBuilder> {
    @BuiltValueField(wireName: r'id')
    String? get id;

    ApiVariableID._();

    static void _initializeBuilder(ApiVariableIDBuilder b) => b;

    factory ApiVariableID([void updates(ApiVariableIDBuilder b)]) = _$ApiVariableID;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiVariableID> get serializer => _$ApiVariableIDSerializer();
}

class _$ApiVariableIDSerializer implements StructuredSerializer<ApiVariableID> {
    @override
    final Iterable<Type> types = const [ApiVariableID, _$ApiVariableID];

    @override
    final String wireName = r'ApiVariableID';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiVariableID object,
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
    ApiVariableID deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiVariableIDBuilder();

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

