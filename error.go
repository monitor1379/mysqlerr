package mysqlerr

/*
 * @Date: 2020-10-31 14:37:15
 * @LastEditors: monitor1379
 * @LastEditTime: 2020-10-31 14:45:20
 */

type Error interface {
	Code() int
	Message() string
	Error() string
}
