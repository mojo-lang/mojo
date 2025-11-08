package org.mojolang.mojo.core;

import java.lang.reflect.Type;
import java.lang.reflect.ParameterizedType;
import java.util.List;

import com.google.protobuf.InvalidProtocolBufferException;
import com.google.protobuf.Message;
import com.alibaba.fastjson2.writer.ObjectWriter;
import com.alibaba.fastjson2.JSONWriter;
import com.google.protobuf.MessageOrBuilder;

public class JsonObjectWriter implements ObjectWriter  {
    public static final JsonObjectWriter INSTANCE = new JsonObjectWriter();

    @Override
    public void write(JSONWriter jsonWriter, java.lang.Object object, java.lang.Object fieldName, Type fieldType, long features) {
        if (object == null) {
            jsonWriter.writeNull();
            return;
        }

        fieldType = object.getClass();
        try {
            if (fieldType instanceof Class && MessageOrBuilder.class.isAssignableFrom((Class<?>) fieldType)) {
                final MessageOrBuilder value = (MessageOrBuilder) object;
                jsonWriter.writeRaw(JsonFormat.printer().print(value));
            } else if (fieldType instanceof ParameterizedType) {
                final ParameterizedType type = (ParameterizedType) fieldType;

                if (List.class.isAssignableFrom((Class<?>) type.getRawType())) {
                    final Type argument = type.getActualTypeArguments()[0];
                    if (Message.class.isAssignableFrom((Class<?>) argument)) {
                        final List<Message> messageList = (List<Message>) object;
                        jsonWriter.writeString(JsonFormat.printer().print(messageList));
                    } else {
                        jsonWriter.writeString("[]");
                    }
                }
//                else if (Map.class.isAssignableFrom((Class<?>) type.getRawType())) {
//                    final Type[] arguments = type.getActualTypeArguments();
//                    if (arguments.length == 2) {
//                        final Type keyType = arguments[0], valueType = arguments[1];
//
//                        if (Message.class.isAssignableFrom((Class<?>) valueType)) {
//                            Map<?, Message> messageMap = (Map<?, Message>) object;
//
//                            final String toStr = JsonFormat.printer().print(messageMap);
//
//                            jsonWriter.writeString(toStr);
//                        }
//                    } else {
//                        jsonWriter.writeString("{}");
//                    }
//                }
            }
        } catch (InvalidProtocolBufferException e) {
            throw new RuntimeException(e);
        }
    }
}