package errno

import (
	"errors"
	"fmt"
)

type Errno struct {
	Code    int64
	Message string
}

func (e Errno) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func NewErrno(code int64, msg string) Errno {
	return Errno{
		Code:    code,
		Message: msg,
	}
}

func (e Errno) WithMessage(msg string) Errno {
	e.Message = msg
	return e
}

func (e Errno) WithFormat(temp, msg string) Errno {
	e.Message = fmt.Sprintf(temp, msg)
	return e
}

func ConvertErr(err error) Errno {
	if err == nil {
		return Success
	}
	errno := Errno{}
	if errors.As(err, &errno) {
		return errno
	}
	s := SystemErr
	s.Message = err.Error()
	return s
}
