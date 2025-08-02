package util

import (
	"fmt"
	"strings"
)

func RemoveInt64ByMap(slc []int64) []int64 {
	var result []int64          //存放返回的不重复切片
	tempMap := map[int64]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

func RemoveIntByMap(slc []int) []int {
	var result []int          //存放返回的不重复切片
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

func RemoveStringByMap(slc []string) []string {
	var result []string          //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

func HideEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		username := parts[0]
		domain := parts[1]

		// 如果用户名长度小于等于2，不进行隐藏处理
		if len(username) <= 2 {
			return email
		}

		// 隐藏用户名的中间部分字符，保留第一个和最后一个字符
		hiddenUsername := string(username[0])
		for i := 1; i < len(username)-1; i++ {
			hiddenUsername += "*"
		}
		hiddenUsername += string(username[len(username)-1])

		// 构建隐藏后的邮箱地址
		hiddenEmail := hiddenUsername + "@" + domain
		return hiddenEmail
	}

	// 如果邮箱地址不符合格式，不进行隐藏处理
	return email
}

func HidePhoneNumber(phoneNumber string) string {
	// 确保手机号码至少有8个字符（例如：1234567890）
	if len(phoneNumber) >= 8 {
		// 获取手机号码的前三位和后四位
		prefix := phoneNumber[:3]
		suffix := phoneNumber[len(phoneNumber)-4:]
		// 构建隐藏后的手机号码
		hiddenPhoneNumber := fmt.Sprintf("%s****%s", prefix, suffix)
		return hiddenPhoneNumber
	}
	// 如果手机号码不符合最小长度要求，不进行隐藏处理
	return phoneNumber
}
