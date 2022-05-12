/**
 * @Author zhanglu
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package zh_CN

var Msg map[int]string

func init() {
	Msg = map[int]string{
		500: "rpc 服务错误",
		999: "参数错误，请检查",
		102: "密码错误，【%d】次重试失败后，账号将被锁定【%d】分钟，锁定期间账号将无法登录",
		101: "用户不存在",
		107: "用户已被冻结",
		108: "账号已被锁定，锁定剩余【%.0f】分钟",
		119: "密码错误",
	}
}
