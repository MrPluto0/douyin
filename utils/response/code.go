package response

import "fmt"

// It's also the common content in Response, see `/app/define/xxx.go`
type Errno struct {
	Code int    `json:"status_code"`
	Msg  string `json:"status_msg"`
}

/* Server Code(1) + Module Code(2) + Detail Code(3) */
var (
	// common error
	OK = &Errno{0, "OK"}

	ErrValidation = &Errno{10001, "Validation failed."}
	ErrDatabase   = &Errno{10002, "Database error."}
	ErrToken      = &Errno{10003, "Error occurred while signing the JSON web token."}

	// user error
	ErrUserNotFound = &Errno{20101, "user doesn't exist."}
)

func (e *Errno) Error() string {
	return fmt.Sprintf("code: %v, msg: %v", e.Code, e.Msg)
}
