package httpd

import (
	"encoding/json"
	"fmt"
)

// 返回类型
type Response struct {
	Msg interface{} `msg`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Set(msg interface{}) {
	r.Msg = msg
}

func (r *Response) Marshal() json.RawMessage {
	rsp, err := json.Marshal(r)
	if err != nil {
		fmt.Errorf("response is error %v", err)
		return nil
	}
	return rsp
}
