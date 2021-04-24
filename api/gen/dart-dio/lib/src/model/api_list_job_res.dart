//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:openapi/src/model/api_job.dart';
import 'package:built_collection/built_collection.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_list_job_res.g.dart';



abstract class ApiListJobRes implements Built<ApiListJobRes, ApiListJobResBuilder> {
    @BuiltValueField(wireName: r'jobs')
    BuiltList<ApiJob>? get jobs;

    ApiListJobRes._();

    static void _initializeBuilder(ApiListJobResBuilder b) => b;

    factory ApiListJobRes([void updates(ApiListJobResBuilder b)]) = _$ApiListJobRes;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiListJobRes> get serializer => _$ApiListJobResSerializer();
}

class _$ApiListJobResSerializer implements StructuredSerializer<ApiListJobRes> {
    @override
    final Iterable<Type> types = const [ApiListJobRes, _$ApiListJobRes];

    @override
    final String wireName = r'ApiListJobRes';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiListJobRes object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.jobs != null) {
            result
                ..add(r'jobs')
                ..add(serializers.serialize(object.jobs,
                    specifiedType: const FullType(BuiltList, [FullType(ApiJob)])));
        }
        return result;
    }

    @override
    ApiListJobRes deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiListJobResBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
                case r'jobs':
                    result.jobs.replace(serializers.deserialize(value,
                        specifiedType: const FullType(BuiltList, [FullType(ApiJob)])) as BuiltList<ApiJob>);
                    break;
            }
        }
        return result.build();
    }
}

