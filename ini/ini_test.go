package ini

import (
	"log"
	"testing"
	"time"
)

var i = New("./test.ini")

func TestIni_Add(t *testing.T) {
	i.Add("Test", "key", "val")
	i.Add("Test", "key", "val")
	i.Add("Test", "key1", "val1")
	i.Save()
}

func TestIni_Get(t *testing.T) {
	if i.Get("Test", "key") != "val" {
		t.Error()
	}
}

func TestIni_Set(t *testing.T) {
	i.Set("Test", "key", "333")
	if i.Get("Test", "key") != "333" {
		t.Error()
	}
}

func TestIni_Marshal(t *testing.T) {
	type Note struct {
		Content string
		Cities  []string
	}

	type Person struct {
		Name string
		Age  int `ini:"age"`
		Male bool
		Born time.Time
		Note
		Created time.Time `ini:"-"`
	}
	p := new(Person)
	i.Marshal(&p)
	log.Println(p)
}

func TestIni_Unmarshal(t *testing.T) {
	type Embeded struct {
		Dates  []time.Time `delim:"|" comment:"Time data"`
		Places []string    `ini:"places,omitempty"`
		None   []int       `ini:",omitempty"`
	}

	type Author struct {
		Name      string `ini:"NAME"`
		Male      bool
		Age       int `comment:"Author's age"`
		GPA       float64
		NeverMind string `ini:"-"`
		*Embeded  `comment:"Embeded section"`
	}
	a := &Author{"Unknwon", true, 21, 2.8, "",
		&Embeded{
			[]time.Time{time.Now(), time.Now()},
			[]string{"HangZhou", "Boston"},
			[]int{},
		}}
	i.Unmarshal(a)
}
