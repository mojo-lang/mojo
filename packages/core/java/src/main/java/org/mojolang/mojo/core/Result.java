package org.mojolang.mojo.core;

import com.alibaba.fastjson2.annotation.JSONField;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import lombok.Data;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Result<T> {
    private int code;
    private String message;

    @JSONField(ordinal = 10)
    private T data;

    public Result<T> setError(Error error) {
        this.code = error.getCode().getCode();
        this.message = error.getMessage();
        return this;
    }
    public Result<T> setError(ErrorCode code) {
        this.code = code.getCode();
        this.message = code.getName();
        return this;
    }
    public Result<T> setError(Integer code) {
        this.code = code;
        return this;
    }
    public Result<T> setError(Integer code, String msg) {
        this.code = code;
        this.message = msg;
        return this;
    }

    public Result<T> setData(T data) {
        this.data = data;
        return this;
    }

    public static <T> Result<T> success() {
        return build(ErrorCodes.SUCCESS,null);
    }

    public static <T> Result<T> success(T data) {
        return build(ErrorCodes.SUCCESS, data);
    }

    public static <T> Result<T> fail() {
        return build(ErrorCodes.BAD_REQUEST, null);
    }

    public static <T> Result<T> fail(T data) {
        return build(ErrorCodes.BAD_REQUEST, data);
    }

    public static <T> Result<T> fail(int code) {
        return build(code, null);
    }

    public static <T> Result<T> fail(int code, T data) {
        return build(code, data);
    }

    public static <T> Result<T> fail(int code, String msg) {
        return build(code, msg, null);
    }

    public static <T> Result<T> fail(int code, String msg, T data) {
        return build(code, msg, data);
    }

    public static <T> Result<T> fail(ErrorException exception) {
        return fail(exception, null);
    }

    public static <T> Result<T> fail(ErrorException exception, T data) {
        return new Result<T>().setError(exception.toError()).setData(data);
    }

    public static <T> Result<T> build(ErrorCode code, T data) {
        return new Result<T>().setError(code).setData(data);
    }

    public static <T> Result<T> build(Integer code, T data) {
        return new Result<T>().setError(code).setData(data);
    }

    public static <T> Result<T> build(Integer code, String msg, T data) {
        return new Result<T>().setError(code, msg).setData(data);
    }
}