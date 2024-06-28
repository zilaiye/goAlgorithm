package main

import (
	"goAlgorithm/skiplist"
	"math/rand"
)

func main() {
	sl := skiplist.NewSkipList(5)
	//sl.Insert(1)
	//sl.Insert(22)
	//sl.Insert(3)
	//fmt.Println(sl.String())
	//sl.Delete(3)
	//fmt.Println(sl.String())
	//sl.Insert(5)
	//sl.Insert(76)
	//fmt.Println(sl.String())
	//sl.Search(5)
	//sl.PrintList()
	for i := 0; i < 100; i++ {
		v := rand.Intn(100)
		if sl.Search(v) == nil {
			sl.Insert(v)
		}
	}
	sl.PrintList()
}
