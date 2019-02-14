// Code generated by protoc-gen-validate
// source: envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto
// DO NOT EDIT!!!

package v2alpha1

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

// Validate checks the field values on RateLimit with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RateLimit) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetDomain()) < 1 {
		return RateLimitValidationError{
			Field:  "Domain",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if m.GetStage() > 10 {
		return RateLimitValidationError{
			Field:  "Stage",
			Reason: "value must be less than or equal to 10",
		}
	}

	if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitValidationError{
				Field:  "Timeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for FailureModeDeny

	return nil
}

// RateLimitValidationError is the validation error returned by
// RateLimit.Validate if the designated constraints aren't met.
type RateLimitValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RateLimitValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimit.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RateLimitValidationError{}
