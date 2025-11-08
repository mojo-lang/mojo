package org.mojolang.mojo.core;

public final class Nulls {
    private static final Null value = Null.newBuilder().build();

    private Nulls() {}

    public static Null instance() {
        return value;
    }
}
