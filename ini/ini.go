package ini

import (
	"gopkg.in/ini.v1"
	"log"
)

type Ini struct {
	Path string
	Cfg  *ini.File
}

// New 创建对象
func New(path string) *Ini {
	return &Ini{
		Path: path,
		Cfg:  func() *ini.File { f, _ := ini.LooseLoad(path); return f }(),
	}
}

// Add 在内存中的指定Section下添加key和value(注意：此操作仅在内存中进行，调用Save方法后保存进文件)
func (i *Ini) Add(section, key, val string) {
	if !i.Cfg.HasSection(section) {
		_, err := i.Cfg.NewSection(section)
		if err != nil {
			log.Println("[ini] NewSection err:", err)
			return
		}
	}
	s, err := i.Cfg.GetSection(section)
	if err != nil {
		log.Println("[ini] GetSection err:", err)
		return
	}
	_, err = s.NewKey(key, val)
	if err != nil {
		log.Println("[ini] NewKey err:", err)
		return
	}
}

// Set 指定Section下，若key存在，则修改；若key不存在，则创建
func (i *Ini) Set(section, key, val string) {
	i.Cfg.Section(section).Key(key).SetValue(val)
	i.Save()
}

// Get 获取指定Section下key所对应的value值
func (i *Ini) Get(section, key string) string {
	return i.Cfg.Section(section).Key(key).String()
}

// Save 将内存中的ini结构存入文件
func (i *Ini) Save() {
	err := i.Cfg.SaveTo(i.Path)
	if err != nil {
		log.Println("[ini] SaveTo failed:", err)
	}
}

// Marshal 读取文件映射为既定结构
func (i *Ini) Marshal(v interface{}) bool {
	err := i.Cfg.MapTo(v)
	if err != nil {
		log.Println("[ini] Marshal failed:", err)
		return false
	}
	return true
}

// Unmarshal 数据结构保存进文件
func (i *Ini) Unmarshal(v interface{}) bool {
	err := ini.ReflectFrom(i.Cfg, v)
	if err != nil {
		log.Println("[ini] Unmarshal failed:", err)
		return false
	}
	i.Save()
	return true
}
