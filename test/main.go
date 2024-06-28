package main

import (
	"fmt"
	"goAlgorithm/skiplist"
)

func main() {
	sl := skiplist.NewSkipList(5)
	sl.Insert(1)
	sl.Insert(22)
	sl.Insert(3)
	fmt.Println(sl.String())
	sl.Delete(3)
	fmt.Println(sl.String())
	sl.Insert(5)
	sl.Insert(76)
	fmt.Println(sl.String())
	sl.Search(5)
}
