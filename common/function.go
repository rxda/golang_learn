package common

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

func pwd_md5(password string,salt string) string{
	s:="dclm"+password+salt
	h := md5.New()
	h.Write([]byte(s)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)

}

func trimall(str string) string {
	return strings.Replace(str, " ", "", -1)
}

func get_rand_str(length int) string{
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%"
	bytes := []byte(chars)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func check_mobile_number(mobile string){
	regex := "^[a-z0-9A-Z\u4e00-\u9fa5]+$";
	return str.matches(regex);

}