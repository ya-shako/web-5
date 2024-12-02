package main

import "fmt"

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
    defer close(outputStream)
    var previousValue string
    for value := range inputStream {
        if value != previousValue {
            outputStream <- value
            previousValue = value
        }
    }
}

func main() {
    inputStream := make(chan string)
    outputStream := make(chan string)

    go removeDuplicates(inputStream, outputStream)

    go func() {
        inputStream <- "a"
        inputStream <- "a"
        inputStream <- "b"
        inputStream <- "b"
        inputStream <- "c"
        inputStream <- "a"
        close(inputStream)
    }()

    for value := range outputStream {
        fmt.Println(value) // Ожидаемый вывод: a, b, c, a
    }
}
