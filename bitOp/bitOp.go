package bitOp

import "log"

// 位运算，组合属性
const (
	Ldate         = 1 << iota //0000 0001
	Ltime                     //0000 0010
	Lmicriseconds             //0000 0100
)

// Ldate | Ltime => 0000 0011 => 3
// (Ldate | Ltime) & Ldate => (0000 0011) & (0000 0001) => (0000 0001) => Ldate => 1

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
