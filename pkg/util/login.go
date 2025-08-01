package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
	"time"
)

// CreateToken 创建token
// @uid 用户id
func CreateToken(uid int64) string {

	claims := Claims{
		Uid: uid,
	}

	token, err := JwtEncode(claims, 24*3600, nil)
	if err != nil {
		return ""
	}
	return token
}

// VerifyToken 验证token
func VerifyToken(token string) (int64, error) {

	claims, err := JwtDecode(token, nil)
	if err != nil {
		return 0, err
	}
	return claims.Uid, nil
}

func CreatePwd(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func GetAuthMemberId(c *gin.Context) int64 {
	aid, ok := c.Get("aid")
	if !ok {
		return 0
	}

	if aid == nil {
		return 0
	}
	return aid.(int64)
}
