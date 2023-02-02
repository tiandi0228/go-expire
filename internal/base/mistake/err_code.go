package mistake

const (
	// ErrSuccess - 200: OK.
	ErrSuccess = 100001
	// ErrUnknown - 500: Internal server error.
	ErrUnknown = 100002
	// ErrBind - 400: Error occurred while binding the request body to the struct.
	ErrBind = 100003
	// ErrValidation - 400: Validation failed.
	ErrValidation = 100004
	// ErrTokenInvalid - 401: Token invalid.
	ErrTokenInvalid = 100005
	// ErrPageNotFound - 404: Page not found.
	ErrPageNotFound = 100006
)

// common: database errors.
const (
	// ErrDatabase - 500: Database error.
	ErrDatabase = 100101
)

// common: authorization and authentication errors.
const (
	// ErrEncrypt - 401: Error occurred while encrypting the user password.
	ErrEncrypt = 100201
	// ErrSignatureInvalid - 401: Signature is invalid.
	ErrSignatureInvalid = 100202
	// ErrExpired - 401: Token expired.
	ErrExpired = 100203
	// ErrInvalidAuthHeader - 401: Invalid authorization header.
	ErrInvalidAuthHeader = 100204
	// ErrMissingHeader - 401: The `Authorization` header was empty.
	ErrMissingHeader = 100205
	// ErrPasswordIncorrect - 401: Password was incorrect.
	ErrPasswordIncorrect = 100206
	// ErrPermissionDenied - 403: Permission denied.
	ErrPermissionDenied = 100207
)

// common: encode/decode errors.
const (
	// ErrEncodingFailed - 500: Encoding failed due to an error with the data.
	ErrEncodingFailed = 100301
	// ErrDecodingFailed - 500: Decoding failed due to an error with the data.
	ErrDecodingFailed = 100302
	// ErrInvalidJSON - 500: Data is not valid JSON.
	ErrInvalidJSON = 100303
	// ErrEncodingJSON - 500: JSON data could not be encoded.
	ErrEncodingJSON = 100304
	// ErrDecodingJSON - 500: JSON data could not be decoded.
	ErrDecodingJSON = 100305
	// ErrInvalidYaml - 500: Data is not valid Yaml.
	ErrInvalidYaml = 100306
	// ErrEncodingYaml - 500: Yaml data could not be encoded.
	ErrEncodingYaml = 100307
	// ErrDecodingYaml - 500: Yaml data could not be decoded.
	ErrDecodingYaml = 100308
)
