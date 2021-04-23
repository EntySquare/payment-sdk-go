package util

import "encoding/json"

type MessageError struct {
	errorCode int
	errorType string
	msg       string
}

func (e *MessageError) Error() string {
	return e.msg
}
func (e *MessageError) SetMsg(msg string) {
	e.msg = msg
}
func (e *MessageError) getMsg() string {
	return e.msg
}
func OutputJson(messageError MessageError) (errJson []byte, err error) {
	errJson, err = json.Marshal(messageError)
	if err != nil {
		msgError := NewMsgError(3, "wrong error struct")
		return nil, msgError
	}
	return errJson, err
}
func OutputString(messageError MessageError) (errString string, err error) {
	errJson, err := json.Marshal(messageError)
	if err != nil {
		msgError := NewMsgError(3, "wrong error struct")
		return "", msgError
	}
	errString = "ERROR MESSAGE: "
	for i := 0; i < len(errJson); i++ {
		errString += string(errJson[i])
	}
	return errString, err
}
func NewMsgError(code int, msg string) *MessageError {
	err := MessageError{
		errorCode: code,
		msg:       msg,
	}
	switch err.errorCode {
	case 0:
		err.errorType = "math Error"
	case 1:
		err.errorType = "rpc-link Error"
	case 2:
		err.errorType = "dataBase Error"
	case 3:
		err.errorType = "read-json Error"
	case 4:
		err.errorType = "type-trans Error"
	case 5:
		err.errorType = "server Error"
	case 9:
		err.errorType = "unknown Error"

	}
	return &err
}
