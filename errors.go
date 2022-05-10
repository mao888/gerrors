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
	//if temp, ok := e.(*BaseError); ok {
	//	if temp.Code != 0 && temp.Msg != "" {
	//		return temp.Code, temp.Msg
	//	} else {
	//		return Resp(temp.Err)
	//	}
	//}
	//return 0, ""
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
