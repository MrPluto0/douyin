package define

import (
	"douyin/utils/response"

	"github.com/dlclark/regexp2"
)

// Login API
type LoginReq struct {
	Username string `form:"username" binding:"required,lte=32" url:"username"`
	Password string `form:"password" binding:"required,lte=32" url:"password"`
}

func (r *LoginReq) Validate() (bool, error) {
	reg := regexp2.MustCompile(`(?=.*\d)(?=.*[a-z])(?=.*[A-Z])`, 0)
	return reg.MatchString(r.Password)
}

type LoginRes struct {
	response.Errno        // common struct + composition
	UserId         uint   `json:"user_id,omitempty"`
	Token          string `json:"token,omitempty"`
}

// Register API, same as Login API
type RegisterReq LoginReq

func (r *RegisterReq) Validate() (bool, error) {
	reg := regexp2.MustCompile(`(?=.*\d)(?=.*[a-z])(?=.*[A-Z])`, 0)
	return reg.MatchString(r.Password)
}

type RegisterRes LoginRes
