package mistake

import "sync"

var codeErrorInfoMapping = map[int]string{}
var codeErrorInfoMappingMux = &sync.Mutex{}

func init() {
	RegisterCodeErrorInfo(ErrSuccess, "OK")
	RegisterCodeErrorInfo(ErrUnknown, "Internal server error")
	RegisterCodeErrorInfo(ErrBind, "Parameter format error")
	RegisterCodeErrorInfo(ErrValidation, "Parameter validation error")
	RegisterCodeErrorInfo(ErrTokenInvalid, "Invalid token")
	RegisterCodeErrorInfo(ErrPageNotFound, "No Found")
	RegisterCodeErrorInfo(ErrDatabase, "Internal server data error")
	RegisterCodeErrorInfo(ErrEncrypt, "Error occurred while encrypting the user password")
	RegisterCodeErrorInfo(ErrSignatureInvalid, "Signature is invalid")
	RegisterCodeErrorInfo(ErrExpired, "Token expired")
	RegisterCodeErrorInfo(ErrInvalidAuthHeader, "Invalid authorization header")
	RegisterCodeErrorInfo(ErrMissingHeader, "The `Authorization` header was empty")
	RegisterCodeErrorInfo(ErrPasswordIncorrect, "Password was incorrect")
	RegisterCodeErrorInfo(ErrPermissionDenied, "Permission denied")
	RegisterCodeErrorInfo(ErrEncodingFailed, "Encoding failed due to an error with the data")
	RegisterCodeErrorInfo(ErrDecodingFailed, "Decoding failed due to an error with the data")
	RegisterCodeErrorInfo(ErrInvalidJSON, "Data is not valid JSON")
	RegisterCodeErrorInfo(ErrEncodingJSON, "JSON data could not be encoded")
	RegisterCodeErrorInfo(ErrDecodingJSON, "JSON data could not be decoded")
	RegisterCodeErrorInfo(ErrInvalidYaml, "Data is not valid Yaml")
	RegisterCodeErrorInfo(ErrEncodingYaml, "Yaml data could not be encoded")
	RegisterCodeErrorInfo(ErrDecodingYaml, "Yaml data could not be decoded")
}

// RegisterCodeErrorInfo 注册错误编码对应的错误信息
func RegisterCodeErrorInfo(code int, errorInfo string) {
	codeErrorInfoMappingMux.Lock()
	defer codeErrorInfoMappingMux.Unlock()
	if _, ok := codeErrorInfoMapping[code]; ok {
		panic("error code duplicated")
	}
	codeErrorInfoMapping[code] = errorInfo
}

// GetErrInfo 通过 code 换错误信息
func GetErrInfo(code int) string {
	return codeErrorInfoMapping[code]
}
