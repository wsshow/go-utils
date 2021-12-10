package ini

import (
	"gopkg.in/ini.v1"
	"log"
)

type Ini struct {
	Path string
	Cfg  *ini.File
}

func New(path string) *Ini {
	return &Ini{
		Path: path,
		Cfg:  func() *ini.File { f, _ := ini.LooseLoad(path); return f }(),
	}
}

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

func (i *Ini) Set(section, key, val string) {
	i.Cfg.Section(section).Key(key).SetValue(val)
	i.Save()
}

func (i *Ini) Get(section, key string) string {
	return i.Cfg.Section(section).Key(key).String()
}

func (i *Ini) Save() {
	err := i.Cfg.SaveTo(i.Path)
	if err != nil {
		log.Println("[ini] SaveTo failed:", err)
	}
}

func (i *Ini) Marshal(v interface{}) bool {
	err := i.Cfg.MapTo(v)
	if err != nil {
		log.Println("[ini] Marshal failed:", err)
		return false
	}
	return true
}

func (i *Ini) Unmarshal(v interface{}) bool {
	err := ini.ReflectFrom(i.Cfg, v)
	if err != nil {
		log.Println("[ini] Unmarshal failed:", err)
		return false
	}
	i.Save()
	return true
}
