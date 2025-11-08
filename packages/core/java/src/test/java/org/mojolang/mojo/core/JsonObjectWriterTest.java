package org.mojolang.mojo.core;

import com.alibaba.fastjson2.JSON;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import java.util.List;
import java.util.ArrayList;

import static com.google.common.truth.Truth.assertThat;

@RunWith(JUnit4.class)
public class JsonObjectWriterTest {
    @Test
    public void testSimpleWriter() {
        JsonObjectRegister.init();

        StringValue value = StringValue.newBuilder().setVal("xx").build();
        String out = JSON.toJSONString(value);
        assertThat(out).isEqualTo("\"xx\"");
    }
    @Test
    public void testResult() {
        JsonObjectRegister.init();

        Result<Object> result = new Result<>();
        result.setData(Objects.of("key", ValueUtil.of("value")));

        String out = JSON.toJSONString(result);
        assertThat(out).isNotEmpty();
    }

    @Test
    public void testPagination() {
        JsonObjectRegister.init();

        Pagination<Object> pagination = new Pagination<>();

        List<Object> list = new ArrayList<>();
        Object obj1 = Objects.of("key1", ValueUtil.of("value1"));
        Object obj2 = Objects.of("key2", ValueUtil.of("value2"));
        list.add(obj1);
        list.add(obj2);
        pagination.setData(list);

        String out = JSON.toJSONString(pagination);
        assertThat(out).isNotEmpty();
    }
}
