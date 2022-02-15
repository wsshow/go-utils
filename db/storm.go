package db

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/asdine/storm/v3"
	"log"
)

type User struct {
	Name string `storm:"id"`
	Pwd  string
}

func contains(db *storm.DB, name string) bool {
	var u User
	if err := db.One("Name", name, &u); err == nil {
		log.Println(name, "had existed")
		return true
	}
	return false
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func WorkFlow() {
	boltDb, err := storm.Open("my.db")
	if err != nil {
		log.Println("storm.Open error:", err)
	}
	defer boltDb.Close()
	// 写入数据
	user := User{
		Name: "zhangsan",
		Pwd:  MD5("qwer"),
	}
	if !contains(boltDb, user.Name) {
		if err = boltDb.Save(&user); err != nil {
			log.Println("boltDb.Save error:", err)
			return
		}
	}

	user = User{
		Name: "lisi",
		Pwd:  MD5("qwert"),
	}
	if !contains(boltDb, user.Name) {
		if err = boltDb.Save(&user); err != nil {
			log.Println("boltDb.Save error:", err)
			return
		}
	}
	// 查询数据
	var us []User
	err = boltDb.All(&us)
	log.Println(us)
}
