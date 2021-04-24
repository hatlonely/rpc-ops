//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'api_run_ops_req.g.dart';



abstract class ApiRunOpsReq implements Built<ApiRunOpsReq, ApiRunOpsReqBuilder> {
    @BuiltValueField(wireName: r'repositoryID')
    String? get repositoryID;

    @BuiltValueField(wireName: r'variableID')
    String? get variableID;

    @BuiltValueField(wireName: r'version')
    String? get version;

    @BuiltValueField(wireName: r'environment')
    String? get environment;

    @BuiltValueField(wireName: r'task')
    String? get task;

    @BuiltValueField(wireName: r'args')
    String? get args;

    ApiRunOpsReq._();

    static void _initializeBuilder(ApiRunOpsReqBuilder b) => b;

    factory ApiRunOpsReq([void updates(ApiRunOpsReqBuilder b)]) = _$ApiRunOpsReq;

    @BuiltValueSerializer(custom: true)
    static Serializer<ApiRunOpsReq> get serializer => _$ApiRunOpsReqSerializer();
}

class _$ApiRunOpsReqSerializer implements StructuredSerializer<ApiRunOpsReq> {
    @override
    final Iterable<Type> types = const [ApiRunOpsReq, _$ApiRunOpsReq];

    @override
    final String wireName = r'ApiRunOpsReq';

    @override
    Iterable<Object?> serialize(Serializers serializers, ApiRunOpsReq object,
        {FullType specifiedType = FullType.unspecified}) {
        final result = <Object?>[];
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
        if (object.environment != null) {
            result
                ..add(r'environment')
                ..add(serializers.serialize(object.environment,
                    specifiedType: const FullType(String)));
        }
        if (object.task != null) {
            result
                ..add(r'task')
                ..add(serializers.serialize(object.task,
                    specifiedType: const FullType(String)));
        }
        if (object.args != null) {
            result
                ..add(r'args')
                ..add(serializers.serialize(object.args,
                    specifiedType: const FullType(String)));
        }
        return result;
    }

    @override
    ApiRunOpsReq deserialize(Serializers serializers, Iterable<Object?> serialized,
        {FullType specifiedType = FullType.unspecified}) {
        final result = ApiRunOpsReqBuilder();

        final iterator = serialized.iterator;
        while (iterator.moveNext()) {
            final key = iterator.current as String;
            iterator.moveNext();
            final Object? value = iterator.current;
            switch (key) {
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
                case r'environment':
                    result.environment = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'task':
                    result.task = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
                case r'args':
                    result.args = serializers.deserialize(value,
                        specifiedType: const FullType(String)) as String;
                    break;
            }
        }
        return result.build();
    }
}

