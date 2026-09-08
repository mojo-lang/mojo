package org.mojolang.mojo.core;

import static org.hamcrest.Matchers.equalTo;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import com.alibaba.fastjson2.JSONObject;
import org.junit.jupiter.api.Test;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.mock.web.MockHttpServletResponse;
import org.springframework.test.web.servlet.MockMvc;

import org.junit.Before;
import org.junit.runner.RunWith;
import org.mockito.MockitoAnnotations;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.http.MediaType;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.context.web.WebAppConfiguration;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

@SpringBootApplication
class TestApplication {
    public static void main(String[] args) {
        SpringApplication.run(TestApplication.class, args);
    }
}

@SpringBootTest
@AutoConfigureMockMvc
public class JsonControllerTest {
    @Autowired
    private MockMvc mvc;

    @Test
    public void simple() throws Exception {
        mvc.perform(get("/simple").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string(equalTo("{\n" +
                        "  \"Hello\": \" World\"\n" +
                        "}")));
    }

    @Test
    public void simpleResult() throws Exception {
        JsonObjectRegister.init();
        mvc.perform(get("/simple/result").accept(MediaType.APPLICATION_JSON_UTF8))
                .andExpect(status().isOk())
                .andExpect(content().string(equalTo("{\"code\":200,\"message\":\"OK\",\"data\":{\n  \"Hello\": \" World\"\n}}")));
    }

    @Test
    public void simplePagination() throws Exception {
        JsonObjectRegister.init();
        mvc.perform(get("/simple/pagination").accept(MediaType.APPLICATION_JSON))
                .andExpect(status().isOk())
                .andExpect(content().string(equalTo("{\"code\":200,\"message\":\"OK\",\"totalCount\":100,\"nextPageToken\":\"1\",\"data\":[{\n  \"Hello\": \" World\"\n}]}")));
    }

    @Test
    public void requestBody() throws Exception {
        mvc.perform(post("/request-body")
                        .contentType(MediaType.APPLICATION_JSON_VALUE)
                        .accept(MediaType.APPLICATION_JSON)
                        .content("{\"name\":\"world\"}".getBytes())
                )
                .andExpect(status().isOk())
                .andExpect(content().string(equalTo("{\n  \"Hello\": \"world\"\n}")));
    }

    @Test
    public void requestListBody() throws Exception {
        mvc.perform(post("/request-list-body")
                        .contentType(MediaType.APPLICATION_JSON_VALUE)
                        .accept(MediaType.APPLICATION_JSON)
                        .content("[{\"name\":\"world\"},{\"name\":\"universe\"}]".getBytes())
                )
                .andExpect(status().isOk())
                .andExpect(content().string(equalTo("[{\n  \"Hello\": \"world\"\n},{\n  \"Hello\": \"universe\"\n}]")));
    }
//
//    @Test
//    public void mixedReq() throws Exception {
//        mvc.perform(
//                        MockMvcRequestBuilders
//                                .post("/mixed-req")
//                                .contentType(MediaType.APPLICATION_JSON_VALUE)
//                                .accept(MediaType.APPLICATION_JSON)
//                                .content("{\"req\":{\"name\":\"world\"}}".getBytes())
//                )
//                .andExpect(status().isOk())
//                .andExpect(content().string(equalTo("{\"message\":\"hello world\"}")));
//    }
//
//    @Test
//    public void mixedReqUnexpectedJson() throws Exception {
//        mvc.perform(
//                        MockMvcRequestBuilders
//                                .post("/mixed-req")
//                                .contentType(MediaType.APPLICATION_JSON_VALUE)
//                                .accept(MediaType.APPLICATION_JSON)
//                                .content("{\"req\":\"wrong\"".getBytes())
//                )
//                .andExpect(status().isBadRequest());
//        // .andExpect(content().string(equalTo("{\"message\":\"hello world\"}")));
//    }
//
//    @Test
//    public void mixed() throws Exception {
//        mvc.perform(
//                        MockMvcRequestBuilders
//                                .post("/mixed")
//                                .contentType(MediaType.APPLICATION_JSON_VALUE)
//                                .accept(MediaType.APPLICATION_JSON)
//                                .content("{\"req\":{\"name\":\"world\"}}".getBytes())
//                )
//                .andExpect(status().isOk())
//                .andExpect(content().string(equalTo("{\"res\":{\"message\":\"hello world\"}}")));
//    }
//
//    @Test
//    public void deep() throws Exception {
//        mvc.perform(
//                        MockMvcRequestBuilders
//                                .post("/deep")
//                                .contentType(MediaType.APPLICATION_JSON_VALUE)
//                                .accept(MediaType.APPLICATION_JSON)
//                                .content("{\"req\":{\"name\":\"world\"}}".getBytes())
//                )
//                .andExpect(status().isOk())
//                .andExpect(content().string(equalTo("{\"res\":{\"message\":\"hello world\"}}")));
//    }
//
//    @Test
//    public void deepWithWrongRequestJson() throws Exception {
//        mvc.perform(
//                        MockMvcRequestBuilders
//                                .post("/deep")
//                                .contentType(MediaType.APPLICATION_JSON_VALUE)
//                                .accept(MediaType.APPLICATION_JSON)
//                                .content("{\"req\":{\"name\":\"world".getBytes())
//                )
//                .andExpect(status().isBadRequest());
//    }
}