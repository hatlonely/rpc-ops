//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_job.g.dart';



abstract class ApiJob implements Built<ApiJob, ApiJobBuilder> {
    @BuiltValueField(wireName: r'id')
    String? get id;

    @BuiltValueField(wireName: r'seq')
    int? get seq;

    @BuiltValueField(wireName: r'state')
    String? get state;

    @BuiltValueField(wireName: r'repositoryID')
    String? get repositoryID;

    @BuiltValueField(wireName: r'variableID')
    String? get variableID;

    @BuiltValueField(wireName: r'version')
    String? get version;

    @BuiltValueField(wireName: r'createAt')
    String? get createAt;

    @BuiltValueField(wireName: r'updateAt')
    String? get updateAt;

    @BuiltValueField(wireName: r'scheduleAt')
    String? get scheduleAt;

    @BuiltValueField(wireName: r'elapseSeconds')
    int? get elapseSeconds;

    @BuiltValueField(wireName: r'output')
    String? get output;

    ApiJob._();

    static void _initializeBuilder(ApiJobBuilder b) => b;

    factory ApiJob([void updates(ApiJobBuilder b)]) = _$ApiJob;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiJob> get serializer => _$ApiJobSerializer();
}

class _$ApiJobSerializer implements StructuredSerializer<ApiJob> {
    @override
    final Iterable<Type> types = const [ApiJob, _$ApiJob];

    @override
    final String wireName = r'ApiJob';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiJob object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
        if (object.id != null) {
            result
                ..add(r'id')
                ..add(serializers.serialize(object.id,
                    specifiedType: const FullType(String)));
        }
        if (object.seq != null) {
            result
                ..add(r'seq')
                ..add(serializers.serialize(object.seq,
                    specifiedType: const FullType(int)));
        }
        if (object.state != null) {
            result
                ..add(r'state')
                ..add(serializers.serialize(object.state,
                    specifiedType: const FullType(String)));
        }
        if (object.repositoryID != null) {
            result
                ..add(r'repositoryID')
                ..add(serializers.serialize(object.repositoryID,
                    specifiedType: const FullType(String)));
        }
        if (object.variableID != null) {
            result
                ..add(r'variableID')
                ..add(serializers.serialize(object.variableID,
                    specifiedType: const FullType(String)));
        }
        if (object.version != null) {
            result
                ..add(r'version')
                ..add(serializers.serialize(object.version,
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
        if (object.scheduleAt != null) {
            result
                ..add(r'scheduleAt')
                ..add(serializers.serialize(object.scheduleAt,
                    specifiedType: const FullType(String)));
        }
        if (object.elapseSeconds != null) {
            result
                ..add(r'elapseSeconds')
                ..add(serializers.serialize(object.elapseSeconds,
                    specifiedType: const FullType(int)));
        }
        if (object.output != null) {
            result
                ..add(r'output')
                ..add(serializers.serialize(object.output,
                    specifiedType: const FullType(String)));
        }
        return result;
    }

    @override
    ApiJob deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiJobBuilder();

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
                case r'seq':
                    result.seq = serializers.deserialize(value,
                        specifiedType: const FullType(int)) as int;
                    break;
                case r'state':
                    result.state = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'repositoryID':
                    result.repositoryID = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'variableID':
                    result.variableID = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'version':
                    result.version = serializers.deserialize(value,
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
                case r'scheduleAt':
                    result.scheduleAt = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'elapseSeconds':
                    result.elapseSeconds = serializers.deserialize(value,
                        specifiedType: const FullType(int)) as int;
                    break;
                case r'output':
                    result.output = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
            }
        }
        return result.build();
    }
}

