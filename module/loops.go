package module

import (
	"fmt"
	"time"
)

func Loops() {

	for {

		for {

			func() {
				fmt.Println("Старт")
				defer time.Sleep(5 * time.Second) //не срабатывает на break, continue
				// defer <-time.After(3 * time.Second) function must be invoked in defer statementsyntax

			}()
			fmt.Println("Финиш")

			goto EndLoop
			// break

			

		}
	}
EndLoop:
}
