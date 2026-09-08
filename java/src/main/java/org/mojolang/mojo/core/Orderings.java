package org.mojolang.mojo.core;

import com.google.common.base.Optional;
import com.google.protobuf.Descriptors;

import java.util.Arrays;

public final class Orderings {
    private static final String ORDERING_FIELD_SEPARATOR = ",";
    private static final String ORDER_SORT_SEPARATOR = " ";

    private Orderings() {
    }

    public static Ordering fromString(String value) {
        return fromStringList(Arrays.asList(value.split(ORDERING_FIELD_SEPARATOR)));
    }

    public static Ordering fromStringList(Iterable<String> paths) {
        Ordering.Builder builder = Ordering.newBuilder();
        for (String path : paths) {
            String[] parts = path.split(ORDER_SORT_SEPARATOR);
            int sort = 0;
            if (parts.length == 2) {
                switch (parts[1].trim().toLowerCase()) {
                    case "asc":
                        sort = 1;
                    case "desc":
                        sort = 2;
                }

                builder.addOrders(Ordering.Order.newBuilder()
                        .setField(parts[0].trim())
                        .setSort(Ordering.Sort.forNumber(sort))
                        .build());
            }
        }
        return builder.build();
    }
}
