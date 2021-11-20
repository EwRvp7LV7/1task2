package module

import (
	"fmt"
	//"strings"
)

// https://go.dev/blog/slices-intro
// В Go нет ссылок. Есть только указатели.
// Ссылка -- ссылается на ячейку памяти. Указатель -- указывает на объект в глобальной таблице объектов (нужна для сборщика мусора).
// Если объект в памяти перемещается -- изменяется ссылка хранимая в указателе на перемещаемый объект.
// Само значение указателя в глобальной таблице (сиречь -- номер строки) -- остаётся неизменным.

//эксперименты с передачей слайса в функцию
//функция делает свой экземпляр слайса, у которого с входным остается общая часть
//но! при изменении размеров внутреннего слайса размер внешнего не меняется
func CheckArrSize() {
	arr := make([]byte, 16)
	fmt.Println("hello01", len(arr), arr) //hello01 16 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	cutArrSize(arr)
	fmt.Println("hello02", len(arr), arr) //hello02 16 [1 1 1 1 1 1 1 1 0 0 0 0 0 0 0 0]
}
func cutArrSize(arr []byte) {
	for len(arr) > 0+8 {
		arr[0] = 1 // remembered in external array
		arr = arr[1:]
		fmt.Println("hello05", len(arr))
	}
	fmt.Println("hello10", len(arr), arr) //hello10 8 [0 0 0 0 0 0 0 0]

	for i := 0; i < 32; i++ {
		if len(arr) == i {
			arr = append(arr, 0) // dont pass in external array
		}
		// arr[i] = 1
		fmt.Println("hello15", len(arr))
	}
	fmt.Println("hello20", len(arr), arr) //hello20 32 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

type Circle1 struct {
	x, y, r float64
}

//значения передаются!
func CheckStruct() {
	// c := new(Circle1)
	//c := &Circle1{x: 1, y: 1}
	 c := &Circle1{1, 1, 1}//приоритетный способ инициализации, так как в предыдущем потерялся параметр r и компилятор не ругается!
	fmt.Println("hello10", c)
	c.cutStruct()
	fmt.Println("hello20", c)

}
func (c *Circle1) cutStruct() {
	c.x = 5
	// c.y = 6
	c.r = 7
}

// Reversed!
func SliceReverse() {
	tab := []int{1, 2, 3}
	fmt.Println(tab)
	// Results in [1 2 3]
	reverse(tab)
	fmt.Println(tab)
	// Results in [3 2 1]
}

func reverse(tab []int) {
	for i, j := 0, len(tab)-1; i < j; i, j = i+1, j-1 {
		tab[i], tab[j] = tab[j], tab[i]
	}
}

// reversed!
func ArrayReverse() {
	tab := [3]int{1, 2, 3}
	fmt.Println(tab)
	// Results in [1 2 3]
	reverseArray(&tab)
	fmt.Println(tab)
	// Results in [3 2 1]
}

func reverseArray(tab *[3]int) {
	for i, j := 0, len(*tab)-1; i < j; i, j = i+1, j-1 {
		tab[i], tab[j] = tab[j], tab[i]
	}
}

// // Not reversed!
// func ArrayReverse() {
// 	tab := [3]int{1, 2, 3}
// 	fmt.Println(tab)
// 	// Results in [1 2 3]
// 	reverseArray(tab)
// 	fmt.Println(tab)
// 	// Results in [3 2 1]
// }

// func reverseArray(tab [3]int) {
// 	for i, j := 0, len(tab)-1; i < j; i, j = i+1, j-1 {
// 		tab[i], tab[j] =  tab[j], tab[i]
// 	}
// }

// // Not reversed! ======================
// func StringReverse() {
// 	tab := `[]int{1, 2, 3}`
// 	fmt.Println(tab)
// 	// Results in [1 2 3]
// 	reverseStr(tab)
// 	fmt.Println(tab)
// 	// Results in [3 2 1]
// }

// func reverseStr(tab string) {
// 	ra := []rune(tab)
// 	for i, j := 0, len(ra)-1; i < j; i, j = i+1, j-1 {
// 		ra[i], ra[j] = ra[j], ra[i]
// 	}

// 	tab = string(ra) // не возвращается из функции, даже компиляятор ругается
// }
// // Not reversed! ======================

// reversed!
func RunesReverse() {
	tab := []rune(`[]int{1, 2, 3}`)
	fmt.Println(tab)
	// Results in [1 2 3]
	reverseRune(tab)
	fmt.Println(tab)
	// Results in [3 2 1]
}

func reverseRune(tab []rune) {
	for i, j := 0, len(tab)-1; i < j; i, j = i+1, j-1 {
		tab[i], tab[j] = tab[j], tab[i]
	}
}
