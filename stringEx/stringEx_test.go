package stringEx

import (
	"log"
	"testing"
)

var str = NewString("123qwe...")

func TestString_Contain(t *testing.T) {
	log.Println(str.Contain("123"), str.Length())
	log.Println(str.ReplaceAll("123", "789").String())
	log.Println(str.ReplaceAll("123", "789").Contain("789"))
}

func TestString_Concat(t *testing.T) {
	str.Concat("666", "777")
	log.Println(str)
}
