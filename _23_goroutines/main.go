package main

import (
	"fmt"
	"sync"
	// "time"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Task :", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)

		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}

	// time.Sleep(time.Second * 2)

	wg.Wait()

}
