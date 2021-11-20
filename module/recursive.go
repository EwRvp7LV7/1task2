package module

import (
	"fmt"
	"strings"
)

func factorial(n uint) uint {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
func Recursive() {

	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(6)) // 720
}

func fibonaci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
func Recursive1() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonaci(i))
	}
}

func Recursive3() {

	str := "fgfwfhh"
	//replace repeated runes on blanks
	rs := []rune(str)
	for i := 0; i < len(rs); i++ {
		for j := i + 1; j < len(rs); j++ {
			if rs[j] == rs[i] {
				rs[j] = ' '
			}
		}
	}
	// remove blanks
	str = strings.ReplaceAll(string(rs), " ", "")
	// fmt.Println(str)
	recursiveStr("", str)
}

func recursiveStr(variants string, s string) {

	for _, r := range s {

		//fmt.Println(variants+string(r))
		//fmt.Println(strings.ReplaceAll(s, string(r), ""))
		recursiveStr(variants+string(r), strings.ReplaceAll(s, string(r), ""))
	}

	if len(s) == 0 {
		fmt.Println(variants)
	}
}
