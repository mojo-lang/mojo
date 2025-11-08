package org.mojolang.mojo.core;

import com.alibaba.fastjson2.JSON;
import com.alibaba.fastjson2.JSONObject;
import com.google.protobuf.MessageOrBuilder;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

import java.util.ArrayList;
import java.util.List;

import static com.google.common.truth.Truth.assertThat;

@RunWith(JUnit4.class)
public class JsonObjectReaderTest {
    @Test
    public void testSimpleParser() {
        JsonObjectRegister.init();

        StringValue value = JSON.parseObject("\"xx\"", StringValue.class);
        assertThat(value.getVal()).isEqualTo("xx");
    }

    @Test
    public void testResult() {
        JsonObjectRegister.init();

        String json = "{\"result\":{\"code\": 200, \"message\":\"OK\", \"data\":{\"key\":\"value\"}}}";

        ResultWrap result = JSON.parseObject(json, ResultWrap.class);
        assertThat(result).isNotNull();
    }

    private static class ResultWrap {
        public Result<org.mojolang.mojo.core.Object> result;
    }

    private static class PaginationWrap {
        public Pagination<org.mojolang.mojo.core.Object> pagination;
    }

    @Test
    public void testPagination() {
        JsonObjectRegister.init();

        String json = "{\"pagination\":{\"code\": 200, \"message\":\"OK\", \"data\":[{\"key\":\"value\"}]}}";

        PaginationWrap pagination = JSON.parseObject(json, PaginationWrap.class);
        assertThat(pagination).isNotNull();
    }
}
