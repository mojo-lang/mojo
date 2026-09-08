package org.mojolang.mojo.core;

import java.util.HashMap;
import java.util.Map;

public final class ErrorCodes {
    private static final Map<Integer, ErrorCode> codes = new HashMap<Integer, ErrorCode>();

    public static ErrorCode SUCCESS = ErrorCode.newBuilder()
            .setCode(0)
            .setName("OK")
            .setHttpStatusCode(200).build();

    public static ErrorCode BAD_REQUEST = ErrorCode.newBuilder()
            .setCode(400)
            .setHttpStatusCode(400)
            .setName("BAD_REQUEST")
            .setDescription("The request could not be understood by the server due to malformed syntax.")
            .build();

    public static ErrorCode INTERNAL_ERROR = ErrorCode.newBuilder()
            .setCode(500)
            .setHttpStatusCode(500)
            .setName("INTERNAL_ERROR")
            .setDescription("This means that some invariants expected by the underlying system have been broken.")
            .build();
    public static ErrorCode UNIMPLEMENTED = ErrorCode.newBuilder()
            .setCode(501)
            .setHttpStatusCode(501)
            .setName("UNIMPLEMENTED")
            .setDescription("The operation is not implemented or is not supported/enabled in this service.")
            .build();

    static {
        codes.put(0, SUCCESS);
        codes.put(400, BAD_REQUEST);
        codes.put(500, INTERNAL_ERROR);
        codes.put(501, UNIMPLEMENTED);
    }

    /**
     * build a new ErrorCode if the code is not predefined.
     *
     * @param code the code number
     * @return ErrorCode
     */
    public static ErrorCode build(int code) {
        if (codes.containsKey(code)) {
            return codes.get(code);
        } else {
            return ErrorCode.newBuilder().setCode(code).build();
        }
    }
}
