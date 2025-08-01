package util

import (
	"fmt"
	"strconv"
	"strings"
)

func SplitToInt64List(str string, sep string) []int64 {
	var i64List []int64
	if str == "" {
		return i64List
	}
	strList := strings.Split(str, sep)
	if len(strList) == 0 {
		return i64List
	}
	for _, item := range strList {
		if item == "" {
			continue
		}
		val, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			// logs.CtxError(ctx, "ParseInt fail, err=%v, str=%v, sep=%v", err, str, sep) // 此处打印出log报错信息
			continue
		}
		i64List = append(i64List, val)
	}
	return i64List
}

func Int64ListToStr(arr []int64, sep string) string {
	var temp = make([]string, len(arr))
	for k, v := range arr {
		temp[k] = fmt.Sprintf("%d", v)
	}
	return strings.Join(temp, sep)
}
