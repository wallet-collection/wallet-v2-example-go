package util

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {

	str := []byte("12fff我是ww.topgoer.com的站长枯藤")
	pwd, _ := EnPwdCode(str)
	fmt.Println(pwd)
	bytes, _ := DePwdCode(pwd)
	fmt.Println(string(bytes))

}
