// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: testdata/ping-service/internal/conf/config.conf.proto

package conf

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on ServiceConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServiceConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServiceConfigMultiError, or
// nil if none found.
func (m *ServiceConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPingService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ServiceConfigValidationError{
					field:  "PingService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ServiceConfigValidationError{
					field:  "PingService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPingService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServiceConfigValidationError{
				field:  "PingService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ServiceConfigMultiError(errors)
	}

	return nil
}

// ServiceConfigMultiError is an error wrapping multiple validation errors
// returned by ServiceConfig.ValidateAll() if the designated constraints
// aren't met.
type ServiceConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceConfigMultiError) AllErrors() []error { return m }

// ServiceConfigValidationError is the validation error returned by
// ServiceConfig.Validate if the designated constraints aren't met.
type ServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceConfigValidationError) ErrorName() string { return "ServiceConfigValidationError" }

// Error satisfies the builtin error interface
func (e ServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceConfigValidationError{}

// Validate checks the field values on ServiceConfig_PingService with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ServiceConfig_PingService) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServiceConfig_PingService with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ServiceConfig_PingServiceMultiError, or nil if none found.
func (m *ServiceConfig_PingService) ValidateAll() error {
	return m.validate(true)
}

func (m *ServiceConfig_PingService) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetKey() != "" {

	}

	if len(errors) > 0 {
		return ServiceConfig_PingServiceMultiError(errors)
	}

	return nil
}

// ServiceConfig_PingServiceMultiError is an error wrapping multiple validation
// errors returned by ServiceConfig_PingService.ValidateAll() if the
// designated constraints aren't met.
type ServiceConfig_PingServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServiceConfig_PingServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServiceConfig_PingServiceMultiError) AllErrors() []error { return m }

// ServiceConfig_PingServiceValidationError is the validation error returned by
// ServiceConfig_PingService.Validate if the designated constraints aren't met.
type ServiceConfig_PingServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceConfig_PingServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceConfig_PingServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceConfig_PingServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceConfig_PingServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceConfig_PingServiceValidationError) ErrorName() string {
	return "ServiceConfig_PingServiceValidationError"
}

// Error satisfies the builtin error interface
func (e ServiceConfig_PingServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServiceConfig_PingService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceConfig_PingServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceConfig_PingServiceValidationError{}