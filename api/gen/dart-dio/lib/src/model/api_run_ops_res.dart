//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_run_ops_res.g.dart';



abstract class ApiRunOpsRes implements Built<ApiRunOpsRes, ApiRunOpsResBuilder> {
    @BuiltValueField(wireName: r'jobID')
    String? get jobID;

    ApiRunOpsRes._();

    static void _initializeBuilder(ApiRunOpsResBuilder b) => b;

    factory ApiRunOpsRes([void updates(ApiRunOpsResBuilder b)]) = _$ApiRunOpsRes;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiRunOpsRes> get serializer => _$ApiRunOpsResSerializer();
}

class _$ApiRunOpsResSerializer implements StructuredSerializer<ApiRunOpsRes> {
    @override
    final Iterable<Type> types = const [ApiRunOpsRes, _$ApiRunOpsRes];

    @override
    final String wireName = r'ApiRunOpsRes';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiRunOpsRes object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.jobID != null) {
            result
                ..add(r'jobID')
                ..add(serializers.serialize(object.jobID,
                    specifiedType: const FullType(String)));
        }
        return result;
    }

    @override
    ApiRunOpsRes deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiRunOpsResBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'jobID':
                    result.jobID = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
            }
        }
        return result.build();
    }
}

