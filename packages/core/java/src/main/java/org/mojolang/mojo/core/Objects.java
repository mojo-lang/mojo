package org.mojolang.mojo.core;

import java.util.HashMap;
import java.util.Map;
import java.util.List;
import java.util.ArrayList;

/** Utilities to help create {@code google.protobuf.Struct} messages. */
public final class Objects {

    public static org.mojolang.mojo.core.Object of(Map<String, java.lang.Object> values) {
        org.mojolang.mojo.core.Object.Builder builder = Object.newBuilder();
        for (String key : values.keySet()) {
            builder.putVals(key, ValueUtil.of(values.get(key)));
        }
        return builder.build();
    }

    public static List<org.mojolang.mojo.core.Object> ofList(List<Map<String, java.lang.Object>> valueList) {
        List<org.mojolang.mojo.core.Object> list = new ArrayList<>();
        for (Map<String, java.lang.Object> value : valueList) {
            list.add(of(value));
        }
        return list;
    }

    /**
     * Returns a struct containing the key-value pair.
     */
    public static org.mojolang.mojo.core.Object of(String k1, Value v1) {
        return Object.newBuilder().putVals(k1, v1).build();
    }

    /**
     * Returns a struct containing each of the key-value pairs.
     *
     * <p>Providing duplicate keys is undefined behavior.
     */
    public static org.mojolang.mojo.core.Object of(String k1, Value v1, String k2, Value v2) {
        return Object.newBuilder().putVals(k1, v1).putVals(k2, v2).build();
    }

    /**
     * Returns a struct containing each of the key-value pairs.
     *
     * <p>Providing duplicate keys is undefined behavior.
     */
    public static org.mojolang.mojo.core.Object of(String k1, Value v1, String k2, Value v2, String k3, Value v3) {
        return Object.newBuilder().putVals(k1, v1).putVals(k2, v2).putVals(k3, v3).build();
    }

    public static Map<String, java.lang.Object> to(org.mojolang.mojo.core.Object obj) {
        Map<String, Value> values = obj.getValsMap();

        Map<String, java.lang.Object> result = new HashMap<>();
        for (String key : values.keySet()) {
            result.put(key, ValueUtil.to(values.get(key)));
        }

        return result;
    }

    public static List<Map<String, java.lang.Object>> toList(List<org.mojolang.mojo.core.Object> objects) {
        List<Map<String, java.lang.Object>> list = new ArrayList<>();
        for (org.mojolang.mojo.core.Object obj : objects) {
            list.add(to(obj));
        }
        return list;
    }

    private Objects() {}
}
