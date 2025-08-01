package util

import "strings"

const (
	BASE            = "678U5CH2LW43AEPXDFZKYQRMNIJB9STV"
	SpotOrderBASE   = "998U5CH2LW43AEPXTFZKYSRMNIJB9STV"
	FutureOrderBASE = "999U5CH2LW43AEPXDFZKYBRMNIJB9SUV"
	DECIMAL         = 32
	PAD             = "G"
	LEN             = 6
)

// IdToCode id转code
func IdToCode(uid int64) string {
	id := uid
	mod := int64(0)
	res := ""
	for id != 0 {
		mod = id % DECIMAL
		id = id / DECIMAL
		res += string(BASE[mod])
	}
	resLen := len(res)
	if resLen < LEN {
		res += PAD
		for i := 0; i < LEN-resLen-1; i++ {
			res += string(BASE[(int(uid)+i)%DECIMAL])
		}
	}
	return res
}

func CodeToId(code string) int64 {
	res := int64(0)
	lenCode := len(code)
	baseArr := []byte(BASE)       // 字符串进制转换为byte数组
	baseRev := make(map[byte]int) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}
	// 查找补位字符的位置
	isPad := strings.Index(code, PAD)
	if isPad != -1 {
		lenCode = isPad
	}
	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == PAD {
			continue
		}
		index, ok := baseRev[code[i]]
		if !ok {
			return 0
		}
		b := int64(1)
		for j := 0; j < r; j++ {
			b *= DECIMAL
		}
		res += int64(index) * b
		r++
	}
	return res
}

// IdToSpotOrderId id转订单号
func IdToSpotOrderId(uid int64) string {
	id := uid
	mod := int64(0)
	res := ""
	for id != 0 {
		mod = id % DECIMAL
		id = id / DECIMAL
		res += string(SpotOrderBASE[mod])
	}
	resLen := len(res)
	if resLen < LEN {
		res += PAD
		for i := 0; i < LEN-resLen-1; i++ {
			res += string(SpotOrderBASE[(int(uid)+i)%DECIMAL])
		}
	}
	return res
}

func SpotOrderIdToId(code string) int64 {
	res := int64(0)
	lenCode := len(code)
	baseArr := []byte(SpotOrderBASE) // 字符串进制转换为byte数组
	baseRev := make(map[byte]int)    // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}
	// 查找补位字符的位置
	isPad := strings.Index(code, PAD)
	if isPad != -1 {
		lenCode = isPad
	}
	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == PAD {
			continue
		}
		index, ok := baseRev[code[i]]
		if !ok {
			return 0
		}
		b := int64(1)
		for j := 0; j < r; j++ {
			b *= DECIMAL
		}
		res += int64(index) * b
		r++
	}
	return res
}

// IdToFutureOrderId id转订单号
func IdToFutureOrderId(uid int64) string {
	id := uid
	mod := int64(0)
	res := ""
	for id != 0 {
		mod = id % DECIMAL
		id = id / DECIMAL
		res += string(FutureOrderBASE[mod])
	}
	resLen := len(res)
	if resLen < LEN {
		res += PAD
		for i := 0; i < LEN-resLen-1; i++ {
			res += string(FutureOrderBASE[(int(uid)+i)%DECIMAL])
		}
	}
	return res
}

func FutureOrderIdToId(code string) int64 {
	res := int64(0)
	lenCode := len(code)
	baseArr := []byte(FutureOrderBASE) // 字符串进制转换为byte数组
	baseRev := make(map[byte]int)      // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}
	// 查找补位字符的位置
	isPad := strings.Index(code, PAD)
	if isPad != -1 {
		lenCode = isPad
	}
	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == PAD {
			continue
		}
		index, ok := baseRev[code[i]]
		if !ok {
			return 0
		}
		b := int64(1)
		for j := 0; j < r; j++ {
			b *= DECIMAL
		}
		res += int64(index) * b
		r++
	}
	return res
}
