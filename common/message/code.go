/**
 * @Author zhanglu
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package message

import (
	en_US "qiyaowu-go-zero/common/message/en-US"
	zh_CN "qiyaowu-go-zero/common/message/zh-CN"
)

const OK = 0
const RPC_SERVER_ERROR = 500
const PARAM_ERROR = 999
const WEB_ERROR_LOGIN_PASSWORD_ERR = 102
const WEB_ERROR_LOGIN_USER_NOTEXIST = 101
const WEB_ERROR_LOGIN_USER_NOTENABLED = 107
const WEB_ERROR_VERIFY_CODE = 108
const WEB_ERROR_LOGIN_PASSWORD_FIX = 119

func Message(code int, Language string) string {
	message := ""
	switch Language {
	case "zh-CN":
		message = zh_CN.Msg[code]
		break
	case "en-US":
		message = en_US.Msg[code]
		break
	}
	return message
}
