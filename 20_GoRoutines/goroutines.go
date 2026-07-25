package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // defer keyword makes this line execute at end of this function
	fmt.Println(id)
}

func main() {
	var wg sync.WaitGroup

	for i := range 11 {
		wg.Add(1)
		go task(i, &wg) // adding go pushes this execution task to a scheduler which then schedules it to run on another internal thread
		// if you run only go task and then go ahead and do nothing, and program ends then the output may not be printed because the program is exited before output is printed
		// that is why we either wait for it to finish
	}

	wg.Wait()

}
