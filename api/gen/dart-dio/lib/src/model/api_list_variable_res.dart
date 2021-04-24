//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_collection/built_collection.dart';
import 'package:api/src/model/api_variable.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_list_variable_res.g.dart';



abstract class ApiListVariableRes implements Built<ApiListVariableRes, ApiListVariableResBuilder> {
    @BuiltValueField(wireName: r'variables')
    BuiltList<ApiVariable>? get variables;

    ApiListVariableRes._();

    static void _initializeBuilder(ApiListVariableResBuilder b) => b;

    factory ApiListVariableRes([void updates(ApiListVariableResBuilder b)]) = _$ApiListVariableRes;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiListVariableRes> get serializer => _$ApiListVariableResSerializer();
}

class _$ApiListVariableResSerializer implements StructuredSerializer<ApiListVariableRes> {
    @override
    final Iterable<Type> types = const [ApiListVariableRes, _$ApiListVariableRes];

    @override
    final String wireName = r'ApiListVariableRes';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiListVariableRes object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.variables != null) {
            result
                ..add(r'variables')
                ..add(serializers.serialize(object.variables,
                    specifiedType: const FullType(BuiltList, [FullType(ApiVariable)])));
        }
        return result;
    }

    @override
    ApiListVariableRes deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiListVariableResBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'variables':
                    result.variables.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltList, [FullType(ApiVariable)])) as BuiltList<ApiVariable>);
                    break;
            }
        }
        return result.build();
    }
}

