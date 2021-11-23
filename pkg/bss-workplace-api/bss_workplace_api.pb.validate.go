// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ozonmp/bss_workplace_api/v1/bss_workplace_api.proto

package bss_workplace_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Workplace with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Workplace) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Size

	if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkplaceValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkplaceValidationError is the validation error returned by
// Workplace.Validate if the designated constraints aren't met.
type WorkplaceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkplaceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkplaceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkplaceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkplaceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkplaceValidationError) ErrorName() string { return "WorkplaceValidationError" }

// Error satisfies the builtin error interface
func (e WorkplaceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkplace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkplaceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkplaceValidationError{}

// Validate checks the field values on CreateWorkplaceV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateWorkplaceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if m.GetSize() <= 0 {
		return CreateWorkplaceV1RequestValidationError{
			field:  "Size",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateWorkplaceV1RequestValidationError is the validation error returned by
// CreateWorkplaceV1Request.Validate if the designated constraints aren't met.
type CreateWorkplaceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWorkplaceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWorkplaceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWorkplaceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWorkplaceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWorkplaceV1RequestValidationError) ErrorName() string {
	return "CreateWorkplaceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWorkplaceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWorkplaceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWorkplaceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWorkplaceV1RequestValidationError{}

// Validate checks the field values on CreateWorkplaceV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateWorkplaceV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for WorkplaceId

	return nil
}

// CreateWorkplaceV1ResponseValidationError is the validation error returned by
// CreateWorkplaceV1Response.Validate if the designated constraints aren't met.
type CreateWorkplaceV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWorkplaceV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWorkplaceV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWorkplaceV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWorkplaceV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWorkplaceV1ResponseValidationError) ErrorName() string {
	return "CreateWorkplaceV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWorkplaceV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWorkplaceV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWorkplaceV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWorkplaceV1ResponseValidationError{}

// Validate checks the field values on DescribeWorkplaceV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeWorkplaceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetWorkplaceId() <= 0 {
		return DescribeWorkplaceV1RequestValidationError{
			field:  "WorkplaceId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeWorkplaceV1RequestValidationError is the validation error returned
// by DescribeWorkplaceV1Request.Validate if the designated constraints aren't met.
type DescribeWorkplaceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeWorkplaceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeWorkplaceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeWorkplaceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeWorkplaceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeWorkplaceV1RequestValidationError) ErrorName() string {
	return "DescribeWorkplaceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeWorkplaceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeWorkplaceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeWorkplaceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeWorkplaceV1RequestValidationError{}

// Validate checks the field values on DescribeWorkplaceV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeWorkplaceV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeWorkplaceV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeWorkplaceV1ResponseValidationError is the validation error returned
// by DescribeWorkplaceV1Response.Validate if the designated constraints
// aren't met.
type DescribeWorkplaceV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeWorkplaceV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeWorkplaceV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeWorkplaceV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeWorkplaceV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeWorkplaceV1ResponseValidationError) ErrorName() string {
	return "DescribeWorkplaceV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeWorkplaceV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeWorkplaceV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeWorkplaceV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeWorkplaceV1ResponseValidationError{}

// Validate checks the field values on ListWorkplacesV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListWorkplacesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Offset

	if m.GetLimit() <= 0 {
		return ListWorkplacesV1RequestValidationError{
			field:  "Limit",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// ListWorkplacesV1RequestValidationError is the validation error returned by
// ListWorkplacesV1Request.Validate if the designated constraints aren't met.
type ListWorkplacesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWorkplacesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWorkplacesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWorkplacesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWorkplacesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWorkplacesV1RequestValidationError) ErrorName() string {
	return "ListWorkplacesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListWorkplacesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWorkplacesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWorkplacesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWorkplacesV1RequestValidationError{}

// Validate checks the field values on ListWorkplacesV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListWorkplacesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListWorkplacesV1ResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListWorkplacesV1ResponseValidationError is the validation error returned by
// ListWorkplacesV1Response.Validate if the designated constraints aren't met.
type ListWorkplacesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWorkplacesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWorkplacesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWorkplacesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWorkplacesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWorkplacesV1ResponseValidationError) ErrorName() string {
	return "ListWorkplacesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListWorkplacesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWorkplacesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWorkplacesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWorkplacesV1ResponseValidationError{}

// Validate checks the field values on RemoveWorkplaceV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveWorkplaceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetWorkplaceId() <= 0 {
		return RemoveWorkplaceV1RequestValidationError{
			field:  "WorkplaceId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveWorkplaceV1RequestValidationError is the validation error returned by
// RemoveWorkplaceV1Request.Validate if the designated constraints aren't met.
type RemoveWorkplaceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveWorkplaceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveWorkplaceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveWorkplaceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveWorkplaceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveWorkplaceV1RequestValidationError) ErrorName() string {
	return "RemoveWorkplaceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveWorkplaceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveWorkplaceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveWorkplaceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveWorkplaceV1RequestValidationError{}

// Validate checks the field values on RemoveWorkplaceV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveWorkplaceV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveWorkplaceV1ResponseValidationError is the validation error returned by
// RemoveWorkplaceV1Response.Validate if the designated constraints aren't met.
type RemoveWorkplaceV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveWorkplaceV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveWorkplaceV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveWorkplaceV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveWorkplaceV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveWorkplaceV1ResponseValidationError) ErrorName() string {
	return "RemoveWorkplaceV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveWorkplaceV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveWorkplaceV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveWorkplaceV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveWorkplaceV1ResponseValidationError{}

// Validate checks the field values on UpdateWorkplaceV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateWorkplaceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateWorkplaceV1RequestValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateWorkplaceV1RequestValidationError is the validation error returned by
// UpdateWorkplaceV1Request.Validate if the designated constraints aren't met.
type UpdateWorkplaceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateWorkplaceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateWorkplaceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateWorkplaceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateWorkplaceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateWorkplaceV1RequestValidationError) ErrorName() string {
	return "UpdateWorkplaceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateWorkplaceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateWorkplaceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateWorkplaceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateWorkplaceV1RequestValidationError{}

// Validate checks the field values on UpdateWorkplaceV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateWorkplaceV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Updated

	return nil
}

// UpdateWorkplaceV1ResponseValidationError is the validation error returned by
// UpdateWorkplaceV1Response.Validate if the designated constraints aren't met.
type UpdateWorkplaceV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateWorkplaceV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateWorkplaceV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateWorkplaceV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateWorkplaceV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateWorkplaceV1ResponseValidationError) ErrorName() string {
	return "UpdateWorkplaceV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateWorkplaceV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateWorkplaceV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateWorkplaceV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateWorkplaceV1ResponseValidationError{}

// Validate checks the field values on WorkplaceEvent with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *WorkplaceEvent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for EventType

	// no validation rules for EventStatus

	if v, ok := interface{}(m.GetWorkplace()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkplaceEventValidationError{
				field:  "Workplace",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WorkplaceEventValidationError is the validation error returned by
// WorkplaceEvent.Validate if the designated constraints aren't met.
type WorkplaceEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkplaceEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkplaceEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkplaceEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkplaceEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkplaceEventValidationError) ErrorName() string { return "WorkplaceEventValidationError" }

// Error satisfies the builtin error interface
func (e WorkplaceEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkplaceEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkplaceEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkplaceEventValidationError{}
