package module

import (
	"fmt"
)

// замыкания (closures)
// Замыкание - это функция, которая ссылается к переменным вне ее тела. 
func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func Closures() {
  pos, neg := adder(), adder()//здесь у каждой adder своя sum
  for i := 0; i < 10; i++ {
    fmt.Println(
      pos(i),
      neg(-2*i),
    )
  }
}
