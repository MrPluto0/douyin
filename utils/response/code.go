package response

import "fmt"

// It's also the common content in Response, see `/app/define/xxx.go`
type Errno struct {
	Code   int    `json:"status_code"`
	Msg    string `json:"status_msg"`
	Detail string `json:"status_detail,omitempty"`
}

/* Server Code(1) + Module Code(2) + Detail Code(3) */
var (
	// common error
	OK = &Errno{Code: 0, Msg: "OK"}

	ErrValidation = &Errno{Code: 10001, Msg: "Validation failed."}
	ErrDatabase   = &Errno{Code: 10002, Msg: "Database error."}
	ErrToken      = &Errno{Code: 10003, Msg: "Error occurred while signing the JSON web token."}

	// user error
	ErrUserNotFound = &Errno{Code: 20101, Msg: "user doesn't exist."}
)

func (e *Errno) Error() string {
	return fmt.Sprintf("code: %v, msg: %v, detail: %v", e.Code, e.Msg, e.Detail)
}

func (e *Errno) Extend(err error) Errno {
	return Errno{
		Code:   e.Code,
		Msg:    e.Msg,
		Detail: err.Error(),
	}
}
