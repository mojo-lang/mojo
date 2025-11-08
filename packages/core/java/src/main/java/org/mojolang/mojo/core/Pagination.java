package org.mojolang.mojo.core;

import com.alibaba.fastjson2.annotation.JSONField;
import lombok.Data;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Pagination<T> {
    @JSONField(ordinal = 0)
    private int code;

    @JSONField(ordinal = 1)
    private String message;

    @JSONField(ordinal = 2)
    private int totalCount;

    @JSONField(ordinal = 3)
    private String nextPageToken;

    @JSONField(ordinal = 10)
    private List<T> data;

    public Pagination<T> setError(Error error) {
        this.code = error.getCode().getCode();
        this.message = error.getMessage();
        return this;
    }
    public Pagination<T> setError(ErrorCode code) {
        this.code = code.getCode();
        this.message = code.getName();
        return this;
    }
    public Pagination<T> setError(Integer code) {
        this.code = code;
        return this;
    }
    public Pagination<T> setError(Integer code, String msg) {
        this.code = code;
        this.message = msg;
        return this;
    }

    public Pagination<T> setTotalCount(int count) {
        this.totalCount = count;
        return this;
    }

    public Pagination<T> setNextPageToken(String token) {
        this.nextPageToken = token;
        return this;
    }

    public Pagination<T> setData(List<T> data) {
        this.data = data;
        return this;
    }

    public static <T> Pagination<T> fail() {
        return build(ErrorCodes.BAD_REQUEST, null, 0, "");
    }

    public static <T> Pagination<T> fail(int code) {
        return build(code, null, 0, "");
    }

    public static <T> Pagination<T> fail(int code, String msg) {
        return build(code, msg, null, 0, "");
    }

    public static <T> Pagination<T> fail(ErrorException exception) {
        return build(exception, null, 0, "");
    }

    public static <T> Pagination<T> build(ErrorCode code, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<T>().setError(code).setData(data).setTotalCount(totalCount).setNextPageToken(nextPageToken);
    }

    public static <T> Pagination<T> build(Integer code, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<T>().setError (code).setData(data).setTotalCount(totalCount).setNextPageToken(nextPageToken);
    }

    public static <T> Pagination<T> build(Integer code, String msg, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<T>().setError(code, msg).setData(data).setTotalCount(totalCount).setNextPageToken(nextPageToken);
    }

    public static <T> Pagination<T> build(ErrorException exception, List<T> data, int totalCount, String nextPageToken) {
        return new Pagination<T>().setError(exception.toError()).setData(data).setTotalCount(totalCount).setNextPageToken(nextPageToken);
    }
}