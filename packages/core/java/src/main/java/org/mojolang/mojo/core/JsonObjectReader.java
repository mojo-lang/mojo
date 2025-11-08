package org.mojolang.mojo.core;

import java.io.IOException;
import java.util.List;
import java.lang.reflect.Type;
import java.lang.reflect.ParameterizedType;

import com.alibaba.fastjson2.JSON;
import com.alibaba.fastjson2.reader.ObjectReader;
import com.alibaba.fastjson2.JSONReader;
import com.google.protobuf.MessageOrBuilder;

import com.google.protobuf.Message;

public class JsonObjectReader implements ObjectReader<java.lang.Object> {
    public static final JsonObjectReader INSTANCE = new JsonObjectReader();

    @Override
    public java.lang.Object readObject(JSONReader jsonReader, Type fieldType, java.lang.Object fieldName, long features) {
        if (jsonReader.nextIfNull()) {
            return null;
        }

        final String value = JSON.toJSONString(jsonReader.readAny());
        try {
            if (fieldType instanceof Class && Message.class.isAssignableFrom((Class<?>) fieldType)) {
                return JsonFormat.parser().fromJSON(value, (Class) fieldType);
            }

            if (fieldType instanceof ParameterizedType) {
                final ParameterizedType type = (ParameterizedType) fieldType;

                if (List.class.isAssignableFrom((Class<?>) type.getRawType())) {
                    final Type argument = type.getActualTypeArguments()[0];
                    if (Message.class.isAssignableFrom((Class<?>) argument)) {
                        return JsonFormat.parser().fromJSON(value, (Class) argument);
                    }
                }

//                if (Map.class.isAssignableFrom((Class<?>) type.getRawType())) {
//                    final Type[] arguments = type.getActualTypeArguments();
//                    if (arguments.length == 2) {
//                        final Type keyType = arguments[0], valueType = arguments[1];
//                        if (Message.class.isAssignableFrom((Class<?>) valueType)) {
//                            return (T) ProtobufUtils.toBeanMap(value, (Class) keyType, (Class) valueType);
//                        }
//                    }
//                }
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        return null;
    }
}
