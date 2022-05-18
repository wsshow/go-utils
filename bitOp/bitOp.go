package bitOp

import "log"

// 位运算，组合属性
const (
	Ldate = 1 << iota
	Ltime
	Lmicriseconds
)

func FormatHeader(flag int) {
	if (flag & Ldate) != 0 {
		log.Println("Ldate is enable")
	}
	if (flag & Ltime) != 0 {
		log.Println("Ltime is enable")
	}
	if (flag & Lmicriseconds) != 0 {
		log.Println("Lmicriseconds is enable")
	}
}
