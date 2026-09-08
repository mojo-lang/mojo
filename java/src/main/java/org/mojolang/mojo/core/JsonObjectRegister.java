package org.mojolang.mojo.core;

import com.alibaba.fastjson2.JSON;
import com.google.protobuf.Message;
import com.google.protobuf.NullValue;
import org.springframework.stereotype.Component;

import java.lang.reflect.Type;

@Component
public class JsonObjectRegister {
    public static void init() {
        register(Message.class);

        register(Any.class);
        register(BoolValue.class);
        register(BytesValue.class);
        register(Date.class);
        register(DateTime.class);
        register(Duration.class);
        register(ErrorCode.class);
        register(FieldMask.class);
        register(Float32Value.class);
        register(Float64Value.class);
        register(Int32Value.class);
        register(Int64Value.class);
        register(NullValue.class);
        register(org.mojolang.mojo.core.Object.class);
        register(Ordering.class);
        register(StringValue.class);
        register(Timestamp.class);
        register(UInt32Value.class);
        register(UInt64Value.class);
        register(Url.class);
        register(Value.class);
        register(Values.class);
        register(Version.class);
    }

    public static void register(Type type) {
        JSON.register(type, JsonObjectReader.INSTANCE);
        JSON.register(type, JsonObjectWriter.INSTANCE);
    }
}