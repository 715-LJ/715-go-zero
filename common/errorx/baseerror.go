/**
 * @Author zhanglu
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package errorx

import message2 "qiyaowu-go-zero/common/message"

var Language string

const defaultCode = 0
const defaultMessage = "ok"

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type CodeErrorResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewCodeError(code int, message string) error {
	return &CodeError{Code: code, Message: message}
}

func NewDefaultError(code int) error {
	message := message2.Message(code, Language)
	return NewCodeError(code, message)
}

func (e *CodeError) Error() string {
	return e.Message
}
func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
		Data:    map[string]interface{}{},
	}
}
