package org.mojolang.mojo.core;


import com.alibaba.fastjson2.JSONWriter;
import com.google.common.collect.Lists;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

import com.alibaba.fastjson2.support.config.FastJsonConfig;
import com.alibaba.fastjson2.support.spring.http.converter.FastJsonHttpMessageConverter;

import java.util.Arrays;
import java.util.List;

@EnableWebMvc
@Configuration
class WebMvcConfig implements WebMvcConfigurer {

//    @Override
//    public void extendMessageConverters(List<HttpMessageConverter<?>> converters) {
//        JsonHttpMessageConverter converter =
//                new JsonHttpMessageConverter();
//
//        converter.setSupportedMediaTypes(
//                Arrays.asList(
//                        MediaType.APPLICATION_JSON,
//                        MediaType.APPLICATION_JSON_UTF8
//                )
//        );
//        converters.add(0, converter);
//
//
//    }

    @Override
    public void extendMessageConverters(List<HttpMessageConverter<?>> converters) {
        FastJsonConfig fastJsonConfig = new FastJsonConfig();
        fastJsonConfig.setWriterFeatures(JSONWriter.Feature.WriteMapNullValue);

        FastJsonHttpMessageConverter fastJsonHttpMessageConverter = new FastJsonHttpMessageConverter();
        fastJsonHttpMessageConverter.setFastJsonConfig(fastJsonConfig);
        fastJsonHttpMessageConverter.setSupportedMediaTypes(Lists.newArrayList(MediaType.APPLICATION_JSON_UTF8));
        converters.add(0, fastJsonHttpMessageConverter);
    }

//    @Override
//    public void configureMessageConverters(List<HttpMessageConverter<?>> converters) {
//
//        JsonHttpMessageConverter converter =
//                new JsonHttpMessageConverter();
//
//        converter.setSupportedMediaTypes(
//                Arrays.asList(
//                        MediaType.APPLICATION_JSON,
//                        MediaType.APPLICATION_JSON_UTF8
//                )
//        );
//        converters.add(converter);
//    }
}