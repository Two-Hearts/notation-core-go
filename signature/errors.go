// Copyright The Notary Project Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signature

import "fmt"

// SignatureIntegrityError is used when the signature associated is no longer
// valid.
type SignatureIntegrityError struct {
	Err error
}

// Error returns the formatted error message.
func (e *SignatureIntegrityError) Error() string {
	return fmt.Sprintf("signature is invalid. Error: %s", e.Err.Error())
}

// Unwrap unwraps the internal error.
func (e *SignatureIntegrityError) Unwrap() error {
	return e.Err
}

// InvalidSignatureError is used when Signature envelope is invalid.
type InvalidSignatureError struct {
	Msg string
}

// Error returns the error message or the default message if not provided.
func (e InvalidSignatureError) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "signature envelope format is invalid"
}

// UnsupportedSignatureFormatError is used when Signature envelope is not supported.
type UnsupportedSignatureFormatError struct {
	MediaType string
}

// Error returns the formatted error message.
func (e *UnsupportedSignatureFormatError) Error() string {
	return fmt.Sprintf("signature envelope format with media type %q is not supported", e.MediaType)
}

// SignatureNotFoundError is used when signature envelope is not present.
type SignatureNotFoundError struct{}

func (e SignatureNotFoundError) Error() string {
	return "signature envelope is not present"
}

// SignatureAuthenticityError is used when signature is not generated using
// trusted certificates.
type SignatureAuthenticityError struct{}

// Error returns the default error message.
func (e *SignatureAuthenticityError) Error() string {
	return "the signature's certificate chain does not contain any trusted certificate"
}

// UnsupportedSigningKeyError is used when a signing key is not supported.
type UnsupportedSigningKeyError struct {
	Msg string
}

// Error returns the error message or the default message if not provided.
func (e UnsupportedSigningKeyError) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "signing key is not supported"
}

// InvalidArgumentError is used when an argument to a function is invalid.
type InvalidArgumentError struct {
	Param string
	Err   error
}

// Error returns the error message.
func (e *InvalidArgumentError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%q param is invalid. Error: %s", e.Param, e.Err.Error())
	}
	return fmt.Sprintf("%q param is invalid", e.Param)
}

// Unwrap returns the unwrapped error.
func (e *InvalidArgumentError) Unwrap() error {
	return e.Err
}

// InvalidSignRequestError is used when SignRequest is invalid.
type InvalidSignRequestError struct {
	Msg string
}

// Error returns the error message or the default message if not provided.
func (e *InvalidSignRequestError) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "SignRequest is invalid"
}

// UnsupportedSignatureAlgoError is used when signing algo is not supported.
type UnsupportedSignatureAlgoError struct {
	Alg string
}

// Error returns the formatted error message.
func (e *UnsupportedSignatureAlgoError) Error() string {
	return fmt.Sprintf("signature algorithm %q is not supported", e.Alg)
}

// SignatureEnvelopeNotFoundError is used when signature envelope is not present.
type SignatureEnvelopeNotFoundError struct{}

// Error returns the default error message.
func (e *SignatureEnvelopeNotFoundError) Error() string {
	return "signature envelope is not present"
}

// DuplicateKeyError is used when repeated key name found.
type DuplicateKeyError struct {
	Key string
}

// Error returns the formatted error message.
func (e *DuplicateKeyError) Error() string {
	return fmt.Sprintf("repeated key: %q exists.", e.Key)
}

// TimestampError is any error related to RFC3161 Timestamp.
type TimestampError struct {
	Msg    string
	Detail error
}

// Error returns the formatted error message.
func (e *TimestampError) Error() string {
	if e.Msg != "" && e.Detail != nil {
		return fmt.Sprintf("timestamp: %s. Error: %s", e.Msg, e.Detail.Error())
	}
	if e.Msg != "" {
		return fmt.Sprintf("timestamp: %s", e.Msg)
	}
	if e.Detail != nil {
		return fmt.Sprintf("timestamp: %s", e.Detail.Error())
	}
	return "timestamp error"
}

// Unwrap returns the detail error of e.
func (e *TimestampError) Unwrap() error {
	return e.Detail
}
