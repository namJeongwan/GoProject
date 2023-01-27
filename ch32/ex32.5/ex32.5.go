package main

import (
	"fmt"
	"sort"
)

type CustomSlice []string

func (cs CustomSlice) Len() int {
	return len(cs)
}

func (cs CustomSlice) Less(i, j int) bool {
	return cs[i] < cs[j]
}

func (cs CustomSlice) Swap(i, j int) {

}

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)

	sort.Sort(CustomSlice{})
}
