package org.mojolang.mojo.core;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import java.util.Map;

import static com.google.common.truth.Truth.assertThat;

@RunWith(JUnit4.class)
public class ValueUtilTest {

    @Test
    public void testValueConvert() {
        org.mojolang.mojo.core.Object obj = Object.newBuilder()
                .putVals("foo", Value.newBuilder().setPositiveVal(12).build())
                .putVals("bar", Value.newBuilder().setDoubleVal(10.12).build())
                .putVals("f", Value.newBuilder().setBoolVal(true).build())
                .putVals("b", Value.newBuilder().setNegativeVal(10).build())
                .putVals("baz", Value.newBuilder().setStringVal("foobar").build()).build();
        Map<String, java.lang.Object> map = Objects.to(obj);

        org.mojolang.mojo.core.Object o = Objects.of(map);
        assertThat(obj.getValsCount()).isEqualTo(o.getValsCount());
        assertThat(obj.getValsMap().get("foo").getPositiveVal()).isEqualTo(o.getValsMap().get("foo").getPositiveVal());
    }
}
