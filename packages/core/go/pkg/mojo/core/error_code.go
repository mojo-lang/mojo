package core

import "google.golang.org/protobuf/proto"

const ErrorCodeTypeName = "ErrorCode"
const ErrorCodeTypeFullName = "mojo.core.ErrorCode"

func NewErrorCode(value int32) *ErrorCode {
	if ec, ok := errorCodeIndex[value]; ok {
		return proto.Clone(ec).(*ErrorCode)
	}
	return &ErrorCode{Code: value}
}
