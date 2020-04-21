package http

import (
	"bufio"
	"errors"
	"strings"
)

type Header map[string]string

type Request struct {
	Method string
	URL string

	Proto      string // "HTTP/1.0"

	Header Header
	Body []byte
}

var InvalidRequest = errors.New("invalid request")


func ParseRequest(content string) (*Request, error) {
	req := new(Request)

	// parse request line
	scanner := bufio.NewScanner(strings.NewReader(content))
	if scanner.Scan() {
		info := strings.Split(scanner.Text(), " ")
		if len(info) != 3 {
			return req, InvalidRequest
		}

		req.Method = info[0]
		req.URL = info[1]
		req.Proto = info[2]
	}

	// parse header
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		idx := strings.Index(line, " ")
		if idx == -1 {
			return req, InvalidRequest
		}
		k, v := line[:idx], line[idx+1:]
		if req.Header == nil {
			req.Header = make(Header)
		}
		req.Header[k] = v
	}

	// parse body
	for scanner.Scan() {
		req.Body = append(req.Body, scanner.Bytes()...)
	}

	return req, nil

}
