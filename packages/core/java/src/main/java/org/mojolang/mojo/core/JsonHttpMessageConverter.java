package org.mojolang.mojo.core;

import com.alibaba.fastjson2.JSON;
import com.alibaba.fastjson2.JSONException;
import com.alibaba.fastjson2.JSONPObject;
import com.alibaba.fastjson2.support.config.FastJsonConfig;
import com.google.protobuf.MessageOrBuilder;
import org.springframework.core.ResolvableType;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpInputMessage;
import org.springframework.http.HttpOutputMessage;
import org.springframework.http.MediaType;
import org.springframework.http.converter.AbstractHttpMessageConverter;
import org.springframework.http.converter.GenericHttpMessageConverter;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.http.converter.HttpMessageNotWritableException;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.lang.reflect.ParameterizedType;
import java.lang.reflect.Type;
import java.lang.reflect.TypeVariable;
import java.util.List;

import com.google.protobuf.Message;
import com.google.protobuf.MessageOrBuilder;

/**
 * Fastjson for Spring MVC Converter.
 *
 * @author Victor.Zxy
 * @see AbstractHttpMessageConverter
 * @see GenericHttpMessageConverter
 * @since 2.0.2
 */
public class JsonHttpMessageConverter
        extends AbstractHttpMessageConverter<java.lang.Object>
        implements GenericHttpMessageConverter<java.lang.Object> {
    public static final MediaType APPLICATION_JAVASCRIPT = new MediaType("application", "javascript");
    public static final MediaType APPLICATION_JSON = new MediaType("application", "json");

    /**
     * with fastJson config
     */
    private FastJsonConfig config = new FastJsonConfig();

    /**
     * Can serialize/deserialize all types.
     */
    public JsonHttpMessageConverter() {
        super(MediaType.ALL);
    }

    /**
     * @return the fastJsonConfig.
     */
    public FastJsonConfig getFastJsonConfig() {
        return config;
    }

    /**
     * @param fastJsonConfig the fastJsonConfig to set.
     */
    public void setFastJsonConfig(FastJsonConfig fastJsonConfig) {
        this.config = fastJsonConfig;
    }

    @Override
    protected boolean supports(Class<?> clazz) {
        return true;
    }

    @Override
    public boolean canRead(Type type, Class<?> contextClass, MediaType mediaType) {
        return super.canRead(contextClass, mediaType);
    }

    @Override
    public boolean canWrite(Type type, Class<?> clazz, MediaType mediaType) {
        return super.canWrite(clazz, mediaType);
    }

    @Override
    public java.lang.Object read(Type type, Class<?> contextClass, HttpInputMessage inputMessage) throws IOException, HttpMessageNotReadableException {
        return readType(getType(type, contextClass), inputMessage);
    }

    @Override
    public void write(java.lang.Object o, Type type, MediaType contentType, HttpOutputMessage outputMessage) throws IOException, HttpMessageNotWritableException {
        // support StreamingHttpOutputMessage in spring4.0+
        super.write(o, contentType, outputMessage);
    }

    @Override
    protected java.lang.Object readInternal(Class<? extends java.lang.Object> clazz, HttpInputMessage inputMessage) throws IOException, HttpMessageNotReadableException {
        return readType(getType(clazz, null), inputMessage);
    }

    private java.lang.Object readType(Type type, HttpInputMessage inputMessage) {
        try (ByteArrayOutputStream baos = new ByteArrayOutputStream()) {
            InputStream in = inputMessage.getBody();

            byte[] buf = new byte[1024 * 64];
            for (; ; ) {
                int len = in.read(buf);
                if (len == -1) {
                    break;
                }

                if (len > 0) {
                    baos.write(buf, 0, len);
                }
            }


            if (type instanceof Class && Message.class.isAssignableFrom((Class<?>) type)) {
                return JsonFormat.parser().fromJSON(baos.toString(), (Class<? extends Message>) type);
            }
            if (type instanceof ParameterizedType) {
                final ParameterizedType paramType = (ParameterizedType) type;

                if (List.class.isAssignableFrom((Class<?>) paramType.getRawType())) {
                    final Type argument = paramType.getActualTypeArguments()[0];
                    if (Message.class.isAssignableFrom((Class<?>) argument)) {
                        return  JsonFormat.parser().listFromJSON (baos.toString(), (Class<? extends Message>) argument);
                    }
                }
//                if (Pagination.class.isAssignableFrom((Class<?>) paramType.getRawType())) {
//                    final Type argument = paramType.getActualTypeArguments()[0];
//                    if (Message.class.isAssignableFrom((Class<?>) argument)) {
//                        return JsonFormat.parser().paginationFromJSON(baos.toString(), (Class<? extends  Message>) argument);
//                    }
//                }
//
//                if (Result.class.isAssignableFrom((Class<?>) paramType.getRawType())) {
//                    final Type argument = paramType.getActualTypeArguments()[0];
//                    if (Message.class.isAssignableFrom((Class<?>) argument)) {
//                        return JsonFormat.parser().resultFromJSON(baos.toString(), (Class<? extends  Message>) argument);
//                    }
//                }
            }

            byte[] bytes = baos.toByteArray();
            return JSON.parseObject(bytes, type, config.getDateFormat(), config.getReaderFilters(), config.getReaderFeatures());
        } catch (JSONException ex) {
            throw new HttpMessageNotReadableException("JSON parse error: " + ex.getMessage(), ex, inputMessage);
        } catch (IOException ex) {
            throw new HttpMessageNotReadableException("I/O error while reading input message", ex, inputMessage);
        }
    }

    @Override
    protected void writeInternal(java.lang.Object object, HttpOutputMessage outputMessage) throws IOException, HttpMessageNotWritableException {
        HttpHeaders headers = outputMessage.getHeaders();

        try (ByteArrayOutputStream baos = new ByteArrayOutputStream()) {
            int contentLength = 0;
            if (object instanceof String && JSON.isValidObject((String) object)) {
                byte[] strBytes = ((String) object).getBytes(config.getCharset());
                contentLength = strBytes.length;
                outputMessage.getBody().write(strBytes, 0, strBytes.length);
            } else if (object instanceof byte[] && JSON.isValid((byte[]) object)) {
                byte[] strBytes = (byte[]) object;
                contentLength = strBytes.length;
                outputMessage.getBody().write(strBytes, 0, strBytes.length);
            } else {
                if (object instanceof JSONPObject) {
                    headers.setContentType(APPLICATION_JAVASCRIPT);
                }

                String out = "";
                if (object instanceof MessageOrBuilder) {
                    out = JsonFormat.printer().print((MessageOrBuilder) object);
                } else if (object instanceof List) {
                    out = JsonFormat.printer().print((List<? extends MessageOrBuilder>) object);
                }
//                else if (object instanceof Result) {
//                    out = JsonFormat.printer().print((Result<? extends MessageOrBuilder>) object);
//                } else if (object instanceof Pagination) {
//                    out = JsonFormat.printer().print((Pagination<? extends MessageOrBuilder>) object);
//                }

                byte[] strBytes = out.getBytes(config.getCharset());
                contentLength = strBytes.length;
                outputMessage.getBody().write(strBytes, 0, strBytes.length);

                if (out.isEmpty()) {
                    contentLength = JSON.writeTo(
                            baos, object,
                            config.getDateFormat(),
                            config.getWriterFilters(),
                            config.getWriterFeatures()
                    );
                }
            }

            if (headers.getContentLength() < 0 && config.isWriteContentLength()) {
                headers.setContentLength(contentLength);
            }
            baos.writeTo(outputMessage.getBody());
        } catch (JSONException ex) {
            throw new HttpMessageNotWritableException("Could not write JSON: " + ex.getMessage(), ex);
        } catch (IOException ex) {
            throw new HttpMessageNotWritableException("I/O error while writing output message", ex);
        }
    }

    protected Type getType(Type type, Class<?> contextClass) {
        if (Spring4TypeResolvableHelper.isSupport()) {
            return Spring4TypeResolvableHelper.getType(type, contextClass);
        }
        return type;
    }

    private static class Spring4TypeResolvableHelper {
        private static boolean hasClazzResolvableType;

        static {
            try {
                Class.forName("org.springframework.core.ResolvableType");
                hasClazzResolvableType = true;
            } catch (ClassNotFoundException e) {
                hasClazzResolvableType = false;
            }
        }

        private static boolean isSupport() {
            return hasClazzResolvableType;
        }

        private static Type getType(Type type, Class<?> contextClass) {
            if (contextClass != null) {
                ResolvableType resolvedType = ResolvableType.forType(type);
                if (type instanceof TypeVariable) {
                    ResolvableType resolvedTypeVariable = resolveVariable((TypeVariable) type, ResolvableType.forClass(contextClass));
                    if (resolvedTypeVariable != ResolvableType.NONE) {
                        return resolvedTypeVariable.resolve();
                    }
                } else if (type instanceof ParameterizedType && resolvedType.hasUnresolvableGenerics()) {
                    ParameterizedType parameterizedType = (ParameterizedType) type;
                    Class<?>[] generics = new Class[parameterizedType.getActualTypeArguments().length];
                    Type[] typeArguments = parameterizedType.getActualTypeArguments();

                    for (int i = 0; i < typeArguments.length; ++i) {
                        Type typeArgument = typeArguments[i];
                        if (typeArgument instanceof TypeVariable) {
                            ResolvableType resolvedTypeArgument = resolveVariable((TypeVariable) typeArgument, ResolvableType.forClass(contextClass));
                            if (resolvedTypeArgument != ResolvableType.NONE) {
                                generics[i] = resolvedTypeArgument.resolve();
                            } else {
                                generics[i] = ResolvableType.forType(typeArgument).resolve();
                            }
                        } else {
                            generics[i] = ResolvableType.forType(typeArgument).resolve();
                        }
                    }

                    return ResolvableType.forClassWithGenerics(resolvedType.getRawClass(), generics).getType();
                }
            }

            return type;
        }

        private static ResolvableType resolveVariable(TypeVariable<?> typeVariable, ResolvableType contextType) {
            ResolvableType resolvedType;
            if (contextType.hasGenerics()) {
                resolvedType = ResolvableType.forType(typeVariable, contextType);
                if (resolvedType.resolve() != null) {
                    return resolvedType;
                }
            }

            ResolvableType superType = contextType.getSuperType();
            if (superType != ResolvableType.NONE) {
                resolvedType = resolveVariable(typeVariable, superType);
                if (resolvedType.resolve() != null) {
                    return resolvedType;
                }
            }
            for (ResolvableType ifc : contextType.getInterfaces()) {
                resolvedType = resolveVariable(typeVariable, ifc);
                if (resolvedType.resolve() != null) {
                    return resolvedType;
                }
            }
            return ResolvableType.NONE;
        }
    }
}
