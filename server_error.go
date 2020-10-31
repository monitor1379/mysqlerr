package mysqlerr

/*
 * @Date: 2020-10-31 14:45:01
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-10-31 15:12:57
 */

import "fmt"

var _ error = (*ServerError)(nil)
var _ Error = (*ServerError)(nil)

// ServerError includes an error code, SQLSTATE value, and message string,
// as described in Error Message Sources and Elements.
// These elements are available as described in Error Information Interfaces.
// Server Error Message Refence: https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
type ServerError struct {
	code     int
	sqlState string
	message  string
}

func NewServerError(code int, message string) error {
	return &ServerError{code: code, message: message}
}

func (s *ServerError) Error() string {
	return fmt.Sprintf("Error %d: %s", s.code, s.message)
}

func (s *ServerError) Code() int {
	return s.code
}

func (s *ServerError) Message() string {
	return s.message
}
