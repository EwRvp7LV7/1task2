package module

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func pinger(c chan string) {
// 	for i := 0; ; i++ {
// 		fmt.Println("ping")
// 		c <- "pong"
// 	}
// }
// func printer(c chan string) {
// 	for {
// 		msg := <-c
// 		fmt.Println(msg)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func fibonacci2(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select { //
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func fibonacci1(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

// func sum(s []int, c chan int) {
// 	fmt.Println(s)
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum // send sum to c
// }

func Channels() {
	// var c chan int //is nil, blocked in and out  - deadlock!
	// dataChannel := make(chan string)
	//    fmt.Println(<-dataChannel)  //  - deadlock!

	// dataChannel := make(chan string, 3)
	// dataChannel <- "Some Sample Data"
	// dataChannel <- "Some Other Sample Data"
	// dataChannel <- "Buffered Channel"
	// // dataChannel <- "Buffered Channel2"//deadlock!
	// fmt.Println(<-dataChannel)
	// fmt.Println(<-dataChannel)
	// fmt.Println(<-dataChannel)
	// // fmt.Println(<-dataChannel)//deadlock!

	// var c chan string = make(chan string)
	// go pinger(c)
	// go printer(c)
	// var input string
	// fmt.Scanln(&input)

	// s := []int{7, 2, 8, -9, 4, 0}
	// c := make(chan int)
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)//Почему то запускается вперед предыдущей даже если их поменять местами
	// x, y := <-c, <-c // receive from c
	// fmt.Println(x, y, x+y)

	// A sender can close a channel to indicate that no more values will be sent

	// c := make(chan int, 10)
	// go fibonacci1(cap(c), c)
	// for i := range c {
	// 	fmt.Println(i)
	// }

	// c := make(chan int)
	// go fibonacci1(15, c)
	// for i := range c { //может крутиться в range пока не close
	// 	fmt.Println(i)
	// }

	// c := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		// выводит сообщения из fibonacci, который крутится в цикле
	// 		fmt.Println(<-c)
	// 	}
	// 	// завершает цикл fibonacci
	// 	quit <- 0
	// }()
	// fibonacci2(c, quit)

	// go func(s string) { fmt.Println(s) }("Hello inline")

	// tick := time.Tick(100 * time.Millisecond)
	// boom := time.After(500 * time.Millisecond)
	// for {
	// 	select {
	// 	case <-tick:
	// 		fmt.Println("tick.")
	// 	case <-boom:
	// 		fmt.Println("BOOM!")
	// 		return
	// 	default:
	// 		fmt.Println("    .")
	// 		time.Sleep(50 * time.Millisecond)
	// 	}
	// }

	// race condition(состоянии гонки). Горутины имеют независимый стек, следовательно нет необходимости в обмене данными между ними. Но, иногда, необходимо использовать общие данные между несколькими горутинами. В этом случае несколько горутин пытаются взаимодействовать с данными в общей области памяти, что иногда приводит к непредсказуемому результату.

	// https://habr.com/ru/post/490336/

	// <-time.After(2 * time.Second)

	/* 	//до исполнения блокируются как отправитель так и получатель
	   	fmt.Println("main() started")
	   	c := make(chan string)

	   	go func(c chan string) {
	   		time.Sleep(100 * time.Millisecond)
	   		fmt.Println("Hello " + <-c + "!")
	   	}(c)
	   	fmt.Println("main() run1!")

	   	c <- "John"

	   	// close(c) // closing channel
	   	// c <- "Mike"//panic: send on closed channel

	   	fmt.Println("main() stopped") */

	/* 	fmt.Println("closed() example")
	   	c := make(chan int)

	   	go func(c chan int) {
	   		for i := 0; i <= 9; i++ {
	   			c <- i * i
	   		}

	   		close(c) // close channel
	   	}(c)

	   	// // periodic block/unblock of main goroutine until chanel closes
	   	// for {
	   	// 	val, ok := <-c
	   	// 	if !ok {
	   	// 		fmt.Println(val, ok, "<-- loop broke!")
	   	// 		break // exit break loop
	   	// 	} else {
	   	// 		fmt.Println(val, ok)
	   	// 	}
	   	// }

	   	for val := range c {
	   		fmt.Println(val)
	   	}

	   	fmt.Println("main() stopped") */

	/*     c := make(chan int, 3)
	       c <- 1
	       c <- 2
	       c <- 3
	       // c <- 4 deadlock!
	       close(c)

	       // iteration terminates after receiving 3 values
	       for elem := range c {
	           fmt.Println(elem)
	       } */

	/* 	fmt.Println("buffered channel example")
	   	c := make(chan int, 3)

	   	go func(c chan int) {
	   		num := <-c
	   		fmt.Println(num * num)
	   	}(c)

	   	c <- 1
	   	c <- 2
	   	c <- 3
	   	//пролетает
	   	c <- 4 // ждет своей очереди (выполнения 1) и пролетает дальше

	   	fmt.Println("main() stopped") */

	/* 	fmt.Println("main() started")
	   	c := make(chan int, 3)

	   	go func(c chan int) {
	   		time.Sleep(100 * time.Millisecond)
	   		for i := 0; i <= 3; i++ {
	   			num := <-c
	   			fmt.Println("func1", num*num)
	   		}
	   	}(c)

	   	go func(c chan int) {
	   		time.Sleep(100 * time.Millisecond)
	   		for i := 0; i <= 3; i++ {
	   			num := <-c
	   			fmt.Println("func2", num*num)
	   		}
	   	}(c)

	   	c <- 1
	   	c <- 2
	   	c <- 3
	   	c <- 4
	   	c <- 5
	   	c <- 6
	   	c <- 7
	   	fmt.Println(<-c)
	   	fmt.Println("main() stopped") */

	/* 	fmt.Println("two gorutines")

	   	// squareChan := make(<-chan int)
	   	squareChan := make(chan int)
	   	cubeChan := make(chan int)

	   	// go func(c <-chan int) {
	   	go func(c chan int) {
	   		time.Sleep(3 * time.Second)
	   		num := <-c
	   		fmt.Println("num * num", num*num)
	   		c <- num * num
	   	}(squareChan)

	   	go func(c chan int) {
	   		time.Sleep(7 * time.Second)
	   		num := <-c
	   		fmt.Println("num * num * num", num*num*num)
	   		c <- num * num * num
	   	}(cubeChan)

	   	testNum := 3
	   	fmt.Println("[main] sent testNum to squareChan")

	   	squareChan <- testNum //start 1st func square
	   	//func writing back to channels

	   	cubeChan <- testNum//start 2nd func cube
	   	//func writing back to channels

	   	squareVal, cubeVal := <-squareChan, <-cubeChan//wait, send 9 and 27
	   	sum := squareVal + cubeVal

	   	fmt.Println("[main] sum of square and cube of", testNum, " is", sum) */

	fmt.Println("WaitGroup")
	var wg sync.WaitGroup // create waitgroup (empty struct)

	for i := 1; i <= 3; i++ {

		wg.Add(1) // increment counter

		go func(wg *sync.WaitGroup, instance int) {
			time.Sleep(2 * time.Second)
			fmt.Println("Service called on instance", instance)

			wg.Done() // decrement counter
		}(&wg, i)
	}

	wg.Wait() // blocks here
	fmt.Println("WaitGroup")

}

//================================
// worker than make squares
func sqrWorker(tasks <-chan int, results chan<- int, id int) {
	for num := range tasks {
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %v] Sending result by worker %v\n", id, id)
		results <- num * num
	}
}
func WorkersPool() {
	fmt.Println("WorkersPool() started")

	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// launching 3 worker goroutines
	for i := 0; i < 3; i++ {
		go sqrWorker(tasks, results, i)
	}

	// passing 5 tasks
	for i := 0; i < 5; i++ {
		tasks <- i * 2 // non-blocking as buffer capacity is 10
	}

	fmt.Println("[main] Wrote 5 tasks")

	// closing tasks
	close(tasks)

	// receving results from all workers
	for i := 0; i < 5; i++ {
		result := <-results // blocking because buffer is empty
		fmt.Println("[main] Result", i, ":", result)
	}

	fmt.Println("WorkersPool() stopped")
}

