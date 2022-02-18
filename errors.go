package gerrors

import (
	"fmt"
)

type BaseError struct {
	Code int
	Msg  string
	Err  error
	*stack
}

func (e *BaseError) Error() string { return e.listMsg(0) }
func (e *BaseError) Unwrap() error { return e.Err }
func (e *BaseError) Cause() error  { return e.Err }

func (e *BaseError) listMsg(sept int) string {
	var msg = e.Msg
	frame := e.stackTrace()[0]
	if temp, ok := e.Err.(*BaseError); ok {
		msg = fmt.Sprintf("\n\t #%d %s %s:%d %s ", sept, msg, frame.file(), frame.line(), temp.listMsg(sept+1))
	} else {
		msg = fmt.Sprintf("\n\t #%d %s %s:%d \n\t #err %s ",
			sept, msg, frame.file(), frame.line(), fmt.Sprintf("%s", e.Err.Error()))
	}
	return msg
}

func New(code int, msg, format string, args ...interface{}) error {
	return &BaseError{
		Code:  code,
		Msg:   msg,
		Err:   fmt.Errorf(format, args...),
		stack: callers(),
	}
}

func Wrap(err error, msg string) *BaseError {
	if err == nil {
		return nil
	}
	if e, ok := err.(*BaseError); ok {
		return &BaseError{
			Err:   err,
			Code:  e.Code,
			Msg:   msg,
			stack: callers(),
		}
	}
	return &BaseError{
		Msg:   msg,
		Err:   err,
		stack: callers(),
	}
}

func WrapCode(err error, code int, msg string) error {
	return &BaseError{
		Code:  code,
		Msg:   msg,
		Err:   err,
		stack: callers(),
	}
}

func Err(e error) error {
	if temp, ok := e.(*BaseError); ok {
		return Err(temp.Err)
	}
	return e
}
