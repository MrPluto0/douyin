package define

import "douyin/utils/response"

/*  Two Interfaces for XXXReq and XXXRes */
type IReq interface {
	Validate() // not required
}

type IRes interface {
	response.Errno // required
}
