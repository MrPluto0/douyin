package define

import "douyin/utils/response"

/*  Two Interfaces for XXXReq and XXXRes */
type IReq interface {
	Validate()
}

type IRes interface {
	response.Errno
}
