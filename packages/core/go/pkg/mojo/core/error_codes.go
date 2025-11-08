package core

var errorCodeIndex map[int32]*ErrorCode
var errorCodeNameIndex map[string]*ErrorCode

func init() {
	errorCodeIndex = make(map[int32]*ErrorCode)
	errorCodeNameIndex = make(map[string]*ErrorCode)

	addIndex := func(codes ...*ErrorCode) {
		for _, code := range codes {
			errorCodeIndex[code.Code] = code
			errorCodeNameIndex[code.Name] = code
		}
	}

	addIndex(BadRequest, InvalidArgument, MalformedSyntax, FailedPrecondition, OutOfRange, Unauthenticated, PermissionDenied,
		NotFound, AlreadyExists, Aborted, ResourceExhausted, Cancelled, UnknownError, DeadlineExceeded,
		Unimplemented, InternalError,
		Unavailable, DataLoss)
}

var BadRequest = &ErrorCode{
	Code:           400,
	Name:           "BAD_REQUEST",
	Description:    "The request could not be understood by the server due to malformed syntax. ",
	HttpStatusCode: 400,
}

var InvalidArgument = &ErrorCode{
	Code:           3,
	Name:           "INVALID_ARGUMENT",
	Description:    "The client specified an invalid argument.",
	HttpStatusCode: 400,
}

var MalformedSyntax = &ErrorCode{
	Code:           5,
	Name:           "MALFORMED_SYNTAX",
	Description:    "The syntax of the requested string is malformed.",
	HttpStatusCode: 400,
}

var FailedPrecondition = &ErrorCode{
	Code:           9,
	Name:           "FAILED_PRECONDITION",
	Description:    "The operation was rejected because the system is not in a state required for the operation's execution.",
	HttpStatusCode: 400,
}

var OutOfRange = &ErrorCode{
	Code:           11,
	Name:           "OUT_OF_RANGE",
	Description:    "The operation was attempted past the invalid range.",
	HttpStatusCode: 400,
}

var Unauthenticated = &ErrorCode{
	Code:           401,
	Name:           "UNAUTHENTICATED",
	Description:    "The request does not have valid authentication credentials for the operation.",
	HttpStatusCode: 401,
}

var PermissionDenied = &ErrorCode{
	Code:           403,
	Name:           "PERMISSION_DENIED",
	Description:    "The caller does not have permission to execute the specified operation.",
	HttpStatusCode: 403,
}

var NotFound = &ErrorCode{
	Code:           404,
	Name:           "NOT_FOUND",
	Description:    "Some requested entity (e.g., file or directory) was not found.",
	Document:       nil,
	HttpStatusCode: 404,
}

var AlreadyExists = &ErrorCode{
	Code:           6,
	Name:           "ALREADY_EXISTS",
	Description:    "The entity that a client attempted to create (e.g., file or directory) already exists.",
	HttpStatusCode: 409,
}

var Aborted = &ErrorCode{
	Code:           10,
	Name:           "ABORTED",
	Description:    "The operation was aborted, typically due to a concurrency issue such as a sequencer check failure or transaction abort.",
	HttpStatusCode: 409,
}

var ResourceExhausted = &ErrorCode{
	Code:           429,
	Name:           "RESOURCE_EXHAUSTED",
	Description:    "Some resource has been exhausted, perhaps a per-user quota, or perhaps the entire file system is out of space.",
	HttpStatusCode: 429,
}

var Cancelled = &ErrorCode{
	Code:           499,
	Name:           "CANCELLED",
	Description:    "The operation was cancelled, typically by the caller.",
	Document:       nil,
	HttpStatusCode: 499,
}

var UnknownError = &ErrorCode{
	Code:           2,
	Name:           "UNKNOWN_ERROR",
	Description:    "Unknown error.  For example, this error may be returned when a `Status` Code received from another address space belongs to an error space that is not known in this address space.",
	HttpStatusCode: 500,
}

var InternalError = &ErrorCode{
	Code:           500,
	Name:           "INTERNAL_ERROR",
	Description:    "Internal errors.  This means that some invariants expected by the underlying system have been broken.",
	HttpStatusCode: 500,
}

var DataLoss = &ErrorCode{
	Code:           15,
	Name:           "DATA_LOSS",
	Description:    "Unrecoverable data loss or corruption.",
	HttpStatusCode: 500,
}

var Unimplemented = &ErrorCode{
	Code:           501,
	Name:           "UNIMPLEMENTED",
	Description:    "The operation is not implemented or is not supported/enabled in this service.",
	HttpStatusCode: 501,
}

var Unavailable = &ErrorCode{
	Code:           503,
	Name:           "UNAVAILABLE",
	Description:    "The service is currently unavailable.",
	HttpStatusCode: 503,
}

var DeadlineExceeded = &ErrorCode{
	Code:           504,
	Name:           "DEADLINE_EXCEEDED",
	Description:    "The deadline expired before the operation could complete.",
	HttpStatusCode: 504,
}