//================================
// Медленная функция
func sleepRandom(fromFunction string, ch chan int) {
	// Отложенная функция очистки
	defer func() { fmt.Println(fromFunction, "sleepRandom complete") }()

	// Выполним медленную задачу
	// В качестве примера,
	// «заснем» на рандомное время в мс
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleeptime := randomNumber + 100

	fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms")
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleeptime, "ms")

	// Напишем в канал, если он был передан
	if ch != nil {
		ch <- sleeptime
	}
}

// Функция, выполняющая медленную работу с использованием контекста
// Заметьте, что контекст - это первый аргумент
func sleepRandomContext(ctx context.Context, ch chan bool) {

	// Выполнение (прим. пер.: отложенное выполнение) действий по очистке
	// Созданных контекстов больше нет
	// Следовательно, отмена не требуется
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()

	// Создаем канал
	sleeptimeChan := make(chan int)

	// Запускаем выполнение медленной задачи в горутине
	// Передаем канал для коммуникаций
	go sleepRandom("sleepRandomContext", sleeptimeChan)

	// Используем select для выхода по истечении времени жизни контекста
	select {
	case <-ctx.Done():
		// Если контекст отменен, выбирается этот случай
		// Это случается, если заканчивается таймаут doWorkContext или
		// doWorkContext или main вызывает cancelFunction
		// Высвобождаем ресурсы, которые больше не нужны из-за прерывания работы
		// Посылаем сигнал всем горутинам, которые должны завершиться (используя каналы)
		// Обычно вы посылаете что-нибудь в канал,
		// ждете выхода из горутины, затем возвращаетесь
		// Или используете группы ожидания вместо каналов для синхронизации
		fmt.Println("sleepRandomContext: Time to return")

	case sleeptime := <-sleeptimeChan:
		// Этот вариант выбирается, когда работа завершается до отмены контекста
		fmt.Println("Slept for ", sleeptime, "ms")
	}
}

