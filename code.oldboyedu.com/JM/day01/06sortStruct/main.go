package main

import (
	"fmt"
	"sort"
)

type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string   //名称
	Kind HeroKind //种类
}

type Heros []*Hero

func (s Heros) Len() int {
	return len(s)
}

func (s Heros) Less(i, j int) bool {
	if s[i].Kind != s[j].Kind {
		//首先按照种类排序
		return s[i].Kind < s[j].Kind
	}

	//种类相同按照姓名排名
	return s[i].Name < s[j].Name
}

func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	hero1 := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Mage},
		&Hero{"关羽", Tank},
	}

	sort.Sort(hero1)
	for _, v := range hero1 {
		fmt.Println(v)
	}
}
