package core

import (
	"database/sql/driver"
	"errors"

	"google.golang.org/protobuf/reflect/protoreflect"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type BadRequestError struct {
	*basicError
}

func NewBadRequestError(format string, arguments ...interface{}) *BadRequestError {
	return &BadRequestError{newBasicError(BadRequest, format, arguments...)}
}

func IsBadRequestError(err error) bool {
	return errors.Is(err, &BadRequestError{})
}

func (*BadRequestError) Is(err error) bool {
	if _, ok := err.(*BadRequestError); ok {
		return true
	}
	return false
}

type InvalidArgumentError struct {
	*basicError
}

func NewInvalidArgumentError(format string, arguments ...interface{}) *InvalidArgumentError {
	return &InvalidArgumentError{newBasicError(InvalidArgument, format, arguments...)}
}

func IsInvalidArgumentError(err error) bool {
	return errors.Is(err, &InvalidArgumentError{})
}

func (*InvalidArgumentError) Is(err error) bool {
	if _, ok := err.(*InvalidArgumentError); ok {
		return true
	}
	return false
}

type MalformedSyntaxError struct {
	*basicError
}

func NewMalformedSyntaxError(format string, arguments ...interface{}) *MalformedSyntaxError {
	return &MalformedSyntaxError{newBasicError(MalformedSyntax, format, arguments...)}
}

func IsMalformedSyntaxError(err error) bool {
	return errors.Is(err, &MalformedSyntaxError{})
}

func (*MalformedSyntaxError) Is(err error) bool {
	if _, ok := err.(*MalformedSyntaxError); ok {
		return true
	}
	return false
}

type FailedPreconditionError struct {
	*basicError
}

func NewFailedPreconditionError(format string, arguments ...interface{}) *FailedPreconditionError {
	return &FailedPreconditionError{newBasicError(FailedPrecondition, format, arguments...)}
}

func IsFailedPreconditionError(err error) bool {
	return errors.Is(err, &FailedPreconditionError{})
}

func (*FailedPreconditionError) Is(err error) bool {
	if _, ok := err.(*FailedPreconditionError); ok {
		return true
	}
	return false
}

type OutOfRangeError struct {
	*basicError
}

func NewOutOfRangeError(format string, arguments ...interface{}) *OutOfRangeError {
	return &OutOfRangeError{newBasicError(OutOfRange, format, arguments...)}
}

func IsOutOfRangeError(err error) bool {
	return errors.Is(err, &OutOfRangeError{})
}

func (*OutOfRangeError) Is(err error) bool {
	if _, ok := err.(*OutOfRangeError); ok {
		return true
	}
	return false
}

type UnauthenticatedError struct {
	*basicError
}

func NewUnauthenticatedError(format string, arguments ...interface{}) *UnauthenticatedError {
	return &UnauthenticatedError{newBasicError(Unauthenticated, format, arguments...)}
}

func IsUnauthenticatedError(err error) bool {
	return errors.Is(err, &UnauthenticatedError{})
}

func (*UnauthenticatedError) Is(err error) bool {
	if _, ok := err.(*UnauthenticatedError); ok {
		return true
	}
	return false
}

type PermissionDeniedError struct {
	*basicError
}

func NewPermissionDeniedError(format string, arguments ...interface{}) *PermissionDeniedError {
	return &PermissionDeniedError{newBasicError(PermissionDenied, format, arguments...)}
}

func IsPermissionDeniedError(err error) bool {
	return errors.Is(err, &PermissionDeniedError{})
}

func (*PermissionDeniedError) Is(err error) bool {
	if _, ok := err.(*PermissionDeniedError); ok {
		return true
	}
	return false
}

type NotFoundError struct {
	*basicError
}

func NewNotFoundError(format string, arguments ...interface{}) *NotFoundError {
	return &NotFoundError{newBasicError(NotFound, format, arguments...)}
}

func IsNotFoundError(err error) bool {
	return errors.Is(err, &NotFoundError{})
}

func (*NotFoundError) Is(err error) bool {
	if _, ok := err.(*NotFoundError); ok {
		return true
	}
	return false
}

type AlreadyExistsError struct {
	*basicError
}

func NewAlreadyExistsError(format string, arguments ...interface{}) *AlreadyExistsError {
	return &AlreadyExistsError{newBasicError(AlreadyExists, format, arguments...)}
}

func IsAlreadyExistsError(err error) bool {
	return errors.Is(err, &AlreadyExistsError{})
}

func (*AlreadyExistsError) Is(err error) bool {
	if _, ok := err.(*AlreadyExistsError); ok {
		return true
	}
	return false
}

type AbortedError struct {
	*basicError
}

func NewAbortedError(format string, arguments ...interface{}) *AbortedError {
	return &AbortedError{newBasicError(Aborted, format, arguments...)}
}

func IsAbortedError(err error) bool {
	return errors.Is(err, &AbortedError{})
}

func (*AbortedError) Is(err error) bool {
	if _, ok := err.(*AbortedError); ok {
		return true
	}
	return false
}

type ResourceExhaustedError struct {
	*basicError
}

func NewResourceExhaustedError(format string, arguments ...interface{}) *ResourceExhaustedError {
	return &ResourceExhaustedError{newBasicError(ResourceExhausted, format, arguments...)}
}

func IsResourceExhaustedError(err error) bool {
	return errors.Is(err, &ResourceExhaustedError{})
}

func (*ResourceExhaustedError) Is(err error) bool {
	if _, ok := err.(*ResourceExhaustedError); ok {
		return true
	}
	return false
}

type CancelledError struct {
	*basicError
}

func NewCancelledError(format string, arguments ...interface{}) *CancelledError {
	return &CancelledError{newBasicError(Cancelled, format, arguments...)}
}

func IsCancelledError(err error) bool {
	return errors.Is(err, &CancelledError{})
}

func (*CancelledError) Is(err error) bool {
	if _, ok := err.(*CancelledError); ok {
		return true
	}
	return false
}

type UnknownErrorError struct {
	*basicError
}

func NewUnknownError(format string, arguments ...interface{}) *UnknownErrorError {
	return &UnknownErrorError{newBasicError(UnknownError, format, arguments...)}
}

func IsUnknownError(err error) bool {
	return errors.Is(err, &UnknownErrorError{})
}

func (*UnknownErrorError) Is(err error) bool {
	if _, ok := err.(*UnknownErrorError); ok {
		return true
	}
	return false
}

type InternalErrorError struct {
	*basicError
}

func NewInternalError(format string, arguments ...interface{}) *InternalErrorError {
	return &InternalErrorError{newBasicError(InternalError, format, arguments...)}
}

func IsInternalError(err error) bool {
	return errors.Is(err, &InternalErrorError{})
}

func (*InternalErrorError) Is(err error) bool {
	if _, ok := err.(*InternalErrorError); ok {
		return true
	}
	return false
}

type DataLossError struct {
	*basicError
}

func NewDataLossError(format string, arguments ...interface{}) *DataLossError {
	return &DataLossError{newBasicError(DataLoss, format, arguments...)}
}

func IsDataLossError(err error) bool {
	return errors.Is(err, &DataLossError{})
}

func (*DataLossError) Is(err error) bool {
	if _, ok := err.(*DataLossError); ok {
		return true
	}
	return false
}

type UnimplementedError struct {
	*basicError
}

func NewUnimplementedError(format string, arguments ...interface{}) *UnimplementedError {
	return &UnimplementedError{newBasicError(Unimplemented, format, arguments...)}
}

func IsUnimplementedError(err error) bool {
	return errors.Is(err, &UnimplementedError{})
}

func (*UnimplementedError) Is(err error) bool {
	if _, ok := err.(*UnimplementedError); ok {
		return true
	}
	return false
}

type UnavailableError struct {
	*basicError
}

func NewUnavailableError(format string, arguments ...interface{}) *UnavailableError {
	return &UnavailableError{newBasicError(Unavailable, format, arguments...)}
}

func IsUnavailableError(err error) bool {
	return errors.Is(err, &UnavailableError{})
}

func (*UnavailableError) Is(err error) bool {
	if _, ok := err.(*UnavailableError); ok {
		return true
	}
	return false
}

type DeadlineExceededError struct {
	*basicError
}

func NewDeadlineExceededError(format string, arguments ...interface{}) *DeadlineExceededError {
	return &DeadlineExceededError{newBasicError(DeadlineExceeded, format, arguments...)}
}

func IsDeadlineExceededError(err error) bool {
	return errors.Is(err, &DeadlineExceededError{})
}

func (*DeadlineExceededError) Is(err error) bool {
	if _, ok := err.(*DeadlineExceededError); ok {
		return true
	}
	return false
}

func init() {
	RegisterJSONFieldEncoder("core.basicError", "Code", &ErrorCodeStringCodec{IsFieldPointer: true})
	RegisterJSONFieldDecoder("core.basicError", "Code", &ErrorCodeStringCodec{IsFieldPointer: true})
}

type basicError Error

func newBasicError(code *ErrorCode, message string, arguments ...interface{}) *basicError {
	return (*basicError)(NewError(code, message, arguments...))
}

func (e *basicError) Error() string {
	return (*Error)(e).Error()
}

func (e *basicError) Unwrap() error {
	return (*Error)(e)
}

func (e *basicError) ToError() *Error {
	return (*Error)(e)
}

func (e *basicError) StatusCode() int {
	return (*Error)(e).StatusCode()
}

func (e *basicError) AddDetail(detail interface{}) *basicError {
	return (*basicError)((*Error)(e).AddDetail(detail))
}

func (e *basicError) Value() (driver.Value, error) {
	if e != nil {
		return (*Error)(e).Value()
	}
	return nil, nil
}

func (e *basicError) Scan(src interface{}) error {
	if e != nil && src != nil {
		return (*Error)(e).Scan(src)
	}
	return nil
}

func (e *basicError) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return e.ToError().GormDBDataType(db, field)
}

func (e *basicError) GormDataType() string {
	return e.ToError().GormDataType()
}

func (e *basicError) Reset() {
	(*Error)(e).Reset()
}

func (e *basicError) String() string {
	return (*Error)(e).String()
}

func (*basicError) ProtoMessage() {}

func (e *basicError) ProtoReflect() protoreflect.Message {
	return (*Error)(e).ProtoReflect()
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (e *basicError) Descriptor() ([]byte, []int) {
	return (*Error)(e).Descriptor()
}

func (e *basicError) GetCode() *ErrorCode {
	return (*Error)(e).GetCode()
}

func (e *basicError) GetMessage() string {
	return (*Error)(e).GetMessage()
}

func (e *basicError) GetDetails() []*Any {
	return (*Error)(e).GetDetails()
}

func (e *basicError) ContainsDetailType(i interface{}) bool {
	return (*Error)(e).ContainsDetailType(i)
}
