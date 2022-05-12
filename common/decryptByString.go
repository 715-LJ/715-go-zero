package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"qsase/models"
	"regexp"
)

const (
	secretKey = "qeghj65nvpav0dh7"
)

func DecryptByString(str string) string {
	data := []byte(secretKey)
	has := md5.Sum(data)
	md5Str := fmt.Sprintf("%x", has)
	ivStr := md5Str[0:16]
	originData, _ := base64.StdEncoding.DecodeString(str)
	iv := []byte(ivStr)
	cipherBlock, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		logx.Error(err)
	}
	cipher.NewCBCDecrypter(cipherBlock, iv).CryptBlocks(originData, originData)
	return string(PKCS5UnPadding(originData))
	
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	if length-unpadding < 0 {
		return []byte("")
	}
	return src[:(length - unpadding)]
}

/*
 * hash 加密密码
 * @param string pwd 待加密的明文密码
 */
func HashEncode(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

/*
 * 验证 hash 密码
 * @param string hashedPwd 已加密的hash密码
 * @param string sourcePwd 确认密码
 */
func ComparePasswords(hashedPwd string, sourcePwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(sourcePwd))
	if err != nil {
		return false
	}
	return true
}
/*
 * 密码复杂度验证
 * @param string cypher 密码
 */
func CheckPasswordComplex(cypher string) bool{
	passwordPolicyModel := models.NewPasswordPolicyModel()
	limit := passwordPolicyModel.GetItem()
	if limit != nil{
		cypherdLength := len(cypher)
		if cypherdLength < limit.MinChar{
			return false
		}
		if limit.IncludeSpecialChar == 1{
			pattern := "^[A-Za-z0-9]+$"
			reg := regexp.MustCompile(pattern)
			if reg.MatchString(cypher){
				return false
			}
		}
	}
	return true
}
