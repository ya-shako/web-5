package main

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
    resultChan := make(chan int)

    go func() {
        defer close(resultChan) // Закрываем канал, чтобы освободить ресурсы.
        select {
        case val := <-firstChan: // Если получили значение из первого канала.
            resultChan <- val * val
        case val := <-secondChan: // Если получили значение из второго канала.
            resultChan <- val * 3
        case <-stopChan: // Если получили сигнал завершения.
            return
        }
    }()

    return resultChan
}

func main() {
	firstChan := make(chan int)
    secondChan := make(chan int)
    stopChan := make(chan struct{})

    resultChan := calculator(firstChan, secondChan, stopChan)

    // Пример использования:
    go func() {
        firstChan <- 4 // Отправляем значение в первый канал.
    }()

    fmt.Println(<-resultChan) // Ожидаем результат: 16 (4^2).

    go func() {
        secondChan <- 3 // Отправляем значение во второй канал.
    }()

    fmt.Println(<-resultChan) // Ожидаем результат: 9 (3*3).

    go func() {
        close(stopChan) // Отправляем сигнал завершения.
    }()

    _, ok := <-resultChan
    if !ok {
        fmt.Println("Канал закрыт, работа завершена.") // Ожидаем завершение работы.
    }
}
