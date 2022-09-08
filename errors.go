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
func (e *BaseError) clone() *BaseError {
	return &BaseError{
		Code:  e.Code,
		Msg:   e.Msg,
		Err:   e.Err,
		stack: e.stack,
	}
}
func (e *BaseError) listMsg(sept int) string {
	var msg = e.Msg
	if e.stack == nil {
		return msg
	}
	frame := e.stackTrace()[0]
	if temp, ok := e.Err.(*BaseError); ok {
		msg = fmt.Sprintf("\n #%d %s %s %s ", sept, msg, frame, temp.listMsg(sept+1))
	} else {
		errMsg := "nil"
		if e.Err != nil {
			errMsg = e.Err.Error()
		}
		msg = fmt.Sprintf("\n #%d %s %s \n #e %s ",
			sept, msg, frame, errMsg)
	}
	return msg
}

func New(code int, msg string) error {
	return &BaseError{
		Code:  code,
		Msg:   msg,
		Err:   nil,
		stack: callers(),
	}
}

func NewCodeMsg(code int, msg string) error {
	return &BaseError{
		Code:  code,
		Msg:   msg,
		Err:   nil,
		stack: nil,
	}
}

func AddStack(e error) error {
	err, ok := e.(*BaseError)
	if !ok || err.stack != nil {
		return err
	}
	stackErr := err.clone()
	stackErr.stack = callers()
	return stackErr
}

func Wrap(err error, msg string) error {
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

func Resp(e error) (int, string) {
	temps := make([]*BaseError, 0)
	for {
		temp, ok := e.(*BaseError)
		if !ok {
			break
		}
		if temp.Code == 0 || temp.Msg == "" {
			e = temp.Err
			continue
		}
		temps = append(temps, temp)
		if temp.Err == nil {
			break
		}
		e = temp.Err
	}
	if len(temps) > 0 {
		return temps[len(temps)-1].Code, temps[len(temps)-1].Msg
	}
	return 0, ""
}

func Err(e error) error {
	if e == nil {
		return nil
	}
	if temp, ok := e.(*BaseError); ok {
		return Err(temp.Err)
	}
	return e
}
