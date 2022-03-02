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
var _qrcodegenerator_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// ValidateFields checks the field values on QRCodeFormat with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *QRCodeFormat) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = QRCodeFormatFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "name":

			if utf8.RuneCountInString(m.GetName()) > 100 {
				return QRCodeFormatValidationError{
					field:  "name",
					reason: "value length must be at most 100 runes",
				}
			}

		case "description":

			if utf8.RuneCountInString(m.GetDescription()) > 200 {
				return QRCodeFormatValidationError{
					field:  "description",
					reason: "value length must be at most 200 runes",
				}
			}

		case "field_mask":

			if v, ok := interface{}(m.GetFieldMask()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return QRCodeFormatValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return QRCodeFormatValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// QRCodeFormatValidationError is the validation error returned by
// QRCodeFormat.ValidateFields if the designated constraints aren't met.
type QRCodeFormatValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QRCodeFormatValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QRCodeFormatValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QRCodeFormatValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QRCodeFormatValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QRCodeFormatValidationError) ErrorName() string { return "QRCodeFormatValidationError" }

// Error satisfies the builtin error interface
func (e QRCodeFormatValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQRCodeFormat.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QRCodeFormatValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QRCodeFormatValidationError{}

// ValidateFields checks the field values on QRCodeFormats with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *QRCodeFormats) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = QRCodeFormatsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "formats":

			for key, val := range m.GetFormats() {
				_ = val

				if utf8.RuneCountInString(key) > 36 {
					return QRCodeFormatsValidationError{
						field:  fmt.Sprintf("formats[%v]", key),
						reason: "value length must be at most 36 runes",
					}
				}

				if !_QRCodeFormats_Formats_Pattern.MatchString(key) {
					return QRCodeFormatsValidationError{
						field:  fmt.Sprintf("formats[%v]", key),
						reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
					}
				}

				if v, ok := interface{}(val).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return QRCodeFormatsValidationError{
							field:  fmt.Sprintf("formats[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return QRCodeFormatsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// QRCodeFormatsValidationError is the validation error returned by
// QRCodeFormats.ValidateFields if the designated constraints aren't met.
type QRCodeFormatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QRCodeFormatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QRCodeFormatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QRCodeFormatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QRCodeFormatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QRCodeFormatsValidationError) ErrorName() string { return "QRCodeFormatsValidationError" }

// Error satisfies the builtin error interface
func (e QRCodeFormatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQRCodeFormats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QRCodeFormatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QRCodeFormatsValidationError{}

var _QRCodeFormats_Formats_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on GetQRCodeFormatRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetQRCodeFormatRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetQRCodeFormatRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "format_id":

			if utf8.RuneCountInString(m.GetFormatId()) > 36 {
				return GetQRCodeFormatRequestValidationError{
					field:  "format_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_GetQRCodeFormatRequest_FormatId_Pattern.MatchString(m.GetFormatId()) {
				return GetQRCodeFormatRequestValidationError{
					field:  "format_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
				}
			}

		default:
			return GetQRCodeFormatRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetQRCodeFormatRequestValidationError is the validation error returned by
// GetQRCodeFormatRequest.ValidateFields if the designated constraints aren't met.
type GetQRCodeFormatRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetQRCodeFormatRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetQRCodeFormatRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetQRCodeFormatRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetQRCodeFormatRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetQRCodeFormatRequestValidationError) ErrorName() string {
	return "GetQRCodeFormatRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetQRCodeFormatRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetQRCodeFormatRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetQRCodeFormatRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetQRCodeFormatRequestValidationError{}

var _GetQRCodeFormatRequest_FormatId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on GenerateEndDeviceQRCodeRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *GenerateEndDeviceQRCodeRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GenerateEndDeviceQRCodeRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "format_id":

			if utf8.RuneCountInString(m.GetFormatId()) > 36 {
				return GenerateEndDeviceQRCodeRequestValidationError{
					field:  "format_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_GenerateEndDeviceQRCodeRequest_FormatId_Pattern.MatchString(m.GetFormatId()) {
				return GenerateEndDeviceQRCodeRequestValidationError{
					field:  "format_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
				}
			}

		case "end_device":

			if m.GetEndDevice() == nil {
				return GenerateEndDeviceQRCodeRequestValidationError{
					field:  "end_device",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetEndDevice()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GenerateEndDeviceQRCodeRequestValidationError{
						field:  "end_device",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "image":

			if v, ok := interface{}(m.GetImage()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GenerateEndDeviceQRCodeRequestValidationError{
						field:  "image",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GenerateEndDeviceQRCodeRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GenerateEndDeviceQRCodeRequestValidationError is the validation error
// returned by GenerateEndDeviceQRCodeRequest.ValidateFields if the designated
// constraints aren't met.
type GenerateEndDeviceQRCodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenerateEndDeviceQRCodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenerateEndDeviceQRCodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenerateEndDeviceQRCodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenerateEndDeviceQRCodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenerateEndDeviceQRCodeRequestValidationError) ErrorName() string {
	return "GenerateEndDeviceQRCodeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GenerateEndDeviceQRCodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenerateEndDeviceQRCodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenerateEndDeviceQRCodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenerateEndDeviceQRCodeRequestValidationError{}

var _GenerateEndDeviceQRCodeRequest_FormatId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on GenerateQRCodeResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GenerateQRCodeResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GenerateQRCodeResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "text":
			// no validation rules for Text
		case "image":

			if v, ok := interface{}(m.GetImage()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GenerateQRCodeResponseValidationError{
						field:  "image",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GenerateQRCodeResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GenerateQRCodeResponseValidationError is the validation error returned by
// GenerateQRCodeResponse.ValidateFields if the designated constraints aren't met.
type GenerateQRCodeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenerateQRCodeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenerateQRCodeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenerateQRCodeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenerateQRCodeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenerateQRCodeResponseValidationError) ErrorName() string {
	return "GenerateQRCodeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GenerateQRCodeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenerateQRCodeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenerateQRCodeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenerateQRCodeResponseValidationError{}

// ValidateFields checks the field values on EntityOnboardingData with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EntityOnboardingData) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = EntityOnboardingDataFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "format_id":
			// no validation rules for FormatId
		case "data":
			if m.Data == nil {
				return EntityOnboardingDataValidationError{
					field:  "data",
					reason: "value is required",
				}
			}
			if len(subs) == 0 {
				subs = []string{
					"end_device_tempate",
				}
			}
			for name, subs := range _processPaths(subs) {
				_ = subs
				switch name {
				case "end_device_tempate":
					w, ok := m.Data.(*EntityOnboardingData_EndDeviceTempate)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetEndDeviceTempate()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return EntityOnboardingDataValidationError{
								field:  "end_device_tempate",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				}
			}
		default:
			return EntityOnboardingDataValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// EntityOnboardingDataValidationError is the validation error returned by
// EntityOnboardingData.ValidateFields if the designated constraints aren't met.
type EntityOnboardingDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityOnboardingDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityOnboardingDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityOnboardingDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityOnboardingDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityOnboardingDataValidationError) ErrorName() string {
	return "EntityOnboardingDataValidationError"
}

// Error satisfies the builtin error interface
func (e EntityOnboardingDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntityOnboardingData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityOnboardingDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityOnboardingDataValidationError{}

// ValidateFields checks the field values on ParseQRCodeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ParseQRCodeRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ParseQRCodeRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "format_id":

			if utf8.RuneCountInString(m.GetFormatId()) > 36 {
				return ParseQRCodeRequestValidationError{
					field:  "format_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_ParseQRCodeRequest_FormatId_Pattern.MatchString(m.GetFormatId()) {
				return ParseQRCodeRequestValidationError{
					field:  "format_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$\"",
				}
			}

		case "qr_code":

			if l := len(m.GetQrCode()); l < 10 || l > 1024 {
				return ParseQRCodeRequestValidationError{
					field:  "qr_code",
					reason: "value length must be between 10 and 1024 bytes, inclusive",
				}
			}

		default:
			return ParseQRCodeRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ParseQRCodeRequestValidationError is the validation error returned by
// ParseQRCodeRequest.ValidateFields if the designated constraints aren't met.
type ParseQRCodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParseQRCodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParseQRCodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParseQRCodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParseQRCodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParseQRCodeRequestValidationError) ErrorName() string {
	return "ParseQRCodeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ParseQRCodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParseQRCodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParseQRCodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParseQRCodeRequestValidationError{}

var _ParseQRCodeRequest_FormatId_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")

// ValidateFields checks the field values on ParseQRCodeResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ParseQRCodeResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ParseQRCodeResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "entity_onboarding_data":

			if v, ok := interface{}(m.GetEntityOnboardingData()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ParseQRCodeResponseValidationError{
						field:  "entity_onboarding_data",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return ParseQRCodeResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ParseQRCodeResponseValidationError is the validation error returned by
// ParseQRCodeResponse.ValidateFields if the designated constraints aren't met.
type ParseQRCodeResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParseQRCodeResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParseQRCodeResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParseQRCodeResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParseQRCodeResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParseQRCodeResponseValidationError) ErrorName() string {
	return "ParseQRCodeResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ParseQRCodeResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParseQRCodeResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParseQRCodeResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParseQRCodeResponseValidationError{}

// ValidateFields checks the field values on
// GenerateEndDeviceQRCodeRequest_Image with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *GenerateEndDeviceQRCodeRequest_Image) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GenerateEndDeviceQRCodeRequest_ImageFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "image_size":

			if val := m.GetImageSize(); val < 10 || val > 1000 {
				return GenerateEndDeviceQRCodeRequest_ImageValidationError{
					field:  "image_size",
					reason: "value must be inside range [10, 1000]",
				}
			}

		default:
			return GenerateEndDeviceQRCodeRequest_ImageValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GenerateEndDeviceQRCodeRequest_ImageValidationError is the validation error
// returned by GenerateEndDeviceQRCodeRequest_Image.ValidateFields if the
// designated constraints aren't met.
type GenerateEndDeviceQRCodeRequest_ImageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GenerateEndDeviceQRCodeRequest_ImageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GenerateEndDeviceQRCodeRequest_ImageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GenerateEndDeviceQRCodeRequest_ImageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GenerateEndDeviceQRCodeRequest_ImageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GenerateEndDeviceQRCodeRequest_ImageValidationError) ErrorName() string {
	return "GenerateEndDeviceQRCodeRequest_ImageValidationError"
}

// Error satisfies the builtin error interface
func (e GenerateEndDeviceQRCodeRequest_ImageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGenerateEndDeviceQRCodeRequest_Image.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GenerateEndDeviceQRCodeRequest_ImageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GenerateEndDeviceQRCodeRequest_ImageValidationError{}
