/*
 * @Date: 2020-10-31 14:37:15
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-10-31 14:43:22
 */

package mysqlerr

import "fmt"

type Error interface {
	Code() int
	Reason() string
	Error() string
}

var _ error = (*ServerError)(nil)
var _ Error = (*ServerError)(nil)

func NewServerError(code int, reason string) error {
	return &ServerError{code: code, reason: reason}
}

type ServerError struct {
	code   int
	reason string
}

func (s *ServerError) Error() string {
	return fmt.Sprintf("Error %d: %s", s.code, s.reason)
}

func (s *ServerError) Code() int {
	return s.code
}

func (s *ServerError) Reason() string {
	return s.reason
}