// Вспомогательная функция, которая в реальности может использоваться для разных целей
// Здесь она просто вызывает одну функцию
// В данном случае, она могла бы быть в main
func doWorkContext(ctx context.Context) {

	// От контекста с функцией отмены создаём производный контекст с тайм-аутом
	// Таймаут 150 мс
	// Все контексты, производные от этого, завершатся через 150 мс
	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)

	// Функция отмены для освобождения ресурсов после завершения функции
	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()

	// Создаем канал и вызываем функцию контекста
	// Можно также использовать группы ожидания для этого конкретного случая,
	// поскольку мы не используем возвращаемое значение, отправленное в канал
	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)

	// Используем select для выхода при истечении контекста
	select {
	case <-ctx.Done():
		// Этот случай выбирается, когда переданный в качестве аргумента контекст уведомляет о завершении работы
		// В данном примере это произойдёт, когда в main будет вызвана cancelFunction
		fmt.Println("doWorkContext: Time to return")

	case <-ch:
		// Этот вариант выбирается, когда работа завершается до отмены контекста
		fmt.Println("sleepRandomContext returned")
	}
}

func GoRoutine() {
	// Создаем контекст background
	ctx := context.Background()
	// Производим контекст с отменой
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	// Отложенная отмена высвобождает все ресурсы
	// для этого и производных от него контекстов
	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunction()
	}()

	// Отмена контекста после случайного тайм-аута
	// Если это происходит, все производные от него контексты должны завершиться
	go func() {
		sleepRandom("Main", nil)
		cancelFunction()
		fmt.Println("Main Sleep complete. canceling context")
	}()

	// Выполнение работы
	doWorkContext(ctxWithCancel)

}
