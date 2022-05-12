/**
 * @Author zhanglu
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package en_US

var Msg map[int]string

func init() {
	Msg = map[int]string{
		500: "rpc server error",
		999: "Parameter error, please check",
		102: "If the password is incorrect and the retry fails for 【%d】, the account will be locked for 【%d】 minutes. During the lock, the account cannot be logged in to",
		101: "user does not exist",
		107: "User has been frozen",
		108: "The account has been locked. There are 【%.0f】 minutes left to lock",
		119: "Password error",
	}
}
