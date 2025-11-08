package org.mojolang.mojo.core;

public class ErrorException extends Exception {
    private final Error error;

    public Error toError() {
        return error;
    }

    public ErrorException(ErrorCode code) {
        super(code.getDescription());

        this.error = Error.newBuilder().setCode(code).build();
    }

    public ErrorException(Error error) {
        super(error.getMessage());
        this.error = error;
    }

    public ErrorException(Throwable throwable) {
        super(throwable.getMessage());
        this.error = Error.newBuilder().setCode(ErrorCodes.INTERNAL_ERROR).setMessage(throwable.getMessage()).build();
    }

    public ErrorException(ErrorCode code, String message) {
        super(message);
        this.error = Error.newBuilder().setCode(code).setMessage(message).build();
    }

    public int getCode() {
        return error.getCode().getCode();
    }
}
