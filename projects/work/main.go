package main

import (
	"fmt"
	"time"
	"sync"
)

func work() {
	time.Sleep(time.Millisecond * 50)
	fmt.Println("done")
}

func main() {
	var wg sync.WaitGroup

	// Увеличиваем счетчик WaitGroup на 10.
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения работы.
			work()          // Вызываем функцию work.
		}()
	}

	// Ждем завершения всех горутин.
	wg.Wait()
}
