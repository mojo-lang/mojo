// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

package org.mojolang.mojo.core;

import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.List;

/**
 * Utilities to help create {@code google.protobuf.Value} messages.
 */
public final class ValueUtil {

    private static final Value NULL_VALUE =
            Value.newBuilder().setNullVal(Null.newBuilder().build()).build();

    public static Value ofNull() {
        return NULL_VALUE;
    }

    /**
     * Returns a Value object with number set to value.
     */
    public static Value of(boolean value) {
        return Value.newBuilder().setBoolVal(value).build();
    }

    public static Value of(long value) {
        if (value >= 0) {
            return Value.newBuilder().setPositiveVal(value).build();
        } else {
            return Value.newBuilder().setNegativeVal(-value).build();
        }
    }

    /**
     * Returns a Value object with number set to value.
     */
    public static Value of(double value) {
        return Value.newBuilder().setDoubleVal(value).build();
    }

    /**
     * Returns a Value object with string set to value.
     */
    public static Value of(String value) {
        return Value.newBuilder().setStringVal(value).build();
    }

    /**
     * Returns a Value object with struct set to value.
     */
    public static Value of(org.mojolang.mojo.core.Object value) {
        return Value.newBuilder().setObjectVal(value).build();
    }

    /**
     * Returns a Value with ListValue set to value.
     */
    public static Value of(Values value) {
        return Value.newBuilder().setValuesVal(value).build();
    }

    /**
     * Returns a Value with ListValue set to the appending the result of calling {@link #of} on each
     * element in the iterable.
     */
    public static Value of(Iterable<Value> values) {
        return Value.newBuilder().setValuesVal(Values.newBuilder().addAllVals(values)).build();
    }

    public static Value of(java.lang.Object value) {
        if (java.util.Objects.isNull(value)) {
            return ofNull();
        }

        if (value instanceof Integer) {
            return of(((Integer) value).longValue());
        } else if (value instanceof Short) {
            return of(((Short) value).longValue());
        } else if (value instanceof Long) {
            return of(((Long) value).longValue());
        } else if (value instanceof BigDecimal) {
            return of(((BigDecimal) value).doubleValue());
        } else if (value instanceof Boolean) {
            return of(((Boolean) value).booleanValue());
        } else if (value instanceof Double) {
            return of(((Double) value).doubleValue());
        } else if (value instanceof Float) {
            return of(((Float) value).floatValue());
        } else if (value instanceof String) {
            return of((String) value);
        } else if (value instanceof Map) {
            Object.Builder builder = Object.newBuilder();
            for (String key : ((Map<String, java.lang.Object>) value).keySet()) {
                builder.putVals(key, of(((Map<?, ?>) value).get(key)));
            }
            return of(builder.build());
        } else if (value instanceof List) {
            Values.Builder builder = Values.newBuilder();
            for (java.lang.Object val : (List<java.lang.Object>) value) {
                builder.addVals(of(val));
            }
            return of(builder.build());
        } else if (value instanceof Value) {
            return (Value) value;
        } else if (value instanceof Values) {
            return of((Values) value);
        } else if (value instanceof org.mojolang.mojo.core.Object) {
            return of((org.mojolang.mojo.core.Object) value);
        } else {
            System.out.println("unknown type, or yourself type");
            return ofNull();
        }
    }

    public static java.lang.Object to(Value value) {
        if (value.hasNullVal()) {
            return null;
        } else if (value.hasBoolVal()) {
            return value.getBoolVal();
        } else if (value.hasPositiveVal()) {
            return value.getPositiveVal();
        } else if (value.hasNegativeVal()) {
            return -value.getNegativeVal();
        } else if (value.hasDoubleVal()) {
            return value.getDoubleVal();
        } else if (value.hasStringVal()) {
            return value.getStringVal();
        } else if (value.hasValuesVal()) {
            Values values = value.getValuesVal();
            List<java.lang.Object> vs = new ArrayList<>();
            for (Value val : values.getValsList()) {
                vs.add(to(val));
            }
            return vs;
        } else if (value.hasObjectVal()) {
            Map<String, Value> obj = value.getObjectVal().getValsMap();
            Map<String, java.lang.Object> vs = new HashMap<>();
            for (String key : obj.keySet()) {
                vs.put(key, to(obj.get(key)));
            }
            return vs;
        }

        return null;
    }

    private ValueUtil() {
    }
}