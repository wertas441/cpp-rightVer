// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _join_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on JoinRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *JoinRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = JoinRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "raw_payload":

			if len(m.GetRawPayload()) != 23 {
				return JoinRequestValidationError{
					field:  "raw_payload",
					reason: "value length must be 23 bytes",
				}
			}

		case "payload":

			if v, ok := interface{}(m.GetPayload()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return JoinRequestValidationError{
						field:  "payload",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "dev_addr":
			// no validation rules for DevAddr
		case "selected_mac_version":
			// no validation rules for SelectedMacVersion
		case "net_id":
			// no validation rules for NetId
		case "downlink_settings":

			if v, ok := interface{}(&m.DownlinkSettings).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return JoinRequestValidationError{
						field:  "downlink_settings",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "rx_delay":

			if _, ok := RxDelay_name[int32(m.GetRxDelay())]; !ok {
				return JoinRequestValidationError{
					field:  "rx_delay",
					reason: "value must be one of the defined enum values",
				}
			}

		case "cf_list":

			if v, ok := interface{}(m.GetCfList()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return JoinRequestValidationError{
						field:  "cf_list",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "correlation_ids":

			for idx, item := range m.GetCorrelationIds() {
				_, _ = idx, item

				if utf8.RuneCountInString(item) > 100 {
					return JoinRequestValidationError{
						field:  fmt.Sprintf("correlation_ids[%v]", idx),
						reason: "value length must be at most 100 runes",
					}
				}

			}

		case "consumed_airtime":

			if v, ok := interface{}(m.GetConsumedAirtime()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return JoinRequestValidationError{
						field:  "consumed_airtime",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return JoinRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// JoinRequestValidationError is the validation error returned by
// JoinRequest.ValidateFields if the designated constraints aren't met.
type JoinRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JoinRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JoinRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JoinRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JoinRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JoinRequestValidationError) ErrorName() string { return "JoinRequestValidationError" }

// Error satisfies the builtin error interface
func (e JoinRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJoinRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JoinRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JoinRequestValidationError{}

// ValidateFields checks the field values on JoinResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *JoinResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = JoinResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "raw_payload":

			if l := len(m.GetRawPayload()); l < 17 || l > 33 {
				return JoinResponseValidationError{
					field:  "raw_payload",
					reason: "value length must be between 17 and 33 bytes, inclusive",
				}
			}

		case "session_keys":

			if m.GetSessionKeys() == nil {
				return JoinResponseValidationError{
					field:  "session_keys",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetSessionKeys()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return JoinResponseValidationError{
						field:  "session_keys",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "lifetime":

			if v, ok := interface{}(m.GetLifetime()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return JoinResponseValidationError{
						field:  "lifetime",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "correlation_ids":

			for idx, item := range m.GetCorrelationIds() {
				_, _ = idx, item

				if utf8.RuneCountInString(item) > 100 {
					return JoinResponseValidationError{
						field:  fmt.Sprintf("correlation_ids[%v]", idx),
						reason: "value length must be at most 100 runes",
					}
				}

			}

		default:
			return JoinResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// JoinResponseValidationError is the validation error returned by
// JoinResponse.ValidateFields if the designated constraints aren't met.
type JoinResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JoinResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JoinResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JoinResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JoinResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JoinResponseValidationError) ErrorName() string { return "JoinResponseValidationError" }

// Error satisfies the builtin error interface
func (e JoinResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJoinResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JoinResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JoinResponseValidationError{}
