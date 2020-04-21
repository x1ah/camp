package http

import "fmt"

type Response struct {
	Proto string
	StatusCode int
	StatusMsg string
	Header Header
	Body []byte
}

func WriteResponse(resp *Response) []byte {
	res := []byte(fmt.Sprintf("%s %d %s\n", resp.Proto, resp.StatusCode, resp.StatusMsg))
	for k, v := range resp.Header {
		// urlencode
		res = append(res, []byte(fmt.Sprintf("%s %s\n", k, v))...)
	}
	res = append(res, []byte("\r\n")...)
	res = append(res, resp.Body...)
	return res
}
