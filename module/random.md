## Don't use Pseudorandom number generator rand.Intn()!

Each time you set the same seed, you get the same sequence. So of course if you're setting the seed to the time in a fast loop, you'll probably call it with the same seed many times. [stackoverflow](https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator)

``` golang
package main

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	//*if we seed with rand.Seed(time.Now().UnixNano())
	//and then another initializer calls rand.Seed(1)
	//our seed will get overridden, and that definitely isn’t what we want
	//fmt.Println(customRand.Intn(100))
	arg := os.Args[1]

	ai := []int{} //array former data 4 compare
	ao := []int{} //array for new data
	c := 0

    //It gets data array from file with custom prefix and compares with current data array.
	file, err := os.OpenFile(arg+"a.gob", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("will create new file")
	} else {
		decoder := gob.NewDecoder(file)
		err := decoder.Decode(&ai)
		if err != nil {
			fmt.Printf("decoder ", err)
		}
	}
	defer	file.Close()

	customSource := rand.NewSource(time.Now().UnixNano())
	customRand := rand.New(customSource)

	for i := 0; i < 1000; i++ {
		rnd := 0
		switch arg {
		case "0":
			rnd = rand.Intn(100) //don't use!
		case "1":
			rnd = customRand.Intn(100)
		case "2":
			customSource2 := rand.NewSource(time.Now().UnixNano())
			customRand2 := rand.New(customSource2)
			rnd = customRand2.Intn(100)
		default:
			rand.Seed(time.Now().UnixNano())
			rnd = rand.Intn(100)
		}

		ao = append(ao, rnd)

		if len(ai) > 0 {
			//fmt.Printf("%3d, %5d\n", ai[i], rnd)

			if rnd != ai[i] {
				c++
			}
		}

	}

	fmt.Println("not match:", c)

	file, _ = os.Create(arg + "a.gob")
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(ao)
	if err != nil {
		fmt.Printf("encoder ", err)
	}

}
```
