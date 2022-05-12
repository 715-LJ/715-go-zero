package common
//
//import (
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/zeromicro/go-zero/rest/httpx"
//	"net/http"
//	"strconv"
//	"strings"
//)
//var whiteList []string = []string{
//	"/login/login",
//}
//
//const  jwtkey  = "qsaseJwt"
////鉴权中间件
//func Auth(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Add("X-Middleware", "auth")
//		uri := r.RequestURI
//		fmt.Println("uri:",uri)
//		//默认不在
//		isInWhiteList := false
//		//判断请求是否包含白名单中的元素
//		for _, v := range whiteList {
//			if strings.Contains(uri, v) {
//				isInWhiteList = true
//			}
//		}
//		//如果在白名单里面直接通过
//		if isInWhiteList {
//			next(w, r)
//			return
//		}
//		//否则获取前端header 里面的X-Token字段,这个就是token
//		token := r.Header.Get("Authorization")
//		token = strings.Replace(token,"Bearer ","",1)
//
//		//工具类见util\jwttoken.go
//		_, err, _ := parseToken(token,w,r)
//		//如果有错直接返回error
//		if err != nil {
//			w.WriteHeader(http.StatusUnauthorized)
//			httpx.Error(w, err)
//			return
//		}
//		//没报错就继续
//		next(w, r)
//	}
//}
//
//// parseToken 解析token
//func parseToken(tk string,w http.ResponseWriter ,r *http.Request) (c *jwt.StandardClaims, err error, msg string) {
//	token, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
//		_, ok := token.Method.(*jwt.SigningMethodHMAC)
//		if !ok {
//			return nil, fmt.Errorf("不合法的token格式: %v", token.Header["alg"])
//		}
//		return []byte(jwtkey), nil
//	})
//
//
//	// 检测合法
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if !ok || !token.Valid {
//		return nil, fmt.Errorf("[parseToken] 不合法的token"), "fail token1"
//	}
//	host := r.Host
//	iss := claims["iss"].(string)
//	aud := claims["aud"].(string)
//	if iss != host{
//		fmt.Println("服务端ip错误")
//		return nil, fmt.Errorf("Authentication failed"),""
//	}
//	remoteAddr := strings.Split(r.RemoteAddr,":")
//	clientIp := remoteAddr[0] //客户端ip
//	if aud != clientIp{
//		fmt.Println("客户端ip错误")
//		return nil, fmt.Errorf("Authentication failed"),""
//	}
//	return mapClaimToJwClaim(claims), nil, "ok"
//}
//
//// 把jwt的claim转成claims
//func mapClaimToJwClaim(claims jwt.MapClaims) *jwt.StandardClaims {
//	uid := strconv.FormatInt(int64(claims["uid"].(float64)), 10)
//	fmt.Println("iss:",claims["iss"])
//	fmt.Println("aud:",claims["aud"])
//	jC := &jwt.StandardClaims{
//		Id: uid,
//		//Subject:
//	}
//
//	return jC
//}