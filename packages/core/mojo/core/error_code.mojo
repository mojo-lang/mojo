// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/// Error
///
/// ErrorCode is usually static and predefined in a specific domain
/// "system.121"
/// "system.501:NOT_SUPPORT"
/// code `[100..599]` is reserved by duplicated with http status code
/// code `0` is also reserved because of it usually has success meaning
@format('{domain.}{code}{:name}')
@encoding_as_struct
type ErrorCode {
    ///
    code: Int32 @1 @required

    /// The name of the error code. This is a constant value that identifies the
    /// error code. Error code name are unique within a particular
    /// domain of errors. This should be at most 63 characters and match
    /// /[A-Z0-9_]+/.
    name: String @2

    /// system, runtime, ...
    domain: String @3
    
    /// a detail description for the code
    description: String @5
    
    /// the api document url for this error code
    document: Url @4

    /// the http status code, which is error code mapping to
    http_status_code: Int32 @10
}

/// The canonical error codes for gRPC APIs.
///
///
/// Sometimes multiple error codes may apply.  Services should return
/// the most specific error code that applies.  For example, prefer
/// `OUT_OF_RANGE` over `FAILED_PRECONDITION` if both codes apply.
/// Similarly prefer `NOT_FOUND` or `ALREADY_EXISTS` over `FAILED_PRECONDITION`.
const {
    /// Not an error; returned on success
    ///
    /// HTTP Mapping: 200 OK
    ok = ErrorCode{code: 0, name: "OK", http_status_code: 200}

    /// The operation was cancelled, typically by the caller.
    ///
    /// HTTP Mapping: 499 Client Closed Request
    cancelled = ErrorCode{code: 499, name: "CANCELLED", http_status_code: 499}

    /// Unknown error.  For example, this error may be returned when
    /// a `Status` value received from another address space belongs to
    /// an error space that is not known in this address space.  Also
    /// errors raised by APIs that do not return enough error information
    /// may be converted to this error.
    ///
    /// HTTP Mapping: 500 Internal Server Error
    unknown_error = ErrorCode{code: 2, name: "UNKNOWN_ERROR", http_status_code: 500}

    /// The client specified an invalid argument.  Note that this differs
    /// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments
    /// that are problematic regardless of the state of the system
    /// (e.g., a malformed file name).
    ///
    /// HTTP Mapping: 400 Bad Request
    invalid_argument = ErrorCode{code: 3, name: "INVALID_ARGUMENT", http_status_code: 400}

    /// The deadline expired before the operation could complete. For operations
    /// that change the state of the system, this error may be returned
    /// even if the operation has completed successfully.  For example, a
    /// successful response from a server could have been delayed long
    /// enough for the deadline to expire.
    ///
    /// HTTP Mapping: 504 Gateway Timeout
    deadline_exceeded = ErrorCode{code: 504, name: "DEADLINE_EXCEEDED", http_status_code: 504}

    /// Some requested entity (e.g., file or directory) was not found.
    ///
    /// Note to server developers: if a request is denied for an entire class
    /// of users, such as gradual feature rollout or undocumented whitelist,
    /// `NOT_FOUND` may be used. If a request is denied for some users within
    /// a class of users, such as user-based access control, `PERMISSION_DENIED`
    /// must be used.
    ///
    /// HTTP Mapping: 404 Not Found
    not_found = ErrorCode{code: 404, name: "NOT_FOUND", http_status_code: 404}

    /// The entity that a client attempted to create (e.g., file or directory)
    /// already exists.
    ///
    /// HTTP Mapping: 409 Conflict
    already_exists = ErrorCode{code: 6, name: "ALREADY_EXISTS", http_status_code: 409}

    /// The caller does not have permission to execute the specified
    /// operation. `PERMISSION_DENIED` must not be used for rejections
    /// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
    /// instead for those errors). `PERMISSION_DENIED` must not be
    /// used if the caller can not be identified (use `UNAUTHENTICATED`
    /// instead for those errors). This error code does not imply the
    /// request is valid or the requested entity exists or satisfies
    /// other pre-conditions.
    ///
    /// HTTP Mapping: 403 Forbidden
    permission_denied = ErrorCode{code: 403, name: "PERMISSION_DENIED", http_status_code: 403}

    /// The request does not have valid authentication credentials for the
    /// operation.
    ///
    /// HTTP Mapping: 401 Unauthorized
    unauthenticated = ErrorCode{code: 401, name: "UNAUTHENTICATED", http_status_code: 401}

    /// Some resource has been exhausted, perhaps a per-user quota, or
    /// perhaps the entire file system is out of space.
    ///
    /// HTTP Mapping: 429 Too Many Requests
    resource_exhausted = ErrorCode{code: 429, name: "RESOURCE_EXHAUSTED", http_status_code: 429}

    /// The operation was rejected because the system is not in a state
    /// required for the operation's execution.  For example, the directory
    /// to be deleted is non-empty, an rmdir operation is applied to
    /// a non-directory, etc.
    ///
    /// Service implementors can use the following guidelines to decide
    /// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
    ///  (a) Use `UNAVAILABLE` if the client can retry just the failing call.
    ///  (b) Use `ABORTED` if the client should retry at a higher level
    ///      (e.g., when a client-specified test-and-set fails, indicating the
    ///      client should restart a read-modify-write sequence).
    ///  (c) Use `FAILED_PRECONDITION` if the client should not retry until
    ///      the system state has been explicitly fixed.  E.g., if an "rmdir"
    ///      fails because the directory is non-empty, `FAILED_PRECONDITION`
    ///      should be returned since the client should not retry unless
    ///      the files are deleted from the directory.
    ///
    /// HTTP Mapping: 400 Bad Request
    failed_precondition = ErrorCode{code: 9, name: "FAILED_PRECONDITION", http_status_code: 400}

    /// The operation was aborted, typically due to a concurrency issue such as
    /// a sequencer check failure or transaction abort.
    ///
    /// See the guidelines above for deciding between `FAILED_PRECONDITION`,
    /// `ABORTED`, and `UNAVAILABLE`.
    ///
    /// HTTP Mapping: 409 Conflict
    aborted = ErrorCode{code: 10, name: "ABORTED", http_status_code: 409}

    /// The operation was attempted past the valid range.  E.g., seeking or
    /// reading past end-of-file.
    ///
    /// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
    /// be fixed if the system state changes. For example, a 32-bit file
    /// system will generate `INVALID_ARGUMENT` if asked to read at an
    /// offset that is not in the range [0,2^32-1], but it will generate
    /// `OUT_OF_RANGE` if asked to read from an offset past the current
    /// file size.
    ///
    /// There is a fair bit of overlap between `FAILED_PRECONDITION` and
    /// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
    /// error) when it applies so that callers who are iterating through
    /// a space can easily look for an `OUT_OF_RANGE` error to detect when
    /// they are done.
    ///
    /// HTTP Mapping: 400 Bad Request
    out_of_range = ErrorCode{code: 11, name: "OUT_OF_RANGE", http_status_code: 400}

    /// The operation is not implemented or is not supported/enabled in this
    /// service.
    ///
    /// HTTP Mapping: 501 Not Implemented
    unimplemented = ErrorCode{code: 501, name: "UNIMPLEMENTED", http_status_code: 501}

    /// Internal errors.  This means that some invariants expected by the
    /// underlying system have been broken.  This error code is reserved
    /// for serious errors.
    ///
    /// HTTP Mapping: 500 Internal Server Error
    internal_error = ErrorCode{code: 500, name: "INTERNAL_ERROR", http_status_code: 500}

    /// The service is currently unavailable.  This is most likely a
    /// transient condition, which can be corrected by retrying with
    /// a backoff. Note that it is not always safe to retry
    /// non-idempotent operations.
    ///
    /// See the guidelines above for deciding between `FAILED_PRECONDITION`,
    /// `ABORTED`, and `UNAVAILABLE`.
    ///
    /// HTTP Mapping: 503 Service Unavailable
    unavailable = ErrorCode{code: 503, name: "UNAVAILABLE", http_status_code: 503}

    /// Unrecoverable data loss or corruption.
    ///
    /// HTTP Mapping: 500 Internal Server Error
    data_loss = ErrorCode{code: 15, name: "DATA_LOSS", http_status_code: 500}
}
